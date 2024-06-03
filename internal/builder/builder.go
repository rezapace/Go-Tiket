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

	userRepository := repository.NewUserRepository(db)
	loginService := service.NewLoginService(userRepository)
	tokenService := service.NewTokenService(cfg)

	// Create and initialize userService
	userService := service.NewUserService(userRepository)

	authHandler := handler.NewAuthHandler(registrationService, loginService, tokenService)

	// Update the line below with the additional TicketHandler argument
	return router.PublicRoutes(authHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB, midtransClient snap.Client) []*router.Route {
	// Create a user handler
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// Menggunakan PrivateRoutes dengan kedua handler
	return router.PrivateRoutes(userHandler)
}
