package auth

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// AuthModel ...
type AuthModel struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;unique;default:gen_random_uuid()"`

	UserID    string `gorm:"NOT NULL"`
	Token     string `gorm:"NOT NULL"`
	Blacklist bool   `gorm:"NOT NULL"`
	Kind      string `gorm:"NOT NULL"`

	CreatedAt time.Time  `gorm:"default:now()"`
	UpdatedAt time.Time  `gorm:"default:now()"`
	DeletedAt *time.Time `sql:"index"`
}

// TableName Set table name
func (AuthModel) TableName() string {
	return "auth"
}

// To ...
func (am *AuthModel) To() *Auth {
	return &Auth{
		ID: am.ID.String(),

		UserID:    am.UserID,
		Token:     am.Token,
		Blacklist: am.Blacklist,
		Kind:      am.Kind,

		CreatedAt: am.CreatedAt,
		UpdatedAt: am.UpdatedAt,
	}
}

// From ...
func (am *AuthModel) From(a *Auth) error {
	if a.ID != "" {
		id, err := uuid.FromString(a.ID)
		if err != nil {
			return err
		}
		am.ID = id
	}

	am.CreatedAt = a.CreatedAt
	am.UpdatedAt = a.UpdatedAt

	am.UserID = a.UserID
	am.Token = a.Token
	am.Blacklist = a.Blacklist
	am.Kind = a.Kind

	return nil
}

// RunMigrations ...
func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&AuthModel{})
}
