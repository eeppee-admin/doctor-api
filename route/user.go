package route

import (
	"signin_and_signup/controller"

	"github.com/gin-gonic/gin"
)

// 注册路由
func SetupUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/api/users")
	{
		userGroup.POST("/register", controller.RegisterUser)
		userGroup.POST("/login", controller.LoginUser)
	}
}
