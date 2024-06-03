package entity

import (
    "time"
)

// Topup struct menyimpan data top-up
type Topup struct {
    ID        string    `json:"id" db:"id"`
    UserID    int       `json:"user_id" db:"user_id"`
    Amount    int       `json:"amount" db:"amount"`
    Status    int       `json:"status" db:"status"`
    SnapURL   string    `json:"snap_url" db:"snap_url"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// TopupStatus struct menyimpan data status top-up
type NewTopup struct {
    ID        string    `json:"id" db:"id"`
    UserID    int       `json:"user_id" db:"user_id"`
    Amount    int       `json:"amount" db:"amount"`
    Status    int       `json:"status" db:"status"`
    SnapURL   string    `json:"snap_url" db:"snap_url"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}