package users

import (
	"time"
)

// User ...
type User struct {
	ID string `json:"id,omitempty" db:"id"`

	Email    string `json:"email,omitempty" db:"email"`
	Name     string `json:"name,omitempty" db:"name"`
	Password string `json:"password,omitempty" db:"password"`
	Verified bool   `json:"verified" db:"verified"`

	CreatedAt time.Time  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

// Service ...
type Service interface {
	GetByID(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	Create(*User) error
	List() ([]*User, error)
	Update(*User) error
	Delete(*User) error
	VerifyPassword(email string, password string) error
}

// Query ...
type Query struct {
	ID    string
	Email string
}

// Events ...
type Events struct {
	BeforeCreate func() error
	AfterCreate  func() error

	// TODO(ca): implements all events
}
