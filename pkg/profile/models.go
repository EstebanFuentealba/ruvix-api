package profile

import (
	"time"

	"github.com/cagodoy/ruvix-api/internal/database"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Model ...
type Model struct {
	database.Base

	UserID        string `gorm:"not null;"`
	Fingerprint   string
	Age           int    `gorm:"not null;"`
	Birth         int    `gorm:"not null;"`
	MaritalStatus string `gorm:"not null;"`
	Childs        int    `gorm:"not null;DEFAULT:0"`
	Gender        string `gorm:"not null;"`
}

// TableName Set table name
func (Model) TableName() string {
	return "profiles"
}

// To ...
func (m *Model) To() *Profile {
	p := &Profile{
		UserID:        m.UserID,
		Fingerprint:   m.Fingerprint,
		Age:           m.Age,
		Birth:         m.Birth,
		MaritalStatus: m.MaritalStatus,
		Childs:        m.Childs,
		Gender:        m.Gender,
	}

	p.ID = m.Base.ID.String()
	p.CreatedAt = m.Base.CreatedAt.Unix()
	p.UpdatedAt = m.Base.UpdatedAt.Unix()

	return p
}

// From ...
func (m *Model) From(p *Profile) error {
	m.Base = database.Base{}

	if p.ID != "" {
		id, err := uuid.FromString(p.ID)
		if err != nil {
			return err
		}
		m.Base.ID = id
	}

	if p.CreatedAt != 0 {
		m.Base.CreatedAt = time.Unix(p.CreatedAt, 0)
	}

	if p.UpdatedAt != 0 {
		m.Base.UpdatedAt = time.Unix(p.UpdatedAt, 0)
	}

	m.UserID = p.UserID
	m.Fingerprint = p.Fingerprint
	m.Age = p.Age
	m.Birth = p.Birth
	m.MaritalStatus = p.MaritalStatus
	m.Childs = p.Childs
	m.Gender = p.Gender

	return nil
}

// RunMigrations ...
func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&Model{})
}
