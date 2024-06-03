package handler

import (
	"Ticketing/entity"
	"Ticketing/internal/service"
	"net/http"

	"Ticketing/internal/http/validator"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// TicketHandler handles HTTP requests related to tickets.
type TicketHandler struct {
	ticketService service.TicketUseCase
}

// NewTicketHandler creates a new instance of TicketHandler.
func NewTicketHandler(ticketService service.TicketUseCase) *TicketHandler {
	return &TicketHandler{ticketService}
}

// GetAllTicket
func (h *TicketHandler) GetAllTickets(c echo.Context) error {
	tickets, err := h.ticketService.GetAllTickets(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": tickets,
	})
}

// CreateTicket
func (h *TicketHandler) CreateTicket(c echo.Context) error {
	var input struct {
		Title       string    `json:"title" validate:"required"`
		Description string    `json:"description" validate:"required"`
		Image       string    `json:"image"`
		Location    string    `json:"location"`
		Date        time.Time `json:"date"`
		Status	  	string    `json:"status"`
		Price       float64   `json:"price"`
		Quota       int       `json:"quota"`
		Terjual	 	int       `json:"terjual"`
		Category    string    `json:"category"`
	}

	// Input validation
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	// Convert input.Date to a string with the desired format
	dateStr := input.Date.Format("2006-01-02T15:04:05Z")

	// Create a Ticket object
	ticket := entity.Ticket{
		Title:       input.Title,
		Description: input.Description,
		Image:       input.Image,
		Location:    input.Location,
		Date:        dateStr, // Assign the formatted date string
		Status:		 input.Status,
		Price:       int64(input.Price),
		Quota:       int64(input.Quota),
		Terjual:     int64(input.Terjual),
		Category:    input.Category,
		CreatedAt:   time.Now(),
	}

	// Call the ticketService to create the ticket
	err := h.ticketService.CreateTicket(c.Request().Context(), &ticket)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	// Return a success message
	return c.JSON(http.StatusCreated, "Ticket created successfully")
}

// GetTicket handles the retrieval of a ticket by ID.
// untuk sesudah login
func (h *TicketHandler) GetTicket(c echo.Context) error {
	idStr := c.Param("id")                     // assuming the ID is passed as a URL parameter as a string
	id, err := strconv.ParseInt(idStr, 10, 64) // Convert the string to int64
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
	}

	ticket, err := h.ticketService.GetTicket(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"id":          ticket.ID,
			"title":       ticket.Title,
			"description": ticket.Description,
			"image":       ticket.Image,
			"location":    ticket.Location,
			"date":        ticket.Date,
			"price":       ticket.Price,
			"quota":       ticket.Quota,
			"category":    ticket.Category,
			"created":     ticket.CreatedAt,
		},
	})
}

// UpdateTicket handles the update of an existing ticket.
func (h *TicketHandler) UpdateTicket(c echo.Context) error {
	var input struct {
		ID          int64     `param:"id" validate:"required"`
		Title       string    `json:"title" validate:"required"`
		Description string    `json:"description" validate:"required"`
		Image       string    `json:"image"`
		Location    string    `json:"location"`
		Date        time.Time `json:"date"`
		Price       float64   `json:"price"`
		Quota       int       `json:"quota"`
		Category    string    `json:"category"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	// Convert input.Date to a formatted string
	dateStr := input.Date.Format("2006-01-02T15:04:05Z")

	// Create a Ticket object
	ticket := entity.Ticket{
		ID:          input.ID, // Assuming ID is already of type int64
		Title:       input.Title,
		Description: input.Description,
		Image:       input.Image,
		Location:    input.Location,
		Date:        dateStr,            // Assign the formatted date string
		Price:       int64(input.Price), // Convert Price to int64 if needed
		Quota:       int64(input.Quota), // Convert Quota to int64 if needed
		Category:    input.Category,
	}

	err := h.ticketService.UpdateTicket(c.Request().Context(), &ticket)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Ticket updated successfully",
		"data": map[string]interface{}{
			"id":          ticket.ID,
			"title":       ticket.Title,
			"description": ticket.Description,
			"image":       ticket.Image,
			"location":    ticket.Location,
			"date":        ticket.Date,
			"price":       ticket.Price,
			"quota":       ticket.Quota,
			"category":    ticket.Category,
			"update":      ticket.UpdatedAt,
		},
	})
}

// DeleteTicket handles the deletion of a ticket by ID.
func (h *TicketHandler) DeleteTicket(c echo.Context) error {
	var input struct {
		ID int64 `param:"id" validate:"required"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	err := h.ticketService.DeleteTicket(c.Request().Context(), input.ID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Ticket deleted successfully",
	})
}

