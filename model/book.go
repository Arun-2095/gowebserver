package model

import (
	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null;type:varchar(256)"`
	Author string `json:"arthor" gorm:"not null;type:varchar(256)"`
	Genre  string `json:"genre" gorm:"not null;type:varchar(256)"`
}
