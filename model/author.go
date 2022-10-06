package model

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null;type:varchar(256)"`
	Country string `json:"country" gorm:"not null;type:varchar(256)"`
	DOB     string `json:"dob" gorm:"type:DATE; not null"`
	Gender  string `json:"gender" gorm:"type:varchar(10);not null"`
}
