package service

import (
	"context"
	"time"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
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
	if err := s.repository.CreateEvent(ctx, event); err != nil {
		return err
	}
	return nil
}

func (s *EventService) UpdateEvent(ctx context.Context, event *entity.Event) error {
	if err := s.repository.UpdateEvent(ctx, event); err != nil {
		return err
	}
	return nil
}

func (s *EventService) DeleteEvent(ctx context.Context, id int64) error {
	if err := s.repository.DeleteEvent(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *EventService) FindAllEvent(ctx context.Context) ([]*entity.Event, error) {
	events, err := s.repository.FindAllEvent(ctx)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *EventService) FindEventByID(ctx context.Context, id int64) (*entity.Event, error) {
	event, err := s.repository.FindEventByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (s *EventService) SearchEvent(ctx context.Context, keyword string) ([]*entity.Event, error) {
	events, err := s.repository.SearchEvent(ctx, keyword)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *EventService) FilterEventByPrice(ctx context.Context, min, max string) ([]*entity.Event, error) {
	events, err := s.repository.FilterEventByPrice(ctx, min, max)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *EventService) FilterEventByDate(ctx context.Context, startDate, endDate time.Time) ([]*entity.Event, error) {
	events, err := s.repository.FilterEventByDate(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *EventService) FilterEventByLocation(ctx context.Context, location string) ([]*entity.Event, error) {
	events, err := s.repository.FilterEventByLocation(ctx, location)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *EventService) FilterEventByAvailable(ctx context.Context, available bool) ([]*entity.Event, error) {
	events, err := s.repository.FilterEventByAvailable(ctx, available)
	if err != nil {
		return nil, err
	}
	return events, nil
}
