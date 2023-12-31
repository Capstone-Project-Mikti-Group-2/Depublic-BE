package repository

import (
	"context"
	"time"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (r *EventRepository) CreateEvent(ctx context.Context, event *entity.Event) error {
	if err := r.db.WithContext(ctx).Create(event).Error; err != nil {
		return err
	}
	return nil
}

func (r *EventRepository) UpdateEvent(ctx context.Context, event *entity.Event) error {
	query := r.db.WithContext(ctx).Model(event).Where("id = ?", event.ID).Updates(event)

	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *EventRepository) DeleteEvent(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Event{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *EventRepository) FindAllEvent(ctx context.Context) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0)
	err := r.db.WithContext(ctx).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) FindEventByID(ctx context.Context, id int64) (*entity.Event, error) {
	event := new(entity.Event)
	err := r.db.WithContext(ctx).Where("id = ?", id).First(event).Error
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (r *EventRepository) GetEvent(ctx context.Context, id int64) (*entity.Event, error) {
	event := new(entity.Event)
	result := r.db.WithContext(ctx).First(&event, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return event, nil
}

func (r *EventRepository) FilterEventByPrice(ctx context.Context, min, max string) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0)
	err := r.db.WithContext(ctx).Where("price >= ? And price <= ?", min, max).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) FilterEventByDate(ctx context.Context, startDate, endDate time.Time) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0)
	err := r.db.WithContext(ctx).Where("start_date >= ? And end_date <= ?", startDate, endDate).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) FilterEventByLocation(ctx context.Context, location string) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0)
	err := r.db.WithContext(ctx).Where("location ILIKE ?", "%"+location+"%").Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) SearchEvent(ctx context.Context, keyword string) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0)
	err := r.db.WithContext(ctx).Where("name ILIKE ?", "%"+keyword+"%").Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) FilterEventByAvailable(ctx context.Context, available bool) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0)
	err := r.db.WithContext(ctx).Where("available = ?", available).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) SortEventByExpensive(ctx context.Context) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0)
	err := r.db.WithContext(ctx).Order("price DESC").Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) SortEventByCheapest(ctx context.Context) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0)
	err := r.db.WithContext(ctx).Order("price ASC").Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) SortEventByNewest(ctx context.Context) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0)
	err := r.db.WithContext(ctx).Order("created_at DESC").Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}
