package profile

import "database/sql/driver"

// Status ...
type Status string

const (
	pending   Status = "pending"
	completed Status = "completed"
	reject    Status = "reject"
)

// Scan ...
func (e *Status) Scan(value interface{}) error {
	*e = Status(value.([]byte))
	return nil
}

// Value ...
func (e Status) Value() (driver.Value, error) {
	return string(e), nil
}
