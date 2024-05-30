package services

import (
	"context"
	"github.com/EraldCaka/prizz-backend/internal/types"
)

type UserService interface {
	Create(ctx context.Context, userReq *types.UserCreateRequest) (string, error)
	GetAll(ctx context.Context) ([]*types.User, error)
	GetByID(ctx context.Context, userID string) (*types.User, error)
	Update(ctx context.Context, userID string, userReq *types.UserRequest) error
	Delete(ctx context.Context, userID string) error
	Login(ctx context.Context, userReq *types.UserRequest) (string, error)
	GetUserByToken(ctx context.Context, tokenString string) (*types.User, error)
}
