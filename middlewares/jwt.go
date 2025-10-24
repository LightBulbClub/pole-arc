package middlewares

import (
	"github.com/LightBulbClub/pole-arc/utils"

	"github.com/gofiber/fiber/v2"
)

// AuthRequired 是一个 Fiber 中间件，用于验证 JWT
func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "未提供认证令牌"})
		}

		tokenString := ""
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "无效的认证令牌格式"})
		}

		claims, err := utils.ParseJWT(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "认证令牌无效或已过期"})
		}

		// 将用户 ID 存储在 Fiber 上下文中，以便后续处理程序使用
		c.Locals("userID", claims.UserID)
		return c.Next()
	}
}
