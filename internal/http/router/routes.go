package router

import (
	"Ticketing/internal/http/handler"

	"github.com/labstack/echo/v4"
)

const (
	Admin = "Admin"
	Buyer = "Buyer"
)

var (
	allRoles  = []string{Admin, Buyer}
	onlyAdmin = []string{Admin}
	onlyBuyer = []string{Buyer}
)

// membuat struct route
type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
	Role    []string
}

// membuat fungsi untuk mengembalikan route
// pada func ini perlu login krna private
func PublicRoutes(
	authHandler *handler.AuthHandler,
	TicketHandler *handler.TicketHandler,
	BlogHandler *handler.BlogHandler,
	transactionHandler *handler.TransactionHandler) []*Route {
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
			Method:  echo.GET,
			Path:    "/public/blog",
			Handler: BlogHandler.GetAllBlogs,
		},
		{
			Method:  echo.GET,
			Path:    "/public/ticket",
			Handler: TicketHandler.GetAllTickets,
		},
		{
			Method:  echo.GET,
			Path:    "/blog",
			Handler: BlogHandler.GetAllBlogs,
		},
		{
			Method:  echo.GET,
			Path:    "/blog/:id",
			Handler: BlogHandler.GetBlog,
		},
		{
			Method:  echo.GET,
			Path:    "/blog/search/:search",
			Handler: BlogHandler.SearchBlog,
		},
		{
			Method:  echo.GET,
			Path:    "/ticket/:id",
			Handler: TicketHandler.GetTicket,
		},
		{
			Method:  echo.GET,
			Path:    "/ticket",
			Handler: TicketHandler.GetAllTickets,
		},
		//filter ticket by location
		{
			Method:  echo.GET,
			Path:    "/ticket/location/:location",
			Handler: TicketHandler.FilterTicket,
		},
		// filter ticket by category
		{
			Method:  echo.GET,
			Path:    "/ticket/category/:category",
			Handler: TicketHandler.FilterTicketByCategory,
		},
		// filter ticket by range time (start - end)
		{
			Method:  echo.GET,
			Path:    "/ticket/range/:start/:end",
			Handler: TicketHandler.FilterTicketByRangeTime,
		},
		// filter ticket by price (min - max)
		{
			Method:  echo.GET,
			Path:    "/ticket/price/:min/:max",
			Handler: TicketHandler.FilterTicketByPrice,
		},
		//sortir tiket dari yang terbaru
		{
			Method:  echo.GET,
			Path:    "/ticket/terbaru",
			Handler: TicketHandler.SortTicketByNewest,
		},
		//sortir tiket dari yang termahal
		{
			Method:  echo.GET,
			Path:    "/ticket/most-expensive",
			Handler: TicketHandler.SortTicketByMostExpensive,
		},
		//sortir tiket dari yang termurah
		{
			Method:  echo.GET,
			Path:    "/ticket/cheapest",
			Handler: TicketHandler.SortTicketByCheapest,
		},
		// filter ticket by most bought
		{
			Method:  echo.GET,
			Path:    "/ticket/most-bought",
			Handler: TicketHandler.SortTicketByMostBought,
		},
		// ticket yang masih tersedia
		{
			Method:  echo.GET,
			Path:    "/ticket/available",
			Handler: TicketHandler.SortTicketByAvailable,
		},
		{
			Method:  echo.GET,
			Path:    "/ticket/search/:search",
			Handler: TicketHandler.SearchTicket,
		},
		{
			Method:  echo.POST,
			Path:    "/transactions/webhook",
			Handler: transactionHandler.WebHookTransaction,
		},
		// {
		// 	Method:  echo.POST,
		// 	Path:    "/users/register/buyer",
		// 	Handler: authHandler.BuyerCreateAccount,
		// },
	}
}

