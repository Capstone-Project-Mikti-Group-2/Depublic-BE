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
	TranscationHandler *handler.TransactionHandler, EventHandler *handler.EventHandler,
) []*Route {
	allRoutes := []*Route{}
	routeSlices := [][]*Route{
		{
			{ //Public Register and Login
				Method:  echo.POST,
				Path:    "/login",
				Handler: authHandler.Login,
			},
			{
				Method:  echo.POST,
				Path:    "/register",
				Handler: authHandler.Registration,
			},
			{ //Public Webhook
				Method:  echo.POST,
				Path:    "/transaction/webhook",
				Handler: TranscationHandler.WebHookTransaction,
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
	}
	for _, routes := range routeSlices {
		allRoutes = append(allRoutes, routes...)
	}

	return allRoutes
}

func PrivateRoutes(
	UserHandler *handler.UserHandler,
	ProfileHandler *handler.ProfileHandler,
	EventHandler *handler.EventHandler,
	TransactionHandler *handler.TransactionHandler, TicketHandler *handler.TicketHandler, TopUpHandler *handler.TopUpHandler, NotificationHandler *handler.NotificationHandler,
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
				Roles:   onlyAdmin,
			},
			{
				Method:  echo.PUT,
				Path:    "/users/profile",
				Handler: UserHandler.UpdateSelfUser,
				Roles:   allRoles,
			},
			{
				Method:  echo.DELETE,
				Path:    "/users",
				Handler: UserHandler.DeleteUser,
				Roles:   onlyAdmin,
			},
			{
				Method:  echo.DELETE,
				Path:    "/users/profile/:id",
				Handler: UserHandler.DeleteAccount,
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
			{
				Method:  echo.POST,
				Path:    "/users/logout",
				Handler: UserHandler.Logout,
				Roles:   allRoles,
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
			{
				Method:  echo.POST,
				Path:    "/transactions",
				Handler: TransactionHandler.CreateTransaction,
				Roles:   allRoles,
			},
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
			{ //topup Routes
				Method:  echo.POST,
				Path:    "/topup",
				Handler: TopUpHandler.CreateTopup,
				Roles:   allRoles,
			},
			{
				Method:  echo.POST,
				Path:    "/users/topup",
				Handler: TopUpHandler.UserTopup,
				Roles:   allRoles,
			},
			{
				Method:  echo.POST,
				Path:    "/users/input-saldo",
				Handler: UserHandler.InputSaldo,
				Roles:   allRoles,
			},
		},
		{
			{
				Method:  echo.POST,
				Path:    "/transaction",
				Handler: TransactionHandler.CreateTransaction,
				Roles:   allRoles,
			},
		},
		{
			{
				Method:  echo.POST,
				Path:    "/notification",
				Handler: NotificationHandler.CreateNotification,
				Roles:   allRoles,
			},
			{
				Method:  echo.GET,
				Path:    "/notifications",
				Handler: NotificationHandler.GetAllNotification,
				Roles:   allRoles,
			},
			{
				Method:  echo.GET,
				Path:    "/users/notifications",
				Handler: NotificationHandler.UserGetNotification,
				Roles:   allRoles,
			},
		},
	}

	for _, routes := range routeSlices {
		allRoutes = append(allRoutes, routes...)
	}

	return allRoutes
}
