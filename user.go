package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// database variable
var DB *gorm.DB

// error variable
var err error

// database url (whoever is testing it, don't forget to change this to your local connection)
const DNS = "root:feynman65@tcp(127.0.0.1:3306)/social_tap?charset=utf8mb4&parseTime=True&loc=Local"

// creating a struct(class) to store the different data types in order for us to be able to save these data in the databse
type User struct {
	gorm.Model `json:"-"`
	UserName   string `json:"username"`
	FirstName  string `json: "firstname"`
	LastName   string `json: "lastname"`
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
}

//defining the handler functions

// getting all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// getting one user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Vars returns the route variables for the current request
	var user User
	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
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