// membuat fungsi untuk mengembalikan route
// pada func ini tdk perlu login krna public
func PrivateRoutes(
	UserHandler *handler.UserHandler,
	TicketHandler *handler.TicketHandler,
	BlogHandler *handler.BlogHandler,
	OrderHandler *handler.OrderHandler,
	NotificationHandler *handler.NotificationHandler,
	transactionHandler *handler.TransactionHandler,
	TopupHandler *handler.TopupHandler) []*Route {
	return []*Route{
		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: UserHandler.CreateUser,
			Role:    allRoles,
		},

		{
			Method:  echo.GET,
			Path:    "/users",
			Handler: UserHandler.GetAllUser,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.PUT,
			Path:    "/users/:id",
			Handler: UserHandler.UpdateUser,
			Role:    allRoles,
		},

		{
			Method:  echo.GET,
			Path:    "/users/:id",
			Handler: UserHandler.GetUserByID,
			Role:    allRoles,
		},

		{
			Method:  echo.DELETE,
			Path:    "/users/:id",
			Handler: UserHandler.DeleteUser,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.POST,
			Path:    "/ticket",
			Handler: TicketHandler.CreateTicket,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.GET,
			Path:    "/ticketa",
			Handler: TicketHandler.GetAllTickets,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.PUT,
			Path:    "/ticket/:id",
			Handler: TicketHandler.UpdateTicket,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.DELETE,
			Path:    "/ticket/:id",
			Handler: TicketHandler.DeleteTicket,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.POST,
			Path:    "/blog",
			Handler: BlogHandler.CreateBlog,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.PUT,
			Path:    "/blog/:id",
			Handler: BlogHandler.UpdateBlog,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.DELETE,
			Path:    "/blog/:id",
			Handler: BlogHandler.DeleteBlog,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.POST,
			Path:    "/order",
			Handler: OrderHandler.CreateOrder,
			Role:    allRoles,
		},

		{
			Method:  echo.GET,
			Path:    "/order",
			Handler: OrderHandler.GetAllOrders,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.GET,
			Path:    "/order/:id",
			Handler: OrderHandler.GetOrderByUserID,
			Role:    allRoles,
		},

		// create notification
		{
			Method:  echo.POST,
			Path:    "/notification",
			Handler: NotificationHandler.CreateNotification,
			Role:    onlyAdmin,
		},

		// get all notification
		{
			Method:  echo.GET,
			Path:    "/notifications",
			Handler: NotificationHandler.GetAllNotification,
			Role:    allRoles,
		},

		// topup
		{
			Method:  echo.POST,
			Path:    "/topup",
			Handler: TopupHandler.CreateTopup,
			Role:    allRoles,
		},

		// delete user self
		// {
		// 	Method:  echo.DELETE,
		// 	Path:    "/users/self/:id",
		// 	Handler: UserHandler.DeleteUserSelf,
		// 	Role:    allRoles,
		// },

		// getprofile
		{
			Method:  echo.GET,
			Path:    "/users/profile",
			Handler: UserHandler.GetProfile,
			Role:    allRoles,
		},

		// update profile
		{
			Method:  echo.PUT,
			Path:    "/users/profile",
			Handler: UserHandler.UpdateProfile,
			Role:    allRoles,
		},
		{
			Method:  echo.GET,
			Path:    "/users/balance",
			Handler: UserHandler.GetUserBalance,
			Role:    onlyBuyer,
		},
		{
			Method:  echo.DELETE,
			Path:    "/users/deleteprofile",
			Handler: UserHandler.DeleteAccount,
			Role:    allRoles,
		},
		//UserCreateOrder
		{
			Method:  echo.POST,
			Path:    "user/order",
			Handler: OrderHandler.UserCreateOrder,
			Role:    onlyBuyer,
		},
		//GetOrderHistory
		{
			Method:  echo.GET,
			Path:    "user/order",
			Handler: OrderHandler.GetOrderHistory,
			Role:    onlyBuyer,
		},
		//UserGetNotification
		{
			Method:  echo.GET,
			Path:    "user/notification",
			Handler: NotificationHandler.UserGetNotification,
			Role:    allRoles,
		},
		{
			Method:  echo.POST,
			Path:    "/user/topup",
			Handler: TopupHandler.UserTopup,
			Role:    onlyBuyer,
		},
		{
			Method:  echo.POST,
			Path:    "/user/upgrade",
			Handler: UserHandler.UpgradeSaldo,
			Role:    onlyBuyer,
		},
		{
			Method:  echo.POST,
			Path:    "/user/logout",
			Handler: UserHandler.UserLogout,
			Role:    allRoles,
		},
		{
			Method:  echo.POST,
			Path:    "/transactions",
			Handler: transactionHandler.CreateOrder,
			Role:    allRoles,
		},
		{
			Method:  echo.GET,
			Path:    "/transactions/history",
			Handler: transactionHandler.HistoryTransaction,
			Role:    allRoles,
		},
		// {
		// 	Method:  http.MethodPost,
		// 	Path:    "/diagnostic-ai",
		// 	Handler: UserHandler.DiagnosticAI,
		// 	Role:    allRoles,
		// },
	}
}

//NOTE :
//MENGAPA TERDAPAT 2 FUNC DIATAS? YAITU PUBLIC DAN PRIVATE
//KAREN DI SERVER.GO KITA BUAT GROUP API, DAN KITA MEMBAGI ROUTE YANG PERLU LOGIN DAN TIDAK PERLU LOGIN
// YAITU PUBLIC DAN PRIVATE

//note ;
//untuk menjalankan nya setelah port 8080 ditambahin /api/v1
// karna di server.go kita membuat group API
