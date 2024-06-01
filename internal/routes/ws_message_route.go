package routes

import (
	"github.com/EraldCaka/prizz-backend/internal/ws"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func WsMessageRoute(wsHandler *ws.Handler, route fiber.Router) {
	route.Get("/ws/comment:/taskID", websocket.New(wsHandler.ConnectAndComment))

}
