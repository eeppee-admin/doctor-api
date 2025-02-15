package controller

import (
	"log"
	"net/http"
	"signin_and_signup/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetPatient(c *gin.Context) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/prod?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection established.")

	// 自动迁移表结构
	err = db.AutoMigrate(&model.PatientV2{}, &model.MedicalRecordV2{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Seeding data...")

	id := c.Param("id")
	var patient model.PatientV2
	if err := db.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	c.JSON(http.StatusOK, patient)
}

func GetMedicalRecords(c *gin.Context) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/prod?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection established.")

	patientID := c.Param("id")
	var records []model.MedicalRecordV2
	if err := db.Where("patient_id = ?", patientID).Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch records"})
		return
	}
	c.JSON(http.StatusOK, records)
}
