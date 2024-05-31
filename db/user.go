package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/EraldCaka/prizz-backend/internal/types"
	"github.com/EraldCaka/prizz-backend/util"
	"log"
)

func (pg *Postgres) CreateUser(ctx context.Context, u *types.UserCreateRequest) (string, error) {
	query := "INSERT INTO public.users (username, password, department, role) VALUES ($1, $2, $3, $4) RETURNING id"
	var userID string
	err := pg.db.QueryRow(ctx, query, u.Username, u.Password, u.Role, u.Department).Scan(&userID)
	if err != nil {
		log.Printf("Unable to insert user: %v\n", err)
		return "", err
	}
	return userID, nil
}

func (pg *Postgres) GetUserByID(ctx context.Context, userID string) (*types.User, error) {
	query := "SELECT * FROM public.users WHERE id = $1"
	row := pg.db.QueryRow(ctx, query, userID)

	var user types.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Department)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with ID %s not found", userID)
		}
		log.Printf("Error scanning user data: %v\n", err)
		return nil, err
	}
	return &user, nil
}

func (pg *Postgres) GetUserByName(ctx context.Context, username string) (*types.User, error) {
	query := "SELECT * FROM public.users WHERE username = $1"
	row := pg.db.QueryRow(ctx, query, username)

	var user types.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Department)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with ID %s not found", username)
		}
		log.Printf("Error scanning user data: %v\n", err)
		return nil, err
	}
	return &user, nil
}

func (pg *Postgres) GetUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User

	query := "SELECT * FROM public.users"
	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error querying users: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user types.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Department)
		if err != nil {
			log.Printf("Error scanning commit row: %v\n", err)
			continue
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over user rows: %v\n", err)
		return nil, err
	}
	return users, nil
}

func (pg *Postgres) UpdateUser(ctx context.Context, userID string, u *types.UserRequest) error {
	query := "UPDATE public.users SET username=$1, password=$2 WHERE id=$3"
	_, err := pg.db.Exec(ctx, query, u.Username, u.Password, userID)
	if err != nil {
		log.Printf("Unable to update user: %v\n", err)
		return err
	}
	return nil
}

func (pg *Postgres) DeleteUser(ctx context.Context, userID string) error {
	_, err := pg.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("user with ID %s not found", userID)
	}
	query := "DELETE FROM public.users WHERE id=$1"
	_, err = pg.db.Exec(ctx, query, userID)
	if err != nil {
		log.Printf("Unable to delete user: %v\n", err)
		return err
	}
	return nil
}

func (pg *Postgres) Login(ctx context.Context, u *types.UserRequest) (string, error) {
	query := "SELECT id, username, password, role, email FROM public.users WHERE username = $1"
	row := pg.db.QueryRow(ctx, query, u.Username)
	var user types.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Department)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("user with username %s not found", u.Username)
		}
		log.Printf("Error scanning user data: %v\n", err)
		return "", err
	}

	if user.ID == "" {
		return "", fmt.Errorf("user with username %s not found", u.Username)
	}

	if err := util.CheckPassword(u.Password, user.Password); err != nil {
		return "", fmt.Errorf("invalid password for user %s", u.Username)
	}

	return user.ID, nil
}
