package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//_ "image/png"
)

// database variable
var DB *gorm.DB

// error variable
var err error

// to protect the password we generate an environmental variable
var pswd = os.Getenv("MYSQL_PASSWORD")
var dbName = os.Getenv("DBNAME")
var DNS = "root:" + pswd + "@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

// creating a struct(class) to store the different data types in order for us to be able to save these data in the databse
type User struct {
	gorm.Model     `json:"-"`
	UserEmail      string  `json:"user_email"`
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	StickerCode    string  `json:"sticker_code"`
	BioText        string  `json:"bio_text"`
	ProfilePicture string  `json:"profile_picture"`
	SocialList     Socials `json:"social_list"`
}

type UnregisteredCodes struct {
	gorm.Model
	StickerCode string `json:"sticker_code"`
}

type SocialType struct {
	SocialName string `json:"name"`
	Status     bool   `json:"status"`
	URL        string `json:"url"`
}

type Socials struct {
	gorm.Model `json:"-"`
	UserID     uint       `json:"-"`
	Facebook   SocialType `gorm:"embedded;embeddedPrefix:facebook_" json:"facebook"`
	Snapchat   SocialType `gorm:"embedded;embeddedPrefix:snapchat_" json:"snapchat"`
	Instagram  SocialType `gorm:"embedded;embeddedPrefix:instagram_" json:"instagram"`
}


// once the model is created we are going to initialize the database and enable automigration
func InitialMigration() {
	//here we will define our database details
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&UnregisteredCodes{})
	DB.AutoMigrate(&Socials{})
}

//defining the handler functions

// the /{id} route
func IDRoute(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var unregistered UnregisteredCodes
	//first, query the unregistered codes table
	unregisteredCheck := DB.First(&unregistered, "sticker_code = ?", params["id"])
	//if it's not in the unregistered table, query the registered users table in the database
	if unregisteredCheck.Error != nil {
		var registered User
		registeredCheck := DB.First(&registered, "sticker_code = ?", params["sticker_code"])
		if registeredCheck.Error != nil { // if code doesn't exist
			json.NewEncoder(w).Encode("User Not Found")
		} else { // if the code exists
			json.NewEncoder(w).Encode(registered)
		}
	} else { // if the code is in the unregistered table
		json.NewEncoder(w).Encode(unregistered)
	}
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["sticker_code"]
	var registered User
	registeredCheck := DB.First(&registered, "sticker_code = ?", code)
	if registeredCheck.Error != nil { // if user doesn't exist
		//the user just created their account
		//first, check if it's a valid code in the unregistered table
		var unregistered UnregisteredCodes
		unregisteredCheck := DB.First(&unregistered, "sticker_code = ?", code)
		if unregisteredCheck.Error != nil {
			json.NewEncoder(w).Encode("Invalid Request")
		} else {
			CreateUser(w, r)
		}
	} else { // if user exists
		json.NewEncoder(w).Encode(registered)
	}
}

func AddCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user UnregisteredCodes
	//decoding the data from the body of the request
	// data we're getting we are decoding it with a reference to user

	json.NewDecoder(r.Body).Decode(&user)
	//to save the data in the database
	DB.Create(&user)

	//to parse the data back to the browser, w=response writer
	json.NewEncoder(w).Encode(user)
}


// getting all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)

}

// getting one user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Vars returns the route variables for the current request
	//var user User
	//DB.First(&user, params["id"])
	user, err := findUserByCode(params["id"])
	if err != nil {

	}
	json.NewEncoder(w).Encode(user)

}

// enhanced function to get a user
func GetUserEnhanced(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var user User
	result := DB.Where("sticker_code = ?", id).First(&user)
	if result.Error != nil {
		log.Printf("database error: %v", result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		log.Printf("JSON marshal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}



func GetSocial(w http.ResponseWriter, r *http.Request) {

}

func AddUserSocial(w http.ResponseWriter, r *http.Request) {
	foundUser, err := findUserByCode("hello")
	if err != nil {

	}
	json.NewEncoder(w).Encode(foundUser)

}

// creating a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	//decoding the data from the body of the request
	//whatever data we're getting we are decoding it with a reference to user

	json.NewDecoder(r.Body).Decode(&user)
	//to save the data in the database
	DB.Create(&user)
	//to parse the data back to the browser, w=response writer
	json.NewEncoder(w).Encode(user)
}

// updating information about a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

// deleting a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The user is deleted successfully!")
}

func UpdateProfilePicture(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	r.ParseMultipartForm(10 * 1024 * 1024) // Limit 10 MB

	file, handler, err := r.FormFile("profile_picture")

	//for file name: handler.Filename
	//for file size: handler.Size
	//for file type: handler.Header.Get("Content-Type")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println("File name: ", handler.Filename)
	
	fileName := "user-id-" + params["id"] + "-*.jpg"
	filePath := filepath.Join("profile_pictures", fileName)

	// check if file exists
	files, err := filepath.Glob(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// delete file if it exists
	for _, f := range files {
		err = os.Remove(f)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// upload picture
	// to store it in the directory temporary:
	tempFile, err2 := ioutil.TempFile("profile-pictures", "user-id-" + params["id"] + ".jpg")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer tempFile.Close()

	fileBytes, err3 := ioutil.ReadAll(file)
	if err3 != nil {
		fmt.Println(err3)
	}
	tempFile.Write(fileBytes)
}

func findUserByCode(code string) (User, error) {
	var foundUser User
	err := DB.Model(&User{}).Preload("SocialList").Find(&foundUser, "sticker_code = ?", code).Error
	return foundUser, err
}

// deleting a code
func DeleteCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user UnregisteredCodes
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The code is deleted successfully.")
}




