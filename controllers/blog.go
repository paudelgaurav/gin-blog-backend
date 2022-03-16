package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/paudelgaurav/gin-blog-backend/models"
)

// get all blogs

func GetAllBlogs(c *gin.Context) {
	var books []models.Blog
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}
