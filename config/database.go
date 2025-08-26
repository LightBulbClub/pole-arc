package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	// 从TOML配置加载数据库连接字符串
	if AppConfig == nil {
		log.Fatal("配置未加载，请先调用 LoadConfig()")
	}
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Singapore",
		AppConfig.Database.Host,
		AppConfig.Database.User,
		AppConfig.Database.Password,
		AppConfig.Database.Name,
		AppConfig.Database.Port,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 打印 SQL 日志
	})

	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}

	log.Println("成功连接到 PostgreSQL 数据库!")
}
