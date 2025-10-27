package controllers

import (
	"chatserver/modules/auth/dto"
	"chatserver/modules/auth/usecases"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Authcontrollers struct {
	useacses  usecases.AuthUsecases
	validator *validator.Validate
}

func NewAuthcontrollers(useacses usecases.AuthUsecases) *Authcontrollers {
	return &Authcontrollers{
		useacses:  useacses,
		validator: validator.New(),
	}
}

func (cc *Authcontrollers) Register(c *fiber.Ctx) error {
	var req dto.RegisterReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err,
			"success": false,
		})
	}

	if err := cc.validator.Struct(&req); err != nil {
		validationErrs := err.(validator.ValidationErrors)
		return c.Status(400).JSON(fiber.Map{
			"message": validationErrs.Error(),
			"success": false,
		})
	}

	if err := cc.useacses.Register(&req); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
			"success": false,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}

func (cc *Authcontrollers) Login(c *fiber.Ctx) error {
	var req dto.LoginReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err,
			"success": false,
		})
	}

	if err := cc.validator.Struct(&req); err != nil {
		validationErrs := err.(validator.ValidationErrors)
		return c.Status(400).JSON(fiber.Map{
			"message": validationErrs.Error(),
			"success": false,
		})
	}

	accesstoken, err := cc.useacses.Login(&req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
			"success": false,
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:   "token",
		Value:  accesstoken,
		MaxAge: int((700000 * time.Hour).Seconds()),
	})

	return c.JSON(fiber.Map{
		"success": true,
	})
}
