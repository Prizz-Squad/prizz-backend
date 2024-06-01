package services

import (
	"context"
	"github.com/EraldCaka/prizz-backend/db"
	"github.com/EraldCaka/prizz-backend/internal/types"
)

type TicketService interface {
	CreateTicket(ctx context.Context, ticketRequest *types.CreateTicketRequest) (string, error)
	GetAllTickets(ctx context.Context) ([]*types.Ticket, error)
	UpdateTicketDepartment(ctx context.Context, ticketID string, d *types.UpdateDepartment) error
	UpdateTicketStatus(ctx context.Context, ticketID string, d *types.UpdateStatus) error
	UpdateTicketDescription(ctx context.Context, ticketID string, d *types.UpdateDescription) error
	UpdateTicketPostDate(ctx context.Context, ticketID string, d *types.UpdatePostDate) error
}
type TicketStore struct {
	db *db.Postgres
}

func NewTicketService(db *db.Postgres) *TicketStore {
	return &TicketStore{db: db}
}

func (s *TicketStore) CreateTicket(ctx context.Context, ticketRequest *types.CreateTicketRequest) (string, error) {
	return s.db.CreateTicket(ctx, ticketRequest)
}

func (s *TicketStore) GetAllTickets(ctx context.Context) ([]*types.Ticket, error) {
	return s.db.GetTickets(ctx)
}

func (s *TicketStore) UpdateTicketDepartment(ctx context.Context, ticketID string, d *types.UpdateDepartment) error {
	return s.db.UpdateTicketDepartment(ctx, ticketID, d)
}

func (s *TicketStore) UpdateTicketStatus(ctx context.Context, ticketID string, d *types.UpdateStatus) error {
	return s.db.UpdateTicketStatus(ctx, ticketID, d)
}

func (s *TicketStore) UpdateTicketDescription(ctx context.Context, ticketID string, d *types.UpdateDescription) error {
	return s.db.UpdateTicketDescription(ctx, ticketID, d)
}

func (s *TicketStore) UpdateTicketPostDate(ctx context.Context, ticketID string, d *types.UpdatePostDate) error {
	return s.db.UpdateTicketPostDate(ctx, ticketID, d)
}
