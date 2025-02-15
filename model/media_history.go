package model

import (
	"gorm.io/gorm"
)

// type MedicalRecord struct {
// 	//gorm.Model
// 	PatientID        string `gorm:"uniqueIndex;not null"`
// 	Name             string `gorm:"not null"`
// 	Age              int
// 	CurrentCondition string
// }

// type MedicalHistory struct {
// 	//gorm.Model
// 	RecordID    string `gorm:"uniqueIndex;not null"`
// 	PatientID   string `gorm:"not null"`
// 	Action      string
// 	Description string
// 	Timestamp   time.Time
// }

type PatientV2 struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Age     int
	Gender  string
	Address string
}

type MedicalRecordV2 struct {
	gorm.Model
	PatientID  uint   `gorm:"not null"`
	Diagnosis  string `gorm:"type:text"`
	Treatment  string `gorm:"type:text"`
	RecordDate string
}
