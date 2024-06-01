package handlers

import (
	"fmt"
	"github.com/EraldCaka/prizz-backend/internal/services"
	"github.com/EraldCaka/prizz-backend/internal/types"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type ReportHistory interface {
	CreateReport(ctx *fiber.Ctx) error
	GetAllReport(ctx *fiber.Ctx) error
	GetReportHistory(ctx *fiber.Ctx) error
	GetTotalCountByUserId(ctx *fiber.Ctx) error
}

type ReportHistoryHandler struct {
	reportService services.ReportService
}

func NewReportHistory(reportService services.ReportService) *ReportHistoryHandler {
	return &ReportHistoryHandler{reportService: reportService}
}

func (th *ReportHistoryHandler) CreateReport(ctx *fiber.Ctx) error {
	var report *types.ReportRequest
	if err := ctx.BodyParser(&report); err != nil {
		return types.ErrBadRequest()
	}
	createrReport, err := th.reportService.CreateReport(ctx.Context(), report)

	if err != nil {
		fmt.Println(err)
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't create the report"))
	}
	return ctx.JSON(createrReport)

}

func (th *ReportHistoryHandler) GetAllReport(ctx *fiber.Ctx) error {
	allReports, err := th.reportService.GetAllReports(ctx.Context())
	if err != nil {
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't get the report history"))
	}
	return ctx.JSON(allReports)
}

func (th *ReportHistoryHandler) GetTotalCountByUserId(ctx *fiber.Ctx) error {
	var userID = ctx.Query("id")
	taskReport, err := th.reportService.GetAllHoursReportsByUserId(ctx.Context(), userID)
	if err != nil {
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't get report"))
	}
	return ctx.JSON(taskReport)
}

func (th *ReportHistoryHandler) GetReportHistory(ctx *fiber.Ctx) error {
	startDateStr := ctx.Query("startDate")
	endDateStr := ctx.Query("endDate")
	userId := ctx.Query("userId")
	tasksHistory, err := th.reportService.GetReportHistory(ctx.Context(), startDateStr, endDateStr, userId)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't retrieve report history"))
	}
	return ctx.JSON(tasksHistory)
}
