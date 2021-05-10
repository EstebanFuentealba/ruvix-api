package savings

import (
	"math/rand"
	"strconv"

	ruvixapi "github.com/cagodoy/ruvix-api"
	uuid "github.com/satori/go.uuid"
)

// FactoryCreateInstitution ...
func FactoryCreateInstitution(addr string, options ruvixapi.ClientOptions) (*Institution, *Institution, error) {
	sc, err := NewClient(addr, options)
	if err != nil {
		return nil, nil, err
	}

	randomUUID, err := uuid.NewV4()
	if err != nil {
		return nil, nil, err
	}

	before := &Institution{
		Name: "fake_name_" + randomUUID.String(),
		Slug: "fake_slug_" + randomUUID.String(),
	}

	after, err := sc.CreateInstitution(before)
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// FactoryCreateInstitutionWithAccounts ...
func FactoryCreateInstitutionWithAccounts(addr string, options ruvixapi.ClientOptions) (*Institution, *Institution, error) {
	sc, err := NewClient(addr, options)
	if err != nil {
		return nil, nil, err
	}

	randomUUID, err := uuid.NewV4()
	if err != nil {
		return nil, nil, err
	}

	account := &Account{
		Name: "fake_account_name_" + strconv.Itoa(rand.Intn(999999)),
		Slug: "fake_account_slug_" + strconv.Itoa(rand.Intn(999999)),
	}

	before := &Institution{
		Name:     "fake_name_" + randomUUID.String(),
		Slug:     "fake_slug_" + randomUUID.String(),
		Accounts: []*Account{account, account},
	}

	after, err := sc.CreateInstitution(before)
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// FactoryCreateInstitutionWithAccountsAndInstruments ...
func FactoryCreateInstitutionWithAccountsAndInstruments(addr string, options ruvixapi.ClientOptions) (*Institution, *Institution, error) {
	sc, err := NewClient(addr, options)
	if err != nil {
		return nil, nil, err
	}

	randomUUID, err := uuid.NewV4()
	if err != nil {
		return nil, nil, err
	}

	instrument := &Instrument{
		Name:               "fake_instrument_name_" + strconv.Itoa(rand.Intn(999999)),
		Slug:               "fake_instrument_slug_" + strconv.Itoa(rand.Intn(999999)),
		Return1m:           0.1,
		Return1y:           1.2,
		Return5y:           6,
		Return10y:          12,
		ProjectedWorstCase: 0.1,
		ProjectedAvgCase:   0.2,
		ProjectedBestCase:  0.3,
	}

	account := &Account{
		Name:        "fake_account_name_" + strconv.Itoa(rand.Intn(999999)),
		Slug:        "fake_account_slug_" + strconv.Itoa(rand.Intn(999999)),
		Instruments: []*Instrument{instrument, instrument},
	}

	before := &Institution{
		Name:     "fake_name_" + randomUUID.String(),
		Slug:     "fake_slug_" + randomUUID.String(),
		Accounts: []*Account{account, account},
	}

	after, err := sc.CreateInstitution(before)
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// FactoryListInstitutions ...
func FactoryListInstitutions(addr string, options ruvixapi.ClientOptions) (*Institution, []*Institution, error) {
	sc, err := NewClient(addr, options)
	if err != nil {
		return nil, nil, err
	}

	_, before, err := FactoryCreateInstitutionWithAccountsAndInstruments(addr, options)
	if err != nil {
		return nil, nil, err
	}

	after, err := sc.ListInstitutions()
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}
