package controllers

import (
	"backend/database"
	"backend/models"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const JWTSecret = "secret"

func Register(c *fiber.Ctx) error {
	//parse user input
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//check if user is unregistered
	unregistered := database.DB.Find(&models.UnregisteredUser{}, "code = ?", data["code"])
	if unregistered.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid code",
		})
	}
	//hash password
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	//create user object
	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  password,
		Code:      data["code"],
	}
	//create user in database
	err := database.DB.Create(&user)
	if err.Error != nil {
		return c.Status(400).JSON(err.Error)
	}
	//remove unregistered code
	database.DB.Delete(&models.UnregisteredUser{}, "code = ?", data["code"])

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Login(c *fiber.Ctx) error {
	//check if user is already logged in

	cookieCheck := c.Cookies("jwt")
	_, CheckLoginError := jwt.ParseWithClaims(cookieCheck, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})
	if CheckLoginError == nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "User is Already Logged In",
		})
	}
	//parse user input
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(JWTSecret))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not log in",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func RegisterNewCode(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)
	if user.Role != "admin" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	unregisteredCheck := database.DB.Find(&models.UnregisteredUser{}, "code = ?", data["code"])
	if unregisteredCheck.RowsAffected != 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Code Already Exists",
		})
	}
	registeredCHeck := database.DB.Find(&models.User{}, "code = ?", data["code"])
	if registeredCHeck.RowsAffected != 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Code Already Exists",
		})
	}
	unregistered := models.UnregisteredUser{
		Code: data["code"],
	}
	database.DB.Create(&unregistered)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func IDRoute(c *fiber.Ctx) error {
	code := c.Params("id")
	if code == "" {
		return c.JSON(fiber.Map{
			"message": "User Not Found",
		})
	}
	var user models.User
	err := database.DB.Preload("Socials").First(&user, "code = ?", code)
	if err.RowsAffected == 0 {
		unregistered := database.DB.First(&models.UnregisteredUser{}, "code = ?", code)
		if unregistered.RowsAffected == 0 {
			return c.JSON(fiber.Map{
				"message": "User Not Found",
			})
		}
		return c.JSON(fiber.Map{
			"message": "User Not Registered",
		})
	}
	return c.JSON(user)
}

func UpdateBio(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)
	if user.ID == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	database.DB.Model(&user).Update("bio", data["bio"])
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func UpdateName(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)
	if user.ID == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	database.DB.Model(&user).Update("first_name", data["first_name"])
	database.DB.Model(&user).Update("last_name", data["last_name"])

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
func UpdateEmail(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)
	if user.ID == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if err := database.DB.Model(&user).Update("email", data["email"]); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email Already Exists",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
func UpdatePassword(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)
	if user.ID == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["oldPassword"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data["newPassword"]), 8)
	database.DB.Model(&user).Update("password", string(hashedPassword))
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func RegisterPersonalUser(c *fiber.Ctx) error {

	password, err := bcrypt.GenerateFromPassword([]byte("password"), 14)
	if err != nil {
		return c.JSON(err)
	}
	user := models.User{
		FirstName: "Ian",
		LastName:  "Blasko",
		Email:     "ian@gmail.com",
		Password:  password,
		Role:      "admin",
		Bio:       "This is my bio.",
		Code:      "123456",
		Socials: []models.Social{
			{
				Name:   "Facebook",
				Link:   "https://www.facebook.com/ian.blasko",
				Active: true,
			}, {
				Name:   "Instagram",
				Link:   "https://www.instagram.com/ianblasko/",
				Active: true,
			},
		},
	}
	database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&user)
	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

func GetUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Preload("Socials").First(&user, "id = ?", claims.Issuer)
	if user.ID == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	return c.JSON(user)
}
func IsLoggedIn(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})
	if err != nil {
		return c.JSON(fiber.Map{
			"status": false,
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)
	if user.ID == 0 {
		return c.JSON(fiber.Map{
			"status": false,
		})
	}
	return c.JSON(fiber.Map{
		"status": true,
	})
}
func UpdateSocials(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)
	if user.ID == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	var data []models.Social
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	database.DB.Model(&user).Association("Socials").Replace(&data)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func UpdateProfilePicture(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)
	if user.ID == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	file, err := c.FormFile("profile_picture")
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	//defer file.Close()

	fileName := "user-id-" + strconv.FormatUint(uint64(user.ID), 10) + "-*.jpg"
	filePath := filepath.Join("profile_pictures", fileName)

	// check if file exists
	files, err := filepath.Glob(filePath)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// delete file if it exists
	for _, f := range files {
		err = os.Remove(f)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	// upload picture
	// to store it in the directory temporary:
	tempFile, err2 := ioutil.TempFile("profile_pictures", "user-id-"+strconv.FormatUint(uint64(user.ID), 10)+"-*.jpg")

	if err2 != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err2.Error(),
		})
	}
	defer tempFile.Close()

	fileBytes, err3 := file.Open()
	if err3 != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err3.Error(),
		})
	}
	defer fileBytes.Close()

	fileBytesBytes, err4 := ioutil.ReadAll(fileBytes)
	if err4 != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err4.Error(),
		})
	}
	tempFile.Write(fileBytesBytes)

	err5 := database.DB.Model(&user).Update("picture_url", tempFile.Name())

	if err5.Error != nil {
		return c.JSON(fiber.Map{
			"message": err5.Error,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Successfully updated profile picture",
	})
}
func GetProfilePicture(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.DB.Where("id = ?", id).First(&user)
	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User Not Found",
		})
	}

	// check if file exists
	files, err := filepath.Glob(user.PictureURL)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if len(files) == 0 {
		// return a default picture or an error message
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Profile Picture Not Found",
		})
	}

	// return the first file that matches the pattern
	return c.SendFile(files[0])
}
