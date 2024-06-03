package handler

//NOTE :
// FOLDER INI UNTUK MEMANGGIL SERVICE DAN REPOSITORY
import (
	"Ticketing/common"
	"Ticketing/entity"
	"Ticketing/internal/http/validator"
	"Ticketing/internal/service"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserUsecase
}

// melakukan instace dari user handler
func NewUserHandler(userService service.UserUsecase) *UserHandler {
	return &UserHandler{userService}
}

// func untuk melakukan getAll User
func (h *UserHandler) GetAllUser(ctx echo.Context) error {
	users, err := h.userService.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": users,
	})
}

// func untuk melakukan createUser update versi reza v5 halo
func (h *UserHandler) CreateUser(ctx echo.Context) error {
	var input struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"email"`
		Number   string `json:"number" validate:"min=11,max=13"`
		Roles    string `json:"roles" validate:"oneof=Admin Buyer"`
		Password string `json:"password"`
		Saldo    int64  `json:"saldo"`
	}
	//ini func untuk error checking
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	user := entity.NewUser(input.Name, input.Email, input.Number, input.Roles, input.Password, input.Saldo)
	err := h.userService.CreateUser(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	//kalau retrun nya kaya gini akan tampil pesan "User created successfully"
	return ctx.JSON(http.StatusCreated, "User created successfully")
}

// func untuk melakukan updateUser by id
func (h *UserHandler) UpdateUser(ctx echo.Context) error {
	var input struct {
		ID       int64  `param:"id" validate:"required"`
		Name     string `json:"name"`
		Email    string `json:"email" validate:"email"`
		Number   string `json:"number" validate:"min=11,max=13"`
		Roles    string `json:"roles" validate:"oneof=Admin Buyer"`
		Password string `json:"password"`
		Saldo    int64  `json:"saldo"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user := entity.UpdateUser(input.ID, input.Name, input.Email, input.Number, input.Roles, input.Password, input.Saldo)

	err := h.userService.UpdateUser(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{"success": "succesfully update user"})
}

// func untuk melakukan getUser by id
func (h *UserHandler) GetUserByID(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// Jika tidak dapat mengonversi ID menjadi int64, kembalikan respons error
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
	}
	user, err := h.userService.GetUserByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"id":       user.ID,
			"name":     user.Name,
			"email":    user.Email,
			"number":   user.Number,
			"password": user.Password,
			"created":  user.CreatedAt,
			"updated":  user.UpdatedAt,
		},
	})
}

// DeleteUser func untuk melakukan delete user by id
func (h *UserHandler) DeleteUser(ctx echo.Context) error {
	var input struct {
		ID int64 `param:"id" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	err := h.userService.Delete(ctx.Request().Context(), input.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})
}

// Update User Self
func (h *UserHandler) UpdateProfile(ctx echo.Context) error {
	var input struct {
		ID       int64
		Name     string `json:"name"`
		Email    string `json:"email" validate:"email"`
		Number   string `json:"number" validate:"min=11,max=13"`
		Password string `json:"password"`
		Saldo    int64  `json:"saldo"`
	}

	// Mengambil nilai 'claims' dari JWT token
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user information from claims")
	}

	// Mengisi ID dari klaim ke input
	input.ID = claimsData.ID

	// Update user
	user := entity.UpdateProfile(input.ID, input.Name, input.Email, input.Number, input.Password)

	// Memanggil service untuk update user
	err := h.userService.UpdateProfile(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{"success": "successfully update user"})
}

// get profile
func (h *UserHandler) GetProfile(ctx echo.Context) error {
	// Retrieve user claims from the JWT token
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user information from claims")
	}

	// Fetch user profile using the user ID
	user, err := h.userService.GetProfile(ctx.Request().Context(), claimsData.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	// Return the user profile
	return ctx.JSON(http.StatusOK, user)
}

// Get user balance
func (h *UserHandler) GetUserBalance(ctx echo.Context) error {
	// Retrieve user claims from the JWT token
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user information from claims")
	}

	// Fetch user balance using the user ID
	balance, err := h.userService.GetUserBalance(ctx.Request().Context(), claimsData.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	// Return the user balance
	return ctx.JSON(http.StatusOK, balance.Saldo)
}

// delete account
func (h *UserHandler) DeleteAccount(ctx echo.Context) error {
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user information from claims")
	}

	// Menggunakan ID dari klaim JWT
	idToDelete := claimsData.ID

	err := h.userService.Delete(ctx.Request().Context(), idToDelete)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})
}

// upgrade saldo
func (h *UserHandler) UpgradeSaldo(ctx echo.Context) error {
	// Retrieve user ID from JWT claims
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user information from claims")
	}

	userID := claimsData.ID

	// Fetch current saldo for the user
	currentUser, err := h.userService.GetUserByID(ctx.Request().Context(), userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "unable to fetch user information")
	}

	// Extract input data
	var input struct {
		Saldo int64 `json:"saldo"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	// Add the new saldo to the current saldo
	newSaldo := currentUser.Saldo + input.Saldo

	// Update user saldo
	currentUser.Saldo = newSaldo
	err = h.userService.UpgradeSaldo(ctx.Request().Context(), currentUser)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{"success": "successfully updated user saldo"})
}

