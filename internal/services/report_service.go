package services

import (
	"context"
	"github.com/EraldCaka/prizz-backend/db"
	"github.com/EraldCaka/prizz-backend/internal/types"
)

type ReportService interface {
	CreateReport(ctx context.Context, tr *types.ReportRequest) (string, error)
	GetAllReports(ctx context.Context) ([]*types.Report, error)
	GetReportHistory(ctx context.Context, startTime, endTime string, userID string) ([]*types.Report, error)
	GetAllHoursReportsByUserId(ctx context.Context, userId string) (*types.TotalCountResponse, error)
}

type ReportStore struct {
	db *db.Postgres
}

func NewReportService(db *db.Postgres) *ReportStore {
	return &ReportStore{db: db}
}

func (r ReportStore) GetAllHoursReportsByUserId(ctx context.Context, userId string) (*types.TotalCountResponse, error) {
	return r.db.GetAllHoursReportsByUserId(ctx, userId)
}

func (r *ReportStore) GetAllReports(ctx context.Context) ([]*types.Report, error) {
	return r.db.GetAllReports(ctx)
}

func (r *ReportStore) CreateReport(ctx context.Context, tr *types.ReportRequest) (string, error) {
	return r.db.CreateReport(ctx, tr)
}

func (r *ReportStore) GetReportHistory(ctx context.Context, startTime, endTime string, userID string) ([]*types.Report, error) {
	return r.db.GetReportHistory(ctx, startTime, endTime, userID)
}
