package builder

import (
	"Rental/internal/config"
	"Rental/internal/http/handler"
	"Rental/internal/http/router"
	"Rental/internal/repository"
	"Rental/internal/service"

	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB, midtransClient snap.Client) []*router.Route {
	registrationRepository := repository.NewRegistrationRepository(db)
	registrationService := service.NewRegistrationService(registrationRepository)
	transactionRepository := repository.NewTransactionRepository(db)

	userRepository := repository.NewUserRepository(db)
	loginService := service.NewLoginService(userRepository)
	tokenService := service.NewTokenService(cfg)
	transactionService := service.NewTransactionService(transactionRepository)
	paymentService := service.NewPaymentService(midtransClient)

	// Create and initialize userService
	userService := service.NewUserService(userRepository)

	transactionHandler := handler.NewTransactionHandler(transactionService, paymentService, userService)

	authHandler := handler.NewAuthHandler(registrationService, loginService, tokenService)

	// Update the line below with the additional TicketHandler argument
	return router.PublicRoutes(authHandler,transactionHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB, midtransClient snap.Client) []*router.Route {
	// Create a user handler
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	transactionRepository := repository.NewTransactionRepository(db)
	paymentService := service.NewPaymentService(midtransClient)
	transactionService := service.NewTransactionService(transactionRepository)

	// Create and initialize transactionHandler with userService
	transactionHandler := handler.NewTransactionHandler(transactionService, paymentService, userService)

	TopupRepository := repository.NewTopupRepository(db)
	TopupService := service.NewTopupService(TopupRepository)

	// Create and initialize TopupHandler with TopupService
	TopupHandler := handler.NewTopupHandler(TopupService)

	// Menggunakan PrivateRoutes dengan kedua handler
	return router.PrivateRoutes(userHandler, transactionHandler, TopupHandler)
}
