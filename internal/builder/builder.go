package builder

import (
	"Ticketing/internal/config"
	"Ticketing/internal/http/handler"
	"Ticketing/internal/http/router"
	"Ticketing/internal/repository"
	"Ticketing/internal/service"

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

	BlogRepository := repository.NewBlogRepository(db)
	BlogService := service.NewBlogService(BlogRepository)
	BlogHandler := handler.NewBlogHandler(BlogService)

	ticketRepository := repository.NewTicketRepository(db)
	ticketService := service.NewTicketService(ticketRepository)
	ticketHandler := handler.NewTicketHandler(ticketService)

	authHandler := handler.NewAuthHandler(registrationService, loginService, tokenService)

	// Update the line below with the additional TicketHandler argument
	return router.PublicRoutes(authHandler, ticketHandler, BlogHandler, transactionHandler)
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

	// Create a ticket handler
	ticketRepository := repository.NewTicketRepository(db)
	ticketService := service.NewTicketService(ticketRepository)
	ticketHandler := handler.NewTicketHandler(ticketService)

	// Create a Blog handler
	BlogRepository := repository.NewBlogRepository(db)
	BlogService := service.NewBlogService(BlogRepository)
	BlogHandler := handler.NewBlogHandler(BlogService)

	// Create an order handler
	OrderRepository := repository.NewOrderRepository(db)
	OrderService := service.NewOrderService(OrderRepository)
	OrderHandler := handler.NewOrderHandler(OrderService)

	NotificationRepository := repository.NewNotificationRepository(db)
	NotificationService := service.NewNotificationService(NotificationRepository)
	NotificationHandler := handler.NewNotificationHandler(NotificationService)

	TopupRepository := repository.NewTopupRepository(db)
	TopupService := service.NewTopupService(TopupRepository)

	// Create and initialize TopupHandler with TopupService
	TopupHandler := handler.NewTopupHandler(TopupService)

	// Menggunakan PrivateRoutes dengan kedua handler
	return router.PrivateRoutes(userHandler, ticketHandler, BlogHandler, OrderHandler, NotificationHandler, transactionHandler, TopupHandler)
}
