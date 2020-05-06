package subscriptions

// Statuses ...
const (
	StatusTransactionPending   = "pending"
	StatusTransactionCompleted = "completed"
	StatusTransactionRejected  = "rejected"
	StatusTransactionCanceled  = "canceled"
)

// Providers ...
const (
	ProviderFree             = "free"
	ProviderWebpayPlusNormal = "webpay-plus-normal"
	ProviderWebpayPatpass    = "webpay-patpass"
)

const (
	daysBeforeDueDate = 3
	daysAfterDueDate  = 5
)

// Subscription ...
type Subscription struct {
	ID string `json:"id"`

	Features []*Feature `json:"features,omitempty"`

	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Months int     `json:"months"`

	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt *int64 `json:"-"`
}

// Feature ...
type Feature struct {
	ID string `json:"id"`

	SubscriptionID string `json:"subscription_id"`
	Text           string `json:"text"`

	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt *int64 `json:"-"`
}

// Provider ...
type Provider struct {
	ID string `json:"id"`
}

// Transaction ...
type Transaction struct {
	ID string `json:"id"`

	Subscription *Subscription `json:"subscription,omitempty"`

	UserID         string `json:"user_id"`
	SubscriptionID string `json:"subscription_id"`
	ProviderID     string `json:"provider_id,omitempty"`

	DueDate      int64  `json:"due_date"`
	RemindedAt   int64  `json:"reminded_at"`
	Status       string `json:"status"`
	PaymentToken string `json:"payment_token,omitempty"`
	OrderNumber  string `json:"order_number"`

	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt *int64 `json:"-"`
}

// TransactionMeta ...
type TransactionMeta struct {
	PaymentURL string `json:"payment_url"`
}