// SearchTicket handles the search of a ticket by title.
func (h *TicketHandler) SearchTicket(c echo.Context) error {
	var input struct {
		Search string `param:"search" validate:"required"` //harus pramater search
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	tickets, err := h.ticketService.SearchTicket(c.Request().Context(), input.Search)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": tickets,
	})
}

// filter ticket by location
func (h *TicketHandler) FilterTicket(c echo.Context) error {
	var input struct {
		Location string `param:"location" validate:"required"` //harus pramater search
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	tickets, err := h.ticketService.FilterTicket(c.Request().Context(), input.Location)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": tickets,
	})
}

// filter ticket by category
func (h *TicketHandler) FilterTicketByCategory(c echo.Context) error {
	var input struct {
		Category string `param:"category" validate:"required"` //harus pramater search
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	tickets, err := h.ticketService.FilterTicketByCategory(c.Request().Context(), input.Category)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": tickets,
	})
}

// filter ticket by range time (start - end)
func (h *TicketHandler) FilterTicketByRangeTime(c echo.Context) error {
	var input struct {
		Start string `param:"start" validate:"required"` //harus pramater search
		End   string `param:"end" validate:"required"`   //harus pramater search
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	tickets, err := h.ticketService.FilterTicketByRangeTime(c.Request().Context(), input.Start, input.End)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": tickets,
	})
}

// filter ticket by price (min - max)
func (h *TicketHandler) FilterTicketByPrice(c echo.Context) error {
	var input struct {
		Min string `param:"min" validate:"required"` //harus pramater search
		Max string `param:"max" validate:"required"` //harus pramater search
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	tickets, err := h.ticketService.FilterTicketByPrice(c.Request().Context(), input.Min, input.Max)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": tickets,
	})
}

// sortir tiket dari yang terbaru
func (h *TicketHandler) SortTicketByNewest(c echo.Context) error {
	// Membaca parameter 'sort' dari URL
	sortParam := c.QueryParam("sort")

	// Memastikan bahwa parameter sort adalah 'terbaru'
	if sortParam != "terbaru" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid sort parameter"})
	}

	// Memanggil service untuk mengurutkan tiket
	tickets, err := h.ticketService.SortTicketByNewest(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": tickets,
	})
}

// sortir dari yang termahal
func (h *TicketHandler) SortTicketByMostExpensive(c echo.Context) error {
	// Membaca parameter 'sort' dari URL
	sortParam := c.QueryParam("sort")

	// Memastikan bahwa parameter sort adalah 'termurah'
	if sortParam != "termahal" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid sort parameter"})
	}

	// Memanggil service untuk mengurutkan tiket dari yang termurah
	tickets, err := h.ticketService.SortTicketByMostExpensive(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": tickets,
	})
}

// ticket yang paling banyak dibeli
func (h *TicketHandler) SortTicketByCheapest(c echo.Context) error {
	sortParam := c.QueryParam("sort")

	// Memastikan bahwa parameter sort adalah 'termurah'
	if sortParam != "termurah" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid sort parameter"})
	}

	// Memanggil service untuk mengurutkan tiket dari yang termurah
	tickets, err := h.ticketService.SortTicketByCheapest(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": tickets,
	})
}

// ticket yang paling banyak dibeli
func (h *TicketHandler) SortTicketByMostBought(c echo.Context) error {
	sortParam := c.QueryParam("sort")

	// Memastikan bahwa parameter sort adalah 'terbanyak'
	if sortParam != "terbanyak" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid sort parameter"})
	}

	// Memanggil service untuk mengurutkan tiket dari yang terbanyak
	tickets, err := h.ticketService.SortTicketByMostBought(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": tickets,
	})
}

// ticket yang masih tersedia
func (h *TicketHandler) SortTicketByAvailable(c echo.Context) error {
	sortParam := c.QueryParam("sort")

	// Memastikan bahwa parameter sort adalah 'tersedia'
	if sortParam != "tersedia" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid sort parameter"})
	}

	// Memanggil service untuk mengurutkan tiket dari yang tersedia
	tickets, err := h.ticketService.SortTicketByAvailable(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": tickets,
	})
}