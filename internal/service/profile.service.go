package service

import (
	"context"
	"errors"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	CreateProfile(ctx context.Context, profile *entity.Profile) error
	UpdateProfile(ctx context.Context, profile *entity.Profile) error
	DeleteProfile(ctx context.Context, id int64) error
	GetProfileByID(ctx context.Context, id int64) (*entity.Profile, error)
}

type ProfileUseCase interface {
	CreateProfile(ctx context.Context, profile *entity.Profile) error
	UpdateProfile(ctx context.Context, profile *entity.Profile) error
	DeleteProfile(ctx context.Context, id int64) error
	GetProfileByID(ctx context.Context, id int64) (*entity.Profile, error)
}

type ProfileService struct {
	repository ProfileRepository
}

func NewProfileService(repository ProfileRepository) *ProfileService {
	return &ProfileService{
		repository: repository,
	}
}

func (s *ProfileService) CreateProfile(ctx context.Context, profile *entity.Profile) error {
	return s.repository.CreateProfile(ctx, profile)
}

func (s *ProfileService) UpdateProfile(ctx context.Context, profile *entity.Profile) error {
	existingProfile, err := s.repository.GetProfileByID(ctx, profile.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("profile not found")
		}
		return err
	}

	existingProfile.UserID = profile.UserID
	existingProfile.Address = profile.Address
	existingProfile.Avatar = profile.Avatar

	if err := s.repository.UpdateProfile(ctx, existingProfile); err != nil {
		return err
	}
	return nil
}

func (s *ProfileService) DeleteProfile(ctx context.Context, id int64) error {
	return s.repository.DeleteProfile(ctx, id)
}

func (s *ProfileService) GetProfileByID(ctx context.Context, id int64) (*entity.Profile, error) {
	return s.repository.GetProfileByID(ctx, id)
}
