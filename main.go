package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/paudelgaurav/gin-blog-backend/models"

	"github.com/joho/godotenv"
	"github.com/paudelgaurav/gin-blog-backend/controllers"
	"github.com/paudelgaurav/gin-blog-backend/middlewares"
)

func loadENV() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadENV()
	models.ConnectDatabase()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})
	router.GET("/blogs", controllers.GetAllBlogs)
	router.POST("/blogs", middlewares.AuthorizeJWT(), controllers.CreateBlog)
	router.DELETE("/blog/:id", middlewares.AuthorizeJWT(), controllers.DeleteBlog)
	router.PATCH("/blog/:id", middlewares.AuthorizeJWT(), controllers.UpdateBlog)

	router.GET("/tags", controllers.GetAllTags)

	router.POST("/register", controllers.RegisterUser)
	router.GET("/users", middlewares.AuthorizeJWT(), controllers.GetAllUsers)

	// auth

	router.POST("/auth-token", controllers.GetToken)

	router.Run()
}
