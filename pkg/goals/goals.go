package goals

import "github.com/cagodoy/ruvix-api/pkg/savings"

// Goal define goal
type Goal struct {
	ID string `json:"id"`

	Name string `json:"name"`

	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt *int64 `json:"-"`
}

// RetirementGoal ...
type RetirementGoal struct {
	ID string `json:"id"`

	Goal                  *Goal                           `json:"goal,omitempty"`
	RetirementInstruments []*savings.RetirementInstrument `json:"retirement_instruments,omitempty"`

	UserID            string  `json:"user_id,omitempty"`
	Fingerprint       string  `json:"fingerprint,omitempty"`
	GoalID            string  `json:"goal_id"`
	MonthlySalary     float64 `json:"monthly_salary"`
	MonthlyRetirement float64 `json:"monthly_retirement"`

	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt *int64 `json:"-"`
}

// GoalQuery ...
type GoalQuery struct {
	ID string
}

// RetirementGoalQuery ...
type RetirementGoalQuery struct {
	Fingerprint string
	UserID      string
}
