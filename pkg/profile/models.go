package profile

import (
	"github.com/jmlopezz/uluru-api/database"
)

// Profile ...
type Profile struct {
	database.Base

	UserID        string        `json:"user_id" gorm:"NOT NULL"`
	Age           int           `json:"age" gorm:"NOT NULL"`
	Birth         int           `json:"birth" gorm:"NOT NULL"`
	MaritalStatus MaritalStatus `json:"marital_status" gorm:"NOT NULL"  sql:"enum('single', 'married', 'divorced')"`
	Childs        int           `json:"childs" gorm:"NOT NULL;DEFAULT:0"`
}
