package model

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	//连接数据库
	db, err := gorm.Open(sqlite.Open("../../config/LavaPanel.db"), &gorm.Config{})
	if err != nil {
		panic("[LavaPanel][InitDB]|[Error]:未寻找到数据库LavaPanel.db")
	}
	//数据库自动迁移
	err = db.AutoMigrate(&SystemConfig{})
	if err != nil {
		panic("[LavaPanel][AutoMigrate]|[Error]:迁移数据库异常")
	}
	DB = db
}
