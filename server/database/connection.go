package database

import (
	"backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:rootroot@/social-tap?parseTime=true"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database!")
	}
	connection.AutoMigrate(&models.User{}, &models.UnregisteredUser{}, &models.Social{})

	DB = connection
}
