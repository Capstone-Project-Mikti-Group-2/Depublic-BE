package service

import (
	"context"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
)

type TransactionUseCase interface {
	Create(ctx context.Context, transaction *entity.Transaction) error
	FindByOrderID(ctx context.Context, orderID string) (*entity.Transaction, error)
	FindByUserID(ctx context.Context, userID int64) ([]*entity.Transaction, error)
	UpdateStatus(ctx context.Context, orderID string, status string) error
}

type TransactionReposiotry interface {
	Create(ctx context.Context, transaction *entity.Transaction) error
	FindByOrderID(ctx context.Context, orderID string) (*entity.Transaction, error)
	FindByUserID(ctx context.Context, userID int64) ([]*entity.Transaction, error)
	UpdateStatus(ctx context.Context, orderID string, status string) error
}

type TransactionService struct {
	transactionRepo TransactionReposiotry
}

func NewTransactionService(transactionRepo TransactionReposiotry) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
	}
}
func (s *TransactionService) Create(ctx context.Context, transaction *entity.Transaction) error {
	return s.transactionRepo.Create(ctx, transaction)
}

func (s *TransactionService) FindByOrderID(ctx context.Context, orderID string) (*entity.Transaction, error) {
	return s.transactionRepo.FindByOrderID(ctx, orderID)
}

func (s *TransactionService) FindByUserID(ctx context.Context, userID int64) ([]*entity.Transaction, error) {
	return s.transactionRepo.FindByUserID(ctx, userID)
}

func (s *TransactionService) UpdateStatus(ctx context.Context, orderID string, status string) error {
	return s.transactionRepo.UpdateStatus(ctx, orderID, status)
}


