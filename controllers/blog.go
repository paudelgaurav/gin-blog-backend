package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/paudelgaurav/gin-blog-backend/models"
)

// creating this to validate incomming post data
type CreateBlogInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdateBlogInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// get all blogs

func GetAllBlogs(c *gin.Context) {
	var blogs []models.Blog
	models.DB.Find(&blogs)
	c.JSON(http.StatusOK, gin.H{"data": blogs})
}

//For post request /blogs

func CreateBlog(c *gin.Context) {
	var input CreateBlogInput

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

// updating a blog post

func UpdateBlog(c *gin.Context) {

	var blog models.Blog
	if err := models.DB.First(&blog, "id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "Blog not found"})
	}

	// validate input data
	var input UpdateBlogInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&blog).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": blog})
}
