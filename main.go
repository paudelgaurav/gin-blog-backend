package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/paudelgaurav/gin-blog-backend/models"

	"github.com/paudelgaurav/gin-blog-backend/controllers"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})
	router.GET("/blogs", controllers.GetAllBlogs)
	router.Run()
}
