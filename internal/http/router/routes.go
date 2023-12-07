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

func PublicRoutes(
	authHandler *handler.AuthHandler,
	transcationHandler *handler.TransactionHandler,
) []*Route {
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
		{
			Method:  echo.POST,
			Path:    "/transaction/webhook",
			Handler: transcationHandler.WebHookTransaction,
		},
	}
}

func PrivateRoutes(
	UserHandler *handler.UserHandler,
	ProfileHandler *handler.ProfileHandler,
	EventHandler *handler.EventHandler,
	TicketHandler *handler.TicketHandler,
	TransactionHandler *handler.TransactionHandler,
) []*Route {
	allRoutes := []*Route{}

	routeSlices := [][]*Route{
		{
			{ //users Routes
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
		},
		{
			{ //profile Routes
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
		},
		{
			{ //event Routes
				Method:  echo.POST,
				Path:    "/events",
				Handler: EventHandler.CreateEvent,
				Roles:   onlyAdmin,
			},
			{
				Method:  echo.PUT,
				Path:    "/events/:id",
				Handler: EventHandler.UpdateEvent,
				Roles:   allRoles,
			},
			{
				Method:  echo.DELETE,
				Path:    "/events/:id",
				Handler: EventHandler.DeleteEvent,
				Roles:   onlyAdmin,
			},
			{
				Method:  echo.GET,
				Path:    "/events",
				Handler: EventHandler.FindAllEvent,
				Roles:   allRoles,
			},
			{
				Method:  echo.GET,
				Path:    "/events/:id",
				Handler: EventHandler.FindEventByID,
				Roles:   allRoles,
			},
			{
				Method:  echo.GET,
				Path:    "/events/keyword/:keyword",
				Handler: EventHandler.SearchEvent,
				Roles:   allRoles,
			},
			{
				Method:  echo.GET,
				Path:    "/events/price/:min/:max",
				Handler: EventHandler.FilterEventByPrice,
				Roles:   allRoles,
			},
			{
				Method:  echo.GET,
				Path:    "/events/location/:location",
				Handler: EventHandler.FilterEventByLocation,
				Roles:   allRoles,
			},
			{
				Method:  echo.GET,
				Path:    "/events/available/:available",
				Handler: EventHandler.FilterEventByAvailable,
				Roles:   allRoles,
			},
			{
				Method:  echo.GET,
				Path:    "/events/date/:start_date/:end_date",
				Handler: EventHandler.FilterEventByDate,
				Roles:   allRoles,
			},
			{
				Method:  echo.GET,
				Path:    "/events/cheapest",
				Handler: EventHandler.SortEventByCheapest,
				Roles:   allRoles,
			},
			{
				Method:  echo.GET,
				Path:    "/events/expensive",
				Handler: EventHandler.SortEventByExpensive,
				Roles:   allRoles,
			},
			{
				Method:  echo.GET,
				Path:    "/events/newest",
				Handler: EventHandler.SortEventByNewest,
				Roles:   allRoles,
			},
		},
		{
	
			{ //ticket Routes
				Method:  echo.POST,
				Path:    "/tickets",
				Handler: TicketHandler.CreateTicket,
				Roles:   allRoles,
			},
			{
				Method:  echo.GET,
				Path:    "/tickets",
				Handler: TicketHandler.GetAllticket,
				Roles:   allRoles,
			},
			{
				Method:  echo.GET,
				Path:    "/tickets/:user_id",
				Handler: TicketHandler.GetTicketByUserID,
				Roles:   allRoles,
			},
		},
		{
			{//transaction Routes
				Method:  echo.POST,
				Path:    "/transaction",
				Handler: TransactionHandler.CreateTransaction,
				Roles:   allRoles,
			},
		},
	}

	for _, routes := range routeSlices {
		allRoutes = append(allRoutes, routes...)
	}

	return allRoutes
}
