package hub

import "github.com/gofiber/websocket/v2"

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
	Room string
}

func NewClient(h *Hub, conn *websocket.Conn, room string) *Client {
	return &Client{
		Hub:  h,
		Conn: conn,
		Send: make(chan []byte, 256),
		Room: room,
	}
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		c.Hub.Broadcast <- &Message{Room: c.Room, Data: msg}
	}
}

func (c *Client) WritePump() {
	defer c.Conn.Close()
	for msg := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}
