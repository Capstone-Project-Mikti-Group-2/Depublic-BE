package handler

import (
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/service"
	"github.com/labstack/echo/v4"
)

type midtransHandler struct {
	midtransService service.MidtransService
}

func NewMidtrans(){

}

func (h *midtransHandler) paymentHandlerNotification(ctx echo.Context) error {
	
}