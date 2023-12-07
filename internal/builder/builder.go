package builder

import (
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/config"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/http/handler"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/http/router"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/repository"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/service"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB, midtransClient snap.Client) []*router.Route {
	//registration, user, login, auth
	registrationRepository := repository.NewRegistrationRepository(db)
	registrationService := service.NewRegistrationService(registrationRepository)
	userRepository := repository.NewUserRepository(db)
	loginService := service.NewLoginService(userRepository)
	tokenService := service.NewTokenService(cfg)
	authHandler := handler.NewAuthHandler(registrationService, loginService, tokenService)

	//create payment
	paymentService := service.NewPaymentService(midtransClient)
	userService := service.NewUserService(userRepository)

	//Create transaction handler
	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService, paymentService, userService)

	return router.PublicRoutes(authHandler, transactionHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB, midtransClient snap.Client) []*router.Route {
	//Create user handler
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(cfg, userService)

	//Create profile handler
	profileRepository := repository.NewProfileRepository(db)
	profileService := service.NewProfileService(profileRepository)
	profileHandler := handler.NewProfileHandler(cfg, profileService)

	//create event handler
	eventRepository := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepository)
	eventHandler := handler.NewEventHandler(cfg, eventService)

	//create ticket handler
	ticketRepository := repository.NewTicketRepository(db)
	ticketService := service.NewTicketService(ticketRepository)
	ticketHandler := handler.NewTicketHandler(cfg, ticketService)

	//create payment
	paymentService := service.NewPaymentService(midtransClient)

	//Create transaction handler
	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService, paymentService, userService)

	//Create topup handler
	topupRepository := repository.NewTopupRepository(db)
	topupService := service.NewTopupService(cfg, topupRepository)
	topupHandler := handler.NewTopUpHandler(cfg, topupService)

	//Combine all routes
	return router.PrivateRoutes(userHandler, profileHandler, eventHandler, transactionHandler, ticketHandler, topupHandler)

}
