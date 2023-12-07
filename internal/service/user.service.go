package service

import (
	"context"
	"errors"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id int64) error
	FindAllUser(ctx context.Context) ([]*entity.User, error)
	FindUserByID(ctx context.Context, id int64) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindUserByNumber(ctx context.Context, number string) (*entity.User, error)
	FindUserByUsername(ctx context.Context, username string) (*entity.User, error)
	FindByID(ctx context.Context, id int64) (*entity.User, error)
	UpdateSaldo(ctx context.Context, userID int64, updatedSaldo int64) error
	DeleteAccount(ctx context.Context, email string) error
	UpdateSelfUser(ctx context.Context, user *entity.User) error
	InputSaldo(ctx context.Context, user *entity.User) error
	Logout(ctx context.Context, user *entity.User) error
}
type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id int64) error
	FindAllUser(ctx context.Context) ([]*entity.User, error)
	FindUserByID(ctx context.Context, id int64) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindUserByNumber(ctx context.Context, number string) (*entity.User, error)
	FindUserByUsername(ctx context.Context, username string) (*entity.User, error)
	FindByID(ctx context.Context, id int64) (*entity.User, error)
	UpdateSaldo(ctx context.Context, userID int64, updatedSaldo int64) error
	DeleteAccount(ctx context.Context, email string) error
	UpdateSelfUser(ctx context.Context, user *entity.User) error
	InputSaldo(ctx context.Context, user *entity.User) error
	Logout(ctx context.Context, user *entity.User) error
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
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hasedPassword)
	return s.repository.CreateUser(ctx, user)
}

func (s *UserService) UpdateUser(ctx context.Context, user *entity.User) error {
	if user.Role != "" {
		if user.Role != "Administrator" && user.Role != "User" {
			return errors.New("role must be Administrator or User")
		}
	}
	if user.Password != "" {
		hasedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hasedPassword)
	}
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

func (s *UserService) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	return s.repository.FindByEmail(ctx, email)
}

func (s *UserService) FindUserByNumber(ctx context.Context, number string) (*entity.User, error) {
	return s.repository.FindUserByNumber(ctx, number)
}

func (s *UserService) FindUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	return s.repository.FindUserByUsername(ctx, username)
}

func (s *UserService) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *UserService) UpdateSaldo(ctx context.Context, userID int64, updatedSaldo int64) error {
	return s.repository.UpdateSaldo(ctx, userID, updatedSaldo)
}

func (s *UserService) DeleteAccount(ctx context.Context, email string) error {
	return s.repository.DeleteAccount(ctx, email)
}

func (s *UserService) UpdateSelfUser(ctx context.Context, user *entity.User) error {
	return s.repository.UpdateSelfUser(ctx, user)
}

func (s *UserService) InputSaldo(ctx context.Context, user *entity.User) error {
	return s.repository.InputSaldo(ctx, user)
}

func (s *UserService) Logout(ctx context.Context, user *entity.User) error {
	return s.repository.Logout(ctx, user)
}
