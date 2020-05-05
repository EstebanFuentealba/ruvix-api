package savings

import (
	"github.com/jinzhu/gorm"
)

// InstitutionStore service definition
type InstitutionStore interface {
	List() ([]*Institution, error)
	Create(p *Institution) (*Institution, error)
}

type institutionStoreDB struct {
	DB *gorm.DB
}

// NewInstitutionStore ...
func NewInstitutionStore(db *gorm.DB) InstitutionStore {
	return &institutionStoreDB{
		DB: db,
	}
}

// List ...
func (is *institutionStoreDB) List() ([]*Institution, error) {
	institutionModels := make([]*InstitutionModel, 0)

	// Get institutions
	err := is.DB.Find(&institutionModels).Error
	if err != nil {
		return nil, err
	}

	// Get accounts
	institutions := make([]*Institution, 0)
	for i := 0; i < len(institutionModels); i++ {
		institutionModels[i].Accounts = make([]*AccountModel, 0)
		err := is.DB.Where("institution_id = ?", institutionModels[i].ID).Find(&institutionModels[i].Accounts).Error
		if err != nil {
			return nil, err
		}

		// Get instruments
		for j := 0; j < len(institutionModels[i].Accounts); j++ {
			institutionModels[i].Accounts[j].Instruments = make([]*InstrumentModel, 0)
			err := is.DB.Where("account_id = ?", institutionModels[i].Accounts[j].ID).Find(&institutionModels[i].Accounts[j].Instruments).Error
			if err != nil {
				return nil, err
			}
		}

		institutions = append(institutions, institutionModels[i].To())
	}

	return institutions, nil
}

// Create ...
func (is *institutionStoreDB) Create(i *Institution) (*Institution, error) {
	model := &InstitutionModel{}

	err := model.From(i)
	if err != nil {
		return nil, err
	}

	// Create institution
	err = is.DB.Create(&model).Error
	if err != nil {
		return nil, err
	}

	return model.To(), nil
}

// AccountStore service definition
type AccountStore interface {
	// some actions..
}

type accountStoreDB struct {
	DB *gorm.DB
}

// NewAccountStore ...
func NewAccountStore(db *gorm.DB) AccountStore {
	return &accountStoreDB{
		DB: db,
	}
}

// InstrumentStore service definition
type InstrumentStore interface {
	// some actions..
}

type instrumentStoreDB struct {
	DB *gorm.DB
}

// NewInstrumentStore ...
func NewInstrumentStore(db *gorm.DB) InstrumentStore {
	return &instrumentStoreDB{
		DB: db,
	}
}

// RetirementInstrumentStore service definition
type RetirementInstrumentStore interface {
	// some actions...
}

type retirementInstrumentStoreDB struct {
	DB *gorm.DB
}

// NewRetirementInstrumentStore ...
func NewRetirementInstrumentStore(db *gorm.DB) RetirementInstrumentStore {
	return &retirementInstrumentStoreDB{
		DB: db,
	}
}
