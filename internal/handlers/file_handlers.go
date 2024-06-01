package handlers

import (
	"fmt"
	"net/http"

	"github.com/EraldCaka/prizz-backend/internal/services"
	"github.com/EraldCaka/prizz-backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type File interface {
	CreateFile(ctx *fiber.Ctx) error
	GetFiles(ctx *fiber.Ctx) error
}

type FileHandler struct {
	fileService services.FileService
}

func NewFileHandler(fileService services.FileService) *FileHandler {
	return &FileHandler{fileService: fileService}
}

func (h *FileHandler) CreateFile(ctx *fiber.Ctx) error {
	var file *types.FileCreateRequest
	if err := ctx.BodyParser(&file); err != nil {
		fmt.Println("err", err)
		return types.ErrBadRequest()
	}
	fileID, err := h.fileService.Create(ctx.Context(), file)

	if err != nil {
		fmt.Println(err)
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't create the file"))
	}
	return ctx.JSON(fileID)
}

func (h *FileHandler) GetFiles(ctx *fiber.Ctx) error {
	files, err := h.fileService.GetAll(ctx.Context())
	if err != nil {
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't get the files"))
	}
	return ctx.JSON(files)
}
