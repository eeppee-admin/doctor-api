package fakedata

import (
	"log"
	"signin_and_signup/config"
	modelv2 "signin_and_signup/model_v2"
)

func SeedMedicineV2Data() {
	log.Println("Seeding MedicineV2 data...")
	medicines := []modelv2.MedicineV2{
		{Name: "阿莫西林", Description: "抗生素", Price: 10.5, Stock: 100},
		{Name: "布洛芬", Description: "解热镇痛药", Price: 8.0, Stock: 50},
		{Name: "维生素C", Description: "维生素补充剂", Price: 5.0, Stock: 200},
		{Name: "胰岛素", Description: "糖尿病治疗药物", Price: 20.0, Stock: 30},
	}

	for _, medicine := range medicines {
		config.DB.Create(&medicine)
	}
	log.Println("MedicineV2 data seeding completed.")
}
