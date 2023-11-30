package repository

import (
	"context"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"gorm.io/gorm"
)

type ProfileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{
		db: db,
	}
}

func (r *ProfileRepository) CreateProfile(ctx context.Context, profile *entity.Profile) error {
	if err := r.db.WithContext(ctx).Create(profile).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProfileRepository) UpdateProfile(ctx context.Context, profile *entity.Profile) error {
	query := r.db.WithContext(ctx).Model(profile).Where("id = ?", profile.ID).Updates(profile)

	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *ProfileRepository) DeleteProfile(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Profile{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProfileRepository) GetProfileByID(ctx context.Context, id int64) (*entity.Profile, error) {
	profile := new(entity.Profile)
	err := r.db.WithContext(ctx).Where("id = ?", id).First(profile).Error
	if err != nil {
		return nil, err
	}
	return profile, nil
}
