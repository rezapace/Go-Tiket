package service

import (
	"context"

	"Ticketing/entity"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentUseCase interface {
	CreateTransaction(ctx context.Context, paymentRequest *entity.PaymentRequest) (string, error)
}

type PaymentService struct {
	client snap.Client
}

func NewPaymentService(client snap.Client) *PaymentService {
	return &PaymentService{client: client}
}

func (s *PaymentService) CreateTransaction(ctx context.Context, paymentRequest *entity.PaymentRequest) (string, error) {
	request := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  paymentRequest.OrderID,
			GrossAmt: paymentRequest.Amount,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: paymentRequest.FirstName,
			LName: paymentRequest.LastName,
			Email: paymentRequest.Email,
		},
	}

	snapResponse, err := s.client.CreateTransaction(request)
	if err != nil {
		return "", err
	}

	return snapResponse.RedirectURL, nil
}
