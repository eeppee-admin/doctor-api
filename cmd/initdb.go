package main

import (
	"fmt"
	"log"
	"signin_and_signup/fakedata"
	modelv2 "signin_and_signup/model_v2"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// sqlite做modelv2初始化,但是没有插入数据
func InitDBWithSqlite() {
	// 连接到SQLite数据库（这里使用SQLite作为示例，你可以根据需要替换为其他数据库）
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// 删除所有表
	if err := dropAllTables(db); err != nil {
		log.Fatal("failed to drop tables:", err)
	}

	// 自动迁移所有模型
	if err := db.AutoMigrate(
		&modelv2.UserV2{},
		&modelv2.Patient{},
		&modelv2.Doctor{},
		&modelv2.Visit{},
		&modelv2.Order{},
		&modelv2.LabTest{},
		&modelv2.MedicalRecordV2{},
		&modelv2.MedicineV2{},
		&modelv2.AdmissionNotice{},
		&modelv2.WesternMedicline{},
		&modelv2.ChineseMedicline{},
		&modelv2.ICD_10{},
	); err != nil {
		log.Fatal("failed to auto migrate:", err)
	}
}

// 使用mysql做生产环境
func InitDBWithMySql() {
	// 获取数据库 DSN 字符串
	dsn := "root:123456@tcp(127.0.0.1:3306)/prod?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接到数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接mysql失败: %v", err)
	}
	log.Printf("%v", db)
	// 打印数据库连接成功信息
	fmt.Printf("连接mysql成功: %s\n", dsn)

	// 删除所有表
	if err := dropAllMySqlTables(db); err != nil {
		log.Fatal("failed to drop tables:", err)
	}

	// 自动迁移所有模型
	if err := db.AutoMigrate(
		&modelv2.UserV2{},
		&modelv2.PatientV2{},
		&modelv2.Doctor{},
		&modelv2.Visit{},
		&modelv2.Order{},
		&modelv2.LabTest{},
		&modelv2.MedicalRecordV2{},
		&modelv2.MedicineV2{},
		&modelv2.AdmissionNotice{},
		&modelv2.WesternMedicline{},
		&modelv2.ChineseMedicline{},
		&modelv2.ICD_10{},
	); err != nil {
		log.Fatal("failed to auto migrate:", err)
	}

	fakedata.SeedPatientV2Data()

	fakedata.SeedMedicineRecordV2Data()

	fakedata.SeedMedicineV2Data()
}

// func InitDBWithMongo() {
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/prod?charset=utf8mb4&parseTime=True&loc=Local"
// 	var err error
// 	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v", err)
// 	}
// 	log.Println("Database connection established.")

// 	// 自动迁移表结构
// 	err = db.AutoMigrate(&model.PatientV2{}, &model.MedicalRecordV2{})
// 	if err != nil {
// 		log.Fatalf("Failed to migrate database: %v", err)
// 	}

// 	log.Println("Data seeding completed.")
// }

// db.createUser({
//     user: "root",
//     pwd: "123456",  // 替换为你的密码
//     roles: [{ role: "userAdminAnyDatabase", db: "admin" }]
// });

// 删除Sqlite所有表
func dropAllTables(db *gorm.DB) error {
	// 获取所有表名
	tables, err := db.Migrator().GetTables()
	if err != nil {
		return err
	}

	// 删除所有表
	for _, tableName := range tables {
		// 不要删除这个表
		if tableName == "sqlite_sequence" {
			continue
		}

		if err := db.Migrator().DropTable(tableName); err != nil {
			return err
		}
	}
	return nil
}

// 删除MySql所有表
func dropAllMySqlTables(db *gorm.DB) error {
	// 获取所有表名
	tables, err := db.Migrator().GetTables()
	if err != nil {
		return err
	}

	// 删除所有表
	for _, tableName := range tables {
		if err := db.Migrator().DropTable(tableName); err != nil {
			return err
		}
	}
	return nil
}
