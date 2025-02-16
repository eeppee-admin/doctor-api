package controller

import (
	"net/http"
	"signin_and_signup/config"
	modelv2 "signin_and_signup/model_v2"

	"github.com/gin-gonic/gin"
)

func GetPatient(c *gin.Context) {
	id := c.Param("id")
	var patient modelv2.PatientV2
	if err := config.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	c.JSON(http.StatusOK, patient)
}

func GetMedicalRecords(c *gin.Context) {
	patientID := c.Param("id")
	var records []modelv2.MedicalRecordV2
	if err := config.DB.Where("patient_id = ?", patientID).Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch records"})
		return
	}
	c.JSON(http.StatusOK, records)
}
