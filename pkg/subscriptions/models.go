package subscriptions

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jmlopezz/uluru-api/internal/database"

	uuid "github.com/satori/go.uuid"
)

// SubscriptionModel ...
type SubscriptionModel struct {
	database.Base

	Features []*FeatureModel `gorm:"foreignkey:SubscriptionID"`

	Name   string  `gorm:"NOT NULL"`
	Price  float64 `gorm:"NOT NULL"`
	Months int     `gorm:"NOT NULL"`
}

// TableName Set table name
func (SubscriptionModel) TableName() string {
	return "subscriptions"
}

// To ...
func (sm *SubscriptionModel) To() *Subscription {
	i := &Subscription{
		ID: sm.Base.ID.String(),

		Features: make([]*Feature, 0),

		Name:   sm.Name,
		Price:  sm.Price,
		Months: sm.Months,

		CreatedAt: sm.Base.CreatedAt.Unix(),
		UpdatedAt: sm.Base.UpdatedAt.Unix(),
	}

	for j := 0; j < len(sm.Features); j++ {
		i.Features = append(i.Features, sm.Features[j].To())
	}

	return i
}

// From ...
func (sm *SubscriptionModel) From(s *Subscription) error {
	sm.Base = database.Base{}

	if s.ID != "" {
		id, err := uuid.FromString(s.ID)
		if err != nil {
			return err
		}
		sm.Base.ID = id
	}

	if s.CreatedAt != 0 {
		sm.Base.CreatedAt = time.Unix(s.CreatedAt, 0)
	}

	if s.UpdatedAt != 0 {
		sm.Base.UpdatedAt = time.Unix(s.UpdatedAt, 0)
	}

	featureModels := make([]*FeatureModel, 0)
	for j := 0; j < len(s.Features); j++ {
		featureModel := &FeatureModel{}
		err := featureModel.From(s.Features[j])
		if err != nil {
			return err
		}
		featureModels = append(featureModels, featureModel)
	}
	sm.Features = featureModels

	sm.Name = s.Name
	sm.Price = s.Price
	sm.Months = s.Months

	return nil
}

// FeatureModel ...
type FeatureModel struct {
	database.Base

	SubscriptionID uuid.UUID `gorm:"type:uuid;NOT NULL"`

	Text string `gorm:"NOT NULL"`
}

// TableName Set table name
func (FeatureModel) TableName() string {
	return "features"
}

// To ...
func (fm *FeatureModel) To() *Feature {
	return &Feature{
		ID: fm.Base.ID.String(),

		SubscriptionID: fm.SubscriptionID.String(),
		Text:           fm.Text,

		CreatedAt: fm.Base.CreatedAt.Unix(),
		UpdatedAt: fm.Base.UpdatedAt.Unix(),
	}
}

// From ...
func (fm *FeatureModel) From(f *Feature) error {
	fm.Base = database.Base{}

	if f.ID != "" {
		id, err := uuid.FromString(f.ID)
		if err != nil {
			return err
		}
		fm.Base.ID = id
	}

	if f.CreatedAt != 0 {
		fm.Base.CreatedAt = time.Unix(f.CreatedAt, 0)
	}

	if f.UpdatedAt != 0 {
		fm.Base.UpdatedAt = time.Unix(f.UpdatedAt, 0)
	}

	if f.SubscriptionID != "" {
		subscriptionID, err := uuid.FromString(f.SubscriptionID)
		if err != nil {
			return err
		}

		fm.SubscriptionID = subscriptionID
	}

	fm.Text = f.Text

	return nil
}

// TransactionModel ...
type TransactionModel struct {
	database.Base

	Subscription *SubscriptionModel `gorm:"foreignkey:Base.ID"`

	UserID         string `gorm:"NOT NULL"`
	ProviderID     string
	SubscriptionID uuid.UUID `gorm:"type:uuid;NOT NULL"`

	RemindedAt   time.Time `gorm:"NOT NULL"`
	DueDate      time.Time `gorm:"NOT NULL"`
	Status       string    `gorm:"NOT NULL"`
	PaymentToken string    `gorm:"NOT NULL"`
	OrderNumber  string    `gorm:"NOT NULL;unique"`
}

// TableName Set table name
func (TransactionModel) TableName() string {
	return "transactions"
}

// To ...
func (tm *TransactionModel) To() *Transaction {
	t := &Transaction{
		ID: tm.Base.ID.String(),

		UserID:         tm.UserID,
		ProviderID:     tm.ProviderID,
		SubscriptionID: tm.SubscriptionID.String(),

		RemindedAt:   tm.RemindedAt.Unix(),
		DueDate:      tm.DueDate.Unix(),
		Status:       tm.Status,
		PaymentToken: tm.PaymentToken,
		OrderNumber:  tm.OrderNumber,

		CreatedAt: tm.Base.CreatedAt.Unix(),
		UpdatedAt: tm.Base.UpdatedAt.Unix(),
	}

	if tm.Subscription != nil {
		t.Subscription = tm.Subscription.To()
	}

	return t
}

// From ...
func (tm *TransactionModel) From(t *Transaction) error {
	tm.Base = database.Base{}

	if t.ID != "" {
		id, err := uuid.FromString(t.ID)
		if err != nil {
			return err
		}
		tm.Base.ID = id
	}

	tm.Base.CreatedAt = time.Unix(t.CreatedAt, 0)
	tm.Base.UpdatedAt = time.Unix(t.UpdatedAt, 0)

	tm.DueDate = time.Unix(t.DueDate, 0)
	tm.RemindedAt = time.Unix(t.RemindedAt, 0)

	if t.SubscriptionID != "" {
		subscriptionID, err := uuid.FromString(t.SubscriptionID)
		if err != nil {
			return err
		}

		tm.SubscriptionID = subscriptionID
	}

	if t.Subscription != nil {
		model := &SubscriptionModel{}
		err := model.From(t.Subscription)
		if err != nil {
			return err
		}
		tm.Subscription = model
	}

	tm.UserID = t.UserID
	tm.ProviderID = t.ProviderID
	tm.Status = t.Status
	tm.PaymentToken = t.PaymentToken
	tm.OrderNumber = t.OrderNumber

	return nil
}

// QuerySubscription ...
type QuerySubscription struct {
	ID    string
	Price float64
}

// QueryTransaction ...
type QueryTransaction struct {
	ID             string
	UserID         string
	SubscriptionID string
	ProviderID     string
	Status         string
	DueDate        time.Time
}

// RunMigrations ...
func RunMigrations(db *gorm.DB) error {
	runSubcriptionSeed := false
	if !db.HasTable(&SubscriptionModel{}) {
		err := db.CreateTable(&SubscriptionModel{}).Error
		if err != nil {
			log.Fatalln(err)
		}

		runSubcriptionSeed = true
	}

	if !db.HasTable(&FeatureModel{}) {
		err := db.CreateTable(&FeatureModel{}).Error
		if err != nil {
			log.Fatalln(err)
		}
	}

	if !db.HasTable(&TransactionModel{}) {
		err := db.CreateTable(&TransactionModel{}).Error
		if err != nil {
			log.Fatalln(err)
		}
	}

	if runSubcriptionSeed {
		err := seedSubscription(db)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return nil
}
