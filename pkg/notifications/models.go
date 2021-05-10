package notifications

import (
	"github.com/cagodoy/ruvix-api/internal/database"
)

// Notification
type Notification struct {
	database.Base

	UserID string `json:"user_id" gorm:""`

	Text string `json:"text" gorm:"NOT NULL"`
	View bool   `json:"view" gorm:"NOT NULL"`
}
