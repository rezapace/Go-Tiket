package service

import (
	"context"

	"Ticketing/entity"
)

// TicketUseCase is an interface for ticket-related use cases.
type TicketUseCase interface {
	GetAllTickets(ctx context.Context) ([]*entity.Ticket, error)
	CreateTicket(ctx context.Context, ticket *entity.Ticket) error
	GetTicket(ctx context.Context, id int64) (*entity.Ticket, error)
	UpdateTicket(ctx context.Context, ticket *entity.Ticket) error
	DeleteTicket(ctx context.Context, id int64) error
	SearchTicket(ctx context.Context, search string) ([]*entity.Ticket, error)
	FilterTicket(ctx context.Context, location string) ([]*entity.Ticket, error)
	FilterTicketByCategory(ctx context.Context, category string) ([]*entity.Ticket, error)
	FilterTicketByRangeTime(ctx context.Context, start string, end string) ([]*entity.Ticket, error)
	FilterTicketByPrice(ctx context.Context, min string, max string) ([]*entity.Ticket, error)
	SortTicketByNewest(ctx context.Context) ([]*entity.Ticket, error)
	SortTicketByMostExpensive(ctx context.Context) ([]*entity.Ticket, error)
	SortTicketByCheapest(ctx context.Context) ([]*entity.Ticket, error)
	SortTicketByMostBought(ctx context.Context) ([]*entity.Ticket, error)
	SortTicketByAvailable(ctx context.Context) ([]*entity.Ticket, error)

}

type TicketRepository interface {
	GetAllTickets(ctx context.Context) ([]*entity.Ticket, error)
	CreateTicket(ctx context.Context, ticket *entity.Ticket) error
	GetTicket(ctx context.Context, id int64) (*entity.Ticket, error)
	UpdateTicket(ctx context.Context, ticket *entity.Ticket) error
	SearchTicket(ctx context.Context, search string) ([]*entity.Ticket, error)
	DeleteTicket(ctx context.Context, id int64) error
	FilterTicket(ctx context.Context, location string) ([]*entity.Ticket, error)
	FilterTicketByCategory(ctx context.Context, category string) ([]*entity.Ticket, error)
	FilterTicketByRangeTime(ctx context.Context, start string, end string) ([]*entity.Ticket, error)
	FilterTicketByPrice(ctx context.Context, min string, max string) ([]*entity.Ticket, error)
	SortTicketByNewest(ctx context.Context) ([]*entity.Ticket, error)
	SortTicketByMostExpensive(ctx context.Context) ([]*entity.Ticket, error)
	SortTicketByCheapest(ctx context.Context) ([]*entity.Ticket, error)
	SortTicketByMostBought(ctx context.Context) ([]*entity.Ticket, error)
	SortTicketByAvailable(ctx context.Context) ([]*entity.Ticket, error)
}

// TicketService is responsible for ticket-related business logic.
type TicketService struct {
	Repository TicketRepository
}

// NewTicketService creates a new instance of TicketService.
func NewTicketService(Repository TicketRepository) *TicketService {
	return &TicketService{Repository: Repository}
}

func (s *TicketService) GetAllTickets(ctx context.Context) ([]*entity.Ticket, error) {
	return s.Repository.GetAllTickets(ctx)
}

func (s *TicketService) CreateTicket(ctx context.Context, ticket *entity.Ticket) error {
	return s.Repository.CreateTicket(ctx, ticket)
}

func (s *TicketService) UpdateTicket(ctx context.Context, ticket *entity.Ticket) error {
	return s.Repository.UpdateTicket(ctx, ticket)
}

func (s *TicketService) GetTicket(ctx context.Context, id int64) (*entity.Ticket, error) {
	return s.Repository.GetTicket(ctx, id)
}

func (s *TicketService) DeleteTicket(ctx context.Context, id int64) error {
	return s.Repository.DeleteTicket(ctx, id)
}

// search ticket
func (s *TicketService) SearchTicket(ctx context.Context, search string) ([]*entity.Ticket, error) {
	return s.Repository.SearchTicket(ctx, search)
}

// filter ticket by location
func (s *TicketService) FilterTicket(ctx context.Context, location string) ([]*entity.Ticket, error) {
	return s.Repository.FilterTicket(ctx, location)
}

// filter ticket by category
func (s *TicketService) FilterTicketByCategory(ctx context.Context, category string) ([]*entity.Ticket, error) {
	return s.Repository.FilterTicketByCategory(ctx, category)
}

// filter ticket by range time (start - end)
func (s *TicketService) FilterTicketByRangeTime(ctx context.Context, start string, end string) ([]*entity.Ticket, error) {
	return s.Repository.FilterTicketByRangeTime(ctx, start, end)
}

// filter ticket by price (min - max)
func (s *TicketService) FilterTicketByPrice(ctx context.Context, min string, max string) ([]*entity.Ticket, error) {
	return s.Repository.FilterTicketByPrice(ctx, min, max)
}

// sortir tiket dari yang terbaru
func (s *TicketService) SortTicketByNewest(ctx context.Context) ([]*entity.Ticket, error) {
	return s.Repository.SortTicketByNewest(ctx)
}

// sortir dari yang termahal
func (s *TicketService) SortTicketByMostExpensive(ctx context.Context) ([]*entity.Ticket, error) {
	return s.Repository.SortTicketByMostExpensive(ctx)
}

// sortir tiket dari yang termurah
func (s *TicketService) SortTicketByCheapest(ctx context.Context) ([]*entity.Ticket, error) {
	return s.Repository.SortTicketByCheapest(ctx)
}

// filter ticket by most bought
func (s *TicketService) SortTicketByMostBought(ctx context.Context) ([]*entity.Ticket, error) {
	return s.Repository.SortTicketByMostBought(ctx)
}

// ticket yang masih tersedia
func (s *TicketService) SortTicketByAvailable(ctx context.Context) ([]*entity.Ticket, error) {
	return s.Repository.SortTicketByAvailable(ctx)
}