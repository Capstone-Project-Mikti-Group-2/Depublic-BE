package service

import (
	"context"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
)

type RegistrationUseCase interface {
	Registration(ctx context.Context, user *entity.User) error
}

type RegistrationRepository interface {
	Registration(ctx context.Context, user *entity.User) error
}

type RegistrationService struct {
	repo RegistrationRepository
}

func NewRegistrationService(repo RegistrationRepository) *RegistrationService {
	return &RegistrationService{
		repo: repo,
	}
}

func (s *RegistrationService) Registration(ctx context.Context, user *entity.User) error {
	err := s.repo.Registration(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
