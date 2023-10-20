package models

import (
	"gorm.io/gorm"
)

type ProfileData struct {
	gorm.Model
	Name        string
	Nickname    string
	Phone       string
	Email       string
	Linkedin    string
	Location    string
	Github      []Github        `gorm:"foreignKey:ProfileDataID"`
	Experience  []Experience    `gorm:"foreignKey:ProfileDataID"`
	Certificate []Certificate   `gorm:"foreignKey:ProfileDataID"`
	Educations  []EducationData `gorm:"foreignKey:ProfileDataID"`
	Projects    []Projects      `gorm:"foreignKey:ProfileDataID"`
	Skills      []Skills        `gorm:"foreignKey:ProfileDataID"`
	Website     []Website       `gorm:"foreignKey:ProfileDataID"`
	Languages   []Languages     `gorm:"foreignKey:ProfileDataID"`
	Interests   []Interests     `gorm:"foreignKey:ProfileDataID"`
}

type Interests struct {
	gorm.Model
	ProfileDataID uint
	Interests     string
}

type Languages struct {
	gorm.Model
	ProfileDataID uint
	Languages     string
}

type Experience struct {
	gorm.Model
	ProfileDataID uint
	Name          string
	Description   string
}

type Certificate struct {
	gorm.Model
	ProfileDataID uint
	Name          string
	Description   string
	Link          string
	Date          string
}

type EducationData struct {
	gorm.Model
	ProfileDataID uint
	NameEdu       string
	Description   string
	Date          string
}

type Projects struct {
	gorm.Model
	ProfileDataID uint
	Name          string
	Description   string
	Link          string
	Technologies  []Technologies `gorm:"foreignKey:ProjectsID"`
	Images        []Images       `gorm:"foreignKey:ProjectsID"`
	Date          string
}

type Images struct {
	gorm.Model
	ProjectsID uint
	Images     string
}

type Technologies struct {
	gorm.Model
	ProjectsID   uint
	Technologies string
}

type Skills struct {
	gorm.Model
	ProfileDataID uint
	Name          string
	Description   string
}

type Github struct {
	gorm.Model
	ProfileDataID uint
	Github        string
}

type Website struct {
	gorm.Model
	ProfileDataID uint
	Website       string
}
