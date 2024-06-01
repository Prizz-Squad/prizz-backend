package handlers

import (
	"fmt"
	"github.com/EraldCaka/prizz-backend/internal/services"
	"github.com/EraldCaka/prizz-backend/internal/types"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type TaskHistory interface {
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetTasksHistory(ctx *fiber.Ctx) error
	GetTotalCountByUserId(ctx *fiber.Ctx) error
}

type TaskHistoryHandler struct {
	taskHistoryService services.TaskHistoryService
}

func NewTaskHistoryHandler(taskHistoryService services.TaskHistoryService) *TaskHistoryHandler {
	return &TaskHistoryHandler{taskHistoryService: taskHistoryService}
}

func (th *TaskHistoryHandler) Create(ctx *fiber.Ctx) error {
	var testHistory *types.TaskRequest
	if err := ctx.BodyParser(&testHistory); err != nil {
		fmt.Println(err)
		fmt.Println(ctx.Body())
		return types.ErrBadRequest()
	}
	testHistoryId, err := th.taskHistoryService.Create(ctx.Context(), testHistory)

	if err != nil {
		fmt.Println(err)
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't create the task history "))
	}
	return ctx.JSON(testHistoryId)

}

func (th *TaskHistoryHandler) GetAll(ctx *fiber.Ctx) error {
	tasksHistory, err := th.taskHistoryService.GetAll(ctx.Context())
	if err != nil {
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't get the task history"))
	}
	return ctx.JSON(tasksHistory)
}

func (th *TaskHistoryHandler) GetTotalCountByUserId(ctx *fiber.Ctx) error {
	var userID = ctx.Query("id")
	taskReport, err := th.taskHistoryService.GetAllHoursByUserId(ctx.Context(), userID)
	if err != nil {
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't get report"))
	}
	return ctx.JSON(taskReport)
}

func (th *TaskHistoryHandler) GetTasksHistory(ctx *fiber.Ctx) error {
	startDateStr := ctx.Query("startDate")
	endDateStr := ctx.Query("endDate")
	userId := ctx.Query("userId")
	tasksHistory, err := th.taskHistoryService.GetTaskHistory(ctx.Context(), startDateStr, endDateStr, userId)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't retrieve task history"))
	}
	return ctx.JSON(tasksHistory)
}
