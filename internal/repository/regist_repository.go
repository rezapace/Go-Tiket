package repository

import (
	"Ticketing/entity"
	"context"

	"gorm.io/gorm"
)

type RegistrationRepository struct {
	db *gorm.DB
}

func NewRegistrationRepository(db *gorm.DB) *RegistrationRepository {
	return &RegistrationRepository{
		db: db,
	}
}

func (r *RegistrationRepository) Registration(ctx context.Context, user *entity.User) error {
	//menggunakan db untuk melakukan query ke database
	err := r.db.WithContext(ctx).Create(&user).Error // pada line ini akan melakukan query "INSERT INTO users"
	if err != nil {
		return err
	}
	return nil
}
