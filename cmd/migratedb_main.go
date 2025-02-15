package main

import (
	_ "encoding/json"
	"errors"
	"net/http"

	"signin_and_signup/controller"
	"signin_and_signup/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	_ "gorm.io/gorm/clause"
)

var db *gorm.DB

func main() {

	// 连接数据库
	// var err error
	// db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("failed to connect database")
	// }

	// // 自动迁移
	// db.AutoMigrate(&model.User{})

	InitDBWithSqlite()

	InitDBWithMySql()

	InitDBWithMongo()

	r := gin.Default()

	// 注册路由
	r.POST("/register", registerUser)
	r.POST("/login", loginUser)
	// 病人相关路由
	r.GET("/patients/:id", controller.GetPatient)
	r.GET("/patients/:id/records", controller.GetMedicalRecords)
	// 启动服务器
	r.Run(":8080")
}

func registerUser(c *gin.Context) {
	var newUser model.User
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
	result := db.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!"})
}

func loginUser(c *gin.Context) {
	var login model.User
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查找用户
	var user model.User
	result := db.Where("username = ?", login.Username).First(&user)
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
