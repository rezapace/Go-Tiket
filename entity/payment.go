package entity

type PaymentRequest struct {
	OrderID   string
	Amount    int64
	FirstName string
	LastName  string
	Email     string
}

func NewPaymentRequest(orderID string, amount int64, firstName, lastName, email string) *PaymentRequest {
	return &PaymentRequest{
		OrderID:   orderID,
		Amount:    amount,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}
