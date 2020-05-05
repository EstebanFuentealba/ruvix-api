package goals

import (
	"github.com/jinzhu/gorm"
	"github.com/jmlopezz/uluru-api/pkg/savings"
)

// GoalStore service definition
type GoalStore interface {
	List() ([]*Goal, error)
	Create(p *Goal) (*Goal, error)
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

// List ...
func (gs *goalStoreDB) List() ([]*Goal, error) {
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

// Create ...
func (gs *goalStoreDB) Create(g *Goal) (*Goal, error) {
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

// RetirementGoalStore service definition
type RetirementGoalStore interface {
	Create(p *RetirementGoal) (*RetirementGoal, error)
	GetLast(q *RetirementGoalQuery) (*RetirementGoal, error)
}

type retirementStoreDB struct {
	DB *gorm.DB
}

// NewRetirementGoalStore ...
func NewRetirementGoalStore(db *gorm.DB) RetirementGoalStore {
	return &retirementStoreDB{
		DB: db,
	}
}

// Create ...
func (rs *retirementStoreDB) Create(r *RetirementGoal) (*RetirementGoal, error) {
	retirementModel := &RetirementGoalModel{}

	err := retirementModel.From(r)
	if err != nil {
		return nil, err
	}

	err = rs.DB.Create(&retirementModel).Error
	if err != nil {
		return nil, err
	}

	// Get Goal by GoalID
	// goalModel := &GoalModel{}
	// err = rs.DB.Where("id = ?", retirementModel.GoalID).Find(&goalModel).Error
	// if err != nil {
	// 	return nil, err
	// }
	// retirementModel.Goal = goalModel

	// Get instrument by each retirement_instrument
	// for i := 0; i < len(retirementModel.RetirementInstruments); i++ {
	// 	retirementInstrument := retirementModel.RetirementInstruments[i]
	// 	id := retirementInstrument.InstrumentID

	// 	instrumentModel := &savings.InstrumentModel{}
	// 	err = rs.DB.Where("id = ?", id).Find(&instrumentModel).Error
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	retirementModel.RetirementInstruments[i].Instrument = instrumentModel
	// }

	return retirementModel.To(), nil
}

// GetLast ...
func (rs *retirementStoreDB) GetLast(q *RetirementGoalQuery) (*RetirementGoal, error) {
	retirementModel := &RetirementGoalModel{}

	err := rs.DB.Where("user_id = ?", q.UserID).Order("created_at DESC").Limit(1).First(retirementModel).Error
	if err != nil {
		return nil, err
	}

	retirementModel.Goal = &GoalModel{}
	err = rs.DB.Where("id = ?", retirementModel.GoalID.String()).Order("created_at DESC").Limit(1).First(retirementModel.Goal).Error
	if err != nil {
		return nil, err
	}

	retirementModel.RetirementInstruments = make([]*savings.RetirementInstrumentModel, 0)
	err = rs.DB.Where("retirement_goal_id = ?", retirementModel.Base.ID).Find(&retirementModel.RetirementInstruments).Error
	if err != nil {
		return nil, err
	}

	return retirementModel.To(), nil
}
