package savings

import (
	"github.com/jinzhu/gorm"
)

// InstitutionStore service definition
type InstitutionStore interface {
	ListInstitutions() ([]*Institution, error)
	CreateInstitution(p *Institution) (*Institution, error)
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

// ListInstitutions ...
func (is *institutionStoreDB) ListInstitutions() ([]*Institution, error) {
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

// CreateInstitution ...
func (is *institutionStoreDB) CreateInstitution(i *Institution) (*Institution, error) {
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
