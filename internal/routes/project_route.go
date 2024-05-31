package routes

import (
	"github.com/EraldCaka/prizz-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func ProjectRoutes(projectHandler *handlers.ProjectHandler, route fiber.Router) {
	route.Get("/getProjects", projectHandler.GetProjects)
	route.Post("/project", projectHandler.CreateProject)
}
