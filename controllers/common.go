package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paudelgaurav/gin-blog-backend/models"
)

func GetAllTags(c *gin.Context) {
	var tags []models.Tag
	models.DB.Find(&tags)
	c.JSON(http.StatusOK, gin.H{"data": tags})
}
