package subscriptions

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/teris-io/shortid"
)

const (
	daysAfterDueDate = 3
)

// SubscriptionStore service definition
type SubscriptionStore interface {
	List() ([]*Subscription, error)
	Create(i *Subscription) (*Subscription, error)
	Get(q Query) (*Subscription, error)
	Subscribe(q Query) (*Transaction, error)
	Unsubscribe(q Query) (*Transaction, error)
	Refresh(q Query) (*Transaction, error)
}

type subscriptionStoreDB struct {
	DB *gorm.DB
}

// NewSubscriptionStore ...
func NewSubscriptionStore(db *gorm.DB) SubscriptionStore {
	return &subscriptionStoreDB{
		DB: db,
	}
}

// List ...
func (ss *subscriptionStoreDB) List() ([]*Subscription, error) {
	subscriptionModels := make([]*SubscriptionModel, 0)

	// get subscriptions
	err := ss.DB.Find(&subscriptionModels).Error
	if err != nil {
		return nil, err
	}

	// Get accounts
	subscriptions := make([]*Subscription, 0)
	for i := 0; i < len(subscriptionModels); i++ {
		subscriptionModels[i].Features = make([]*FeatureModel, 0)
		err := ss.DB.Where("subscription_id = ?", subscriptionModels[i].ID).Find(&subscriptionModels[i].Features).Error
		if err != nil {
			return nil, err
		}

		subscriptions = append(subscriptions, subscriptionModels[i].To())
	}

	return subscriptions, nil
}

// Create ...
func (ss *subscriptionStoreDB) Create(i *Subscription) (*Subscription, error) {
	model := &SubscriptionModel{}

	err := model.From(i)
	if err != nil {
		return nil, err
	}

	// Create institution
	err = ss.DB.Create(&model).Error
	if err != nil {
		return nil, err
	}

	return model.To(), nil
}

// Get ...
func (ss *subscriptionStoreDB) Get(q Query) (*Subscription, error) {
	subscriptionModel := &SubscriptionModel{}

	// get subscriptions
	err := ss.DB.Where("id = ?", q.ID).Find(subscriptionModel).Error
	if err != nil {
		return nil, err
	}

	return subscriptionModel.To(), nil
}

// Subscribe ...
func (ss *subscriptionStoreDB) Subscribe(q Query) (*Transaction, error) {
	// get last transaction
	transactionModel := &TransactionModel{}
	err := ss.DB.Where("user_id = ?", q.UserID).Order("created_at DESC").Limit(1).First(transactionModel).Error
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}

	now := time.Now()

	// get subscription
	transactionModel.Subscription = &SubscriptionModel{}
	err = ss.DB.Where("id = ?", q.SubscriptionID).Find(transactionModel.Subscription).Error
	if err != nil {
		return nil, err
	}

	// case: first user's transaction
	if err == nil {
		lastTransaction := transactionModel.To()

		// if status not completed or canceled
		if lastTransaction.Status == StatusTransactionPending || lastTransaction.Status == StatusTransactionRejected {
			err := errors.New("last transaction its pending or rejected status")
			return nil, err
		}

		// if last provider and provider_id its the same
		if lastTransaction.ProviderID == q.ProviderID && lastTransaction.ProviderID == ProviderFree {
			// its same with last transaction
			if lastTransaction.SubscriptionID == q.SubscriptionID && lastTransaction.ProviderID == q.ProviderID {
				err := errors.New("new transaction ​​are the same as the last transaction")
				return nil, err
			}
		}
	}

	// validate subscription.price respect provider_id
	if transactionModel.Subscription.Price != 0 {
		if q.ProviderID == "free" {
			return nil, fmt.Errorf("provider_id cannot be %s when subscription price isn't free", q.ProviderID)
		}
	} else if transactionModel.Subscription.Price == 0 && q.ProviderID != "free" {
		return nil, fmt.Errorf("provider_id cannot be %s when subscription price is free", q.ProviderID)
	}

	orderNumber, err := shortid.Generate()
	if err != nil {
		return nil, err
	}

	// prepare transaction
	t := &Transaction{
		UserID:         q.UserID,
		SubscriptionID: q.SubscriptionID,
		ProviderID:     q.ProviderID,
		OrderNumber:    orderNumber,
		CreatedAt:      now.Unix(),
		UpdatedAt:      now.Unix(),
	}

	if transactionModel.Subscription.Months == 0 {
		t.ProviderID = ProviderFree
		t.Status = StatusTransactionCompleted
		t.DueDate = 0
		t.RemindedAt = 0
	} else {
		t.ProviderID = q.ProviderID
		t.Status = StatusTransactionPending
		t.DueDate = now.AddDate(0, transactionModel.Subscription.Months, 0).Unix()
		t.RemindedAt = now.AddDate(0, transactionModel.Subscription.Months, 0).AddDate(0, 0, -daysAfterDueDate).Unix()
	}

	// parse transaction to model
	model := &TransactionModel{}
	err = model.From(t)
	if err != nil {
		return nil, err
	}

	// create transaction
	err = ss.DB.Create(&model).Error
	if err != nil {
		return nil, err
	}

	model.Subscription = transactionModel.Subscription

	return model.To(), nil
}

