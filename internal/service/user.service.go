package service

import (
	"context"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id int64) error
	FindAllUser(ctx context.Context) ([]*entity.User, error)
	FindUserByID(ctx context.Context, id int64) (*entity.User, error)
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
	FindUserByNumber(ctx context.Context, number string) (*entity.User, error)
	FindUserByUsername(ctx context.Context, username string) (*entity.User, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id int64) error
	FindAllUser(ctx context.Context) ([]*entity.User, error)
	FindUserByID(ctx context.Context, id int64) (*entity.User, error)
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
	FindUserByNumber(ctx context.Context, number string) (*entity.User, error)
	FindUserByUsername(ctx context.Context, username string) (*entity.User, error)
}

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *entity.User) error {
	return s.repository.CreateUser(ctx, user)
}

func (s *UserService) UpdateUser(ctx context.Context, user *entity.User) error {
	return s.repository.UpdateUser(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	return s.repository.DeleteUser(ctx, id)
}

func (s *UserService) FindAllUser(ctx context.Context) ([]*entity.User, error) {
	return s.repository.FindAllUser(ctx)
}

func (s *UserService) FindUserByID(ctx context.Context, id int64) (*entity.User, error) {
	return s.repository.FindUserByID(ctx, id)
}

func (s *UserService) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	return s.repository.FindUserByEmail(ctx, email)
}

func (s *UserService) FindUserByNumber(ctx context.Context, number string) (*entity.User, error) {
	return s.repository.FindUserByNumber(ctx, number)
}

func (s *UserService) FindUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	return s.repository.FindUserByUsername(ctx, username)
}
