package notifications

import (
	"github.com/jmlopezz/uluru-api/internal/database"
)

// Notification
Type Notification struct {
	database.Base
	
	UserID     string   `json:"user_id" gorm:""`

	Text string `json:"text" gorm:"NOT NULL"`
	View bool `json:"view" gorm:"NOT NULL"`
}