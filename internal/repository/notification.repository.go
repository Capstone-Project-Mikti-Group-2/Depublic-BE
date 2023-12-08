package repository

import (
	"context"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{
		db: db,
	}
}

func (r *NotificationRepository) CreateNotification(ctx context.Context, notification *entity.Notification) error {
	result := r.db.WithContext(ctx).Create(&notification)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *NotificationRepository) GetAllNotification(ctx context.Context) ([]*entity.Notification, error) {
	Notifications := make([]*entity.Notification, 0)
	result := r.db.WithContext(ctx).Find(&Notifications)
	if result.Error != nil {
		return nil, result.Error
	}

	return Notifications, nil
}

func (r *NotificationRepository) NotificationIsRead(ctx context.Context, notificationID int) error {
	result := r.db.WithContext(ctx).Model(&entity.Notification{}).Where("id = ?", notificationID).Update("is_read", true)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *NotificationRepository) UserGetNotification(ctx context.Context) ([]*entity.Notification, error) {
	notification := make([]*entity.Notification, 0)

	result := r.db.WithContext(ctx).Where("is_read = ?", false).Find(&notification)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, n := range notification {
		err := r.NotificationIsRead(ctx, n.ID)
		if err != nil {
			return nil, err
		}
	}

	return notification, nil
}
