package auth

import (
	"chatserver/modules/auth/controllers"
	"chatserver/modules/auth/repositories"
	"chatserver/modules/auth/usecases"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRouteAuthRepo(app *fiber.App, db *gorm.DB) {
	repo := repositories.NewAuthRepo(db)
	usecase := usecases.NewAuthUsecases(repo)
	controller := controllers.NewAuthcontrollers(usecase)

	authAPI := app.Group("/api")

	authAPI.Post("/register", controller.Register)
	authAPI.Post("/login", controller.Login)
}
