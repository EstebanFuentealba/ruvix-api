package database

import (
	"github.com/jinzhu/gorm"
	// Gorm internal use
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewPostgres ...
func NewPostgres(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	// defer db.Close()

	return db, nil
}
