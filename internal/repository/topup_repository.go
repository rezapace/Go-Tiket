package repository

import (
	"Ticketing/entity"
	"context"

	"gorm.io/gorm"
)

type TopupRepository interface {
	InsertTopup(ctx context.Context, topup entity.Topup) (entity.Topup, error)
	UserTopup(ctx context.Context, topup entity.Topup) (entity.Topup, error)
	GetUserByID(ctx context.Context, id int) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
}

type topupRepository struct {
	db *gorm.DB
}

func NewTopupRepository(db *gorm.DB) *topupRepository {
	return &topupRepository{db}
}

func (r *topupRepository) InsertTopup(ctx context.Context, topup entity.Topup) (entity.Topup, error) {
	result := r.db.WithContext(ctx).Create(&topup)
	if result.Error != nil {
		return entity.Topup{}, result.Error
	}
	return topup, nil
}

// UserTopup
func (r *topupRepository) UserTopup(ctx context.Context, topup entity.Topup) (entity.Topup, error) {
	result := r.db.WithContext(ctx).Create(&topup)
	if result.Error != nil {
		return entity.Topup{}, result.Error
	}
	return topup, nil
}

// GetUserByID
func (r *topupRepository) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	user := new(entity.User)
	result := r.db.WithContext(ctx).First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// UpdateUser updates the user information
func (r *topupRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", user.ID).
		Updates(user).Error; err != nil {
		return err
	}
	return nil
}
