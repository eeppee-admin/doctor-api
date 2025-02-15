package model

import (
	"gorm.io/gorm"
)

// 简单版本
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

// User 存储系统用户（医生、管理员等）的账号信息
// todo: 添加电话登录功能
