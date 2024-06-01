package services

import (
	"context"

	"github.com/EraldCaka/prizz-backend/db"
	"github.com/EraldCaka/prizz-backend/internal/types"
)

type FileService interface {
	Create(ctx context.Context, fileReq *types.FileCreateRequest) (string, error)
	GetAll(ctx context.Context) ([]*types.File, error)
}
type FileStore struct {
	db *db.Postgres
}

func NewFileService(db *db.Postgres) *FileStore {
	return &FileStore{db: db}
}

func (s *FileStore) Create(ctx context.Context, fileReq *types.FileCreateRequest) (string, error) {

	return s.db.CreateFile(ctx, fileReq)
}

func (s *FileStore) GetAll(ctx context.Context) ([]*types.File, error) {
	// TODO : IMPLEMENT error handling LOGIC
	return s.db.GetFiles(ctx)
}
