package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// NewService ...
func NewService(store Store, events *Events) Service {
	return &UserService{
		Store:  store,
		Events: events,
	}
}

// UserService ...
type UserService struct {
	Store  Store
	Events *Events
}

// GetByID ...
func (us *UserService) GetByID(id string) (*User, error) {
	q := &Query{
		ID: id,
	}

	return us.Store.Get(q)
}

// GetByEmail ...
func (us *UserService) GetByEmail(email string) (*User, error) {
	q := &Query{
		Email: email,
	}

	return us.Store.Get(q)
}

// Create ...
func (us *UserService) Create(u *User) error {
	// before create user event
	if us.Events.BeforeCreate != nil {
		err := us.Events.BeforeCreate()
		if err != nil {
			return err
		}
	}

	err := us.Store.Create(u)
	if err != nil {
		return err
	}

	// after create user event
	if us.Events.AfterCreate != nil {
		err := us.Events.AfterCreate()
		if err != nil {
			return err
		}
	}

	return nil
}

// Update ...
func (us *UserService) Update(u *User) error {
	return us.Store.Update(u)
}

// Delete ...
func (us *UserService) Delete(u *User) error {
	return us.Store.Delete(u)
}

// List ...
func (us *UserService) List() ([]*User, error) {
	return us.Store.List()
}

// VerifyPassword ...
func (us *UserService) VerifyPassword(email string, password string) error {
	user, err := us.GetByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.New("invalid password")
	}

	return nil
}
