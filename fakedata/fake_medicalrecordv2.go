package fakedata

import (
	"signin_and_signup/config"
	modelv2 "signin_and_signup/model_v2"
	"time"
)

func SeedMedicineRecordV2Data() {
	// 创建病历记录
	medicalRecords := []modelv2.MedicalRecordV2{
		{PatientID: 1, Diagnosis: "感冒", Treatment: "多喝水，休息", RecordDate: time.Now().Format("2006-01-02")},
		{PatientID: 1, Diagnosis: "高血压", Treatment: "服用降压药", RecordDate: time.Now().AddDate(0, 0, -30).Format("2006-01-02")},
		{PatientID: 2, Diagnosis: "糖尿病", Treatment: "控制饮食，注射胰岛素", RecordDate: time.Now().AddDate(0, 0, -15).Format("2006-01-02")},
		{PatientID: 3, Diagnosis: "颈椎病", Treatment: "物理治疗", RecordDate: time.Now().AddDate(0, 0, -7).Format("2006-01-02")},
	}

	// 插入病历记录
	for _, record := range medicalRecords {
		config.DB.Create(&record)
	}
}
