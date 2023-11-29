package service

import (
	"context"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, event *entity.Event) error
	UpdateEvent(ctx context.Context, event *entity.Event) error
	DeleteEvent(ctx context.Context, id int64) error
	FindAllEvent(ctx context.Context) ([]*entity.Event, error)
	FindEventByID(ctx context.Context, id int64) (*entity.Event, error)
}

type EventUseCase interface {
	CreateEvent(ctx context.Context, event *entity.Event) error
	UpdateEvent(ctx context.Context, event *entity.Event) error
	DeleteEvent(ctx context.Context, id int64) error
	FindAllEvent(ctx context.Context) ([]*entity.Event, error)
	FindEventByID(ctx context.Context, id int64) (*entity.Event, error)
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
