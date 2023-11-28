package router

import (
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/http/handler"
	"github.com/labstack/echo/v4"
)

const (
	Administrator = "Administrator"
	User          = "User"
)

var (
	allRoles  = []string{Administrator, User}
	onlyAdmin = []string{Administrator}
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
	Roles   []string
}

func PublicRoutes(authHandler *handler.AuthHandler) []*Route {
	return []*Route{
		{
			Method:  echo.POST,
			Path:    "/login",
			Handler: authHandler.Login,
		},
		{
			Method:  echo.POST,
			Path:    "/register",
			Handler: authHandler.Registration,
		},
	}
}

func PrivateRoutes(UserHandler *handler.UserHandler, ProfileHandler *handler.ProfileHandler) []*Route {
	userRoutes := []*Route{
		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: UserHandler.CreateUser,
			Roles:   onlyAdmin,
		},
		{
			Method:  echo.PUT,
			Path:    "/users/:id",
			Handler: UserHandler.UpdateUser,
			Roles:   allRoles,
		},
		{
			Method:  echo.DELETE,
			Path:    "/users/:id",
			Handler: UserHandler.DeleteUser,
			Roles:   onlyAdmin,
		},
		{
			Method:  echo.GET,
			Path:    "/users",
			Handler: UserHandler.FindAllUser,
			Roles:   onlyAdmin,
		},
		{
			Method:  echo.GET,
			Path:    "/users/:id",
			Handler: UserHandler.FindUserByID,
			Roles:   onlyAdmin,
		},
		{
			Method:  echo.GET,
			Path:    "/users/email/:email",
			Handler: UserHandler.FindByEmail,
			Roles:   onlyAdmin,
		},
		{
			Method:  echo.GET,
			Path:    "/users/number/:number",
			Handler: UserHandler.FindUserByNumber,
			Roles:   onlyAdmin,
		},
		{
			Method:  echo.GET,
			Path:    "/users/name/:name",
			Handler: UserHandler.FindUserByUsername,
			Roles:   onlyAdmin,
		},
	}

	profileRoutes := []*Route{
		{
			Method:  echo.GET,
			Path:    "/profile/:id",
			Handler: ProfileHandler.GetProfileByID,
			Roles:   allRoles,
		},
		{
			Method:  echo.PUT,
			Path:    "/profile/:id",
			Handler: ProfileHandler.UpdateProfile,
			Roles:   allRoles,
		},
		{
			Method:  echo.POST,
			Path:    "/profile",
			Handler: ProfileHandler.CreateProfile,
			Roles:   allRoles,
		},
		{
			Method:  echo.DELETE,
			Path:    "/profile/:id",
			Handler: ProfileHandler.DeleteProfile,
			Roles:   allRoles,
		},
	}
	allRoutes := append(userRoutes, profileRoutes...)

	return allRoutes
}
