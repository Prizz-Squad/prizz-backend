package services

import (
	"context"

	"github.com/EraldCaka/prizz-backend/db"
	"github.com/EraldCaka/prizz-backend/internal/types"
)

type ProjectService interface {
	Create(ctx context.Context, projectReq *types.ProjectCreateRequest) (string, error)
	GetAll(ctx context.Context) ([]*types.Project, error)
}
type ProjectStore struct {
	db *db.Postgres
}

func NewProjectService(db *db.Postgres) *ProjectStore {
	return &ProjectStore{db: db}
}

func (s *ProjectStore) Create(ctx context.Context, projectReq *types.ProjectCreateRequest) (string, error) {
	return s.db.CreateProject(ctx, projectReq)
}

func (s *ProjectStore) GetAll(ctx context.Context) ([]*types.Project, error) {
	// TODO : IMPLEMENT error handling LOGIC
	return s.db.GetProjects(ctx)
}
