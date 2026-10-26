package database

import (
	"homework4/config"
	"homework4/models"
	"log"
	
	"github.com/glebarez/sqlite"  // 替换为纯 Go 驱动
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open(config.DBPath),&gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	
	if err := DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}); err != nil{
		log.Fatalf("自动迁移失败: %v", err)
	}
	log.Println("数据库初始化完成")
	
}