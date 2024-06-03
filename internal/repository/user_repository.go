package repository

// NOTE :
// FOLDER INI UNTUK MENANGANI KE BAGIAN DATABASE DAN QUERY
import (
	"context"
	"errors"

	"Ticketing/entity"

	"gorm.io/gorm"
)

// membuat struct untuk dependency
type UserRepository struct {
	db *gorm.DB
}

// membuat constructor untuk dependency
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db}
}

// menampilkan get all user
// menggunakan []*entity.User = karena akan membutuhkan data yg banyak dengan array slice of user.
func (r *UserRepository) GetAll(ctx context.Context) ([]*entity.User, error) {
	//melakukan returtn dari data user itu sendir, sehingga disimpan di variabel users
	users := make([]*entity.User, 0)
	//menggunakan db untuk melakukan query ke database
	err := r.db.WithContext(ctx).Find(&users).Error // pada line ini akan melakukan query "SELECT * FROM users"
	if err != nil {
		return nil, err
	}
	return users, nil

}

// membuat create user
func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	//menggunakan db untuk melakukan query ke database
	err := r.db.WithContext(ctx).Create(&user).Error // pada line ini akan melakukan query "INSERT INTO users"
	if err != nil {
		return err
	}
	return nil
}

// update data user byID
func (r *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", user.ID).
		Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

// get user by id
func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// detele user by id
func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

// get by email
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user with that email not found")
	}
	return user, nil
}

// Update User Self
func (r *UserRepository) UpdateProfile(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", user.ID).
		Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

// update user balance by id
func (r *UserRepository) UpdateUserBalance(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", user.ID).
		Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

// GetProfile retrieves the user profile by ID from the database
func (r *UserRepository) GetProfile(ctx context.Context, userID int64) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserBalance
func (r *UserRepository) GetUserBalance(ctx context.Context, userID int64) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteAccount
func (r *UserRepository) DeleteAccount(ctx context.Context, email string) error {
	if err := r.db.WithContext(ctx).Delete(&entity.User{}, email).Error; err != nil {
		return err
	}
	return nil
}

// upgrade saldo
func (r *UserRepository) UpgradeSaldo(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", user.ID).
		Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

// logout
func (r *UserRepository) UserLogout(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", user.ID).
		Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateSaldo
// UpdateSaldo updates the saldo of a user by ID
func (r *UserRepository) UpdateSaldo(ctx context.Context, userID int64, updatedSaldo int64) error {
	user := &entity.User{ID: userID, Saldo: updatedSaldo}

	if err := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", userID).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

// FindByID
func (r *UserRepository) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// TanyaAI
// func (r *UserRepository) TanyaAI(ctx context.Context, user *entity.User) error {
// 	return nil
// }

//BuyerCreateAccount
// func (r *UserRepository) BuyerCreateAccount(ctx context.Context, user *entity.User) error {
// 	//menggunakan db untuk melakukan query ke database
// 	err := r.db.WithContext(ctx).Create(&user).Error // pada line ini akan melakukan query "INSERT INTO users"
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
