package handler

import (
	"Ticketing/common"
	"Ticketing/entity"
	"Ticketing/internal/http/validator"
	"Ticketing/internal/service"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	OrderService service.OrderUsecase
}

func NewOrderHandler(OrderService service.OrderUsecase) *OrderHandler {
	return &OrderHandler{OrderService}
}

// func untuk create order
func (h *OrderHandler) CreateOrder(ctx echo.Context) error {
	var input struct {
		TicketID int64  `json:"ticket_id" validate:"required"`
		Quantity int64  `json:"quantity" validate:"required"`
		UserID   int64  `json:"user_id" validate:"required"`
		Status   string `json:"status" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	// Mendapatkan informasi saldo pengguna sebelum membuat pesanan
	userBalance, err := h.OrderService.GetUserBalance(ctx.Request().Context(), input.UserID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	// Mengambil informasi tiket dari TicketService untuk mendapatkan harga tiket
	ticketPrice, err := h.OrderService.GetTicketPrice(ctx.Request().Context(), input.TicketID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	// Memeriksa apakah saldo cukup untuk membuat pesanan
	if userBalance < (input.Quantity * ticketPrice) {
		return ctx.JSON(http.StatusUnprocessableEntity, errors.New("insufficient balance"))
	}

	order := entity.NewOrder(input.TicketID, input.Quantity, input.UserID, input.Status)
	err = h.OrderService.CreateOrder(ctx.Request().Context(), order)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	// Mengurangkan saldo pengguna setelah membuat pesanan
	err = h.OrderService.UpdateUserBalance(ctx.Request().Context(), input.UserID, input.Quantity*ticketPrice)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusCreated, "Order created successfully")
}

// Get All Order
func (h *OrderHandler) GetAllOrders(ctx echo.Context) error {
	// Implementasi untuk mendapatkan semua pesanan dari usecase
	orders, err := h.OrderService.GetOrders(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	var orderDetails []map[string]interface{}
	for _, order := range orders {
		ticket, err := h.OrderService.GetTicketByID(ctx.Request().Context(), order.TicketID)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.NewHTTPError(http.StatusInternalServerError, err.Error()))
		}

		orderDetail := map[string]interface{}{
			"user_id": order.UserID,
			"ticket":  ticket,
		}
		orderDetails = append(orderDetails, orderDetail)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":       "Get all orders success",
		"order_details": orderDetails,
	})
}

// get order by user_id
func (h *OrderHandler) GetOrderByUserID(ctx echo.Context) error {
	// Implementasi untuk mendapatkan semua pesanan dari usecase
	orders, err := h.OrderService.GetOrders(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	var orderDetails []map[string]interface{}
	for _, order := range orders {
		ticket, err := h.OrderService.GetTicketByID(ctx.Request().Context(), order.TicketID)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.NewHTTPError(http.StatusInternalServerError, err.Error()))
		}

		orderDetail := map[string]interface{}{
			"user_id": order.UserID,
			"ticket":  ticket,
		}
		orderDetails = append(orderDetails, orderDetail)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":       "Get all orders success",
		"order_details": orderDetails,
	})
}

// user create order
func (h *OrderHandler) UserCreateOrder(ctx echo.Context) error {
	var input struct {
		UserID   int64
		TicketID int64  `json:"ticket_id" validate:"required"`
		Quantity int64  `json:"quantity" validate:"required"`
		Status   string `json:"status" default:"success"`
	}

	// Get JWT token from the context
	token, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, errors.New("missing or invalid token"))
	}

	// Extract claims from the JWT token
	claims, ok := token.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, errors.New("invalid token claims"))
	}

	// Assign UserID from the JWT claims to the input struct
	input.UserID = claims.ID

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	// Remaining code remains the same...

	// Mendapatkan informasi saldo pengguna sebelum membuat pesanan
	userBalance, err := h.OrderService.GetUserBalance(ctx.Request().Context(), input.UserID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	// Mengambil informasi tiket dari TicketService untuk mendapatkan harga tiket
	ticketPrice, err := h.OrderService.GetTicketPrice(ctx.Request().Context(), input.TicketID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	// Memeriksa apakah saldo cukup untuk membuat pesanan
	if userBalance < (input.Quantity * ticketPrice) {
		return ctx.JSON(http.StatusUnprocessableEntity, errors.New("insufficient balance"))
	}

	order := entity.NewOrder(input.TicketID, input.Quantity, input.UserID, input.Status)
	err = h.OrderService.CreateOrder(ctx.Request().Context(), order)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	// Mengurangkan saldo pengguna setelah membuat pesanan
	err = h.OrderService.UpdateUserBalance(ctx.Request().Context(), input.UserID, input.Quantity*ticketPrice)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	// Remaining code remains the same...
	return ctx.JSON(http.StatusCreated, "Order created successfully")
}

// Get order History by jwt
func (h *OrderHandler) GetOrderHistory(ctx echo.Context) error {
	// Get JWT token from the context
	token, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, errors.New("missing or invalid token"))
	}

	// Extract claims from the JWT token
	claims, ok := token.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, errors.New("invalid token claims"))
	}

	// Get all orders by user ID
	orders, err := h.OrderService.GetOrderByUserID(ctx.Request().Context(), claims.ID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	var orderDetails []map[string]interface{}
	for _, order := range orders {
		ticket, err := h.OrderService.GetTicketByID(ctx.Request().Context(), order.TicketID)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.NewHTTPError(http.StatusInternalServerError, err.Error()))
		}

		orderDetail := map[string]interface{}{
			"user_id": order.UserID,
			"ticket":  ticket,
		}
		orderDetails = append(orderDetails, orderDetail)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":       "Get all orders success",
		"order_details": orderDetails,
	})
}
