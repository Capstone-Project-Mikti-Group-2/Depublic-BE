package handler

import (
	"net/http"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/config"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/http/validator"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/service"
	"github.com/labstack/echo/v4"
)

type topUpHandler struct {
	cfg				config.Config
	topUpService 	service.TopUpUseCase
}

func NewTopUp(
	cfg *config.Config,
	topUpService service.TopUpUseCase) *topUpHandler {
	return &topUpHandler{cfg, topUpService}
}

func (h *topUpHandler) InitializeTopUp(ctx echo.Context) error {
	var input entity.TopUpRequest
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	return nil
}