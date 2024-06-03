package entity

import (
	"time"
)

type Motor struct {
	ID                int       `json:"id"`
	OwnerID           int       `json:"owner_id"`
	MotorName         string    `json:"motor_name"`
	LicensePlate      string    `json:"license_plate"`
	RentalPricePerDay float64   `json:"rental_price_per_day"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty"` // Add DeletedAt field
}

// Admin New Motor
func NewMotor(ownerID int, motorName, licensePlate string, rentalPricePerDay float64, status string) *Motor {
	return &Motor{
		OwnerID:           ownerID,
		MotorName:         motorName,
		LicensePlate:      licensePlate,
		RentalPricePerDay: rentalPricePerDay,
		Status:            status,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}

// Admin Update Motor
func UpdateMotor(id, ownerID int, motorName, licensePlate string, rentalPricePerDay float64, status string) *Motor {
	return &Motor{
		ID:                id,
		OwnerID:           ownerID,
		MotorName:         motorName,
		LicensePlate:      licensePlate,
		RentalPricePerDay: rentalPricePerDay,
		Status:            status,
		UpdatedAt:         time.Now(),
	}
}

// Public Register Motor
func RegisterMotor(ownerID int, motorName, licensePlate string, rentalPricePerDay float64) *Motor {
	return &Motor{
		OwnerID:           ownerID,
		MotorName:         motorName,
		LicensePlate:      licensePlate,
		RentalPricePerDay: rentalPricePerDay,
		Status:            "Available",
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}

// User Update Motor by Self
func UpdateMotorProfile(id, ownerID int, motorName, licensePlate string, rentalPricePerDay float64) *Motor {
	return &Motor{
		ID:                id,
		OwnerID:           ownerID,
		MotorName:         motorName,
		LicensePlate:      licensePlate,
		RentalPricePerDay: rentalPricePerDay,
		UpdatedAt:         time.Now(),
	}
}

// Update the return type to be *Motor
func DeleteMotorByID(id int) *Motor {
	return &Motor{
		ID: id,
	}
}

// Update Motor Status
func UpdateMotorStatus(id int, status string) *Motor {
	return &Motor{
		ID:     id,
		Status: status,
	}
}