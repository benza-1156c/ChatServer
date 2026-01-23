package middleware

import (
	"chatserver/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AuthRequired(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "ไม่พบ token กรุณาเข้าสู่ระบบ",
			"success": false,
		})
	}

	claims, err := utils.ParsedToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "token ไม่ถูกต้องหรือหมดอายุ",
			"success": false,
		})
	}

	userIdStr, ok := claims["userId"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "token ไม่ถูกต้อง",
			"success": false,
		})
	}

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "userId ไม่ถูกต้อง",
			"success": false,
		})
	}

	c.Locals("userId", userId)

	return c.Next()
}
