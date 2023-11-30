package service

import (
	"context"
	"errors"
	"time"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"gorm.io/gorm"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, event *entity.Event) error
	UpdateEvent(ctx context.Context, event *entity.Event) error
	DeleteEvent(ctx context.Context, id int64) error
	FindAllEvent(ctx context.Context) ([]*entity.Event, error)
	FindEventByID(ctx context.Context, id int64) (*entity.Event, error)
	SearchEvent(ctx context.Context, keyword string) ([]*entity.Event, error)
	FilterEventByPrice(ctx context.Context, min, max string) ([]*entity.Event, error)
	FilterEventByDate(ctx context.Context, startDate, endDate time.Time) ([]*entity.Event, error)
	FilterEventByLocation(ctx context.Context, location string) ([]*entity.Event, error)
	FilterEventByAvailable(ctx context.Context, available bool) ([]*entity.Event, error)
	SortEventByExpensive(ctx context.Context) ([]*entity.Event, error)
	SortEventByCheapest(ctx context.Context) ([]*entity.Event, error)
	SortEventByNewest(ctx context.Context) ([]*entity.Event, error)
}

type EventUseCase interface {
	CreateEvent(ctx context.Context, event *entity.Event) error
	UpdateEvent(ctx context.Context, event *entity.Event) error
	DeleteEvent(ctx context.Context, id int64) error
	FindAllEvent(ctx context.Context) ([]*entity.Event, error)
	FindEventByID(ctx context.Context, id int64) (*entity.Event, error)
	SearchEvent(ctx context.Context, keyword string) ([]*entity.Event, error)
	FilterEventByPrice(ctx context.Context, min, max string) ([]*entity.Event, error)
	FilterEventByDate(ctx context.Context, startDate, endDate time.Time) ([]*entity.Event, error)
	FilterEventByLocation(ctx context.Context, location string) ([]*entity.Event, error)
	FilterEventByAvailable(ctx context.Context, available bool) ([]*entity.Event, error)
	SortEventByExpensive(ctx context.Context) ([]*entity.Event, error)
	SortEventByCheapest(ctx context.Context) ([]*entity.Event, error)
	SortEventByNewest(ctx context.Context) ([]*entity.Event, error)
}

type EventService struct {
	repository EventRepository
}

func NewEventService(repository EventRepository) *EventService {
	return &EventService{
		repository: repository,
	}
}

func (s *EventService) CreateEvent(ctx context.Context, event *entity.Event) error {
	return s.repository.CreateEvent(ctx, event)
}

func (s *EventService) UpdateEvent(ctx context.Context, event *entity.Event) error {
	existingEvent, err := s.repository.FindEventByID(ctx, event.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("event not found")
		}
		return err
	}

	existingEvent.ID = event.ID
	existingEvent.Name = event.Name
	existingEvent.Description = event.Description
	existingEvent.StartDate = event.StartDate
	existingEvent.EndDate = event.EndDate
	existingEvent.Location = event.Location
	existingEvent.Price = event.Price
	existingEvent.Quantity = event.Quantity
	existingEvent.Available = event.Available
	existingEvent.Image = event.Image

	if err := s.repository.UpdateEvent(ctx, existingEvent); err != nil {
		return err
	}
	return nil
}

func (s *EventService) DeleteEvent(ctx context.Context, id int64) error {
	return s.repository.DeleteEvent(ctx, id)
}

func (s *EventService) FindAllEvent(ctx context.Context) ([]*entity.Event, error) {
	return s.repository.FindAllEvent(ctx)
}

func (s *EventService) FindEventByID(ctx context.Context, id int64) (*entity.Event, error) {
	return s.repository.FindEventByID(ctx, id)
}

func (s *EventService) SearchEvent(ctx context.Context, keyword string) ([]*entity.Event, error) {
	return s.repository.SearchEvent(ctx, keyword)
}

func (s *EventService) FilterEventByPrice(ctx context.Context, min, max string) ([]*entity.Event, error) {
	return s.repository.FilterEventByPrice(ctx, min, max)
}

func (s *EventService) FilterEventByDate(ctx context.Context, startDate, endDate time.Time) ([]*entity.Event, error) {
	return s.repository.FilterEventByDate(ctx, startDate, endDate)
}

func (s *EventService) FilterEventByLocation(ctx context.Context, location string) ([]*entity.Event, error) {
	return s.repository.FilterEventByLocation(ctx, location)
}

func (s *EventService) FilterEventByAvailable(ctx context.Context, available bool) ([]*entity.Event, error) {
	return s.repository.FilterEventByAvailable(ctx, available)
}

func (s *EventService) SortEventByExpensive(ctx context.Context) ([]*entity.Event, error) {
	return s.repository.SortEventByExpensive(ctx)
}

func (s *EventService) SortEventByCheapest(ctx context.Context) ([]*entity.Event, error) {
	return s.repository.SortEventByCheapest(ctx)
}

func (s *EventService) SortEventByNewest(ctx context.Context) ([]*entity.Event, error) {
	return s.repository.SortEventByNewest(ctx)
}
