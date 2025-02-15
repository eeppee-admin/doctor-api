package modelv2

import "gorm.io/gorm"

type MedicineV2 struct {
	gorm.Model
	Name        string  `gorm:"not null;uniqueIndex;size:255"` // 指定键长为 255
	Description string  `gorm:"type:text"`                     // 药品描述
	Price       float64 `gorm:"not null"`                      // 药品价格
	Stock       int     `gorm:"not null"`                      // 库存数量
}
