package router

import (
	"Rental/internal/http/handler"

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
	}
}

// membuat fungsi untuk mengembalikan route
// pada func ini tdk perlu login krna public
func PrivateRoutes(
	UserHandler *handler.UserHandler) []*Route {
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
			Method:  echo.DELETE,
			Path:    "/users/deleteprofile",
			Handler: UserHandler.DeleteAccount,
			Role:    allRoles,
		},

		{
			Method:  echo.POST,
			Path:    "/user/logout",
			Handler: UserHandler.UserLogout,
			Role:    allRoles,
		},

	}
}

//NOTE :
//MENGAPA TERDAPAT 2 FUNC DIATAS? YAITU PUBLIC DAN PRIVATE
//KAREN DI SERVER.GO KITA BUAT GROUP API, DAN KITA MEMBAGI ROUTE YANG PERLU LOGIN DAN TIDAK PERLU LOGIN
// YAITU PUBLIC DAN PRIVATE

//note ;
//untuk menjalankan nya setelah port 8080 ditambahin /api/v1
// karna di server.go kita membuat group API
