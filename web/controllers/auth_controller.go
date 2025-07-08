package controllers

import (
	"github.com/JWCodeWrote/Login_JWT/models"
	"github.com/JWCodeWrote/Login_JWT/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct{ DB *gorm.DB }

// Register godoc
// @Summary  註冊新使用者
// @Tags     Auth
// @Accept   json
// @Produce  json
// @Param    body  body  models.RegisterRequest  true  "帳密"
// @Success  201   {object} map[string]string     "message"
// @Failure  400   {object} map[string]string     "error"
// @Failure  500   {object} map[string]string     "server error"
// @Router   /register [post]
func (a *AuthController) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user := models.User{Email: req.Email, Password: req.Password}
	if err := a.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Email taken"})
		return
	}
	c.JSON(201, gin.H{"message": "Registered"})
}

// Login godoc
// @Summary     使用者登入
// @Description 使用 email 與 password 進行登入，成功後回傳 JWT token
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       credentials  body  models.LoginRequest  true  "登入資訊"
// @Success     200          {object} map[string]string  "包含 access_token"
// @Failure     400          {object} map[string]string  "request error"
// @Failure     401          {object} map[string]string  "authentication failed"
// @Failure     500          {object} map[string]string  "server error"
// @Router      /login [post]
func (a *AuthController) Login(c *gin.Context) {
	var req models.LoginRequest
	c.ShouldBindJSON(&req)
	var u models.User
	if err := a.DB.Where("email = ?", req.Email).First(&u).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}
	if !u.CheckPassword(req.Password) {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}
	token, _ := utils.GenerateToken(u.ID)
	c.JSON(200, gin.H{"access_token": token})
}
