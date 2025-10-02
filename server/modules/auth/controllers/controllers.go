package controllers

import (
	"chatserver/modules/auth/dto"
	"chatserver/modules/auth/usecases"

	"github.com/gofiber/fiber/v2"
)

type Authcontrollers struct {
	useacses usecases.AuthUsecases
}

func NewAuthcontrollers(useacses usecases.AuthUsecases) *Authcontrollers {
	return &Authcontrollers{useacses}
}

func (cc *Authcontrollers) Register(c *fiber.Ctx) error {
	req := &dto.RegisterReq{
		UserName: c.FormValue("username"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	file, err := c.FormFile("image")
	if err != nil {
		req.Avatar = nil
	}

	req.Avatar = file

	if err := cc.useacses.Register(req); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
			"success": false,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}
