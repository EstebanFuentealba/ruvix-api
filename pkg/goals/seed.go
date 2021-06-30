package goals

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"
)

type goalsJSON struct {
	Goals []*Goal `json:"goals"`
}

func seedGoal(db *gorm.DB) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/misc/seed/goal.json", pwd)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var data goalsJSON

	err = json.Unmarshal(file, &data)
	if err != nil {
		return err
	}

	for _, v := range data.Goals {
		model := &GoalModel{}
		err := model.From(v)
		if err != nil {
			return err
		}

		err = db.Save(model).Error
		if err != nil {
			return err
		}
	}

	return nil
}
