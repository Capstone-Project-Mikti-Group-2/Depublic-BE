package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/common"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/entity"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/config"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/http/validator"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	cfg         *config.Config
	userService service.UserUseCase
}

func NewUserHandler(
	cfg *config.Config,
	userService service.UserUseCase) *UserHandler {
	return &UserHandler{cfg, userService}
}

func (h *UserHandler) CreateUser(ctx echo.Context) error {
	var input struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"email"`
		Number   string `json:"number" validate:"required"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"oneof=Administrator User"`
		Saldo    int64  `json:"saldo"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user := entity.NewUser(input.Name, input.Email, input.Number, input.Password, input.Role, input.Saldo)
	err := h.userService.CreateUser(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":    "user created successfully",
		"created_at": user.CreatedAt,
		"user":       user,
	})
}

func (h *UserHandler) UpdateUser(ctx echo.Context) error {
	var input struct {
		ID       int64  `param:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Number   string `json:"number"`
		Password string `json:"password"`
		Role     string `json:"role" validate:"oneof=Administrator User"`
		Saldo    int64  `json:"saldo"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user := entity.UpdateUser(input.ID, input.Name, input.Email, input.Number, input.Password, input.Role, input.Saldo)

	err := h.userService.UpdateUser(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success update user",
		"updated_at": user.UpdatedAt,
		"user":       user,
	})
}

func (h *UserHandler) DeleteUser(ctx echo.Context) error {
	var input struct {
		ID int64 `param:"id" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	err := h.userService.DeleteUser(ctx.Request().Context(), input.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
		"id":      input.ID,
		"deleted": time.Now(),
	})
}

func (h *UserHandler) FindAllUser(ctx echo.Context) error {
	users, err := h.userService.FindAllUser(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	var responseData []map[string]interface{}
	for _, user := range users {
		responseData = append(responseData, map[string]interface{}{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"number":     user.Number,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": responseData,
	})
}

func (h *UserHandler) FindUserByID(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id",
		})
	}

	user, err := h.userService.FindUserByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"number":     user.Number,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		},
	})
}

func (h *UserHandler) FindByEmail(ctx echo.Context) error {
	var input struct {
		Email string `param:"email" form:"email" query:"email" validate:"email"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid email",
		})
	}

	user, err := h.userService.FindByEmail(ctx.Request().Context(), input.Email)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"number":     user.Number,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		},
	})
}

func (h *UserHandler) FindUserByUsername(ctx echo.Context) error {
	var input struct {
		Username string `param:"username" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user, err := h.userService.FindUserByUsername(ctx.Request().Context(), input.Username)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"number":     user.Number,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		},
	})
}

func (h *UserHandler) FindUserByNumber(ctx echo.Context) error {
	var input struct {
		Number string `param:"number" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user, err := h.userService.FindUserByNumber(ctx.Request().Context(), input.Number)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, validator.ValidatorErrors(err))
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"number":     user.Number,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		},
	})
}

func (h *UserHandler) DeleteAccount(ctx echo.Context) error {
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unable to get user claims",
		})
	}

	claimsData, ok := claims.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unable to get user informations",
		})
	}

	idToDelete := claimsData.ID

	err := h.userService.DeleteUser(ctx.Request().Context(), idToDelete)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete account",
		"id":      idToDelete,
	})
}

func (h *UserHandler) UpdateSelfUser(ctx echo.Context) error {
	var input struct {
		Name     string `json:"name"`
		Number   string `json:"number"`
		Email    string `json:"email" validate:"email"`
		Password string `json:"password"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unable to get user claims",
		})
	}

	claimsData, ok := claims.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unable to get user informations",
		})
	}

	user := entity.UpdateSelfUser(claimsData.ID, input.Name, input.Email, input.Number, input.Password)

	err := h.userService.UpdateSelfUser(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"data":    user,
	})
}

func (h *UserHandler) Logout(ctx echo.Context) error {
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unable to get user claims",
		})
	}

	claimsData, ok := claims.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unable to get user informations",
		})
	}

	userID := claimsData.ID

	user := &entity.User{
		ID: userID,
	}

	err := h.userService.Logout(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"message": "unable to logout",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success logout",
	})
}

func (h *UserHandler) InputSaldo(ctx echo.Context) error {
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unable to get user claims",
		})
	}

	claimsData, ok := claims.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unable to get user informations",
		})
	}

	userID := claimsData.ID

	user, err := h.userService.FindByID(ctx.Request().Context(), userID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"message": err.Error(),
		})
	}

	var input struct {
		Saldo int64 `json:"saldo"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	updateSaldo := user.Saldo + input.Saldo

	user.Saldo = updateSaldo
	err = h.userService.InputSaldo(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success input saldo",
		"data":    user,
	})
}
