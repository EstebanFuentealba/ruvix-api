package goals

import (
	urulu "github.com/jmlopezz/uluru-api"
	"github.com/jmlopezz/uluru-api/pkg/savings"
	uuid "github.com/satori/go.uuid"
)

// FactoryCreateGoal ...
func FactoryCreateGoal(addr string, opts urulu.ClientOptions) (*Goal, *Goal, error) {
	gc, err := NewClient(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	randomUUID, err := uuid.NewV4()
	if err != nil {
		return nil, nil, err
	}

	before := &Goal{
		Name: "fake_name_" + randomUUID.String(),
	}

	after, err := gc.Create(before)
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// FactoryListGoals ...
func FactoryListGoals(addr string, opts urulu.ClientOptions) (*Goal, []*Goal, error) {
	_, before, err := FactoryCreateGoal(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	gc, err := NewClient(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	after, err := gc.List()
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// FactoryCreateRetirementGoal ...
func FactoryCreateRetirementGoal(addr string, opts urulu.ClientOptions) (*RetirementGoal, *RetirementGoal, error) {
	_, afterGoal, err := FactoryCreateGoal(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	_, afterInstitution, err := savings.FactoryCreateInstitutionWithAccountsAndInstruments(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	before := &RetirementGoal{
		GoalID:            afterGoal.ID,
		MonthlySalary:     1000.0,
		MonthlyRetirement: 2000.1,
		RetirementInstruments: []*savings.RetirementInstrument{
			&savings.RetirementInstrument{
				InstrumentID:   afterInstitution.ID,
				Percent:        0.1,
				QuotasQuantity: 89.82,
				QuotasPrice:    39566.14,
				QuotasDate:     "22/04/2020",
				Balance:        3553831,
			},
			&savings.RetirementInstrument{
				InstrumentID:   afterInstitution.ID,
				Percent:        0.9,
				QuotasQuantity: 31.82,
				QuotasPrice:    41502.14,
				QuotasDate:     "22/04/2020",
				Balance:        2222221,
			},
		},
	}

	gc, err := NewClient(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	after, err := gc.CreateRetirement(before)
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// FactoryGetRetirementGoal ...
func FactoryGetRetirementGoal(addr string, opts urulu.ClientOptions) (*RetirementGoal, *RetirementGoal, error) {
	_, before, err := FactoryCreateRetirementGoal(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	gc, err := NewClient(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	after, err := gc.GetRetirement()
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}
