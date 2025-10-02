package controllers

import (
	"chatserver/modules/websocket/hub"

	"github.com/gofiber/websocket/v2"
)

func ChatWsHandlerFiber(h *hub.Hub, c *websocket.Conn) {
	client := hub.NewClient(h, c, "1")
	h.Register <- client

	go client.ReadPump()
	go client.WritePump()
}
