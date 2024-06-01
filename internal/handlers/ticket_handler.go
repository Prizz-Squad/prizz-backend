package handlers

import (
	"fmt"
	"github.com/EraldCaka/prizz-backend/internal/services"
	"github.com/EraldCaka/prizz-backend/internal/types"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type Ticket interface {
	CreateTicket(ctx *fiber.Ctx) error
	GetAllTickets(ctx *fiber.Ctx) error
	UpdateTicketDepartment(ctx *fiber.Ctx) error
	UpdateTicketStatus(ctx *fiber.Ctx) error
	UpdateTicketDescription(ctx *fiber.Ctx) error
	UpdateTicketPostDate(ctx *fiber.Ctx) error
}

type TicketHandler struct {
	ticketService services.TicketService
}

func NewTicketHandler(ticketService services.TicketService) *TicketHandler {
	return &TicketHandler{ticketService: ticketService}
}

func (h *TicketHandler) CreateTicket(ctx *fiber.Ctx) error {
	var project *types.CreateTicketRequest
	if err := ctx.BodyParser(&project); err != nil {
		return types.ErrBadRequest()
	}
	ticketID, err := h.ticketService.CreateTicket(ctx.Context(), project)

	if err != nil {
		fmt.Println(err)
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't create the ticket"))
	}
	return ctx.JSON(ticketID)
}

func (h *TicketHandler) GetAllTickets(ctx *fiber.Ctx) error {
	tickets, err := h.ticketService.GetAllTickets(ctx.Context())
	if err != nil {
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't get the tickets"))
	}
	return ctx.JSON(tickets)
}

func (h *TicketHandler) UpdateTicketDepartment(ctx *fiber.Ctx) error {
	var (
		ticketID      = ctx.Query("ticketID")
		departmentStr = ctx.Query("department")
	)
	department, err := strconv.Atoi(departmentStr)
	if err != nil {
		return types.ErrBadRequest()
	}
	if err := h.ticketService.UpdateTicketDepartment(ctx.Context(), ticketID, &types.UpdateDepartment{Department: department}); err != nil {
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't update ticket"))
	}
	return ctx.JSON(ticketID)
}

func (h *TicketHandler) UpdateTicketStatus(ctx *fiber.Ctx) error {
	var (
		ticketID  = ctx.Query("ticketID")
		statusStr = ctx.Query("status")
	)
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		return types.ErrBadRequest()
	}
	if err := h.ticketService.UpdateTicketStatus(ctx.Context(), ticketID, &types.UpdateStatus{Status: status}); err != nil {
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't update ticket"))
	}
	return ctx.JSON(ticketID)
}

func (h *TicketHandler) UpdateTicketDescription(ctx *fiber.Ctx) error {
	var (
		ticketID    = ctx.Query("ticketID")
		description = ctx.Query("description")
	)
	if err := h.ticketService.UpdateTicketDescription(ctx.Context(), ticketID, &types.UpdateDescription{Description: description}); err != nil {
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't update ticket"))
	}
	return ctx.JSON(ticketID)
}

func (h *TicketHandler) UpdateTicketPostDate(ctx *fiber.Ctx) error {
	var (
		ticketID = ctx.Query("ticketID")
		postDate = ctx.Query("postDate")
	)
	if err := h.ticketService.UpdateTicketPostDate(ctx.Context(), ticketID, &types.UpdatePostDate{PostDate: postDate}); err != nil {
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't update ticket"))
	}
	return ctx.JSON(ticketID)
}
