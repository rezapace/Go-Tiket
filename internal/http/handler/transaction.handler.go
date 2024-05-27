package handler

import (
	"net/http"

	"Ticketing/entity"

	"Ticketing/common"
	"Ticketing/internal/http/validator"
	"Ticketing/internal/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionService service.TransactionUseCase
	paymentService     service.PaymentUseCase
	userService        service.UserUsecase
}

func NewTransactionHandler(transactionService service.TransactionUseCase, paymentService service.PaymentUseCase, userService service.UserUsecase) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
		paymentService:     paymentService,
		userService:        userService,
	}
}

func (h *TransactionHandler) CreateOrder(ctx echo.Context) error {
	var input struct {
		OrderID string `json:"order_id" validate:"required"`
		Amount  int64  `json:"amount" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*common.JwtCustomClaims)

	transaction := entity.NewTransaction(input.OrderID, claims.ID, input.Amount, "unpaid")

	err := h.transactionService.Create(ctx.Request().Context(), transaction)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	paymentRequest := entity.NewPaymentRequest(transaction.OrderID, transaction.Amount, claims.Name, "", claims.Email)

	payment, err := h.paymentService.CreateTransaction(ctx.Request().Context(), paymentRequest)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"url_pembayaran": payment})
}

func (h *TransactionHandler) WebHookTransaction(ctx echo.Context) error {
	var input entity.MidtransRequest

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	// Cari transaksi berdasarkan order ID untuk mendapatkan ID pengguna (user ID)
	transaction, err := h.transactionService.FindByOrderID(ctx.Request().Context(), input.OrderID)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	// Tentukan status default
	status := "unpaid"

	// Jika status transaksi adalah "settlement", ubah status menjadi "paid"
	if input.TransactionStatus == "settlement" {
		status = "paid"

		// Update status transaksi di database
		err = h.transactionService.UpdateStatus(ctx.Request().Context(), transaction.OrderID, status)
		if err != nil {
			return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		}

		// Tambahkan saldo ke user jika status transaksi adalah "paid"
		user, err := h.userService.FindByID(ctx.Request().Context(), transaction.UserID)
		if err != nil {
			return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		}

		// Tambahkan saldo ke user
		updatedSaldo := user.Saldo + transaction.Amount
		err = h.userService.UpdateSaldo(ctx.Request().Context(), user.ID, updatedSaldo)
		if err != nil {
			return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		}
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "success"})
}

// history transaction
func (h *TransactionHandler) HistoryTransaction(ctx echo.Context) error {
	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*common.JwtCustomClaims)

	transactions, err := h.transactionService.FindByUserID(ctx.Request().Context(), claims.ID)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, transactions)
}
