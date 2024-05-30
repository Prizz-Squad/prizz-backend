package handlers

import "github.com/gofiber/fiber/v2"

type User interface {
	CreateUser(ctx *fiber.Ctx) error
	GetUsers(ctx *fiber.Ctx) error
	GetUserByID(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
	LoginUser(ctx *fiber.Ctx) error
	LogoutUser(ctx *fiber.Ctx) error
	RoleBaseMiddleware() fiber.Handler
}
