package institutions

import "github.com/jmlopezz/uluru-api/database"

// Institution ...
type Institution struct {
	database.Base

	Name string `json:"name" gorm:"NOT NULL"`
	Type Type   `gorm:"NOT NULL" json:"type" sql:"enum('investment', 'inmobiliary')"`
}

// Account ...
type Account struct {
	database.Base

	InstitutionID Institution `json:"institution_id" gorm:"foreignkey:InstitutionRefer"`
	Name          string      `json:"name" gorm:"NOT NULL"`
}

// Instrument ...
type Instrument struct {
	database.Base

	AccountID Account `json:"account_id" gorm:"foreignkey:AccountRefer"`

	Name      string  `json:"name" gorm:"NOT NULL"`
	Amount    float64 `json:"amount" gorm:"NOT NULL"`
	Percent   float64 `json:"percent" gorm:"NOT NULL"`
	Return1m  float64 `json:"return_1m" gorm:"NOT NULL"`
	Return1y  float64 `json:"return_1y" gorm:"NOT NULL"`
	Return5y  float64 `json:"return_5y" gorm:"NOT NULL"`
	Return10y float64 `json:"return_10y" gorm:"NOT NULL"`
}
