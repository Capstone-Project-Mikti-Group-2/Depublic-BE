package service

import (
	"context"
	"errors"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
)

type TicketRepository interface {
	CreateTicket(ctx context.Context, ticket *entity.Ticket) error
	UpdateTicket(ctx context.Context, event *entity.Event) error
	FindAllTicket(ctx context.Context, eventID int64) (*entity.Event, error)
	FindTicketByID(ctx context.Context, id int64) (*entity.Event, error)
	GetBooking(ctx context.Context) ([]*entity.Ticket, error)
	GetBookingByID(ctx context.Context, UserID int64) ([]*entity.Ticket, error)
	GetUserSaldo(ctx context.Context, userID int64) (int64, error)
	UpdateUserSaldo(ctx context.Context, userID int64, total int64) error
	GetBookingPrice(ctx context.Context, eventID int64) (int64, error)
}

type TicketUseCase interface {
	CreateTicket(ctx context.Context, ticket *entity.Ticket) error
	UpdateTicket(ctx context.Context, event *entity.Event) error
	FindAllTicket(ctx context.Context, eventID int64) (*entity.Event, error)
	FindTicketByID(ctx context.Context, id int64) (*entity.Event, error)
	GetBooking(ctx context.Context) ([]*entity.Ticket, error)
	GetBookingByID(ctx context.Context, UserID int64) ([]*entity.Ticket, error)
	GetUserSaldo(ctx context.Context, userID int64) (int64, error)
	UpdateUserSaldo(ctx context.Context, userID int64, total int64) error
	GetBookingPrice(ctx context.Context, eventID int64) (int64, error)
}

type TicketService struct {
	repository TicketRepository
}

func NewTicketService(repository TicketRepository) *TicketService {
	return &TicketService{
		repository: repository,
	}
}

func (s *TicketService) CreateTicket(ctx context.Context, ticket *entity.Ticket) error {
	event, err := s.repository.FindAllTicket(ctx, ticket.EventID)
	if err != nil {
		return err
	}

	if int64(event.Quantity) < ticket.Quantity {
		return errors.New("ticket is not available")
	}

	ticket.Total = event.Price * int64(ticket.Quantity)

	if err := s.repository.CreateTicket(ctx, ticket); err != nil {
		return err
	}

	event.Quantity -= ticket.Quantity
	if err := s.repository.UpdateTicket(ctx, event); err != nil {
		return err
	}

	if err := s.repository.UpdateUserSaldo(ctx, ticket.UserID, ticket.Total); err != nil {
		return err
	}

	return nil
}

func (s *TicketService) FindAllTicket(ctx context.Context, eventID int64) (*entity.Event, error) {
	return s.repository.FindAllTicket(ctx, eventID)
}

func (s *TicketService) FindTicketByID(ctx context.Context, id int64) (*entity.Event, error) {
	return s.repository.FindTicketByID(ctx, id)
}

func (s *TicketService) UpdateTicket(ctx context.Context, event *entity.Event) error {
	return s.repository.UpdateTicket(ctx, event)
}

func (s *TicketService) GetBooking(ctx context.Context) ([]*entity.Ticket, error) {
	return s.repository.GetBooking(ctx)
}

func (s *TicketService) UpdateUserSaldo(ctx context.Context, userID int64, total int64) error {
	return s.repository.UpdateUserSaldo(ctx, userID, total)
}

func (s *TicketService) GetBookingByID(ctx context.Context, UserID int64) ([]*entity.Ticket, error) {
	return s.repository.GetBookingByID(ctx, UserID)
}

func (s *TicketService) GetBookingPrice(ctx context.Context, eventID int64) (int64, error) {
	return s.repository.GetBookingPrice(ctx, eventID)
}

func (s *TicketService) GetUserSaldo(ctx context.Context, userID int64) (int64, error) {
	return s.repository.GetUserSaldo(ctx, userID)
}
