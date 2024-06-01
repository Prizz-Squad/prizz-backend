package services

import (
	"context"
	"github.com/EraldCaka/prizz-backend/db"
	"github.com/EraldCaka/prizz-backend/internal/types"
)

type TaskHistoryService interface {
	Create(ctx context.Context, tr *types.TaskRequest) (string, error)
	GetAll(ctx context.Context) ([]*types.TaskHistory, error)
	GetTaskHistory(ctx context.Context, startTime, endTime string, userID string) ([]*types.TaskHistory, error)
	GetAllHoursByUserId(ctx context.Context, userId string) (*types.TotalCountResponse, error)
}

type TaskHistoryStore struct {
	db *db.Postgres
}

func NewTaskService(db *db.Postgres) *TaskHistoryStore {
	return &TaskHistoryStore{db: db}
}

func (ts TaskHistoryStore) GetAllHoursByUserId(ctx context.Context, userId string) (*types.TotalCountResponse, error) {
	return ts.db.GetAllHoursByUserId(ctx, userId)
}

func (ts *TaskHistoryStore) GetAll(ctx context.Context) ([]*types.TaskHistory, error) {
	return ts.db.GetAll(ctx)
}

func (ts *TaskHistoryStore) Create(ctx context.Context, tr *types.TaskRequest) (string, error) {
	return ts.db.Create(ctx, tr)
}

func (ts *TaskHistoryStore) GetTaskHistory(ctx context.Context, startTime, endTime string, userID string) ([]*types.TaskHistory, error) {
	return ts.db.GetTaskHistory(ctx, startTime, endTime, userID)
}
