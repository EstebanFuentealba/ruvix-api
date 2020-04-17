package database

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Base struct used in all models
type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;not null;unique;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"`
}

// connection to db method
