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
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// cek database untuk email
	if user == nil {
		return nil, errors.New("user with that email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("incorrect login credentials")
	}

	return user, nil
}
