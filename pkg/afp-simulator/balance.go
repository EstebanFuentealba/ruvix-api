package simulator

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
)

type balanceByCaseResponse struct {
	BestCase  balanceResponse `json:"best_case"`
	AvrgCase  balanceResponse `json:"avrg_case"`
	WorstCase balanceResponse `json:"worst_case"`
}

type balanceHistResponse struct {
	Amount    float64   `json:"amount"`
	Histogram Histogram `json:"histogram"`
}

type balanceResponse struct {
	Balance balanceHistResponse `json:"balance"`
}

func balance(ctx handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		personParam := &PersonParams{}

		if err := json.NewDecoder(r.Body).Decode(personParam); err != nil {
			fmt.Printf("[Simulator][AFP][Estimation][Error] %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		fmt.Printf("[Simulator][AFP][Estimation][Request] person = %v\n", personParam)

		_, err := govalidator.ValidateStruct(personParam)
		if err != nil {
			fmt.Printf("[Simulator][AFP][Estimation][Error] %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		p := personParam.Person
		p.Normalize()

		BestCaseInterest, AvrgCaseInterest, WorstCaseInterest := GetMapFunds(personParam.InterestFunds)

		balance, hist := p.Balance(BestCaseInterest)
		BestCaseResponse := balanceResponse{
			Balance: balanceHistResponse{
				Amount:    balance,
				Histogram: hist,
			},
		}

		balance, hist = p.Balance(AvrgCaseInterest)
		AvrgaseResponse := balanceResponse{
			Balance: balanceHistResponse{
				Amount:    balance,
				Histogram: hist,
			},
		}

		balance, hist = p.Balance(WorstCaseInterest)
		WorstCaseResponse := balanceResponse{
			Balance: balanceHistResponse{
				Amount:    balance,
				Histogram: hist,
			},
		}

		//prepare response
		br := balanceByCaseResponse{
			BestCase:  BestCaseResponse,
			AvrgCase:  AvrgaseResponse,
			WorstCase: WorstCaseResponse,
		}

		res := Response{
			Data: br,
		}

		fmt.Printf("[Simulator][AFP][Estimation][Response] %v\n", res)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Printf("[Simulator][AFP][Estimation] %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
