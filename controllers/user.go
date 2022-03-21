package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paudelgaurav/gin-blog-backend/models"
	"github.com/paudelgaurav/gin-blog-backend/utils"
)

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterUser(ctx *gin.Context) {
	var user models.User
	var input RegisterUserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	duplicate_user := models.DB.Find(&user, "email = ?", input.Email)
	if duplicate_user.RowsAffected > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": "User with this email already exists"})
		return
	}

	utils.HashPassword(&input.Password)

	user = models.User{Name: input.Name, Email: input.Email, Password: input.Password}
	models.DB.Create(&user)

	ctx.JSON(http.StatusCreated, gin.H{"detail": "Registered successfully"})
}

func GetAllUsers(ctx *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}
