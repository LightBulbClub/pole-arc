package controllers

import (
	"github.com/LightBulbClub/rolling-wheel/config"
	"github.com/LightBulbClub/rolling-wheel/models"
	"github.com/LightBulbClub/rolling-wheel/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Register 处理用户注册
func Register(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "请求体解析失败", "error": err.Error()})
	}

	// 密码哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "密码哈希失败", "error": err.Error()})
	}
	user.Password = string(hashedPassword)

	// 保存用户到数据库
	if result := config.DB.Create(&user); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "用户注册失败", "error": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "用户注册成功", "user": user.Username})
}

// Login 处理用户登录
func Login(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "请求体解析失败", "error": err.Error()})
	}

	// 查找数据库中的用户
	var foundUser models.User
	if result := config.DB.Where("username = ?", user.Username).First(&foundUser); result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "用户名或密码不正确"})
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "用户名或密码不正确"})
	}

	// 生成 JWT
	token, err := utils.GenerateJWT(foundUser.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "生成认证令牌失败", "error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "登录成功", "token": token})
}

// ProtectedRoute 是一个受保护的路由示例
func ProtectedRoute(c *fiber.Ctx) error {
	// 从 Locals 中获取用户 ID
	userID := c.Locals("userID").(uint)
	return c.JSON(fiber.Map{"message": "欢迎访问受保护的资源！", "user_id": userID})
}
