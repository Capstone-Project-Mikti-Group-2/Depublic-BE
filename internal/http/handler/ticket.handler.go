package handler

import (
	"net/http"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/config"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/http/validator"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/service"
	"github.com/labstack/echo/v4"
)

type TicketHandler struct {
	cfg           *config.Config
	ticketService service.TicketUseCase
}

func NewTicketHandler(cfg *config.Config, ticketService service.TicketUseCase) *TicketHandler {
	return &TicketHandler{cfg, ticketService}
}

func (h *TicketHandler) CreateTicket(ctx echo.Context) error {
	var input struct {
		EventID  int64 `json:"event_id" validate:"required"`
		UserID   int64 `json:"user_id" validate:"required"`
		Quantity int64 `json:"quantity" validate:"required"`
		IsPaid   bool  `json:"is_paid" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	userSaldo, err := h.ticketService.GetUserSaldo(ctx.Request().Context(), input.UserID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	eventPrice, err := h.ticketService.GetBookingPrice(ctx.Request().Context(), input.EventID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	if userSaldo < eventPrice*input.Quantity {
		return ctx.JSON(http.StatusUnprocessableEntity, "insufficient balance")
	}

	ticket := entity.NewTicket(input.EventID, input.UserID, input.Quantity, input.IsPaid)
	err = h.ticketService.CreateTicket(ctx.Request().Context(), ticket)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	err = h.ticketService.UpdateUserSaldo(ctx.Request().Context(), input.UserID, input.Quantity*eventPrice)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "ticket created successfully",
	})

}

func (h *TicketHandler) GetAllticket(ctx echo.Context) error {
	bookings, err := h.ticketService.GetBooking(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	var bookingDetails []map[string]interface{}
	for _, booking := range bookings {
		ticket, err := h.ticketService.FindTicketByID(ctx.Request().Context(), booking.EventID)
		if err != nil {
			return ctx.JSON(http.StatusUnprocessableEntity, err)
		}

		bookingDetail := map[string]interface{}{
			"user_id": booking.UserID,
			"ticket":  ticket,
		}
		bookingDetails = append(bookingDetails, bookingDetail)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": bookingDetails,
	})
}

func (h *TicketHandler) GetTicketByUserID(ctx echo.Context) error {
	bookings, err := h.ticketService.GetBooking(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	var bookingDetails []map[string]interface{}
	for _, booking := range bookings {
		ticket, err := h.ticketService.FindTicketByID(ctx.Request().Context(), booking.EventID)
		if err != nil {
			return ctx.JSON(http.StatusUnprocessableEntity, err)
		}

		bookingDetail := map[string]interface{}{
			"user_id": booking.UserID,
			"ticket":  ticket,
		}
		bookingDetails = append(bookingDetails, bookingDetail)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": bookingDetails,
	})
}
