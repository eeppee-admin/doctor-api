package main

import (
	"signin_and_signup/config"
	"signin_and_signup/route"

	"github.com/gin-gonic/gin"
)

func main() {

	config.InitDatabase()

	// InitDBWithSqlite()

	InitDBWithMySql()

	r := gin.Default()

	route.SetupUserRoutes(r)
	route.SetupPatientRoutes(r)
	// 注册药品相关路由
	route.SetupMedicineRoutes(r)
	// 启动服务器
	r.Run(":8080")
}

// func registerUser(c *gin.Context) {
// 	var newUser model.User
// 	if err := c.ShouldBindJSON(&newUser); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// 加密密码
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	newUser.Password = string(hashedPassword)

// 	// 创建用户
// 	result := config.DB.Create(&newUser)
// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!"})
// }

// func loginUser(c *gin.Context) {
// 	var login model.User
// 	if err := c.ShouldBindJSON(&login); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// 查找用户
// 	var user model.User
// 	result := config.DB.Where("username = ?", login.Username).First(&user)
// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
// 		return
// 	}

// 	// 验证密码
// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully!"})
// }
