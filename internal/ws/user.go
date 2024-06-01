package ws

import (
	"github.com/EraldCaka/prizz-backend/internal/types"
	"github.com/gofiber/websocket/v2"
	"log"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *types.Message
	ID       string `json:"id"`
	TaskID   string `json:"taskID"`
	Username string `json:"username"`
}

func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) readMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg := &types.Message{
			Contents: string(m),
			TaskId:   c.TaskID,
			UserId:   c.ID,
		}

		hub.Broadcast <- msg
	}
}
