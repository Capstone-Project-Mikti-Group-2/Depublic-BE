package handler

import (
	"net/http"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/config"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TopupHandler struct {
	cfg          *config.Config
	topupService service.TopupUseCase
}

func NewTopupHandler(cfg *config.Config, topupService service.TopupUseCase) *TopupHandler {
	return &TopupHandler{
		cfg:          cfg,
		topupService: topupService,
	}
}

func (h *TopupHandler) CreateTopup(c echo.Context) error {
	var topup entity.TopUp
	if err := c.Bind(&topup); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	ChargeResponse, err := h.topupService.CreateMidtransTopup(c.Request().Context(), topup.ID, int64(topup.Nominal)) // nominal diubah ke int64
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	topup.SnapURL = ChargeResponse.RedirectURL

	newtopup, err := h.topupService.CreateTopUp(c.Request().Context(), topup)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, newtopup)
}

func (h *TopupHandler) UserTopup(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
	}

	claims := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "invalid token",
		})
	}

	userID := int64(claims["user_id"].(float64))

	var topup entity.TopUp
	if err := c.Bind(&topup); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	userSaldo, err := h.topupService.UpdateUserSaldo(c.Request().Context(), topup.Nominal, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	newTopup, err := h.topupService.UserTopup(c.Request().Context(), topup, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "topup success",
		"topup":   newTopup,
		"user":    userSaldo,
	})
}
