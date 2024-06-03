package handler

import (
	"Ticketing/entity"
	"Ticketing/internal/http/validator"
	"Ticketing/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	registrationService service.RegistrationUseCase // untuk regist
	loginService        service.LoginUseCase        //untuk memanggil service yang ngelakuin pengecekan user.
	tokenService        service.TokenUsecase        //untuk memanggil func akses token
	// buyercreateaccountService service.BuyerCreateAccountUseCase
}

// ini func untuk type AuthHandler
func NewAuthHandler(
	registartionService service.RegistrationUseCase,
	loginService service.LoginUseCase,
	tokenService service.TokenUsecase,
	// buyercreateaccountService service.BuyerCreateAccountUseCase,
) *AuthHandler {
	return &AuthHandler{
		registrationService: registartionService,
		loginService:        loginService,
		tokenService:        tokenService,
		// buyercreateaccountService: buyercreateaccountService,
	}
}

// func ini untuk login
func (h *AuthHandler) Login(ctx echo.Context) error {
	//pengecekan request
	var input struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	if err := ctx.Bind(&input); err != nil { // di cek pake validate buat masukin input
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	//untuk manggil login service di folder service
	user, err := h.loginService.Login(ctx.Request().Context(), input.Email, input.Password)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	//untuk manggil token service di folder service
	accessToken, err := h.tokenService.GenerateAccessToken(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	data := map[string]interface{}{
		"access_token": accessToken,
	}
	return ctx.JSON(http.StatusOK, data)
}

// Public Register
func (h *AuthHandler) Registration(ctx echo.Context) error {
	//pengecekan request
	var input struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
		Roles    string `json:"roles" default:"Buyer"`
		Number   string `json:"number" validate:"required,min=11,max=13"`
	}

	if err := ctx.Bind(&input); err != nil { // di cek pake validate buat masukin input
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	//untuk manggil registration service di folder service
	user := entity.Register(input.Email, input.Password, input.Roles, input.Number)
	err := h.registrationService.Registration(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	accessToken, err := h.tokenService.GenerateAccessToken(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":      "User registration successfully",
		"access_token": accessToken,
	})

}

// BuyerCreateAccount
// func (h *AuthHandler) BuyerCreateAccount(ctx echo.Context) error {
// 	// Pengecekan request
// 	var input struct {
// 		Name     string `json:"name" validate:"required"`
// 		Email    string `json:"email" validate:"required,email"`
// 		Number   string `json:"number" validate:"required,min=11,max=13"`
// 		Password string `json:"password" validate:"required,min=8"`
// 	}

// 	if err := ctx.Bind(&input); err != nil { // Di cek pake validate buat masukin input
// 		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
// 	}

// 	// Untuk manggil BuyerCreateAccount service di folder service
// 	user := entity.Register(input.Email, input.Password, "Buyer", input.Number)
// 	err := h.buyercreateaccountService.BuyerCreateAccount(ctx.Request().Context(), user)
// 	if err != nil {
// 		return ctx.JSON(http.StatusUnprocessableEntity, err)
// 	}

// 	accessToken, err := h.tokenService.GenerateAccessToken(ctx.Request().Context(), user)
// 	if err != nil {
// 		return ctx.JSON(http.StatusUnprocessableEntity, err)
// 	}

// 	return ctx.JSON(http.StatusOK, map[string]interface{}{
// 		"message":      "Buyer account created successfully",
// 		"access_token": accessToken,
// 	})
// }


