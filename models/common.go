package models

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name string `json:"name"`
}
