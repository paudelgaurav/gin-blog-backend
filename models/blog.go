package models

type Blog struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    []Tag  `json:"tags" gorm:"many2many:blog_tag;"`
}
