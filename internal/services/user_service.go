package services

import (
	"context"
	"errors"
	"github.com/EraldCaka/prizz-backend/db"
	"github.com/EraldCaka/prizz-backend/internal/types"
	"github.com/EraldCaka/prizz-backend/util"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserStore struct {
	db *db.Postgres
}

func NewUserService(db *db.Postgres) *UserStore {
	return &UserStore{db: db}
}

func (s *UserStore) Create(ctx context.Context, userReq *types.UserCreateRequest) (string, error) {
	userReq.Password, _ = util.HashPassword(userReq.Password)
	return s.db.CreateUser(ctx, userReq)
}

func (s *UserStore) GetAll(ctx context.Context) ([]*types.User, error) {
	// TODO : IMPLEMENT error handling LOGIC
	return s.db.GetUsers(ctx)
}

func (s *UserStore) GetByID(ctx context.Context, userID string) (*types.User, error) {
	// TODO : IMPLEMENT error handling LOGIC
	return s.db.GetUserByID(ctx, userID)
}

func (s *UserStore) Update(ctx context.Context, userID string, userReq *types.UserRequest) error {
	// TODO : IMPLEMENT error handling LOGIC
	if err := s.db.UpdateUser(ctx, userID, userReq); err != nil {
		return err
	}
	return nil
}

func (s *UserStore) Delete(ctx context.Context, userID string) error {
	// TODO : IMPLEMENT error handling LOGIC
	return s.db.DeleteUser(ctx, userID)
}

func (s *UserStore) Login(ctx context.Context, userReq *types.UserRequest) (string, error) {
	return s.db.Login(ctx, userReq)
}

func (s *UserStore) GetUserByToken(ctx context.Context, tokenString string) (*types.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(util.Secret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	userID, ok := claims["id"].(string)
	if !ok {
		return nil, errors.New("invalid token claims: missing user ID")
	}

	expFloat, ok := claims["exp"].(float64)
	if !ok {
		return nil, errors.New("invalid token claims: missing expiration time")
	}
	exp := time.Unix(int64(expFloat), 0)

	if time.Now().After(exp) {
		return nil, errors.New("token has expired")
	}

	user, err := s.db.GetUserByID(ctx, userID)
	if err != nil {
		return nil, errors.New("couldn't get user")
	}

	return user, nil
}
