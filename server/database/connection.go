package database

import (
	"backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:rootroot@/social_tap"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database!")
	}
	connection.AutoMigrate(&models.User{}, &models.UnregisteredUser{})

	DB = connection
}
