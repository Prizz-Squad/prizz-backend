package handlers

import (
	"github.com/EraldCaka/prizz-backend/internal/services"
	"github.com/EraldCaka/prizz-backend/internal/types"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type MessageHandler struct {
	messageService services.MessageService
}

type Messages interface {
	GetMessages(ctx *fiber.Ctx) error
}

func NewMessageHandler(messageService services.MessageService) *MessageHandler {
	return &MessageHandler{messageService: messageService}
}

func (h *MessageHandler) GetMessages(ctx *fiber.Ctx) error {
	messages, err := h.messageService.GetAll(ctx.Context())
	if err != nil {
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't get the users"))
	}
	return ctx.JSON(messages)
}
