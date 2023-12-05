package handler

import (
	"net/http"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/service"
	"github.com/labstack/echo/v4"
)

type midtransHandler struct {
	midtransService service.MidtransService
}

func NewMidtrans(){

}

func (h *midtransHandler) paymentHandlerNotification(ctx echo.Context) error {
	var notificationPayload map[string]interface{}
	if err := ctx.Bind(&notificationPayload); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.midtransService.VerifyPayment(ctx.Request().Context(), notificationPayload)
	return ctx.SendStatus(util.GetHttpStatus(err))
}