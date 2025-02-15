package controller

import (
	"errors"
	"net/http"
	"signin_and_signup/config"
	modelv2 "signin_and_signup/model_v2"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// todo: replace db to config.DB

func RegisterUser(c *gin.Context) {
	var newUser modelv2.UserV2
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newUser.Password = string(hashedPassword)

	// 创建用户
	result := config.DB.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!"})
}

func LoginUser(c *gin.Context) {
	var login modelv2.UserV2
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查找用户
	var user modelv2.UserV2
	result := config.DB.Where("username = ?", login.Username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully!"})
}
