package institutions

import "database/sql/driver"

// Type ...
type Type string

const (
	Investment  Type = "investment"
	Inmobiliary Type = "inmobiliary"
)

// Scan ...
func (e *Type) Scan(value interface{}) error {
	*e = Type(value.([]byte))
	return nil
}

// Value ...
func (e Type) Value() (driver.Value, error) {
	return string(e), nil
}
