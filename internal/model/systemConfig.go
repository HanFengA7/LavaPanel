package model

import "gorm.io/gorm"

type SystemConfig struct {
	gorm.Model
	Key   string `gorm:"column:key; type:varchar(255);" json:"key"`
	Value string `gorm:"column:value; type:varchar(255);" json:"value"`
}

func (table SystemConfig) TableName() string {
	return "SystemConfig"
}
