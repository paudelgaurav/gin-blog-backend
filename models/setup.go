package models

import (
	"github.com/jinzhu/gorm"
	// why _ ?
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to databse")
	}
	database.AutoMigrate(&Blog{}, &Tag{})

	DB = database
}
