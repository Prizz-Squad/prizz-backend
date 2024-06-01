package routes

import (
	"github.com/EraldCaka/prizz-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func ReportHistoryRoute(reportHistoryHandler *handlers.ReportHistoryHandler, route fiber.Router) {

	route.Get("/reportHistory", reportHistoryHandler.GetAllReport)
	route.Post("/createReport", reportHistoryHandler.CreateReport)
	route.Get("/reportHistory/user", reportHistoryHandler.GetTotalCountByUserId)
	route.Get("/reportHistory/totalHours", reportHistoryHandler.GetReportHistory)
}
