package service

import (
	"context"

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
	c := coreapi.Client{}
	c.New("SB-Mid-server-MekbYm9j66D6YG5Wv67xK75R", midtrans.Sandbox)

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
	user, err := s.topupRepository.GetUserByID(ctx, int(userID))
	if err != nil {
		return 0, err
	}

	user.Saldo += amount

	if err := s.topupRepository.UpdateUser(ctx, user); err != nil {
		return 0, err
	}

	return user.Saldo, nil
}

func (s *topupService) UserTopup(ctx context.Context, userID int, topup entity.TopUp) (entity.TopUp, error) {
	return s.topupRepository.UserTopup(ctx, topup)
}
