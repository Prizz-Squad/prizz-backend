package routes

import (
	"github.com/EraldCaka/prizz-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func TaskHistoryRoutes(taskHistoryHandler *handlers.TaskHistoryHandler, route fiber.Router) {

	route.Get("/taskHistory", taskHistoryHandler.GetAll)
	route.Post("/taskHistory/create", taskHistoryHandler.Create)
	route.Get("/taskHistoryByDateAndUser", taskHistoryHandler.GetTasksHistory)
	route.Get("/taskHistory/totalHours", taskHistoryHandler.GetTotalCountByUserId)
}
