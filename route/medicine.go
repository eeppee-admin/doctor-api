package route

import (
	"signin_and_signup/controller"

	"github.com/gin-gonic/gin"
)

// 定义药品相关的路由
func SetupMedicineRoutes(r *gin.Engine) {
	medicineGroup := r.Group("/api/medicines")
	{
		medicineGroup.GET("", controller.GetMedicines)
		medicineGroup.GET(":id", controller.GetMedicine)
		medicineGroup.POST("", controller.CreateMedicine)
		medicineGroup.PUT(":id", controller.UpdateMedicine)
		medicineGroup.DELETE(":id", controller.DeleteMedicine)
	}
}
