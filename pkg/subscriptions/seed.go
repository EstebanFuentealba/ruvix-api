package subscriptions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"
)

type subscriptionsJSON struct {
	Subscriptions []*Subscription `json:"subscriptions"`
}

func seedSubscription(db *gorm.DB) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/misc/seeds/subscriptions.json", pwd)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var data subscriptionsJSON

	err = json.Unmarshal(file, &data)
	if err != nil {
		return err
	}

	for _, v := range data.Subscriptions {
		model := &SubscriptionModel{}
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
