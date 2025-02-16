package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // 定义全局变量

func InitDatabase() {
	// 用户名:root,密码:123456,地址:(本地),数据库prod,
	// 数据库prod要提前创建
	// sql语句: `create database prod;`
	dsn := "root:123456@tcp(127.0.0.1:3306)/prod?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接mysql数据库失败，检查连接字符串: %v", err)
	}
	log.Println("连接mysql成功")
}
