package handler

import (
	"net/http"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/config"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TopUpHandler struct {
	cfg          *config.Config
	topupService service.TopupService
}

func NewTopUpHandler(cfg *config.Config, topupService service.TopupService) *TopUpHandler {
	return &TopUpHandler{
		cfg:          cfg,
		topupService: topupService,
	}
}

func (h *TopUpHandler) CreateTopup(c echo.Context) error {
	var topup entity.TopUp
	if err := c.Bind(&topup); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	ChargeResponse, err := h.topupService.CreateMidtransCharge(topup.ID, int64(topup.Amount))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	topup.SnapURL = ChargeResponse.RedirectURL

	newTopup, err := h.topupService.CreateTopup(c.Request().Context(), topup)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, newTopup)
}

func (h *TopUpHandler) UserTopup(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "jwt token not found",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "invalid token claims",
		})
	}

	userID := int(claims["user_id"].(float64))

	var topup entity.TopUp
	if err := c.Bind(&topup); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	userSaldo, err := h.topupService.UpdateUserSaldo(c.Request().Context(), userID, int64(topup.Amount))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	newTopup, err := h.topupService.UserTopup(c.Request().Context(), userID, topup)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"user_saldo": userSaldo,
		"topup":      newTopup,
	})
}
