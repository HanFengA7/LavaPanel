package model

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	User  string `gorm:"column:key; type:varchar(255);" json:"key"`
	Value string `gorm:"column:value; type:varchar(255);" json:"value"`
}
