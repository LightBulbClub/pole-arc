package main

import (
	"github.com/LightBulbClub/rolling-wheel/config"
	"github.com/LightBulbClub/rolling-wheel/models"
	"github.com/LightBulbClub/rolling-wheel/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("无法加载 .env 文件: %v", err)
	}

	// 连接数据库
	config.ConnectDB()

	// 自动迁移数据库表
	err = config.DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	app := fiber.New()

	// 注册认证路由
	routes.AuthRoutes(app)

	url := os.Getenv("URL")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // 默认端口
	}

	log.Printf("服务器在端口 %s 上运行...", port)
	log.Fatal(app.Listen(url + ":" + port))
}
