package hub

type Message struct {
	Room string
	Data []byte
}

type Hub struct {
	Client     map[string]map[*Client]bool
	Broadcast  chan *Message
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Client:     make(map[string]map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			if _, ok := h.Client[client.Room]; !ok {
				h.Client[client.Room] = make(map[*Client]bool)
			}
			h.Client[client.Room][client] = true

		case client := <-h.Unregister:
			if clients, ok := h.Client[client.Room]; ok {
				if _, exists := clients[client]; exists {
					delete(clients, client)
					close(client.Send)
					if len(clients) == 0 {
						delete(h.Client, client.Room)
					}
				}
			}

		case msg := <-h.Broadcast:
			if clients, ok := h.Client[msg.Room]; ok {
				for client := range clients {
					select {
					case client.Send <- msg.Data:

					default:
						delete(clients, client)
						close(client.Send)
					}
				}
			}

		}
	}
}
