package savings

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jmlopezz/uluru-api/internal/database"
	uuid "github.com/satori/go.uuid"
)

// InstitutionModel ...
type InstitutionModel struct {
	database.Base

	Accounts []*AccountModel `gorm:"foreignkey:InstitutionID"`

	Name string `gorm:"NOT NULL"`
	Slug string `gorm:"NOT NULL;unique"`
}

// TableName Set table name
func (InstitutionModel) TableName() string {
	return "institutions"
}

// To ...
func (im *InstitutionModel) To() *Institution {
	i := &Institution{
		ID: im.Base.ID.String(),

		Accounts: make([]*Account, 0),

		Name: im.Name,
		Slug: im.Slug,

		CreatedAt: im.Base.CreatedAt.Unix(),
		UpdatedAt: im.Base.UpdatedAt.Unix(),
	}

	for j := 0; j < len(im.Accounts); j++ {
		i.Accounts = append(i.Accounts, im.Accounts[j].To())
	}

	return i
}

// From ...
func (im *InstitutionModel) From(i *Institution) error {
	im.Base = database.Base{}

	if i.ID != "" {
		id, err := uuid.FromString(i.ID)
		if err != nil {
			return err
		}
		im.Base.ID = id
	}

	if i.CreatedAt != 0 {
		im.Base.CreatedAt = time.Unix(i.CreatedAt, 0)
	}

	if i.UpdatedAt != 0 {
		im.Base.UpdatedAt = time.Unix(i.UpdatedAt, 0)
	}

	accountModels := make([]*AccountModel, 0)
	for j := 0; j < len(i.Accounts); j++ {
		accountModel := &AccountModel{}
		err := accountModel.From(i.Accounts[j])
		if err != nil {
			return err
		}
		accountModels = append(accountModels, accountModel)
	}
	im.Accounts = accountModels

	im.Name = i.Name
	im.Slug = i.Slug

	return nil
}

// AccountModel ...
type AccountModel struct {
	database.Base

	Instruments []*InstrumentModel `gorm:"foreignkey:AccountID"`

	InstitutionID uuid.UUID `gorm:"type:uuid;NOT NULL"`
	Name          string    `gorm:"NOT NULL"`
	Slug          string    `gorm:"NOT NULL"`
}

// TableName Set table name
func (AccountModel) TableName() string {
	return "accounts"
}

// To ...
func (am *AccountModel) To() *Account {
	a := &Account{
		ID: am.Base.ID.String(),

		Instruments:   make([]*Instrument, 0),
		InstitutionID: am.InstitutionID.String(),
		Name:          am.Name,
		Slug:          am.Slug,

		CreatedAt: am.Base.CreatedAt.Unix(),
		UpdatedAt: am.Base.UpdatedAt.Unix(),
	}

	for j := 0; j < len(am.Instruments); j++ {
		a.Instruments = append(a.Instruments, am.Instruments[j].To())
	}

	return a
}

// From ...
func (am *AccountModel) From(a *Account) error {
	am.Base = database.Base{}

	if a.ID != "" {
		id, err := uuid.FromString(a.ID)
		if err != nil {
			return err
		}
		am.Base.ID = id
	}

	if a.CreatedAt != 0 {
		am.Base.CreatedAt = time.Unix(a.CreatedAt, 0)
	}

	if a.UpdatedAt != 0 {
		am.Base.UpdatedAt = time.Unix(a.UpdatedAt, 0)
	}

	if a.InstitutionID != "" {
		institutionID, err := uuid.FromString(a.InstitutionID)
		if err != nil {
			return err
		}

		am.InstitutionID = institutionID
	}

	instrumentsModels := make([]*InstrumentModel, 0)
	for i := 0; i < len(a.Instruments); i++ {
		instrumentModel := &InstrumentModel{}
		err := instrumentModel.From(a.Instruments[i])
		if err != nil {
			return err
		}
		instrumentsModels = append(instrumentsModels, instrumentModel)
	}
	am.Instruments = instrumentsModels

	am.Name = a.Name
	am.Slug = a.Slug

	return nil
}

// InstrumentModel ...
type InstrumentModel struct {
	database.Base

	AccountID          uuid.UUID `gorm:"type:uuid;NOT NULL"`
	Name               string    `gorm:"NOT NULL"`
	Slug               string    `gorm:"NOT NULL"`
	Return1m           float64   `gorm:"NOT NULL"`
	Return1y           float64   `gorm:"NOT NULL"`
	Return5y           float64   `gorm:"NOT NULL"`
	Return10y          float64   `gorm:"NOT NULL"`
	ProjectedWorstCase float64   `gorm:"NOT NULL"`
	ProjectedAvgCase   float64   `gorm:"NOT NULL"`
	ProjectedBestCase  float64   `gorm:"NOT NULL"`
}

// TableName Set table name
func (InstrumentModel) TableName() string {
	return "intruments"
}

// To ...
func (im *InstrumentModel) To() *Instrument {
	return &Instrument{
		ID: im.Base.ID.String(),

		AccountID:          im.AccountID.String(),
		Name:               im.Name,
		Slug:               im.Slug,
		Return1m:           im.Return1m,
		Return1y:           im.Return1y,
		Return5y:           im.Return5y,
		Return10y:          im.Return10y,
		ProjectedWorstCase: im.ProjectedWorstCase,
		ProjectedAvgCase:   im.ProjectedAvgCase,
		ProjectedBestCase:  im.ProjectedBestCase,

		CreatedAt: im.Base.CreatedAt.Unix(),
		UpdatedAt: im.Base.UpdatedAt.Unix(),
	}
}

