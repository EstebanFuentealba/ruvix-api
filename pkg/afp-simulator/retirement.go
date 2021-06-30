package simulator

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
)

type retirementResponse struct {
	BestCaseRetirement  float64 `json:"best_case_retirement"`
	AvrgCaseRetirement  float64 `json:"arvg_case_retirement"`
	WorstCaseRetirement float64 `json:"worst_case_retirement"`
}

func retirement(ctx handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		personParam := &PersonParams{}

		if err := json.NewDecoder(r.Body).Decode(personParam); err != nil {
			fmt.Printf("[Simulator][AFP][Retirement][Error] %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		fmt.Printf("[Simulator][AFP][Retirement][Request] person = %v\n", personParam)

		_, err := govalidator.ValidateStruct(personParam)
		if err != nil {
			fmt.Printf("[Simulator][AFP][Retirement][Error] %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		p := personParam.Person
		p.Normalize()

		BestCaseInterest, AvrgCaseInterest, WorstCaseInterest := GetMapFunds(personParam.InterestFunds)

		//prepare response
		rr := retirementResponse{
			BestCaseRetirement:  p.Retirement(BestCaseInterest),
			AvrgCaseRetirement:  p.Retirement(AvrgCaseInterest),
			WorstCaseRetirement: p.Retirement(WorstCaseInterest),
		}

		res := Response{
			Data: rr,
		}

		fmt.Printf("[Simulator][AFP][Retirement][Response] %v\n", res)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Printf("[Simulator][AFP][Retirement] %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
