package repository

import (
	"context"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"gorm.io/gorm"
)

type accountBalanceRepository struct {
	db *gorm.DB
}

func NewAccountBalance(db *gorm.DB) *accountBalanceRepository {
	return &accountBalanceRepository{
		db: db,
	}
}

func (r *accountBalanceRepository) FindByUserID(ctx context.Context, userID int64) (*entity.AccountBalance, error) {
	accountBalance := new(entity.AccountBalance)
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(accountBalance).Error
	if err != nil {
		return nil, err
	}
	return accountBalance,nil
}

func (r *accountBalanceRepository) FindByAccountNumber(ctx context.Context, accountNumber string) (*entity.AccountBalance, error) {
	accountBalance := new(entity.AccountBalance)
	err := r.db.WithContext(ctx).Where("account_number = ?", accountNumber).First(accountBalance).Error
	if err != nil {
		return nil, err
	}
	return accountBalance,nil
}

func (r *accountBalanceRepository) Update(ctx context.Context, accountBalance *entity.AccountBalance) error {
	query := r.db.WithContext(ctx).Model(&entity.AccountBalance{}).Where("id = ?", accountBalance.ID)
	if accountBalance.Balance != 0 {
		query = query.Update("balance", accountBalance.Balance)
	}
	if err := query.Error; err != nil {
		return err
	}
	return nil
}