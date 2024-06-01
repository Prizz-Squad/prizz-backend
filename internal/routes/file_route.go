package routes

import (
	"github.com/EraldCaka/prizz-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func FileRoutes(fileHandler *handlers.FileHandler, route fiber.Router) {
	route.Get("/getFiles", fileHandler.GetFiles)
	route.Post("/file", fileHandler.CreateFile)
}
