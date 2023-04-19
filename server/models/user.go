package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	ID         uint     `gorm:"primarykey" json:"id"`
	Email      string   `gorm:"unique" json:"email"`
	Password   []byte   `json:"-"`
	Code       string   `json:"code"`
	Bio        string   `gorm:"default:'No bio'" json:"bio"`
	PictureURL string   `json:"pictureURL"`
	Role       string   `gorm:"default:'user'" json:"role"`
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	Active     bool     `gorm:"default:true" json:"active"`
	Socials    []Social `json:"socials" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`
}

type UnregisteredUser struct {
	gorm.Model `json:"-"`
	ID         uint
	Code       string
}
type Social struct {
	gorm.Model `json:"-"`
	Name       string `json:"name"`
	Link       string `json:"link"`
	Active     bool   `json:"active"`
	UserID     uint   `json:"-"`
}