// From ...
func (im *InstrumentModel) From(i *Instrument) error {
	im.Base = database.Base{}

	if i.ID != "" {
		id, err := uuid.FromString(i.ID)
		if err != nil {
			return err
		}
		im.Base.ID = id
	}

	if i.CreatedAt != 0 {
		im.Base.CreatedAt = time.Unix(i.CreatedAt, 0)
	}

	if i.UpdatedAt != 0 {
		im.Base.UpdatedAt = time.Unix(i.UpdatedAt, 0)
	}

	if i.AccountID != "" {
		accountID, err := uuid.FromString(i.AccountID)
		if err != nil {
			return err
		}

		im.AccountID = accountID
	}

	im.Name = i.Name
	im.Slug = i.Slug
	im.Return1m = i.Return1m
	im.Return1y = i.Return1y
	im.Return5y = i.Return5y
	im.Return10y = i.Return10y
	im.ProjectedWorstCase = i.ProjectedWorstCase
	im.ProjectedAvgCase = i.ProjectedAvgCase
	im.ProjectedBestCase = i.ProjectedBestCase

	return nil
}

// RetirementInstrumentModel ...
type RetirementInstrumentModel struct {
	database.Base

	Instrument *InstrumentModel `gorm:"foreignkey:Base.ID"`

	InstrumentID     uuid.UUID `gorm:"type:uuid;NOT NULL"`
	RetirementGoalID uuid.UUID `gorm:"type:uuid;NOT NULL"`
	UserID           string
	Fingerprint      string
	Percent          float64 `gorm:"NOT NULL"`
	QuotasQuantity   float64 `gorm:"NOT NULL"`
	QuotasDate       string  `gorm:"NOT NULL"`
	QuotasPrice      float64 `gorm:"NOT NULL"`
	Balance          float64 `gorm:"NOT NULL"`
}

// TableName Set table name
func (RetirementInstrumentModel) TableName() string {
	return "retirement_instruments"
}

// To ...
func (rim *RetirementInstrumentModel) To() *RetirementInstrument {
	ri := &RetirementInstrument{
		ID:               rim.Base.ID.String(),
		InstrumentID:     rim.InstrumentID.String(),
		RetirementGoalID: rim.RetirementGoalID.String(),
		UserID:           rim.UserID,
		Fingerprint:      rim.Fingerprint,
		Percent:          rim.Percent,
		QuotasQuantity:   rim.QuotasQuantity,
		QuotasDate:       rim.QuotasDate,
		QuotasPrice:      rim.QuotasPrice,
		Balance:          rim.Balance,

		CreatedAt: rim.Base.CreatedAt.Unix(),
		UpdatedAt: rim.Base.UpdatedAt.Unix(),
	}

	if rim.Instrument != nil {
		ri.Instrument = rim.Instrument.To()
	}

	return ri
}

// From ...
func (rim *RetirementInstrumentModel) From(ri *RetirementInstrument) error {
	rim.Base = database.Base{}

	if ri.ID != "" {
		id, err := uuid.FromString(ri.ID)
		if err != nil {
			return err
		}
		rim.Base.ID = id
	}

	if ri.CreatedAt != 0 {
		rim.Base.CreatedAt = time.Unix(ri.CreatedAt, 0)
	}

	if ri.UpdatedAt != 0 {
		rim.Base.UpdatedAt = time.Unix(ri.UpdatedAt, 0)
	}

	if ri.InstrumentID != "" {
		instrumentID, err := uuid.FromString(ri.InstrumentID)
		if err != nil {
			return err
		}

		rim.InstrumentID = instrumentID
	}

	if ri.RetirementGoalID != "" {
		retirementGoalID, err := uuid.FromString(ri.RetirementGoalID)
		if err != nil {
			return err
		}

		rim.RetirementGoalID = retirementGoalID
	}

	if ri.Instrument != nil {
		model := &InstrumentModel{}
		err := model.From(ri.Instrument)
		if err != nil {
			return err
		}
		rim.Instrument = model
	}

	rim.UserID = ri.UserID
	rim.Fingerprint = ri.Fingerprint
	rim.Percent = ri.Percent
	rim.QuotasQuantity = ri.QuotasQuantity
	rim.QuotasDate = ri.QuotasDate
	rim.QuotasPrice = ri.QuotasPrice
	rim.Balance = ri.Balance

	return nil
}

// RetirementInstrumentQuery ...
type RetirementInstrumentQuery struct {
	ID string
}

// RunMigrations ...
func RunMigrations(db *gorm.DB) error {
	runInstitutionSeed := false
	if !db.HasTable(&InstitutionModel{}) {
		err := db.CreateTable(&InstitutionModel{}).Error
		if err != nil {
			log.Fatalln(err)
		}

		runInstitutionSeed = true
	}

	if !db.HasTable(&AccountModel{}) {
		err := db.CreateTable(&AccountModel{}).Error
		if err != nil {
			log.Fatalln(err)
		}
	}

	if !db.HasTable(&AccountModel{}) {
		err := db.CreateTable(&AccountModel{}).Error
		if err != nil {
			log.Fatalln(err)
		}
	}

	if !db.HasTable(&InstrumentModel{}) {
		err := db.CreateTable(&InstrumentModel{}).Error
		if err != nil {
			log.Fatalln(err)
		}
	}

	if !db.HasTable(&RetirementInstrumentModel{}) {
		err := db.CreateTable(&RetirementInstrumentModel{}).Error
		if err != nil {
			log.Fatalln(err)
		}
	}

	if runInstitutionSeed {
		err := seedInstitution(db)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return nil
}
