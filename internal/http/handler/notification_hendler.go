package handler

import (
	"Ticketing/entity"
	"Ticketing/internal/service"
	"net/http"

	"Ticketing/internal/http/validator"
	"time"

	"github.com/labstack/echo/v4"
)

type NotificationHandler struct {
	notificationService service.NotificationUsecase
}

func NewNotificationHandler(notificationService service.NotificationUsecase) *NotificationHandler {
	return &NotificationHandler{notificationService}
}

// GetAllNotification
func (h *NotificationHandler) GetAllNotification(c echo.Context) error {
	Notifications, err := h.notificationService.GetAllNotification(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": Notifications,
	})
}

// func untuk create notification
func (h *NotificationHandler) CreateNotification(c echo.Context) error {
	var input struct {
		Type      string    `json:"type" validate:"required"`
		Message   string    `json:"message" validate:"required"`
		Is_Read   bool      `json:"is_read"`
		Create_at time.Time `json:"create_at"`
	}

	// Input validation
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	// Create a Notification object
	Notification := entity.Notification{
		Type:      input.Type,
		Message:   input.Message,
		IsRead:    input.Is_Read,
		CreatedAt: time.Now(),
	}

	err := h.notificationService.CreateNotification(c.Request().Context(), &Notification)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, validator.ValidatorErrors(err))
	}

	return c.JSON(http.StatusCreated, Notification)
}

// get notification after get chage value isRead to true and only get notification if isread false
func (h *NotificationHandler) UserGetNotification(c echo.Context) error {
	Notifications, err := h.notificationService.UserGetNotification(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": Notifications,
	})
}
