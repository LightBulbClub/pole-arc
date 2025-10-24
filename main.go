package main

import (
	"log"

	"github.com/LightBulbClub/pole-arc/config"
	"github.com/LightBulbClub/pole-arc/models"
	"github.com/LightBulbClub/pole-arc/routes""

	"github.com/gofiber/fiber/v2"
)

func main() {
	// 加载 TOML 配置文件
	err := config.LoadConfig("config.toml")
	if err != nil {
		log.Fatalf("无法加载配置文件: %v", err)
	}

	// 连接数据库
	config.ConnectDB()

	// 自动迁移数据库表
	err = config.DB.AutoMigrate(&models.Student{}, &models.Teacher{}, &models.AssociationLog{})
	if err != nil {
		return
	}

	app := fiber.New()

	// 注册认证路由
	routes.Routes(app)

	// 从配置获取服务器地址
	serverAddress := config.GetServerAddress()

	log.Printf("服务器在地址 %s 上运行...", serverAddress)
	log.Fatal(app.Listen(serverAddress))
}
