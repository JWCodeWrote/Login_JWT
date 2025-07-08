package main

import (
	"log"
	"os"

	"github.com/JWCodeWrote/Login_JWT/config"
	"github.com/JWCodeWrote/Login_JWT/controllers"
	_ "github.com/JWCodeWrote/Login_JWT/docs"
	"github.com/JWCodeWrote/Login_JWT/models"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title        Login JWT API
// @version      1.0
// @description  Gin + MySQL + JWT 的 API 範例
// @host         localhost:8080
// @BasePath     /
func main() {
	godotenv.Load()
	db := config.InitDB()
	db.AutoMigrate(&models.User{})

	authC := controllers.AuthController{DB: db}
	r := gin.Default()
	r.POST("/register", authC.Register)
	r.POST("/login", authC.Login)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := r.Group("/")
	auth.Use(controllers.JWTAuth())
	auth.GET("/me", func(c *gin.Context) {
		uid, _ := c.Get("userID")
		c.JSON(200, gin.H{"userID": uid})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)

	url := "http://localhost:" + port + "/swagger/index.html"
	log.Printf("Swagger running at %s", url)
}
