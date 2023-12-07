package repository

import (
	"context"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"gorm.io/gorm"
)

type TopupRepository interface {
	InputTopup(ctx context.Context, topup entity.TopUp) (entity.TopUp, error)
	UserTopup(ctx context.Context, topup entity.TopUp) (entity.TopUp, error)
	GetUserByID(ctx context.Context, id int64) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
}

type topupRepository struct {
	db *gorm.DB
}

func (r *topupRepository) InputTopup(ctx context.Context, topup entity.TopUp) (entity.TopUp, error) {
	result := r.db.WithContext(ctx).Create(&topup)
	if result.Error != nil {
		return entity.TopUp{}, result.Error
	}

	return topup, nil
}

func (r *topupRepository) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	user := new(entity.User)
	result := r.db.WithContext(ctx).First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *topupRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *topupRepository) UserTopup(ctx context.Context, topup entity.TopUp) (entity.TopUp, error) {
	result := r.db.WithContext(ctx).Create(&topup)
	if result.Error != nil {
		return entity.TopUp{}, result.Error
	}

	return topup, nil
}
