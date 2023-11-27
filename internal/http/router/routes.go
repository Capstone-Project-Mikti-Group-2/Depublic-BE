package router

import (
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/http/handler"
	"github.com/labstack/echo/v4"
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

func PublicRoutes(authHandler *handler.AuthHandler) []*Route {
	return []*Route{
		{
			Method:  "POST",
			Path:    "/login",
			Handler: authHandler.Login,
		},
		{
			Method:  "POST",
			Path:    "/register",
			Handler: authHandler.Registration,
		},
	}
}

func PrivateRoutes(UserHandler *handler.UserHandler) []*Route {
	return []*Route{
		{
			Method:  "POST",
			Path:    "/users",
			Handler: UserHandler.CreateUser,
		},
		{
			Method:  "PUT",
			Path:    "/users",
			Handler: UserHandler.UpdateUser,
		},
		{
			Method:  "DELETE",
			Path:    "/users/:id",
			Handler: UserHandler.DeleteUser,
		},
		{
			Method:  "GET",
			Path:    "/users",
			Handler: UserHandler.FindAllUser,
		},
		{
			Method:  "GET",
			Path:    "/users/:id",
			Handler: UserHandler.FindUserByID,
		},
		{
			Method:  "GET",
			Path:    "/users/email/:email",
			Handler: UserHandler.FindUserByEmail,
		},
		{
			Method:  "GET",
			Path:    "/users/number/:number",
			Handler: UserHandler.FindUserByNumber,
		},
		{
			Method:  "GET",
			Path:    "/users/name/:name",
			Handler: UserHandler.FindUserByUsername,
		},
	}
}
