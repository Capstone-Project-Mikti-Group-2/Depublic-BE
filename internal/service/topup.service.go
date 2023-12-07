package service

import (
	"context"
	"os"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/config"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/repository"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type TopupService interface {
	CreateTopup(ctx context.Context, topup entity.TopUp) (entity.TopUp, error)
	CreateMidtransCharge(orderID string, amount int64) (*coreapi.ChargeResponse, error)
	UpdateUserSaldo(ctx context.Context, userID int, amount int64) (int64, error)
	UserTopup(ctx context.Context, userID int, topup entity.TopUp) (entity.TopUp, error)
}

type topupService struct {
	config          *config.Config
	topupRepository repository.TopupRepository
}

type TopupRepository interface {
	UserTopup(ctx context.Context, topup entity.TopUp) (entity.TopUp, error)
	UpdateUserSaldo(ctx context.Context, userID int, amount int64) (int64, error)
}

type TopupUseCase interface {
	UserTopup(ctx context.Context, topup entity.TopUp) (entity.TopUp, error)
	UpdateUserSaldo(ctx context.Context, userID int, amount int64) (int64, error)
}

func NewTopupService(config *config.Config, topupRepository repository.TopupRepository) *topupService {
	return &topupService{
		config:          config,
		topupRepository: topupRepository,
	}
}

func (s *topupService) CreateTopup(ctx context.Context, topup entity.TopUp) (entity.TopUp, error) {
	return s.topupRepository.InputTopup(ctx, topup)
}

func (s *topupService) CreateMidtransCharge(order_id string, amount int64) (*coreapi.ChargeResponse, error) {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	clientKey := os.Getenv("MIDTRANS_CLIENT_KEY")

	c := &coreapi.Client{
		ServerKey: serverKey,
		ClientKey: clientKey, // Set the Client Key here
	}

	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  order_id,
			GrossAmt: amount,
		},
	}

	return c.ChargeTransaction(chargeReq)
}

func (s *topupService) UpdateUserSaldo(ctx context.Context, userID int, amount int64) (int64, error) {
	user, err := s.topupRepository.GetUserByID(ctx, int64(userID))
	if err != nil {
		return 0, err
	}

	user.Saldo += amount

	if err := s.topupRepository.UpdateUser(ctx, user); err != nil {
		return 0, err
	}

	return user.Saldo, nil
}
