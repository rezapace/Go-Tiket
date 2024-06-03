package service

//NOTE :
// FOLDER INI UNTUK MENANGANI LOGIC DAN MEMANGGIL REPOSITORY
import (
	"context"

	"Ticketing/entity"
)

// interface untuk service
// untuk memanngil repository
type UserUsecase interface {
	GetAll(ctx context.Context) ([]*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	GetUserByID(ctx context.Context, id int64) (*entity.User, error)
	Delete(ctx context.Context, id int64) error
	UpdateProfile(ctx context.Context, user *entity.User) error
	UpdateUserBalance(ctx context.Context, user *entity.User) error
	GetProfile(ctx context.Context, userID int64) (*entity.User, error)
	GetUserBalance(ctx context.Context, userID int64) (*entity.User, error)
	DeleteAccount(ctx context.Context, email string) error
	UpgradeSaldo(ctx context.Context, user *entity.User) error
	UserLogout(ctx context.Context, user *entity.User) error
	UpdateSaldo(ctx context.Context, userID int64, updatedSaldo int64) error
	FindByID(ctx context.Context, id int64) (*entity.User, error)
	// BuyerCreateAccount(ctx context.Context, user *entity.User) error
}

// interface untuk repository
// untuk memanggil repository
// GetAll = untuk menampilkan semua data user, dan itu harus sama dengan yang ada di repository
type UserRepository interface {
	GetAll(ctx context.Context) ([]*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	GetUserByID(ctx context.Context, id int64) (*entity.User, error)
	Delete(ctx context.Context, id int64) error
	UpdateProfile(ctx context.Context, user *entity.User) error
	UpdateUserBalance(ctx context.Context, user *entity.User) error
	GetProfile(ctx context.Context, userID int64) (*entity.User, error)
	GetUserBalance(ctx context.Context, userID int64) (*entity.User, error)
	DeleteAccount(ctx context.Context, email string) error
	UpgradeSaldo(ctx context.Context, user *entity.User) error
	UserLogout(ctx context.Context, user *entity.User) error
	UpdateSaldo(ctx context.Context, userID int64, updatedSaldo int64) error
	FindByID(ctx context.Context, id int64) (*entity.User, error)
	// BuyerCreateAccount(ctx context.Context, user *entity.User) error
}

// code di line 23 merupakan dependency injection, karena repository tidak langsung di panggil.
// karena repository dipanggil melalui code pada line 18
type UserService struct {
	repository UserRepository
}

// func untuk UserRepository
func NewUserService(repository UserRepository) *UserService {
	return &UserService{repository}
}

// func dibawah ini untuk type user usecase
// ini untuk menampilkan data user
// untuk memanggil repository
func (s *UserService) GetAll(ctx context.Context) ([]*entity.User, error) {
	return s.repository.GetAll(ctx)
}

// func dibawah ini untuk type user usecase
// ini untuk membuat data user
func (s *UserService) CreateUser(ctx context.Context, user *entity.User) error {
	return s.repository.CreateUser(ctx, user)
}

// untuk update data user
func (s *UserService) UpdateUser(ctx context.Context, user *entity.User) error {
	return s.repository.UpdateUser(ctx, user)
}

// untuk get user by id
func (s *UserService) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	return s.repository.GetUserByID(ctx, id)
}

// untuk delete by id
func (s *UserService) Delete(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}

// func update saldo user by id
func (s *UserService) UpdateUserBalance(ctx context.Context, user *entity.User) error {
	return s.repository.UpdateUserBalance(ctx, user)
}

// Update User Self
func (s *UserService) UpdateProfile(ctx context.Context, user *entity.User) error {
	return s.repository.UpdateProfile(ctx, user)
}

// GetProfile retrieves the user profile by ID
func (s *UserService) GetProfile(ctx context.Context, userID int64) (*entity.User, error) {
	return s.repository.GetProfile(ctx, userID)
}

// GetUserBalance
func (s *UserService) GetUserBalance(ctx context.Context, userID int64) (*entity.User, error) {
	return s.repository.GetUserBalance(ctx, userID)
}

// DeleteAccount
func (s *UserService) DeleteAccount(ctx context.Context, email string) error {
	return s.repository.DeleteAccount(ctx, email)
}

// upgrade saldo
func (s *UserService) UpgradeSaldo(ctx context.Context, user *entity.User) error {
	return s.repository.UpgradeSaldo(ctx, user)
}

// logout
func (s *UserService) UserLogout(ctx context.Context, user *entity.User) error {
	return s.repository.UserLogout(ctx, user)
}

// UpdateSaldo updates the saldo of a user by ID
func (s *UserService) UpdateSaldo(ctx context.Context, userID int64, updatedSaldo int64) error {
	return s.repository.UpdateSaldo(ctx, userID, updatedSaldo)
}

// FindByID
func (s *UserService) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	return s.repository.FindByID(ctx, id)
}

// TanyaAI
// func (s *UserService) TanyaAI(ctx context.Context, pertanyaan string) (string, error) {
// 	// Implementasi logika untuk bertanya ke AI di sini
// 	// Misalnya, menggunakan paket go-openai atau alat pemrosesan bahasa alami lainnya
// 	// ...

// 	// Sebagai contoh, kita akan menggunakan jawaban statis
// 	return "Jawaban dari AI untuk pertanyaan: " + pertanyaan, nil
// }

//BuyerCreateAccount
// func (s *UserService) BuyerCreateAccount(ctx context.Context, user *entity.User) error {
// 	return s.repository.CreateUser(ctx, user)
// }
