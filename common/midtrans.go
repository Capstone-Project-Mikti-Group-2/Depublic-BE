package common

import (
	"context"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
)

type MidtransService interface {
	GenerateSnapURL(ctx context.Context, topup *entity.TopUp) error
	VerifyPayment(ctx context.Context, data map[string]interface{}) (bool,error)
}