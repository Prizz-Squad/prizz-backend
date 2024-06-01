package ws

import (
	"fmt"
	"github.com/EraldCaka/prizz-backend/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"net/http"
)

type Handler struct {
	hub *Hub
}
type Room struct {
	ID      string             `json:"id"`
	UserID  string             `json:"userID"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Broadcast  chan *types.Message
	Register   chan *Client
	Unregister chan *Client
	Rooms      map[string]*Room
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan *types.Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Rooms:      make(map[string]*Room),
	}
}
func NewHandler(hub *Hub) *Handler {
	return &Handler{hub: hub}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			if _, ok := h.Rooms[cl.TaskID]; ok {
				r := h.Rooms[cl.TaskID]

				if _, ok := r.Clients[cl.ID]; !ok {
					r.Clients[cl.ID] = cl
				}
			}
		case cl := <-h.Unregister:
			if _, ok := h.Rooms[cl.TaskID]; ok {
				if _, ok := h.Rooms[cl.TaskID].Clients[cl.ID]; ok {
					delete(h.Rooms[cl.TaskID].Clients, cl.ID)
					close(cl.Message)
				}
			}

		case m := <-h.Broadcast:
			if _, ok := h.Rooms[m.TaskId]; ok {
				for _, cl := range h.Rooms[m.TaskId].Clients {
					cl.Message <- m
				}
			}
		}
	}
}

func (h *Handler) ConnectAndComment(c *websocket.Conn) {
	fmt.Println("inside")
	taskID := c.Params("taskID")
	userID := c.Query("userId")
	username := c.Query("username")
	cl := &Client{
		Conn:     c,
		Message:  make(chan *types.Message, 10),
		ID:       userID,
		TaskID:   taskID,
		Username: username,
	}

	go cl.writeMessage()
	go cl.readMessage(h.hub)

}

func (h *Handler) CloseWSConnection(c *fiber.Ctx) {

	if task := h.hub.Rooms[c.Params("taskID")]; task == nil || task.ID == "" {
		c.JSON(http.StatusBadRequest, "bad request(roomID)")
		return
	}
	client := h.hub.Rooms[c.Params("taskID")].Clients[c.Params("userID")]
	if cl := client; cl == nil || cl.ID == "" {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	client.Conn.Close()
	c.JSON(http.StatusOK, "User disconnected successfully")
}
