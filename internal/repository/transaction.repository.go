package repository

import (
	"context"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"gorm.io/gorm"
)

type TransactionReposiotry struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionReposiotry {
	return &TransactionReposiotry{
		db: db,
	}
}

func (r *TransactionReposiotry) Create(ctx context.Context, transaction *entity.Transaction) error {
	if err := r.db.WithContext(ctx).Create(&transaction).Error; err != nil {
		return err
	}
	return nil
}

func (r *TransactionReposiotry) FindByOrderID(ctx context.Context, orderID string) (*entity.Transaction, error) {
	transaction := new(entity.Transaction)
	err := r.db.WithContext(ctx).Where("order_id = ?", orderID).First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *TransactionReposiotry) FindByUserID(ctx context.Context, userID int64) ([]*entity.Transaction, error) {
	transaction := make([]*entity.Transaction, 0)
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *TransactionReposiotry) UpdateStatus(ctx context.Context, orderID string, status string) error {
	err := r.db.WithContext(ctx).Model(&entity.Transaction{}).Where("order_id = ?", orderID).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}