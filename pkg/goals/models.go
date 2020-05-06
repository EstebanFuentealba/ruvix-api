package goals

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jmlopezz/uluru-api/database"
	"github.com/jmlopezz/uluru-api/pkg/savings"
	uuid "github.com/satori/go.uuid"
)

// GoalModel define goal model
type GoalModel struct {
	database.Base

	Name string `json:"name" gorm:"NOT NULL"`
}

// TableName Set table name
func (GoalModel) TableName() string {
	return "goals"
}

// To ...
func (gm *GoalModel) To() *Goal {
	g := &Goal{
		ID: gm.Base.ID.String(),

		Name: gm.Name,

		CreatedAt: gm.Base.CreatedAt.Unix(),
		UpdatedAt: gm.Base.UpdatedAt.Unix(),
	}

	return g
}

// From ...
func (gm *GoalModel) From(m *Goal) error {
	gm.Base = database.Base{}

	if m.ID != "" {
		id, err := uuid.FromString(m.ID)
		if err != nil {
			return err
		}
		gm.Base.ID = id
	}

	if m.CreatedAt != 0 {
		gm.Base.CreatedAt = time.Unix(m.CreatedAt, 0)
	}

	if m.UpdatedAt != 0 {
		gm.Base.UpdatedAt = time.Unix(m.UpdatedAt, 0)
	}

	gm.Name = m.Name

	return nil
}

// RetirementGoalModel ...
type RetirementGoalModel struct {
	database.Base

	Goal                  *GoalModel                           `gorm:"NOT NULL;foreignkey:Base.ID"`
	RetirementInstruments []*savings.RetirementInstrumentModel `gorm:"NOT NULL;foreignkey:RetirementGoalID"`

	UserID            string    `gorm:"NOT NULL;"`
	GoalID            uuid.UUID `gorm:"NOT NULL;"`
	MonthlySalary     float64   `gorm:"NOT NULL;"`
	MonthlyRetirement float64   `gorm:"NOT NULL;"`
}

// To ...
func (rm *RetirementGoalModel) To() *RetirementGoal {
	r := &RetirementGoal{
		ID: rm.Base.ID.String(),

		UserID:            rm.UserID,
		GoalID:            rm.GoalID.String(),
		MonthlySalary:     rm.MonthlySalary,
		MonthlyRetirement: rm.MonthlyRetirement,

		CreatedAt: rm.Base.CreatedAt.Unix(),
		UpdatedAt: rm.Base.UpdatedAt.Unix(),
	}

	if rm.Goal != nil {
		r.Goal = rm.Goal.To()
	}

	if len(rm.RetirementInstruments) > 0 {
		r.RetirementInstruments = make([]*savings.RetirementInstrument, 0)

		for i := 0; i < len(rm.RetirementInstruments); i++ {
			r.RetirementInstruments = append(r.RetirementInstruments, rm.RetirementInstruments[i].To())
		}
	}

	return r
}

// From ...
func (rm *RetirementGoalModel) From(r *RetirementGoal) error {
	rm.Base = database.Base{}

	if r.ID != "" {
		id, err := uuid.FromString(r.ID)
		if err != nil {
			return err
		}
		rm.Base.ID = id
	}

	if r.CreatedAt != 0 {
		rm.Base.CreatedAt = time.Unix(r.CreatedAt, 0)
	}

	if r.UpdatedAt != 0 {
		rm.Base.UpdatedAt = time.Unix(r.UpdatedAt, 0)
	}

	if r.GoalID != "" {
		goalID, err := uuid.FromString(r.GoalID)
		if err != nil {
			return err
		}

		rm.GoalID = goalID
	}

	if r.Goal != nil {
		model := &GoalModel{}
		err := model.From(r.Goal)
		if err != nil {
			return err
		}
		rm.Goal = model
	}

	if len(r.RetirementInstruments) > 0 {
		retirementIntrumentsModels := make([]*savings.RetirementInstrumentModel, 0)
		for i := 0; i < len(r.RetirementInstruments); i++ {
			retirementInstrumentModel := &savings.RetirementInstrumentModel{}
			err := retirementInstrumentModel.From(r.RetirementInstruments[i])
			if err != nil {
				return err
			}
			retirementIntrumentsModels = append(retirementIntrumentsModels, retirementInstrumentModel)
		}
		rm.RetirementInstruments = retirementIntrumentsModels
	}

	rm.UserID = r.UserID
	rm.MonthlySalary = r.MonthlySalary
	rm.MonthlyRetirement = r.MonthlyRetirement

	return nil
}

// TableName Set table name
func (RetirementGoalModel) TableName() string {
	return "retirements"
}

// RunMigrations ...
func RunMigrations(db *gorm.DB) error {
	runGoalSeed := false
	if !db.HasTable(&GoalModel{}) {
		err := db.CreateTable(&GoalModel{}).Error
		if err != nil {
			log.Fatalln(err)
		}

		runGoalSeed = true
	}

	if !db.HasTable(&RetirementGoalModel{}) {
		err := db.CreateTable(&RetirementGoalModel{}).Error
		if err != nil {
			log.Fatalln(err)
		}
	}

	if runGoalSeed {
		err := seedGoal(db)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return nil
}
