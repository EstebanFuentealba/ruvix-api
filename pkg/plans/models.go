package plans

import (
	"time"

	"github.com/jmlopezz/uluru-api/database"
)

// Plan ...
type Plan struct {
	database.Base

	Name  string `json:"name" gorm:"NOT NULL"`
	Price string `json:"price" gorm:"NOT NULL"`
}

// CREATE TABLE plan_features (
// 	id uuid DEFAULT gen_reandom_uuid() UNIQUE,
// 	plan_id uuid NOT NULL references plans(id),
// 	feature_id: uuid NOT NULL references(features)
// );

type Feature struct {
	database.Base

	Description string `json:"description" gorm:"NOT NULL"`
}

// Provider ...
type Provider struct {
	database.Base

	Name string `json:"name" gorm:"NOT NULL"`
}

// Transaction ...
type Transaction struct {
	database.Base

	UserID     string   `json:"user_id" gorm:""`
	ProviderID Provider `json:"provider" gorm:"foreignkey:ProviderRefer"`
	PlanID     Plan     `json:"provider" gorm:"foreignkey:PlanRefer"`

	Amount      float64   `json:"amount" gorm:""`
	Detail      string    `json:"detail" gorm:""`
	RemindedAt  time.Time `json:"reminded_at"`
	DueDate     string    `json:"due_date"`
	Status      Status    `json:"status" gorm:"NOT NULL" sql:"enum('pending', 'completed', 'rejected')"`
	Quota       int       `json:"quota" gorm:"NOT NULL"`
	TotalQuotas int       `json:"total_quotas" gorm:"NOT NULL"`
}
