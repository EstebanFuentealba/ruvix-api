package profile

// MaritalStatus values
const (
	MaritalStatusSingle   = "single"
	MaritalStatusMarried  = "married"
	MaritalStatusDivorced = "divorced"
	MaritalStatusWidowed  = "widowed"
)

// Gender values
const (
	GenderMale   = "male"
	GenderFemale = "female"
)

// Profile ...
type Profile struct {
	ID string `json:"id"`

	UserID        string `json:"user_id"`
	Age           int    `json:"age"`
	Birth         int    `json:"birth"`
	MaritalStatus string `json:"marital_status"`
	Childs        int    `json:"childs"`
	Gender        string `json:"gender"`

	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt *int64 `json:"-"`
}

// Query ...
type Query struct {
	UserID string
}
