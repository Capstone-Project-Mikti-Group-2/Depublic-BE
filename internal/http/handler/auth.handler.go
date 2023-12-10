package handler

import (
	"net/http"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/http/validator"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/service"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	registrationService service.RegistrationUseCase
	loginService        service.LoginUseCase
	tokenService        service.TokenUseCase
}

func NewAuthHandler(registrationService service.RegistrationUseCase, loginService service.LoginUseCase, tokenService service.TokenUseCase) *AuthHandler {
	return &AuthHandler{
		registrationService: registrationService,
		loginService:        loginService,
		tokenService:        tokenService,
	}
}

func (h *AuthHandler) Login(ctx echo.Context) error {
	var input struct {
		Email    string `json:"email" validate:"email"`
		Password string `json:"password" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user, err := h.loginService.Login(ctx.Request().Context(), input.Email, input.Password)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	accessToken, err := h.tokenService.GenerateToken(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	data := map[string]interface{}{
		"token": accessToken,
		"user":  user,
	}
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, data)
}

func (h *AuthHandler) Registration(ctx echo.Context) error {
	var input struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"email"`
		Password string `json:"password" validate:"required"`
		Number   string `json:"number" validate:"required"`
		Role     string `json:"role" validate:"oneof=Administrator User"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user := entity.Register(input.Name, input.Email, input.Password, input.Number, input.Role)
	err := h.registrationService.Registration(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	accessToken, err := h.tokenService.GenerateToken(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	data := map[string]interface{}{
		"nama":   input.Name,
		"email":  input.Email,
		"number": input.Number,
		"token":  accessToken,
	}
	return ctx.JSON(http.StatusOK, data)
}

//completed auth
