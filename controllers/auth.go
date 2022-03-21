package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paudelgaurav/gin-blog-backend/models"
	"github.com/paudelgaurav/gin-blog-backend/utils"
)

type GetTokenInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func GetToken(ctx *gin.Context) {
	var input GetTokenInput
	var userModel models.User

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.First(&userModel, "email = ?", input.Email).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": "no user found"})
		return
	}

	if isPasswordMatched := utils.ComparePassword(userModel.Password, input.Password); isPasswordMatched {
		token := utils.GenerateToken(userModel.ID)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Successfully SignIN", "token": token})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{"detail": "Password doesnot match"})
}
