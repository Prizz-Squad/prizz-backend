package db

import (
	"context"
	"github.com/EraldCaka/prizz-backend/internal/types"
	"log"
)

func (pg *Postgres) CreateTicket(ctx context.Context, t *types.CreateTicketRequest) (string, error) {
	query := "INSERT INTO public.users (priority, status, department, images, user_id, project_id, created_at, description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	var ticketID string
	err := pg.db.QueryRow(ctx, query, t.Priority, t.Status, t.Department, t.Images, t.UserID, t.ProjectID, t.CreatedAt, t.Description).Scan(&ticketID)
	if err != nil {
		log.Printf("Unable to insert the ticket: %v\n", err)
		return "", err
	}
	return ticketID, nil
}

func (pg *Postgres) GetTickets(ctx context.Context) ([]*types.Ticket, error) {
	var tickets []*types.Ticket

	query := "SELECT * FROM public.ticket"
	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error querying tickets: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ticket types.Ticket
		err := rows.Scan(&ticket.ID, &ticket.Priority, &ticket.Status, &ticket.Department, &ticket.Images, &ticket.UserID, &ticket.ProjectID, &ticket.CreatedAt, &ticket.Description)
		if err != nil {
			log.Printf("Error scanning commit row: %v\n", err)
			continue
		}
		tickets = append(tickets, &ticket)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over user rows: %v\n", err)
		return nil, err
	}
	return tickets, nil
}
