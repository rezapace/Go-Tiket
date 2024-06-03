package handler

import (
	"net/http"
	"Ticketing/entity"
	"Ticketing/internal/service"

	"github.com/labstack/echo/v4"
	"strconv"
	"time"
	"Ticketing/internal/http/validator"
)

// BlogHandler handles HTTP requests related to Blogs.
type BlogHandler struct {
	blogService service.BlogUseCase
}

// NewBlogHandler creates a new instance of BlogHandler.
func NewBlogHandler(blogService service.BlogUseCase) *BlogHandler {
	return &BlogHandler{blogService}
}

// GetAllBlog 
func (h *BlogHandler) GetAllBlogs(c echo.Context) error {
	Blogs, err := h.blogService.GetAllBlogs(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": Blogs,
	})
}


// CreateBlog
func (h *BlogHandler) CreateBlog(c echo.Context) error {
	var input struct {
		Title       string    `json:"title" validate:"required"`
		Description string    `json:"description" validate:"required"`
		Date		time.Time `json:"date"`
		Image       string    `json:"image"`
	}

	// Input validation
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	// Convert input.Date to a string with the desired format
	

	// Create a Blog object
	Blog := entity.Blog{
		Title:       input.Title,
		Description: input.Description,
		Image:       input.Image,
		Date:        time.Now().Format("2006-01-02T15:04:05Z"),
	}

	// Call the blogService to create the Blog
	err := h.blogService.CreateBlog(c.Request().Context(), &Blog)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	// Return a success message
	return c.JSON(http.StatusCreated, "Blog created successfully")
}


// GetBlog handles the retrieval of a Blog by ID.
func (h *BlogHandler) GetBlog(c echo.Context) error {
	idStr := c.Param("id") // assuming the ID is passed as a URL parameter as a string
	id, err := strconv.ParseInt(idStr, 10, 64) // Convert the string to int64
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
	}

	Blog, err := h.blogService.GetBlog(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"id":          Blog.ID,
			"title":       Blog.Title,
			"description": Blog.Description,
			"image":       Blog.Image,
			"date":        Blog.Date,
			"created":     Blog.CreatedAt,
		},
	})
}


// UpdateBlog handles the update of an existing Blog.
func (h *BlogHandler) UpdateBlog(c echo.Context) error {
	var input struct {
		ID          int64     `param:"id" validate:"required"`
		Title       string    `json:"title" validate:"required"`
		Description string    `json:"description" validate:"required"`
		Image       string    `json:"image"`
		Date        time.Time `json:"date"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}


	// Convert input.Date to a formatted string
	dateStr := input.Date.Format("2006-01-02T15:04:05Z")

	// Create a Blog object
	Blog := entity.Blog{
		ID:          input.ID,            // Assuming ID is already of type int64
		Title:       input.Title,
		Description: input.Description,
		Image:       input.Image,
		Date:        dateStr,            // Assign the formatted date string
	}



	err := h.blogService.UpdateBlog(c.Request().Context(), &Blog)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Blog updated successfully",
		"Blog":  Blog,
	})
}


// DeleteBlog handles the deletion of a Blog by ID.
func (h *BlogHandler) DeleteBlog(c echo.Context) error {
	var input struct {
		ID int64 `param:"id" validate:"required"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	err := h.blogService.DeleteBlog(c.Request().Context(), input.ID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Blog deleted successfully",
	})
}

// SearchBlog handles the search of a Blog by title.
func (h *BlogHandler) SearchBlog(c echo.Context) error {
	var input struct {
		Search string `param:"search" validate:"required"` //harus pramater search
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	Blogs, err := h.blogService.SearchBlog(c.Request().Context(), input.Search)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": Blogs,
	})
}
