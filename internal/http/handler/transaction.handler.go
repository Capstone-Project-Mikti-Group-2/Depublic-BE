package handler

import (
	"net/http"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/common"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/http/validator"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionService service.TransactionUseCase
	paymentService     service.PaymentUseCase
	userService        service.UserUseCase
}

func NewTransactionHandler(transactionService service.TransactionUseCase, paymentService service.PaymentUseCase, userService service.UserUseCase) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
		paymentService:     paymentService,
		userService:        userService,
	}
}

func (h *TransactionHandler) CreateTransaction(ctx echo.Context) error {
	var input struct {
		Amount int64 `json:"amount" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*common.JwtCustomClaims)

	uniqueID := uuid.New().String()
	OrderID := "topup-" + uniqueID

	transaction := entity.NewTransaction(OrderID, claims.ID, input.Amount, "unpaid")

	err := h.transactionService.Create(ctx.Request().Context(), transaction)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.userService.UpdateSaldo(ctx.Request().Context(), claims.ID, input.Amount)

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

func (h *TransactionHandler) GetTransactionHistoryByUserID(ctx echo.Context) error {
	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*common.JwtCustomClaims)

	users, err := h.transactionService.FindByUserID(ctx.Request().Context(), claims.ID)
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
