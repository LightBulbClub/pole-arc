package routes

import (
	"github.com/LightBulbClub/rolling-wheel/controllers"
	"github.com/LightBulbClub/rolling-wheel/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)

	// 受保护的路由
	api.Get("/protected", middlewares.AuthRequired(), controllers.ProtectedRoute)
}
