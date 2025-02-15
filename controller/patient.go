package controller

import (
	"log"
	"net/http"
	modelv2 "signin_and_signup/model_v2"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// todo: replace db to config.DB

func GetPatient(c *gin.Context) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/prod?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection established.")

	// 自动迁移表结构
	err = db.AutoMigrate(&modelv2.PatientV2{}, &modelv2.MedicalRecordV2{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Seeding data...")

	id := c.Param("id")
	var patient modelv2.PatientV2
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
	var records []modelv2.MedicalRecordV2
	if err := db.Where("patient_id = ?", patientID).Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch records"})
		return
	}
	c.JSON(http.StatusOK, records)
}
