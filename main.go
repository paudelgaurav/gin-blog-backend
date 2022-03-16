package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/paudelgaurav/gin-blog-backend/models"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})
	router.GET("/blogs", getAllBlogs)
	router.Run()
}

// get all blogs

func getAllBlogs(c *gin.Context) {
	var books []models.Blog
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}
