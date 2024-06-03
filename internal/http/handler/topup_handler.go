package handler

import (
	"Ticketing/entity"
	"Ticketing/internal/service"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TopupHandler struct {
	topupService service.TopupService
}

func NewTopupHandler(topupService service.TopupService) *TopupHandler {
	return &TopupHandler{topupService}
}

func (h *TopupHandler) CreateTopup(c echo.Context) error {
	var topup entity.Topup
	if err := c.Bind(&topup); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	chargeResp, err := h.topupService.CreateMidtransCharge(topup.ID, int64(topup.Amount))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	topup.SnapURL = chargeResp.RedirectURL

	// Perhatikan penambahan c.Request().Context() di sini
	newTopup, err := h.topupService.CreateTopup(c.Request().Context(), topup)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, newTopup)
}

// topup saldo by jwt token
func (h *TopupHandler) UserTopup(c echo.Context) error {
	// Get JWT token from the context
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "missing or invalid token"})
	}

	// Extract claims from the JWT token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token claims"})
	}

	// Get user ID from the JWT claims
	userID := int(claims["user_id"].(float64))

	var topup entity.Topup
	if err := c.Bind(&topup); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// Update user saldo
	userSaldo, err := h.topupService.UpdateUserSaldo(c.Request().Context(), userID, int64(topup.Amount))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// Create topup record
	newTopup, err := h.topupService.UserTopup(c.Request().Context(), userID, topup)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"user_saldo": userSaldo,
		"topup_data": newTopup,
	})
}
