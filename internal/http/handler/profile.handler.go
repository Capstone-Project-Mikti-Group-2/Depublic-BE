package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/config"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/http/validator"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/service"
	"github.com/labstack/echo/v4"
)

type ProfileHandler struct {
	cfg            *config.Config
	profileService service.ProfileUseCase
}

func NewProfileHandler(
	cfg *config.Config,
	profileService service.ProfileUseCase) *ProfileHandler {
	return &ProfileHandler{cfg, profileService}
}

func (h *ProfileHandler) CreateProfile(ctx echo.Context) error {
	var input struct {
		UserID  int64  `json:"user_id" validate:"required"`
		Address string `json:"address" validate:"required"`
		Avatar  []byte `json:"avatar"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	profile := entity.NewProfile(input.UserID, input.Address, input.Avatar)
	err := h.profileService.CreateProfile(ctx.Request().Context(), profile)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":    "profile created successfully",
		"created_at": profile.CreatedAt,
	})

}

func (h *ProfileHandler) UpdateProfile(ctx echo.Context) error {
	var input struct {
		ID      int64  `json:"id"`
		Address string `json:"address"`
		Avatar  []byte `json:"avatar"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	profile := entity.UpdateProfile(input.ID, input.Address, input.Avatar)
	err := h.profileService.UpdateProfile(ctx.Request().Context(), profile)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":    "profile updated successfully",
		"updated_at": profile.UpdatedAt,
	})
}

func (h *ProfileHandler) GetProfileByID(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id",
		})
	}
	profile, err := h.profileService.GetProfileByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"id":         profile.ID,
			"user_id":    profile.UserID,
			"address":    profile.Address,
			"avatar":     profile.Avatar,
			"created_at": profile.CreatedAt,
			"updated_at": profile.UpdatedAt,
		},
	})
}

func (h *ProfileHandler) DeleteProfile(ctx echo.Context) error {
	var input struct {
		ID int64 `json:"id" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	err := h.profileService.DeleteProfile(ctx.Request().Context(), input.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete profile",
		"id":      input.ID,
		"deleted": time.Now(),
	})
}
