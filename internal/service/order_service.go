package service

import (
	"Ticketing/entity"
	"context"
	"errors"
)

type OrderUsecase interface {
	CreateOrder(ctx context.Context, order *entity.Order) error
	GetTicket(ctx context.Context, ticketID int64) (*entity.Ticket, error)
	UpdateTicket(ctx context.Context, ticket *entity.Ticket) error
	GetOrders(ctx context.Context) ([]*entity.Order, error)
	GetTicketByID(ctx context.Context, id int64) (*entity.Ticket, error)
	GetOrderByUserID(ctx context.Context, userID int64) ([]*entity.Order, error)
	UpdateUserBalance(ctx context.Context, userID int64, amount int64) error
	GetUserBalance(ctx context.Context, userID int64) (int64, error)
	GetTicketPrice(ctx context.Context, ticketID int64) (int64, error)
	UserCreateOrder(ctx context.Context, order *entity.Order) error
	GetOrderHistory(ctx context.Context, userID int64) ([]*entity.Order, error)
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *entity.Order) error
	GetTicket(ctx context.Context, ticketID int64) (*entity.Ticket, error)
	UpdateTicket(ctx context.Context, ticket *entity.Ticket) error
	GetOrders(ctx context.Context) ([]*entity.Order, error)
	GetTicketByID(ctx context.Context, id int64) (*entity.Ticket, error)
	GetOrderByUserID(ctx context.Context, userID int64) ([]*entity.Order, error)
	UpdateUserBalance(ctx context.Context, userID int64, amount int64) error
	GetUserBalance(ctx context.Context, userID int64) (int64, error)
	GetTicketPrice(ctx context.Context, ticketID int64) (int64, error)
	UserCreateOrder(ctx context.Context, order *entity.Order) error
	GetOrderHistory(ctx context.Context, userID int64) ([]*entity.Order, error)
}

type OrderService struct {
	repository OrderRepository
}

func NewOrderService(repository OrderRepository) *OrderService {
	return &OrderService{repository}
}

// Updated CreateOrder method in OrderService to receive TicketService
func (s *OrderService) CreateOrder(ctx context.Context, order *entity.Order) error {
	// Mendapatkan informasi tiket berdasarkan ID tiket dalam pesanan
	ticket, err := s.repository.GetTicket(ctx, order.TicketID)
	if err != nil {
		return err
	}

	// Memeriksa ketersediaan tiket
	if int64(ticket.Quota) < order.Quantity {
		return errors.New("ticket is not available")
	}

	// Melakukan perhitungan total harga pesanan
	order.Total = ticket.Price * int64(order.Quantity)

	// Membuat pesanan
	if err := s.repository.CreateOrder(ctx, order); err != nil {
		return err
	}

	// Mengurangi ketersediaan tiket
	ticket.Quota -= order.Quantity
	if err := s.repository.UpdateTicket(ctx, ticket); err != nil {
		return err
	}

	// Mengurangi saldo pengguna
	if err := s.repository.UpdateUserBalance(ctx, order.UserID, order.Total); err != nil {
		return err
	}

	return nil
}

// Implementasi fungsi GetTicket
func (s *OrderService) GetTicket(ctx context.Context, ticketID int64) (*entity.Ticket, error) {
	return s.repository.GetTicket(ctx, ticketID)
}

// Implementasi fungsi UpdateTicket
func (s *OrderService) UpdateTicket(ctx context.Context, ticket *entity.Ticket) error {
	return s.repository.UpdateTicket(ctx, ticket)
}

// implementasi fungsi update user balance
func (s *OrderService) UpdateUserBalance(ctx context.Context, userID int64, saldo int64) error {
	return s.repository.UpdateUserBalance(ctx, userID, saldo)
}

func (s *OrderService) GetOrders(ctx context.Context) ([]*entity.Order, error) {
	return s.repository.GetOrders(ctx)
}

func (s *OrderService) GetTicketByID(ctx context.Context, id int64) (*entity.Ticket, error) {
	return s.repository.GetTicketByID(ctx, id)
}

// get order by user_id
func (s *OrderService) GetOrderByUserID(ctx context.Context, userID int64) ([]*entity.Order, error) {
	return s.repository.GetOrderByUserID(ctx, userID)
}

// get user balance
func (s *OrderService) GetUserBalance(ctx context.Context, userID int64) (int64, error) {
	return s.repository.GetUserBalance(ctx, userID)
}

// GetTicketPrice
func (s *OrderService) GetTicketPrice(ctx context.Context, ticketID int64) (int64, error) {
	ticket, err := s.repository.GetTicket(ctx, ticketID)
	if err != nil {
		return 0, err
	}
	return int64(ticket.Price), nil
}

// UserCreateOrder
func (s *OrderService) UserCreateOrder(ctx context.Context, order *entity.Order) error {
	// Mendapatkan informasi tiket berdasarkan ID tiket dalam pesanan
	ticket, err := s.repository.GetTicket(ctx, order.TicketID)
	if err != nil {
		return err
	}

	// Memeriksa ketersediaan tiket
	if int64(ticket.Quota) < order.Quantity {
		return errors.New("ticket is not available")
	}

	// Melakukan perhitungan total harga pesanan
	order.Total = ticket.Price * int64(order.Quantity)

	// Membuat pesanan
	if err := s.repository.CreateOrder(ctx, order); err != nil {
		return err
	}

	// Mengurangi ketersediaan tiket
	ticket.Quota -= order.Quantity
	if err := s.repository.UpdateTicket(ctx, ticket); err != nil {
		return err
	}

	// Mengurangi saldo pengguna
	if err := s.repository.UpdateUserBalance(ctx, order.UserID, order.Total); err != nil {
		return err
	}

	return nil
}

// GetOrderHistory
func (s *OrderService) GetOrderHistory(ctx context.Context, userID int64) ([]*entity.Order, error) {
	return s.repository.GetOrderByUserID(ctx, userID)
}