// logout
func (h *UserHandler) UserLogout(ctx echo.Context) error {
	// Retrieve user claims from the JWT token
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user information from claims")
	}

	userID := claimsData.ID

	// Create a *entity.User instance with the userID
	user := &entity.User{ID: userID}

	// Invalidate the JWT token
	err := h.userService.UserLogout(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "unable to invalidate JWT token")
	}

	return ctx.JSON(http.StatusOK, map[string]string{"success": "successfully logged out"})
}

// func DiagnosticAI(userInput, openAIKey string) (string, error) {
// 	ctx := context.Background()
// 	client := openai.NewClient(openAIKey)
// 	model := openai.GPT3Dot5Turbo
// 	messages := []openai.ChatCompletionMessage{
// 		{
// 			Role:    openai.ChatMessageRoleSystem,
// 			Content: "Siswa menanyakan soal. Jawablah dengan materi tentang soal dan jangan berikan jawaban secara langsung",
// 		},
// 		{
// 			Role:    openai.ChatMessageRoleUser,
// 			Content: userInput,
// 		},
// 	}

// 	resp, err := getCompletionFromMessages(ctx, client, messages, model)
// 	if err != nil {
// 		return "", err
// 	}
// 	answer := resp.Choices[0].Message.Content
// 	return answer, nil
// }

// func getCompletionFromMessages(
// 	ctx context.Context,
// 	client *openai.Client,
// 	messages []openai.ChatCompletionMessage,
// 	model string,
// ) (openai.ChatCompletionResponse, error) {
// 	if model == "" {
// 		model = openai.GPT3Dot5Turbo
// 	}

// 	resp, err := client.CreateChatCompletion(
// 		ctx,
// 		openai.ChatCompletionRequest{
// 			Model:    model,
// 			Messages: messages,
// 		},
// 	)
// 	return resp, err
// }

// TanyaAI
// func (h *UserHandler) TanyaAI(c echo.Context) error {
// 	// Baca data masukan dari body request
// 	var requestBody struct {
// 		Pertanyaan string `json:"pertanyaan"`
// 	}

// 	if err := c.Bind(&requestBody); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
// 	}

// 	// Panggil fungsi AI untuk menjawab pertanyaan
// 	jawaban, err := h.aiService.TanyaAI(context.Background(), requestBody.Pertanyaan)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 	}

// 	// Kirim jawaban ke client
// 	return c.JSON(http.StatusOK, map[string]string{"jawaban": jawaban})
// }

// buyer create account
// func (h *UserHandler) BuyerCreateAccount(ctx echo.Context) error {
// 	var input struct {
// 		Name     string `json:"name" validate:"required"`
// 		Email    string `json:"email" validate:"email"`
// 		Number   string `json:"number" validate:"min=11,max=13"`
// 		Roles    string `json:"roles" default:"Buyer"`
// 		Password string `json:"password"`
// 		Saldo    int64  `json:"saldo" default:"0"`
// 	}
// 	//ini func untuk error checking
// 	if err := ctx.Bind(&input); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
// 	}
// 	user := entity.NewUser(input.Name, input.Email, input.Number, input.Roles, input.Password, input.Saldo)
// 	err := h.userService.CreateUser(ctx.Request().Context(), user)
// 	if err != nil {
// 		return ctx.JSON(http.StatusUnprocessableEntity, err)
// 	}
// 	//kalau retrun nya kaya gini akan tampil pesan "User created successfully"
// 	return ctx.JSON(http.StatusCreated, "User created successfully")
// }
