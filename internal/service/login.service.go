package service

import (
	"context"
	"errors"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase interface {
	Login(ctx context.Context, email, password string) (*entity.User, error)
}

type LoginRepository interface {
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}

type LoginService struct {
	repo LoginRepository
}

func NewLoginService(repo LoginRepository) *LoginService {
	return &LoginService{
		repo: repo,
	}
}

func (s *LoginService) Login(ctx context.Context, email, password string) (*entity.User, error) {
	// Find the user by email
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// Check if user exists
	if user == nil {
		return nil, errors.New("user with that email not found")
	}

	// Compare the hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("incorrect login credentials")
	}

	return user, nil
}
