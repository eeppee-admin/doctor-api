package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

// User 存储系统用户（医生、管理员等）的账号信息
type UserV2 struct {
	gorm.Model
	UserID    uint   `gorm:"primary_key;auto_increment;not null" json:"user_id"`
	Username  string `gorm:"type:varchar(50);not null;unique" json:"username"`
	Password  string `gorm:"type:varchar(100);not null" json:"password"` // 密码应该加密存储
	Phone     string `gorm:"type:varchar(15);unique" json:"phone"`       // 电话号码，可以设置为唯一
	Role      string `gorm:"type:varchar(20);not null" json:"role"`
	IsDeleted int    `gorm:"not null" json:"is_deleted"`
}
