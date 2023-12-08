package service

import (
	"context"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
)

type NotificationUseCase interface {
	CreateNotification(ctx context.Context, notification *entity.Notification) error
	GetAllNotification(ctx context.Context) ([]*entity.Notification, error)
	UserGetNotification(ctx context.Context) ([]*entity.Notification, error)
}

type NotificationRepository interface {
	CreateNotification(ctx context.Context, notification *entity.Notification) error
	GetAllNotification(ctx context.Context) ([]*entity.Notification, error)
	UserGetNotification(ctx context.Context) ([]*entity.Notification, error)
}

type NotificationService struct {
	Repository NotificationRepository
}

func NewNotificationService(repository NotificationRepository) *NotificationService {
	return &NotificationService{
		Repository: repository,
	}
}

func (s *NotificationService) CreateNotification(ctx context.Context, notification *entity.Notification) error {
	return s.Repository.CreateNotification(ctx, notification)
}

func (s *NotificationService) GetAllNotification(ctx context.Context) ([]*entity.Notification, error) {
	return s.Repository.GetAllNotification(ctx)
}

func (s *NotificationService) UserGetNotification(ctx context.Context) ([]*entity.Notification, error) {
	return s.Repository.UserGetNotification(ctx)
}
