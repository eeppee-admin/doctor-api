package fakedata

import (
	"signin_and_signup/config"
	modelv2 "signin_and_signup/model_v2"
)

func SeedPatientV2Data() {
	// 创建病人
	patients := []modelv2.PatientV2{
		{Name: "张三", Age: 30, Gender: "Male", Address: "北京市"},
		{Name: "李四", Age: 25, Gender: "Female", Address: "上海市"},
		{Name: "王五", Age: 40, Gender: "Male", Address: "广州市"},
	}

	// 插入病人
	for _, patient := range patients {
		config.DB.Create(&patient)
	}
}
