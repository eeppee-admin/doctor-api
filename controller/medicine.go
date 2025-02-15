package controller

import (
	"net/http"
	"signin_and_signup/config"
	modelv2 "signin_and_signup/model_v2"

	"github.com/gin-gonic/gin"
)

// 获取所有药品
func GetMedicines(c *gin.Context) {
	var medicines []modelv2.MedicineV2
	_ = medicines
	err := config.DB.Find(&medicines).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch medicines"})
		return
	}
	c.JSON(http.StatusOK, medicines)
}

// 获取单个药品
func GetMedicine(c *gin.Context) {
	id := c.Param("id")

	var medicine modelv2.MedicineV2
	err := config.DB.First(&medicine, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medicine not found"})
		return
	}
	c.JSON(http.StatusOK, medicine)
}

// 创建药品
func CreateMedicine(c *gin.Context) {
	var medicine modelv2.MedicineV2
	if err := c.ShouldBindJSON(&medicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := config.DB.Create(&medicine).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create medicine"})
		return
	}
	c.JSON(http.StatusOK, medicine)
}

// 更新药品
func UpdateMedicine(c *gin.Context) {
	id := c.Param("id")
	var medicine modelv2.MedicineV2
	_ = id
	_ = medicine
	if err := config.DB.First(&medicine, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medicine not found"})
		return
	}
	if err := c.ShouldBindJSON(&medicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := config.DB.Save(&medicine).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update medicine"})
		return
	}
	c.JSON(http.StatusOK, medicine)
}

// 删除药品
func DeleteMedicine(c *gin.Context) {
	id := c.Param("id")
	var medicine modelv2.MedicineV2
	_ = id
	_ = medicine
	if err := config.DB.Delete(&medicine, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete medicine"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Medicine deleted successfully"})
}
