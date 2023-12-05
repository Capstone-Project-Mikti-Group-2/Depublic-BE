package service

import (
	"context"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
)

type AccountBalanceRepository interface {
	FindByUserID(ctx context.Context, userID int64) (*entity.AccountBalance, error)
	FindByAccountNumber(ctx context.Context, accountNumber string) (*entity.AccountBalance, error)
	Update(ctx context.Context, accountBalance *entity.AccountBalance) error
}
