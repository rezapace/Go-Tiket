package entity

type Transaction struct {
	ID      int64
	OrderID string
	UserID  int64
	Amount  int64
	Status  string
}

func NewTransaction(orderID string, userID int64, amount int64, status string) *Transaction {
	return &Transaction{
		OrderID: orderID,
		UserID:  userID,
		Amount:  amount,
		Status:  status,
	}
}
