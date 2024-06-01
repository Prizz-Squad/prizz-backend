package router

import (
	"context"
	"fmt"
	"github.com/EraldCaka/prizz-backend/db"
	"github.com/EraldCaka/prizz-backend/internal/handlers"
	"github.com/EraldCaka/prizz-backend/internal/middleware"
	"github.com/EraldCaka/prizz-backend/internal/routes"
	"github.com/EraldCaka/prizz-backend/internal/services"
	"github.com/EraldCaka/prizz-backend/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var config = fiber.Config{
	ErrorHandler: types.ErrorHandler,
}

var api *fiber.App

func NewRouter() {
	api = fiber.New(config)
	database, err := db.NewPGInstance(context.Background())
	if err != nil {
		fmt.Println(types.NewError(500, fmt.Sprintf("Could not initialize database connection: %s", err)))
		return
	}

	api.Use(cors.New(cors.Config{
		AllowOrigins:     "https://skeleton-fiber.onrender.com", // fe url
		AllowMethods:     "GET, POST, DELETE, PUT",
		AllowHeaders:     "Content-Type",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
		//AllowOriginsFunc: func(origin string) bool {
		//	return true
		//},
	}))

	var (
		userService        = services.NewUserService(database)
		projectService     = services.NewProjectService(database)
		messageService     = services.NewMessageService(database)
		taskHistoryService = services.NewTaskService(database)
	)
	var (
		userHandler        = handlers.NewUserHandler(userService)
		projectHandler     = handlers.NewProjectHandler(projectService)
		messageHandler     = handlers.NewMessageHandler(messageService)
		taskHistoryHandler = handlers.NewTaskHistoryHandler(taskHistoryService)
	)

	var route = api.Group("/prizz/api/v1")
	route.Use(middleware.AuthMiddleware(userHandler.RoleBaseMiddleware()))
	routes.UserRoutes(userHandler, route)
	routes.MessageRoutes(messageHandler, route)
	routes.ProjectRoutes(projectHandler, route)
	routes.TaskHistoryRoutes(taskHistoryHandler, route)

}
func Start(address string) error {
	return api.Listen(address)
}
