package repository

import (
	"context"

	"Ticketing/entity"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(ctx context.Context, transaction *entity.Transaction) error {
	if err := r.db.WithContext(ctx).Create(&transaction).Error; err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepository) FindByOrderID(ctx context.Context, orderID string) (*entity.Transaction, error) {
	transaction := new(entity.Transaction)
	if err := r.db.WithContext(ctx).Where("order_id = ?", orderID).First(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *TransactionRepository) FindByUserID(ctx context.Context, userID int64) ([]*entity.Transaction, error) {
	transactions := make([]*entity.Transaction, 0)
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) UpdateStatus(ctx context.Context, orderID string, status string) error {
	if err := r.db.WithContext(ctx).Model(&entity.Transaction{}).Where("order_id = ?", orderID).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

// HistoryTransaction
func (r *TransactionRepository) HistoryTransaction(ctx context.Context, userID int64) ([]*entity.Transaction, error) {
	transactions := make([]*entity.Transaction, 0)
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
