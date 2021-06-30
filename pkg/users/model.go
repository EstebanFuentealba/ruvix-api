package users

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// UserModel ...
type UserModel struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;unique;default:gen_random_uuid()"`

	Name     string `gorm:"NOT NULL"`
	Email    string `gorm:"NOT NULL;unique"`
	Password string `gorm:"NOT NULL"`
	Verified bool   `gorm:"default:false"`

	CreatedAt time.Time  `gorm:"default:now()"`
	UpdatedAt time.Time  `gorm:"default:now()"`
	DeletedAt *time.Time `sql:"index"`
}

// // BeforeCreate will set a UUID rather than numeric ID.
// func (UserModel) BeforeCreate(scope *gorm.Scope) error {
// 	uuid, err := uuid.NewV4()
// 	if err != nil {
// 		return err
// 	}
// 	return scope.SetColumn("ID", uuid)
// }

// TableName Set table name
func (UserModel) TableName() string {
	return "users"
}

// To ...
func (um *UserModel) To() *User {
	u := &User{
		ID: um.ID.String(),

		Name:     um.Name,
		Email:    um.Email,
		Password: um.Password,

		CreatedAt: um.CreatedAt,
		UpdatedAt: um.UpdatedAt,
	}

	return u
}

// From ...
func (um *UserModel) From(u *User) error {
	if u.ID != "" {
		id, err := uuid.FromString(u.ID)
		if err != nil {
			return err
		}
		um.ID = id
	}

	um.CreatedAt = u.CreatedAt
	um.UpdatedAt = u.UpdatedAt

	um.Name = u.Name
	um.Email = u.Email
	um.Password = u.Password

	return nil
}

// RunMigrations ...
func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&UserModel{})
}
