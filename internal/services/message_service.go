package services

import (
	"context"
	"github.com/EraldCaka/prizz-backend/db"
	"github.com/EraldCaka/prizz-backend/internal/types"
)

type MessageService interface {
	GetAll(ctx context.Context) ([]*types.Message, error)
	Create(ctx context.Context, messReq *types.MessageRequest) (string, error)
}

type MessageStore struct {
	db *db.Postgres
}

func NewMessageService(db *db.Postgres) *MessageStore {
	return &MessageStore{db: db}
}

func (s *MessageStore) Create(ctx context.Context, messReq *types.MessageRequest) (string, error) {
	return s.db.CreateMessage(ctx, messReq)
}

func (s *MessageStore) GetAll(ctx context.Context) ([]*types.Message, error) {
	// TODO : IMPLEMENT error handling LOGIC
	return s.db.GetMessages(ctx)
}
