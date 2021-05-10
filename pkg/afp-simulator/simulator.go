package simulator

const (
	// CollectionFund holds the name of the Funds collection
	CollectionFund = "funds"
)

// Bucket ...
type Bucket struct {
	Year   int     `json:"year"`
	Amount float64 `json:"amount"`
}

// Histogram ...
type Histogram struct {
	Buckets []Bucket `json:"buckets"`
}
