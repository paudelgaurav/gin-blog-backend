package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/paudelgaurav/gin-blog-backend/models"
)

// creating this to validate incomming post data
type CreateBookInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// get all blogs

func GetAllBlogs(c *gin.Context) {
	var books []models.Blog
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

//For post request /blogs

func CreateBlog(c *gin.Context) {
	var input CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//creating new blog
	blog := models.Blog{Title: input.Title, Content: input.Content}
	models.DB.Create(&blog)

	c.JSON(http.StatusOK, gin.H{"data": blog})
}

// For deleting blog post
func DeleteBlog(c *gin.Context) {
	// Get model if exists
	var blog models.Blog
	if err := models.DB.First(&blog, "id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "Blog not found"})
	}
	models.DB.Delete(&blog)
	c.JSON(http.StatusNoContent, gin.H{"deleted": true})
}
