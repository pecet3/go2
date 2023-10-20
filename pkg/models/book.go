package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pecet3/go2/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func initial() {
	config.Connect()
	db = config.GetDB()
}
