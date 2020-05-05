package profile

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Store service definition
type Store interface {
	Create(p *Profile) (*Profile, error)
	Get(q *Query) (*Profile, error)
	Update(p *Profile) (*Profile, error)
}

// storeDB ...
type storeDB struct {
	DB *gorm.DB
}

// NewStore ...
func NewStore(db *gorm.DB) Store {
	return &storeDB{
		DB: db,
	}
}

// Create ...
func (s *storeDB) Create(p *Profile) (*Profile, error) {
	model := &Model{}

	err := model.From(p)
	if err != nil {
		return nil, err
	}

	_, err = s.Get(&Query{
		UserID: p.UserID,
	})
	if err == nil {
		err = errors.New("user already has profile")
		return nil, err
	}

	err = s.DB.Create(model).Error
	if err != nil {
		return nil, err
	}

	return model.To(), nil
}

// Get ...
func (s *storeDB) Get(q *Query) (*Profile, error) {
	model := &Model{}

	err := s.DB.Where("user_id = ?", q.UserID).Order("created_at DESC").Limit(1).First(model).Error
	if err != nil {
		return nil, err
	}

	return model.To(), nil
}

// Update ...
func (s *storeDB) Update(p *Profile) (*Profile, error) {
	model := &Model{}
	model.From(p)

	err := s.DB.Save(model).Error
	if err != nil {
		return nil, err
	}

	return model.To(), nil
}
