package routes

import (
	"github.com/EraldCaka/prizz-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func MessageRoutes(messageHandler *handlers.MessageHandler, route fiber.Router) {
	route.Get("/messages", messageHandler.GetMessages)
}
