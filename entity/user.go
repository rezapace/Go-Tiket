package entity

import (
	"time"
)

type User struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Number    string     `json:"number"`
	Role      string     `json:"role"`
	Balance   float64    `json:"balance"`
	Password  string     `json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
// Admin New User
func NewUser(name, email, number, password, role string, balance int64) *User {
	return &User{
		Name:      name,
		Email:     email,
		Number:    number,
		Role:     role,
		Balance:     balance,
		Password:  password,
		CreatedAt: time.Now(),
	}
}

// Admin Update User
func UpdateUser(id int64, name, email, number, role, password string, balance int64) *User {
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Number:    number,
		Role:     role,
		Password:  password,
		Balance:     balance,
		UpdatedAt: time.Now(),
	}
}

// Public Register
func Register(email, password, role, number string) *User {
	return &User{
		Email:    email,
		Password: password,
		Role:    role,
		Number:   number,
	}
}

// user update by self
func UpdateProfile(id int64, name, email, number, password string) *User {
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Number:    number,
		Password:  password,
		UpdatedAt: time.Now(),
	}
}

// Update the return type to be *User
func DeleteUserSelfByEmail(email string) *User {
	return &User{
		Email:     email,
		DeletedAt: nil,
	}
}

func UpgradeBalance(id int64, balance int64) *User {
	return &User{
		ID:    id,
		Balance: balance,
	}
}

// user logout
func UserLogout(id int64) *User {
	return &User{
		ID: id,
	}
}

// updatesaldo
func UpdateBalance(id int64, balance int64) *User {
	return &User{
		ID:    id,
		Balance: balance,
	}
}
