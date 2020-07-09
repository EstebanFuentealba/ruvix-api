package savings

// Accounts ...
const (
	AccountMandatoryContributions = "mandatory-contributions-account"
	AccountVoluntaryContributions = "voluntary-contributions-account"
	AccountTwo                    = "two-account"
	AccountAgreedDeposits         = "agreed-deposits-account"
)

// Institution ...
type Institution struct {
	ID string `json:"id"`

	Name string `json:"name"`
	Slug string `json:"slug"`

	Accounts []*Account `json:"accounts"`

	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt *int64 `json:"-"`
}

// Account ...
type Account struct {
	ID string `json:"id"`

	Instruments []*Instrument `json:"instruments"`

	InstitutionID string `json:"institution_id"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`

	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt *int64 `json:"-"`
}

// Instrument ...
type Instrument struct {
	ID string `json:"id"`

	AccountID          string  `json:"account_id"`
	Name               string  `json:"name"`
	Slug               string  `json:"slug"`
	Return1m           float64 `json:"return_1m"`
	Return1y           float64 `json:"return_1y"`
	Return5y           float64 `json:"return_5y"`
	Return10y          float64 `json:"return_10y"`
	ProjectedWorstCase float64 `json:"projected_worst_case"`
	ProjectedAvgCase   float64 `json:"projected_avg_case"`
	ProjectedBestCase  float64 `json:"projected_best_case"`

	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt *int64 `json:"-"`
}

// RetirementInstrument ...
type RetirementInstrument struct {
	ID string `json:"id"`

	Instrument *Instrument `json:"retirement_instrument,omitempty"`

	InstrumentID     string  `json:"instrument_id"`
	RetirementGoalID string  `json:"retirement_goal_id"`
	UserID           string  `json:"user_id,omitempty"`
	Fingerprint      string  `json:"fingerprint,omitempty"`
	Percent          float64 `json:"percent"`
	QuotasQuantity   float64 `json:"quotas_quantity"`
	QuotasDate       string  `json:"quotas_date"`
	QuotasPrice      float64 `json:"quotas_price"`
	Balance          float64 `json:"balance"`

	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt *int64 `json:"-"`
}

// // Query ...
// type Query struct {
// 	ID     string
// 	UserID string
// }
