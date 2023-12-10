package handler

import (
	"net/http"
	"time"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/http/validator"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/service"
	"github.com/labstack/echo/v4"
)

type NotificationHandler struct {
	NotificationService service.NotificationUseCase
}

func NewNotificationHandler(notificationService service.NotificationUseCase) *NotificationHandler {
	return &NotificationHandler{
		NotificationService: notificationService,
	}
}

func (h *NotificationHandler) CreateNotification(ctx echo.Context) error {
	var input struct {
		Type    string `json:"type" validate:"required"`
		Content string `json:"content" validate:"required"`
		IsRead  bool   `json:"is_read"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	Notification := entity.Notification{
		Type:      input.Type,
		Content:   input.Content,
		IsRead:    input.IsRead,
		CreatedAt: time.Now(),
	}

	err := h.NotificationService.CreateNotification(ctx.Request().Context(), &Notification)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, validator.ValidatorErrors(err))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create notification",
		"data":    Notification,
	})

}

func (h *NotificationHandler) GetAllNotification(c echo.Context) error {
	Notification, err := h.NotificationService.GetAllNotification(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": Notification,
	})
}

func (h *NotificationHandler) UserGetNotification(c echo.Context) error {
	Notification, err := h.NotificationService.UserGetNotification(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": Notification,
	})
}
