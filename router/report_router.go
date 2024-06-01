package router

import (
	"context"
	"fmt"
	"github.com/EraldCaka/prizz-backend/db"
	"github.com/EraldCaka/prizz-backend/internal/handlers"
	"github.com/EraldCaka/prizz-backend/internal/routes"
	"github.com/EraldCaka/prizz-backend/internal/services"
	"github.com/EraldCaka/prizz-backend/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var reportApi *fiber.App

func NewReportRouter() {
	reportApi = fiber.New(Config)
	database, err := db.NewPGInstance(context.Background())
	if err != nil {
		fmt.Println(types.NewError(500, fmt.Sprintf("Could not initialize database connection: %s", err)))
		return
	}
	reportApi.Use(cors.New(cors.Config{
		AllowOrigins:     "https://skeleton-fiber.onrender.com", // fe url
		AllowMethods:     "GET, POST, DELETE, PUT",
		AllowHeaders:     "Content-Type",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))
	var route = reportApi.Group("/prizz/api/v2")
	reportHistoryService := services.NewReportService(database)
	reportHistoryHandler := handlers.NewReportHistory(reportHistoryService)
	routes.ReportHistoryRoute(reportHistoryHandler, route)

}
func StartReportServer(address string) error {
	return reportApi.Listen(address)
}
