package routes

import (
	"github.com/EraldCaka/prizz-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func TicketRoutes(ticketHandler *handlers.TicketHandler, route fiber.Router) {
	route.Get("/tickets", ticketHandler.GetAllTickets)
	route.Put("/ticketDescription", ticketHandler.UpdateTicketDescription)
	route.Put("/ticketStatus", ticketHandler.UpdateTicketStatus)
	route.Put("/ticketDepartment", ticketHandler.UpdateTicketDepartment)
	route.Put("/ticketPostDate", ticketHandler.UpdateTicketPostDate)
	route.Post("/newTicket", ticketHandler.CreateTicket)
}
