package repository

import (
	"context"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// FindByEmail implements service.LoginRepository.
func (*UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	panic("unimplemented")
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// membuat User Baru (Create User)
func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

// Update User by ID (Update User)
func (r *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	query := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", user.ID)
	if user.Name != "" {
		query = query.Update("name", user.Name)
	}
	if user.Email != "" {
		query = query.Update("email", user.Email)
	}
	if user.Number != "" {
		query = query.Update("number", user.Number)
	}
	if user.Password != "" {
		query = query.Update("password", user.Password)
	}
	if user.Role != "" {
		query = query.Update("role", user.Role)
	}
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

// Delete User by ID (Delete User)
func (r *UserRepository) DeleteUser(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
		return err
	}
	return nil
}

// Get All User (Find All User)
func (r *UserRepository) FindAllUser(ctx context.Context) ([]*entity.User, error) {
	users := make([]*entity.User, 0)
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Get User by ID (Find User by ID)
func (r *UserRepository) FindUserByID(ctx context.Context, id int64) (*entity.User, error) {
	user := new(entity.User)
	err := r.db.WithContext(ctx).Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Get User by Email (Find User by Email)
func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := new(entity.User)
	err := r.db.WithContext(ctx).Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Get User by Username (Find User by Username)
func (r *UserRepository) FindUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	user := new(entity.User)
	err := r.db.WithContext(ctx).Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Get User by Number (Find User by Number)
func (r *UserRepository) FindUserByNumber(ctx context.Context, number string) (*entity.User, error) {
	user := new(entity.User)
	err := r.db.WithContext(ctx).Where("number = ?", number).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil

}
