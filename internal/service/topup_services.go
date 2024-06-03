package service

import (
	"Ticketing/entity"
	"Ticketing/internal/repository"
	"context"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type TopupService interface {
	CreateTopup(ctx context.Context, topup entity.Topup) (entity.Topup, error)
	CreateMidtransCharge(orderID string, amount int64) (*coreapi.ChargeResponse, error)
	UpdateUserSaldo(ctx context.Context, userID int, amount int64) (int64, error)
	UserTopup(ctx context.Context, userID int, topup entity.Topup) (entity.Topup, error)

	// TopupSaldo(ctx context.Context, topup entity.Topup) (entity.Topup, error)
}

type topupService struct {
	topupRepository repository.TopupRepository
}

type TopupRepository interface {
	UserTopup(ctx context.Context, topup entity.Topup) (entity.Topup, error)
	UpdateUserSaldo(ctx context.Context, userID int, amount int64) (int64, error)
}
type TopupUsecase interface {
	UserTopup(ctx context.Context, topup entity.Topup) (entity.Topup, error)
	UpdateUserSaldo(ctx context.Context, userID int, amount int64) (int64, error)
}

func NewTopupService(topupRepository repository.TopupRepository) *topupService {
	return &topupService{topupRepository}
}
func (s *topupService) CreateTopup(ctx context.Context, topup entity.Topup) (entity.Topup, error) {
	return s.topupRepository.InsertTopup(ctx, topup)
}

func (s *topupService) CreateMidtransCharge(orderID string, amount int64) (*coreapi.ChargeResponse, error) {
	c := coreapi.Client{}
	c.New("SB-Mid-server-RGIeiP4-kdQSADSfhbpG42B-", midtrans.Sandbox) // Ganti dengan server key Anda

	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer, // Sesuaikan dengan jenis pembayaran
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: amount,
		},
		// Tambahkan detail lainnya sesuai kebutuhan
	}

	return c.ChargeTransaction(chargeReq)
}

// UserTopup
func (s *topupService) UserTopup(ctx context.Context, userID int, topup entity.Topup) (entity.Topup, error) {
	return s.topupRepository.UserTopup(ctx, topup)
}

// TopupService method for updating user saldo
func (s *topupService) UpdateUserSaldo(ctx context.Context, userID int, amount int64) (int64, error) {
	user, err := s.topupRepository.GetUserByID(ctx, userID)
	if err != nil {
		return 0, err
	}

	// Update user saldo
	user.Saldo += amount

	// Save the updated user information
	if err := s.topupRepository.UpdateUser(ctx, user); err != nil {
		return 0, err
	}

	return user.Saldo, nil
}
