package repository

import (
	"context"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"gorm.io/gorm"
)

type RepositoryTopup struct {
	db *gorm.DB
}

func NewTopUP(db *gorm.DB) *RepositoryTopup {
	return &RepositoryTopup{
		db: db,
	}
}

// FindByID implements entity.TopUpRepository.
func (r *RepositoryTopup) FindByID(ctx context.Context, id string) (*entity.TopUp,error) {
	topup := new(entity.TopUp)
	err := r.db.WithContext(ctx).Where("id = ?", id).First(topup).Error
	if err != nil {
		return nil, err
	}
	return topup, nil
}

// Insert implements entity.TopUpRepository.
func (r *RepositoryTopup) Insert(ctx context.Context, topup *entity.TopUp) error {
	err := r.db.WithContext(ctx).Create(&topup).Error
	if err != nil {
		return err
	}
	return nil
}

// Update implements entity.TopUpRepository.
func (r *RepositoryTopup) Update(ctx context.Context, topup *entity.TopUp) error {
	query := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", topup.ID)
	if topup.Status != 0 {
		query = query.Update("status", topup.Status)
	}
	if topup.SnapURL != "" {
		query = query.Update("snap_url", topup.SnapURL)
	}
	if topup.Amount != 0 {
		query = query.Update("amount", topup.Amount)
	}
	if err := query.Error; err != nil {
		return err
	}
	return nil 

}


