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
.
.
.
.
.
.
.
.
.


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
		FirstName: "Admin",
		LastName:  "Admin",
		Email:     "admin@admin.com",
		Password:  password,
		Role:      "admin",
		BioText:   "This is admin.",
		Code:      "123456",
		Socials: []models.Social{
			{
				Name:   "Facebook",
				Link:   "https://www.facebook.com/admin",
				Active: true,
			}, {
				Name:   "Instagram",
				Link:   "https://www.instagram.com/admin",
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
		if f != "default.jpg" {
			err = os.Remove(f)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"message": err.Error(),
				})
			}
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

	err5 := database.DB.Model(&user).Update("profile_picture", tempFile.Name())

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
	files, err := filepath.Glob(user.ProfilePicture)
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
