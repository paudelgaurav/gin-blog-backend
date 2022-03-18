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
	router.POST("/blogs", controllers.CreateBlog)
	router.DELETE("/blog/:id", controllers.DeleteBlog)
	router.PATCH("/blog/:id", controllers.UpdateBlog)

	router.GET("/tags", controllers.GetAllTags)

	router.Run()
}
