package repository

import (
	"context"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
)

type TopUpRepository interface {
	CreateMidtransTopup(ctx context.Context, OrderID string, nominal int64) (*coreapi.ChargeResponse, error)
	CreateTopUp(ctx context.Context, topup entity.TopUp) (entity.TopUp, error)
	UserTopup(ctx context.Context, topup entity.TopUp, userID int64) (entity.TopUp, error)
	FindUserID(ctx context.Context, id int64) (*entity.User, error)
	UpdateUserTopup(ctx context.Context, user *entity.User) error
}
type topupRepository struct {
	db *gorm.DB
}

func NewTopUpRepository(db *gorm.DB) *topupRepository {
	return &topupRepository{
		db: db,
	}
}

func (r *topupRepository) CreateTopUp(ctx context.Context, topup entity.TopUp) (entity.TopUp, error) {
	if err := r.db.WithContext(ctx).Create(&topup).Error; err != nil {
		return entity.TopUp{}, err
	}
	return topup, nil
}

func (r *topupRepository) UserTopup(ctx context.Context, topup entity.TopUp) (entity.TopUp, error) {
	if err := r.db.WithContext(ctx).Create(&topup).Error; err != nil {
		return entity.TopUp{}, err
	}
	return topup, nil
}

func (r *topupRepository) FindUserID(ctx context.Context, id int64) (*entity.User, error) {
	user := new(entity.User)
	result := r.db.WithContext(ctx).First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *topupRepository) UpdateUserTopup(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", user.ID).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}
