package goals

import (
	"github.com/jinzhu/gorm"
	"github.com/jmlopezz/uluru-api/pkg/savings"
)

// GoalStore service definition
type GoalStore interface {
	ListGoals() ([]*Goal, error)
	CreateGoal(p *Goal) (*Goal, error)
	CreateRetirementGoal(p *RetirementGoal) (*RetirementGoal, error)
	GetLastRetirementGoal(q *RetirementGoalQuery) (*RetirementGoal, error)
	GetLastGuestRetirementGoal(q *RetirementGoalQuery) (*RetirementGoal, error)
}

type goalStoreDB struct {
	DB *gorm.DB
}

// NewGoalStore ...
func NewGoalStore(db *gorm.DB) GoalStore {
	return &goalStoreDB{
		DB: db,
	}
}

// ListGoals ...
func (gs *goalStoreDB) ListGoals() ([]*Goal, error) {
	goalmodels := make([]*GoalModel, 0)

	err := gs.DB.Find(&goalmodels).Error
	if err != nil {
		return nil, err
	}

	goals := make([]*Goal, 0)
	for i := 0; i < len(goalmodels); i++ {
		goals = append(goals, goalmodels[i].To())
	}

	return goals, nil
}

// CreateGoal ...
func (gs *goalStoreDB) CreateGoal(g *Goal) (*Goal, error) {
	model := &GoalModel{}

	err := model.From(g)
	if err != nil {
		return nil, err
	}

	err = gs.DB.Create(model).Error
	if err != nil {
		return nil, err
	}

	return model.To(), nil
}

// CreateRetirementGoal ...
func (gs *goalStoreDB) CreateRetirementGoal(r *RetirementGoal) (*RetirementGoal, error) {
	retirementModel := &RetirementGoalModel{}

	// TODO(ca): check if exist goal id db using r.GoalID

	err := retirementModel.From(r)
	if err != nil {
		return nil, err
	}

	err = gs.DB.Create(&retirementModel).Error
	if err != nil {
		return nil, err
	}

	// Get Goal by GoalID
	// goalModel := &GoalModel{}
	// err = gs.DB.Where("id = ?", retirementModel.GoalID).Find(&goalModel).Error
	// if err != nil {
	// 	return nil, err
	// }
	// retirementModel.Goal = goalModel

	// Get instrument by each retirement_instrument
	// for i := 0; i < len(retirementModel.RetirementInstruments); i++ {
	// 	retirementInstrument := retirementModel.RetirementInstruments[i]
	// 	id := retirementInstrument.InstrumentID

	// 	instrumentModel := &savings.InstrumentModel{}
	// 	err = gs.DB.Where("id = ?", id).Find(&instrumentModel).Error
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	retirementModel.RetirementInstruments[i].Instrument = instrumentModel
	// }

	return retirementModel.To(), nil
}

// GetLastRetirementGoal ...
func (gs *goalStoreDB) GetLastRetirementGoal(q *RetirementGoalQuery) (*RetirementGoal, error) {
	retirementModel := &RetirementGoalModel{}

	err := gs.DB.Where("user_id = ?", q.UserID).Order("created_at DESC").Limit(1).First(retirementModel).Error
	if err != nil {
		return nil, err
	}

	retirementModel.Goal = &GoalModel{}
	err = gs.DB.Where("id = ?", retirementModel.GoalID.String()).Order("created_at DESC").Limit(1).First(retirementModel.Goal).Error
	if err != nil {
		return nil, err
	}

	retirementModel.RetirementInstruments = make([]*savings.RetirementInstrumentModel, 0)
	err = gs.DB.Where("retirement_goal_id = ?", retirementModel.Base.ID).Find(&retirementModel.RetirementInstruments).Error
	if err != nil {
		return nil, err
	}

	return retirementModel.To(), nil
}

// GetLastGuestRetirementGoal ...
func (gs *goalStoreDB) GetLastGuestRetirementGoal(q *RetirementGoalQuery) (*RetirementGoal, error) {
	retirementModel := &RetirementGoalModel{}

	// TODO(ca): check if exist goal id db using r.GoalID

	err := gs.DB.Where("fingerprint = ?", q.Fingerprint).Order("created_at DESC").Limit(1).First(retirementModel).Error
	if err != nil {
		return nil, err
	}

	retirementModel.Goal = &GoalModel{}
	err = gs.DB.Where("id = ?", retirementModel.GoalID.String()).Order("created_at DESC").Limit(1).First(retirementModel.Goal).Error
	if err != nil {
		return nil, err
	}

	retirementModel.RetirementInstruments = make([]*savings.RetirementInstrumentModel, 0)
	err = gs.DB.Where("retirement_goal_id = ?", retirementModel.Base.ID).Find(&retirementModel.RetirementInstruments).Error
	if err != nil {
		return nil, err
	}

	return retirementModel.To(), nil
}
