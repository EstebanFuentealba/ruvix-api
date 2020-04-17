package profile

import "database/sql/driver"

// MaritalStatus ...
type MaritalStatus string

const (
	single   MaritalStatus = "single"
	married  MaritalStatus = "married"
	divorced MaritalStatus = "divorced"
	widowed  MaritalStatus = "widowed"
)

// Scan ...
func (e *MaritalStatus) Scan(value interface{}) error {
	*e = MaritalStatus(value.([]byte))
	return nil
}

// Value ...
func (e MaritalStatus) Value() (driver.Value, error) {
	return string(e), nil
}
