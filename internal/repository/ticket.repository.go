package repository

import (
	"context"
	"errors"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *TicketRepository {
	return &TicketRepository{
		db: db,
	}
}

func (r *TicketRepository) CreateTicket(ctx context.Context, ticket *entity.Ticket) error {
	if err := r.db.WithContext(ctx).Create(&ticket).Error; err != nil {
		return err
	}
	return nil
}

func (r *TicketRepository) UpdateTicket(ctx context.Context, event *entity.Event) error {
	query := r.db.WithContext(ctx).Model(&entity.Event{}).Where("id = ?", event.ID).Updates(&event)

	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *TicketRepository) DeleteTicket(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Ticket{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *TicketRepository) FindAllTicket(ctx context.Context, eventID int64) (*entity.Event, error) {
	event := new(entity.Event)
	err := r.db.WithContext(ctx).Where("id = ?", eventID).First(&event).Error
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (r *TicketRepository) FindTicketByID(ctx context.Context, id int64) (*entity.Event, error) {
	event := new(entity.Event)
	result := r.db.WithContext(ctx).First(&event, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return event, nil
}

func (r *TicketRepository) GetBooking(ctx context.Context) ([]*entity.Ticket, error) {
	booking := make([]*entity.Ticket, 0)
	err := r.db.WithContext(ctx).Preload("Event").Find(&booking).Error
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (r *TicketRepository) GetBookingByID(ctx context.Context, UserID int64) ([]*entity.Ticket, error) {
	booking := make([]*entity.Ticket, 0)
	err := r.db.WithContext(ctx).Where("user_id = ?", UserID).Preload("Event").Find(&booking).Error
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (r *TicketRepository) GetUserSaldo(ctx context.Context, UserID int64) (int64, error) {
	user := new(entity.User)
	if err := r.db.WithContext(ctx).Where("id = ?", UserID).First(user).Error; err != nil {
		return 0, err
	}
	return user.Saldo, nil
}

func (r *TicketRepository) UpdateUserSaldo(ctx context.Context, UserID int64, total int64) error {
	user := new(entity.User)
	if err := r.db.WithContext(ctx).Where("id = ?", UserID).First(user).Error; err != nil {
		return err
	}

	if user.Saldo < total {
		return errors.New("insufficient balance")
	}

	user.Saldo -= total
	if err := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", UserID).Updates(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *TicketRepository) GetBookingPrice(ctx context.Context, eventID int64) (int64, error) {
	event := new(entity.Event)
	if err := r.db.WithContext(ctx).Where("id = ?", eventID).First(event).Error; err != nil {
		return 0, err
	}
	return event.Price, nil
}
