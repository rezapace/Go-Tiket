package entity

import (
	"time"
)

type Notification struct {
	ID         int       `json:"id"`
	Type       string    `json:"type"`
	Message    string    `json:"message"`
	IsRead     bool      `json:"is_read"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

func NewNotification(id int, typ, message string, isRead bool, createdAt, updatedAt, deletedAt time.Time) *Notification {
	return &Notification{
		ID:        id,
		Type:      typ,
		Message:   message,
		IsRead:    isRead,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}