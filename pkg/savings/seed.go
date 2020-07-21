package savings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"
)

type institutionsJSON struct {
	Institutions []*Institution `json:"institutions"`
}

func seedInstitution(db *gorm.DB) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/institutions.json", pwd)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var data institutionsJSON

	err = json.Unmarshal(file, &data)
	if err != nil {
		return err
	}

	for _, v := range data.Institutions {
		model := &InstitutionModel{}
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