// Unsubscribe ...
func (ss *subscriptionStoreDB) Unsubscribe(q Query) (*Transaction, error) {
	// get last transaction
	model := &TransactionModel{}
	err := ss.DB.Where("user_id = ?", q.UserID).Order("updated_at DESC").Limit(1).First(model).Error
	if err != nil {
		return nil, err
	}

	// if transaction is complete, return with no changes
	if model.Status == StatusTransactionCompleted || model.Status == StatusTransactionCanceled {
		err := errors.New("transaction already completed or cancelled")
		return nil, err
	}

	model.Status = StatusTransactionCanceled

	// save transaction
	err = ss.DB.Save(&model).Error
	if err != nil {
		return nil, err
	}

	// get subscription
	model.Subscription = &SubscriptionModel{}
	err = ss.DB.Where("id = ?", q.SubscriptionID).Find(model.Subscription).Error
	if err != nil {
		return nil, err
	}

	return model.To(), nil
}

// Refresh ...
func (ss *subscriptionStoreDB) Refresh(q Query) (*Transaction, error) {
	// get last transaction
	transactionModel := &TransactionModel{}
	err := ss.DB.Where("user_id = ?", q.UserID).Order("created_at DESC").Limit(1).First(transactionModel).Error
	if err != nil {
		return nil, err
	}
	t := transactionModel.To()

	// get subscription
	subscriptionModel := &SubscriptionModel{}
	err = ss.DB.Where("id = ?", q.SubscriptionID).Find(subscriptionModel).Error
	if err != nil {
		return nil, err
	}

	// if transaction is complete, return with no changes
	if t.Status == StatusTransactionCompleted || t.Status == StatusTransactionCanceled {
		err := errors.New("transaction already completed or cancelled")
		return nil, err
	}

	// validate subscription.price respect provider_id
	if subscriptionModel.Price != 0 {
		if q.ProviderID == "free" {
			return nil, fmt.Errorf("provider_id cannot be %s when subscription price isn't free", q.ProviderID)
		}
	} else if subscriptionModel.Price == 0 && q.ProviderID != "free" {
		return nil, fmt.Errorf("provider_id cannot be %s when subscription price is free", q.ProviderID)
	}

	// prepare transaction
	t.ProviderID = q.ProviderID
	if t.Status == StatusTransactionRejected {
		t.Status = StatusTransactionPending
	}

	// parse transaction to model
	model := &TransactionModel{}
	err = model.From(t)
	if err != nil {
		return nil, err
	}

	// save transaction
	err = ss.DB.Save(&model).Error
	if err != nil {
		return nil, err
	}

	model.Subscription = subscriptionModel

	return model.To(), nil
}

// ProviderStore service definition
type ProviderStore interface {
	List() ([]*Provider, error)
}

type providerStoreDB struct {
	DB *gorm.DB
}

// NewProviderStore ...
func NewProviderStore(db *gorm.DB) ProviderStore {
	return &providerStoreDB{
		DB: db,
	}
}

func (ss *providerStoreDB) List() ([]*Provider, error) {
	providers := make([]*Provider, 0)

	providers = append(providers, &Provider{
		ID: ProviderFree,
	}, &Provider{
		ID: ProviderWebpayPlusNormal,
	}, &Provider{
		ID: ProviderWebpayPatpass,
	})

	return providers, nil
}

// TransactionStore service definition
type TransactionStore interface {
	List() ([]*Transaction, error)
	Create(*Transaction) (*Transaction, error)
	Update(*Transaction) (*Transaction, error)
	Last(string) (*Transaction, error)
}

type transactionStoreDB struct {
	DB *gorm.DB
}

// NewTransactionStore ...
func NewTransactionStore(db *gorm.DB) TransactionStore {
	return &transactionStoreDB{
		DB: db,
	}
}

// List ...
func (ts *transactionStoreDB) List() ([]*Transaction, error) {
	transactionModels := make([]*TransactionModel, 0)

	// Get transactions
	err := ts.DB.Find(&transactionModels).Error
	if err != nil {
		return nil, err
	}

	// Parse transactions slice
	transactions := make([]*Transaction, 0)
	for i := 0; i < len(transactionModels); i++ {
		transactions = append(transactions, transactionModels[i].To())
	}

	return transactions, nil
}

// Create ...
func (ts *transactionStoreDB) Create(t *Transaction) (*Transaction, error) {
	// parse transaction to model
	model := &TransactionModel{}
	err := model.From(t)
	if err != nil {
		return nil, err
	}

	// create transaction
	err = ts.DB.Create(&model).Error
	if err != nil {
		return nil, err
	}

	return model.To(), nil
}

// Update ...
func (ts *transactionStoreDB) Update(t *Transaction) (*Transaction, error) {
	// parse transaction to model
	model := &TransactionModel{}
	err := model.From(t)
	if err != nil {
		return nil, err
	}

	// save transaction
	err = ts.DB.Save(&model).Error
	if err != nil {
		return nil, err
	}

	return model.To(), nil
}

// Last ...
func (ts *transactionStoreDB) Last(userID string) (*Transaction, error) {
	model := &TransactionModel{}
	err := ts.DB.Where("user_id = ?", userID).Order("updated_at DESC").Limit(1).First(model).Error
	if err != nil {
		return nil, err
	}

	// get subscription
	model.Subscription = &SubscriptionModel{}
	err = ts.DB.Where("id = ?", model.SubscriptionID).Find(model.Subscription).Error
	if err != nil {
		return nil, err
	}

	return model.To(), nil
}
