package service

import (
	"context"
	"errors"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"golang.org/x/crypto/bcrypt"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user.Password = string(hashedPassword)

	if err := s.repo.Registration(ctx, user); err != nil {
		return err
	}

	return nil
}
