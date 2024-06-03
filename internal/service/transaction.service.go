package service

import (
	"context"

	"Ticketing/entity"
)

type TransactionUseCase interface {
	Create(ctx context.Context, transaction *entity.Transaction) error
	FindByOrderID(ctx context.Context, orderID string) (*entity.Transaction, error)
	FindByUserID(ctx context.Context, userID int64) ([]*entity.Transaction, error)
	UpdateStatus(ctx context.Context, orderID string, status string) error
	HistoryTransaction(ctx context.Context, userID int64) ([]*entity.Transaction, error)
}

type TransactionRepository interface {
	Create(ctx context.Context, transaction *entity.Transaction) error
	FindByOrderID(ctx context.Context, orderID string) (*entity.Transaction, error)
	FindByUserID(ctx context.Context, userID int64) ([]*entity.Transaction, error)
	UpdateStatus(ctx context.Context, orderID string, status string) error
	HistoryTransaction(ctx context.Context, userID int64) ([]*entity.Transaction, error)
}

type TransactionService struct {
	repo TransactionRepository
}

func NewTransactionService(repo TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Create(ctx context.Context, transaction *entity.Transaction) error {
	return s.repo.Create(ctx, transaction)
}

func (s *TransactionService) FindByOrderID(ctx context.Context, orderID string) (*entity.Transaction, error) {
	return s.repo.FindByOrderID(ctx, orderID)
}

func (s *TransactionService) FindByUserID(ctx context.Context, userID int64) ([]*entity.Transaction, error) {
	return s.repo.FindByUserID(ctx, userID)
}

func (s *TransactionService) UpdateStatus(ctx context.Context, orderID string, status string) error {
	return s.repo.UpdateStatus(ctx, orderID, status)
}

// HistoryTransaction
func (s *TransactionService) HistoryTransaction(ctx context.Context, userID int64) ([]*entity.Transaction, error) {
	return s.repo.FindByUserID(ctx, userID)
}
