package routes

import (
	"github.com/LightBulbClub/pole-arc/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	api := app.Group("/api")
	user := api.Group("/user")

	user.Post("/register", Register)
	api.Post("/student/login", StudentLogin)
	api.Post("/teacher/login", TeacherLogin)

	// 受保护的路由
	api.Get("/protected", middlewares.AuthRequired(), ProtectedRoute)

	// 社团相关路由
	association := api.Group("/association")
	association.Post("/logs", middlewares.AuthRequired(), CreateAssociationLog)
}
