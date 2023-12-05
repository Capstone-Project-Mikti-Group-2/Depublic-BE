package service

import (
	"context"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/repository"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type TopUpRepository interface {
	CreateTopup(ctx context.Context, topup entity.TopUp) (entity.TopUp, error)
	UserTopup(ctx context.Context, topup entity.TopUp, userID int64) (entity.TopUp, error)
	CreateMidtransTopup(ctx context.Context, OrderID string, nominal int64) (*coreapi.ChargeResponse, error)
	UpdateUserSaldo(ctx context.Context, nominal int64, userID int64) (int64, error)
	FindUserID(ctx context.Context, id int64) (*entity.User, error)
	UpdateUserTopup(ctx context.Context, user *entity.User) error
}

type TopupUseCase interface {
	CreateTopUp(ctx context.Context, topup entity.TopUp) (entity.TopUp, error)
	UserTopup(ctx context.Context, topup entity.TopUp, userID int64) (entity.TopUp, error)
	CreateMidtransTopup(ctx context.Context, OrderID string, nominal int64) (*coreapi.ChargeResponse, error)
	UpdateUserSaldo(ctx context.Context, nominal int64, userID int64) (int64, error)
	FindUserID(ctx context.Context, id int64) (*entity.User, error)
	UpdateUserTopup(ctx context.Context, user *entity.User) error
}

type TopupService struct {
	repository repository.TopUpRepository
}

func NewTopupService(repository repository.TopUpRepository) *TopupService {
	return &TopupService{
		repository: repository,
	}
}

func (s *TopupService) CreateTopUp(ctx context.Context, topup entity.TopUp) (entity.TopUp, error) {
	return s.repository.CreateTopup(ctx, topup)
}

func (s *TopupService) CreateMidtransTopup(ctx context.Context, OrderID string, nominal int64) (*coreapi.ChargeResponse, error) {
	c := coreapi.Client{}
	c.New("Server-Key", midtrans.Sandbox)

	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  OrderID,
			GrossAmt: nominal,
		},
	}

	return c.ChargeTransaction(chargeReq)
}

func (s *TopupService) UpdateUserSaldo(ctx context.Context, nominal int64, userID int64) (int64, error) {
	user, err := s.repository.FindUserID(ctx, userID)
	if err != nil {
		return 0, err
	}

	user.Saldo += nominal

	if err := s.repository.UpdateUserTopup(ctx, user); err != nil {
		return 0, err
	}

	return user.Saldo, nil
}

func (s *TopupService) UserTopup(ctx context.Context, topup entity.TopUp, userID int64) (entity.TopUp, error) {
	return s.repository.UserTopup(ctx, topup, userID)
}

func (s *TopupService) FindUserID(ctx context.Context, id int64) (*entity.User, error) {
	return s.repository.FindUserID(ctx, id)
}

func (s *TopupService) UpdateUserTopup(ctx context.Context, user *entity.User) error {
	return s.repository.UpdateUserTopup(ctx, user)
}
