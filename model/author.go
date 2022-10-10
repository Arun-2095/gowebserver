package model

import (
	"webserver/config"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null;type:varchar(256)"`
	Country string `json:"country" gorm:"not null;type:varchar(256)"`
	DOB     string `json:"dob" gorm:"type:DATE; not null"`
	Gender  string `json:"gender" gorm:"type:varchar(10);not null"`
}

type AutherQueries struct{}

func (A AutherQueries) GetAuthors() ([]Author, *gorm.DB) {
	var db = config.GetDatabase()
	var authorsList []Author

	result := db.Find(&authorsList)

	return authorsList, result
}

func (A AutherQueries) GetAuthor(name string) (Author, *gorm.DB) {
	var db = config.GetDatabase()
	var author Author

	result := db.Where("name=?", name).Find(&author)

	return author, result
}

func (A AutherQueries) CreateAuthor(newAuthor *Author) (Author, *gorm.DB) {
	var db = config.GetDatabase()
	var author Author

	result := db.Create(newAuthor)

	return author, result
}

func (A AutherQueries) EditAuthor(newAuthor *Author) (Author, *gorm.DB) {
	var db = config.GetDatabase()
	var author Author

	result := db.Where("name=?", newAuthor.Name).Updates(newAuthor)

	return author, result
}
