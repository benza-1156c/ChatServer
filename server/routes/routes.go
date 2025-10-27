package routes

import (
	"chatserver/modules/auth"
	"chatserver/modules/websocket/controllers"
	"chatserver/modules/websocket/hub"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"gorm.io/gorm"
)

func SetUpRoutes(app *fiber.App, h *hub.Hub, db *gorm.DB) {
	auth.NewRouteAuthRepo(app, db)

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		controllers.ChatWsHandlerFiber(h, c)
	}))
}
