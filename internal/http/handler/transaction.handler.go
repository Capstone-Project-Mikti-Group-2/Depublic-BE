package handler

import (
	"net/http"
	"strconv"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/common"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/http/validator"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionService service.TransactionUseCase
	paymentService     service.PaymentUseCase
}

func NewTransactionHandler(transactionService service.TransactionUseCase, paymentService service.PaymentUseCase) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
		paymentService:     paymentService,
	}
}

func (h *TransactionHandler) CreateTransaction(ctx echo.Context) error {
	var input struct {
		OrderID string `json:"order_id" validate:"required"`
		Amount  int64  `json:"amount" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*common.JwtCustomClaims)

	transaction := entity.NewTransaction(input.OrderID, claims.ID, input.Amount, "unpaid")

	err := h.transactionService.Create(ctx.Request().Context(), transaction)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	paymentRequest := entity.NewPaymentRequest(transaction.OrderID, transaction.Amount, claims.Name, "", claims.Email)

	payment, err := h.paymentService.CreateTransaction(ctx.Request().Context(), paymentRequest)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"url_pembayaran": payment})
}

func (h *TransactionHandler) WebHookTransaction(ctx echo.Context) error {
	var input entity.MidtransRequest

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	transaction, err := h.transactionService.FindByOrderID(ctx.Request().Context(), input.OrderID)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	status := "unpaid"

	if input.TransactionStatus == "settlement" {
		status = "paid"
	}

	err = h.transactionService.UpdateStatus(ctx.Request().Context(), transaction.OrderID, status)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "success"})
}

func (h *TransactionHandler) GetTransactionByUserID(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id",
		})
	}

	users, err := h.transactionService.FindByUserID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	var responseData []map[string]interface{}
	for _, user := range users {
		responseData = append(responseData, map[string]interface{}{
			"id":         user.ID,
			"user_id":    user.UserID,
			"order_id":   user.OrderID,
			"amount":     user.Amount,
			"status":     user.Status,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": responseData,
	})
}
