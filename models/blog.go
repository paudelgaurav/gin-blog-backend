package models

type Blog struct {
	ID      uint   `json: "id" gorm:"primary_key"`
	Title   string `json: "title"`
	Content string `json: "content"`
}
