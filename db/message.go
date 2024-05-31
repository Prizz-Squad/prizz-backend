package db

import (
	"context"
	"github.com/EraldCaka/prizz-backend/internal/types"
	"log"
)

func (pg *Postgres) CreateMessage(ctx context.Context, m *types.MessageRequest) (string, error) {
	query := "INSERT INTO public.message (description, user_id, task_id, contents) VALUES ($1, $2, $3, $4) RETURNING id"
	var messID string
	err := pg.db.QueryRow(ctx, query, m.Description, m.UserId, m.TaskId, m.Contents).Scan(&messID)
	if err != nil {
		log.Printf("Unable to insert messages: %v\n", err)
		return "", err
	}
	return messID, nil
}

func (pg *Postgres) GetMessages(ctx context.Context) ([]*types.Message, error) {
	var messages []*types.Message
	query := "SELECT * FROM public.message"
	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error querying messages: %v\n", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var mess types.Message
		err := rows.Scan(&mess.ID, &mess.Description, &mess.UserId, &mess.TaskId, &mess.Contents)
		if err != nil {
			log.Printf("Error scanning commit row: %v\n", err)
			continue
		}
		messages = append(messages, &mess)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over user rows: %v\n", err)
		return nil, err
	}
	return messages, nil
}
