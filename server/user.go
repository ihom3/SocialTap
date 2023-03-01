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
//var pswd = os.Getenv("MYSQL_PASSWORD")

// database url (whoever is testing it, don't forget to change this to your local connection)
const DNS = "root:root@tcp(127.0.0.1:3306)/social_tap?charset=utf8mb4&parseTime=True&loc=Local"

//var DNS = "root:" + pswd + "@tcp(127.0.0.1:3306)/social_tap?charset=utf8mb4&parseTime=True&loc=Local"

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

// func (sla *SocialType) Scan(src interface{}) error {
// 	return json.Unmarshal(src.([]byte), &sla)
// }

// func (sla SocialType) Value() (driver.Value, error) {
// 	val, err := json.Marshal(sla)
// 	return string(val), err
// }

type Socials struct {
	gorm.Model `json:"-"`
	UserID     uint       `json:"-"`
	Facebook   SocialType `gorm:"embedded;embeddedPrefix:facebook_" json:"facebook"`
	Snapchat   SocialType `gorm:"embedded;embeddedPrefix:snapchat_" json:"snapchat"`
	Instagram  SocialType `gorm:"embedded;embeddedPrefix:instagram_" json:"instagram"`
}

//{"social_name":"facebook","active_status": gorm:"default:false","url":"n/a"}

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

// getting all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodGet {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	// if r.Method != "GET" {
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }
	//err
	if err != nil {
		http.Error(w, "invalid", http.StatusBadRequest)
		return
	}

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

func GetSocial(w http.ResponseWriter, r *http.Request) {

}

// func AddUserSocial(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var social Socials
// 	//decoding the data from the body of the request
// 	//whatever data we're getting we are decoding it with a reference to user
// 	json.NewDecoder(r.Body).Decode(&social)
// 	//to save the data in the database
// 	DB.Create(&social)
// 	//to parse the data back to the browser, w=response writer
// 	json.NewEncoder(w).Encode(social)

// }

func AddUserSocial(w http.ResponseWriter, r *http.Request) {
	foundUser, err := findUserByCode("hello")
	if err != nil {

	}
	json.NewEncoder(w).Encode(foundUser)

}

// creating a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var user User
	//decoding the data from the body of the request
	//whatever data we're getting we are decoding it with a reference to user

	// json.NewDecoder(r.Body).Decode(&user)
	// //to save the data in the database
	// DB.Create(&user)
	// //to parse the data back to the browser, w=response writer
	// json.NewEncoder(w).Encode(user)

	user := User{
		FirstName:      "Elsa",
		UserEmail:      "ian.n.",
		LastName:       "Osmani",
		StickerCode:    "efgh",
		BioText:        "Hello my name is Elsa",
		ProfilePicture: "1",
		SocialList: Socials{
			Facebook: SocialType{
				SocialName: "Facebook",
				Status:     true,
				URL:        "/ian",
			},
			Instagram: SocialType{
				SocialName: "Instagram",
				Status:     true,
				URL:        "/ian",
			},
		},
	}
	DB.Create(&user)
	foundUser, err := findUserByCode("hello")
	if err != nil {

	}
	json.NewEncoder(w).Encode(foundUser)

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
	// upload picture
	// to store it in the directory temporary:
	tempFile, err2 := ioutil.TempFile("profile-pictures", "user-*.jpg")
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




