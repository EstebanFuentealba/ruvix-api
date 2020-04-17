package goals

import (
	"github.com/jmlopezz/uluru-api/database"
)

// Goals define goal model
type Goals struct {
	database.Base

	Name string `json:"name" gorm:"NOT NULL"`
}

// Retirement ...
type Retirement {
	database.Base

	GoalID Goal `json:"goal_id" gorm:"NOT NULL;foreignkey:GoalRefer"`
	MonthySalary float64 `json:"monthly_salary" gorm:""`
	MonthlyRetirement float64 `json:"monthly_retirement" gorm:""`
}