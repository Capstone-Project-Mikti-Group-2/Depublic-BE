package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"github.com/google/uuid"
)

type TopUpRepository interface {
	FindByID(ctx context.Context, id string) (*entity.TopUp, error)
	Insert(ctx context.Context, t *entity.TopUp) error
	Update(ctx context.Context, t *entity.TopUp) error
}

type TopUpUseCase interface {
	ConfirmedTopUp(ctx context.Context, id string) error
	InitializeTopUp(ctx context.Context, req entity.TopUpRequest) (entity.TopUpResponse, error)
}

type topUpService struct {
	// notificationService NotificationService
	midtransService MidtransService
	topUpRepository TopUpRepository
	accountRepository AccountRepository
}

func NewTopUp(notificationService NotificationService, 
	midtransService MidtransService,
	topUpRepository TopUpRepository) TopUpUseCase {
	return &topUpService{
		// notificationService: notificationService,
		midtransService: midtransService,
		topUpRepository: topUpRepository,
	}
}

func (s *topUpService) InitializeTopUp(ctx context.Context, req entity.TopUpRequest) (entity.TopUpResponse, error) {
	topUp := entity.TopUp{
		ID: uuid.NewString(),
		UserID: req.UserID,
		Status:0,
		Amount: req.Amount,
	}
	err := s.midtransService.GenerateSnapURL(ctx, &topUp)
	if err != nil {
		return entity.TopUpResponse{}, err
	}

	err = s.topUpRepository.Insert(ctx, &topUp)
	if err != nil {
		return entity.TopUpResponse{}, err
	}

	return entity.TopUpResponse{
		SnapURL: topUp.SnapURL,
	}, nil
}

func (s *topUpService) ConfirmedTopUp(ctx context.Context, id string) error {
	topUp, err := s.topUpRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if topUp == (&entity.TopUp{}) {
		return errors.New("topup request not found")
	}

	account, err := s.accountRepository.FindByUserID(ctx, topUp.UserID)
	if err != nil {
		return err
	}

	// if account == (&entity.Account{}) {
	// 	return entity.ErrAccountNotFound
	// }

	account.Balance += topUp.Amount
	err = s.accountRepository.Update(ctx, account)
	if err != nil {
		return err
	}

	data := map[string]string{
		"amount": fmt.Sprintf("%.2f", topUp.Amount),
	}

	_ = s.notificationService.Insert(ctx, account.UserID, "TOPUP_SUCCESS", data)

	return err
}