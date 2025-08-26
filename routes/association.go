package routes

import (
	"github.com/LightBulbClub/rolling-wheel/config"
	. "github.com/LightBulbClub/rolling-wheel/models"
	"github.com/gofiber/fiber/v2"
)

// CreateAssociationLog 创建社团活动日志
func CreateAssociationLog(c *fiber.Ctx) error {
	log := new(AssociationLog)
	if err := c.BodyParser(log); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "请求体解析失败", "error": err.Error()})
	}

	// 验证必填字段
	if log.AssociationName == "" || log.ActivityTitle == "" || log.ActivityType == "" || log.Organizer == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "社团名称、活动标题、活动类型和组织者为必填字段"})
	}

	// 保存到数据库
	if result := config.DB.Create(&log); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "创建社团活动日志失败", "error": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "社团活动日志创建成功", "log_id": log.ID})
}
