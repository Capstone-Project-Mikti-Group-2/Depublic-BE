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
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
		Location    string `json:"location" validate:"required"`
		Price       int64  `json:"price" validate:"required"`
		Quantity    int64  `json:"quantity" validate:"required"`
		Image       []byte `json:"image"`
		StartDate   string `json:"start_date" validate:"validDate"`
		EndDate     string `json:"end_date" validate:"validDate"`
		Available   bool   `json:"available" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	// Convert string dates to time.Time
	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"start_date": "invalid date format"})
	}

	endDate, err := time.Parse("2006-01-02", input.EndDate)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"end_date": "invalid date format"})
	}

	event := entity.NewEvent(0, input.Name, input.Description, input.Location, input.Price, input.Quantity, input.Available, input.Image, startDate, endDate)
	err = h.eventService.CreateEvent(ctx.Request().Context(), event)
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
		ID          int64  `param:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Location    string `json:"location"`
		Price       int64  `json:"price"`
		Quantity    int64  `json:"quantity"`
		Image       []byte `json:"image"`
		StartDate   string `json:"start_date" validate:"validDate"`
		EndDate     string `json:"end_date" validate:"validDate"`
		Available   bool   `json:"available"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	// Convert string dates to time.Time
	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"start_date": "invalid date format"})
	}

	endDate, err := time.Parse("2006-01-02", input.EndDate)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"end_date": "invalid date format"})
	}

	event := entity.UpdateEvent(input.ID, input.Name, input.Description, input.Location, input.Price, input.Quantity, input.Available, input.Image, startDate, endDate)
	err = h.eventService.UpdateEvent(ctx.Request().Context(), event)
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
		ID int64 `param:"id" validate:"required"`
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

func (h *EventHandler) SearchEvent(ctx echo.Context) error {
	var input struct {
		Keyword string `param:"keyword" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	events, err := h.eventService.SearchEvent(ctx.Request().Context(), input.Keyword)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": events,
	})
}
func (h *EventHandler) FilterEventByPrice(ctx echo.Context) error {
	var input struct {
		Min string `param:"min"`
		Max string `param:"max"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	events, err := h.eventService.FilterEventByPrice(ctx.Request().Context(), input.Min, input.Max)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": events,
	})
}

// this handler still bugging idk why
func (h *EventHandler) FilterEventByDate(ctx echo.Context) error {
	var input struct {
		Start string `param:"start_date" validate:"validDate"`
		End   string `param:"end_date" validate:"validDate"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	startDate, err := time.Parse("2006-01-02", input.Start)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"start_date": "Invalid date format"})
	}

	endDate, err := time.Parse("2006-01-02", input.End)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"end_date": "Invalid date format"})
	}

	event, err := h.eventService.FilterEventByDate(ctx.Request().Context(), startDate, endDate)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": event,
	})
}

func (h *EventHandler) FilterEventByLocation(ctx echo.Context) error {
	var input struct {
		Location string `param:"location"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	event, err := h.eventService.FilterEventByLocation(ctx.Request().Context(), input.Location)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": event,
	})
}

func (h *EventHandler) FilterEventByAvailable(ctx echo.Context) error {
	var input struct {
		Available bool `param:"available"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	event, err := h.eventService.FilterEventByAvailable(ctx.Request().Context(), input.Available)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": event,
	})
}

func (h *EventHandler) SortEventByExpensive(ctx echo.Context) error {
	sort := ctx.QueryParam("sort")

	if sort != "termahal" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid sort order"})
	}

	event, err := h.eventService.SortEventByExpensive(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": event,
	})
}

func (h *EventHandler) SortEventByCheapest(ctx echo.Context) error {
	sort := ctx.QueryParam("sort")

	if sort != "termurah" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid sort order"})
	}

	event, err := h.eventService.SortEventByCheapest(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": event,
	})
}

func (h *EventHandler) SortEventByNewest(ctx echo.Context) error {
	sort := ctx.QueryParam("sort")

	if sort != "terbaru" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid sort order"})
	}

	event, err := h.eventService.SortEventByNewest(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": event,
	})
}
