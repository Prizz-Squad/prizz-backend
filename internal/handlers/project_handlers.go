package handlers

import (
	"fmt"
	"net/http"

	"github.com/EraldCaka/prizz-backend/internal/services"
	"github.com/EraldCaka/prizz-backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type Project interface {
	CreateProject(ctx *fiber.Ctx) error
	GetProjects(ctx *fiber.Ctx) error
}

type ProjectHandler struct {
	projectService services.ProjectService
}

func NewProjectHandler(projectService services.ProjectService) *ProjectHandler {
	return &ProjectHandler{projectService: projectService}
}

func (h *ProjectHandler) CreateProject(ctx *fiber.Ctx) error {
	var project *types.ProjectCreateRequest
	if err := ctx.BodyParser(&project); err != nil {
		return types.ErrBadRequest()
	}
	projectID, err := h.projectService.Create(ctx.Context(), project)

	if err != nil {
		fmt.Println(err)
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't create the project"))
	}
	return ctx.JSON(projectID)
}

func (h *ProjectHandler) GetProjects(ctx *fiber.Ctx) error {
	projects, err := h.projectService.GetAll(ctx.Context())
	if err != nil {
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't get the projects"))
	}
	return ctx.JSON(projects)
}
