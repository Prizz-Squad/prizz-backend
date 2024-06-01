package db

import (
	"context"
	"github.com/EraldCaka/prizz-backend/internal/types"
	"log"
)

func (pg *Postgres) CreateTicket(ctx context.Context, t *types.CreateTicketRequest) (string, error) {
	query := "INSERT INTO public.tasks (priority, status, department, user_id, project_id, created_at,post_date, description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	var ticketID string
	err := pg.db.QueryRow(ctx, query, t.Priority, t.Status, t.Department, t.UserID, t.ProjectID, t.CreatedAt, t.PostDate, t.Description).Scan(&ticketID)
	if err != nil {
		log.Printf("Unable to insert the ticket: %v\n", err)
		return "", err
	}
	return ticketID, nil
}

func (pg *Postgres) GetTickets(ctx context.Context) ([]*types.Ticket, error) {
	var tickets []*types.Ticket

	query := "SELECT * FROM public.tasks"
	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error querying tickets: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ticket types.Ticket
		err := rows.Scan(&ticket.ID, &ticket.Priority, &ticket.Status, &ticket.Department, &ticket.UserID, &ticket.ProjectID, &ticket.CreatedAt, &ticket.PostDate, &ticket.Description)
		if err != nil {
			log.Printf("Error scanning commit row: %v\n", err)
			continue
		}
		tickets = append(tickets, &ticket)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over ticket rows: %v\n", err)
		return nil, err
	}
	return tickets, nil
}

func (pg *Postgres) UpdateTicketDepartment(ctx context.Context, ticketID string, d *types.UpdateDepartment) error {
	query := "UPDATE public.tasks SET department=$1 WHERE id=$2"
	_, err := pg.db.Exec(ctx, query, d.Department, ticketID)
	if err != nil {
		log.Printf("Unable to update ticket: %v\n", err)
		return err
	}
	return nil
}

func (pg *Postgres) UpdateTicketStatus(ctx context.Context, ticketID string, d *types.UpdateStatus) error {
	query := "UPDATE public.tasks SET status=$1 WHERE id=$2"
	_, err := pg.db.Exec(ctx, query, d.Status, ticketID)
	if err != nil {
		log.Printf("Unable to update ticket: %v\n", err)
		return err
	}
	return nil
}

func (pg *Postgres) UpdateTicketDescription(ctx context.Context, ticketID string, d *types.UpdateDescription) error {
	query := "UPDATE public.tasks SET description=$1 WHERE id=$2"
	_, err := pg.db.Exec(ctx, query, d.Description, ticketID)
	if err != nil {
		log.Printf("Unable to update ticket: %v\n", err)
		return err
	}
	return nil
}

func (pg *Postgres) UpdateTicketPostDate(ctx context.Context, ticketID string, d *types.UpdatePostDate) error {
	query := "UPDATE public.tasks SET post_date=$1 WHERE id=$2"
	_, err := pg.db.Exec(ctx, query, d.PostDate, ticketID)
	if err != nil {
		log.Printf("Unable to update ticket: %v\n", err)
		return err
	}
	return nil
}
