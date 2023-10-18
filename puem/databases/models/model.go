package models

import (
	"gorm.io/gorm"
)

type ProfileData struct {
	gorm.Model
	Name     string
	Nickname string
	Phone    string
	Email    string
	Github   string
	Linkedin string
	Websites []Website `gorm:"foreignKey:ProfileDataID"`
	Location string
}

type Website struct {
	gorm.Model
	Website       string
	ProfileDataID uint
}
