package routes

import (
	"github.com/EraldCaka/prizz-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(userHandler *handlers.UserHandler, route fiber.Router) {

	route.Get("/users/:id", userHandler.GetUserByID)
	route.Put("/users/:id", userHandler.UpdateUser)
	route.Delete("/users/:id", userHandler.DeleteUser)
	route.Get("/users", userHandler.GetUsers)
	route.Post("/register", userHandler.CreateUser)
	route.Post("/login", userHandler.LoginUser)
	route.Post("/logout", userHandler.LogoutUser)
}
