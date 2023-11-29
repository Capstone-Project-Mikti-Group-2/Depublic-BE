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

type EventHandler struct {
	cfg          *config.Config
	eventService service.EventUseCase
}

func NewEventHandler(
	cfg *config.Config,
	eventService service.EventUseCase) *EventHandler {
	return &EventHandler{cfg, eventService}
}

func (h *EventHandler) CreateEvent(ctx echo.Context) error {
	var input struct {
		Name        string    `json:"name" validate:"required"`
		Description string    `json:"description" validate:"required"`
		Location    string    `json:"location" validate:"required"`
		Price       int64     `json:"price" validate:"required"`
		Quantity    int64     `json:"quantity" validate:"required"`
		Image       []byte    `json:"image"`
		StartDate   time.Time `json:"start_date" validate:"required"`
		EndDate     time.Time `json:"end_date" validate:"required"`
		Available   bool      `json:"available" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	event := entity.NewEvent(0, input.Name, input.Description, input.Location, input.Price, input.Quantity, input.Available, input.Image, input.StartDate, input.EndDate)
	err := h.eventService.CreateEvent(ctx.Request().Context(), event)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":    "event created successfully",
		"created_at": event.CreatedAt,
	})
}

func (h *EventHandler) UpdateEvent(ctx echo.Context) error {
	var input struct {
		ID          int64     `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Location    string    `json:"location"`
		Price       int64     `json:"price"`
		Quantity    int64     `json:"quantity"`
		Image       []byte    `json:"image"`
		StartDate   time.Time `json:"start_date"`
		EndDate     time.Time `json:"end_date"`
		Available   bool      `json:"available"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	event := entity.UpdateEvent(input.ID, input.Name, input.Description, input.Location, input.Price, input.Quantity, input.Available, input.Image, input.StartDate, input.EndDate)
	err := h.eventService.UpdateEvent(ctx.Request().Context(), event)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":    "event updated successfully",
		"updated_at": event.UpdatedAt,
	})
}

func (h *EventHandler) FindEventByID(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id",
		})
	}
	event, err := h.eventService.FindEventByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"id":          event.ID,
			"name":        event.Name,
			"description": event.Description,
			"location":    event.Location,
			"price":       event.Price,
			"quantity":    event.Quantity,
			"image":       event.Image,
			"start_date":  event.StartDate,
			"end_date":    event.EndDate,
			"available":   event.Available,
			"created_at":  event.CreatedAt,
			"updated_at":  event.UpdatedAt,
		},
	})
}

func (h *EventHandler) FindAllEvent(ctx echo.Context) error {
	events, err := h.eventService.FindAllEvent(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": events,
	})
}

func (h *EventHandler) DeleteEvent(ctx echo.Context) error {
	var input struct {
		ID int64 `json:"id" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	err := h.eventService.DeleteEvent(ctx.Request().Context(), input.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "event deleted successfully",
		"id":      input.ID,
		"deleted": time.Now(),
	})
}
