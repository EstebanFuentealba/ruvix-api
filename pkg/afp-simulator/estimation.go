package simulator

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
)

type estimationResponse struct {
	BestCase  estimationByCaseResponse `json:"best_case"`
	AvrgCase  estimationByCaseResponse `json:"avrg_case"`
	WorstCase estimationByCaseResponse `json:"worst_case"`
}

type estimationByCaseResponse struct {
	Retirement float64             `json:"retirement"`
	Balance    balanceHistResponse `json:"balance"`
}

// GetMapFunds ...
func GetMapFunds(interestFunds InterestFunds) (map[string]float64, map[string]float64, map[string]float64) {
	var BestCase = map[string]float64{
		"A": interestFunds.FundA[0],
		"B": interestFunds.FundB[0],
		"C": interestFunds.FundC[0],
		"D": interestFunds.FundD[0],
		"E": interestFunds.FundE[0],
	}
	var AvrgCase = map[string]float64{
		"A": interestFunds.FundA[1],
		"B": interestFunds.FundB[1],
		"C": interestFunds.FundC[1],
		"D": interestFunds.FundD[1],
		"E": interestFunds.FundE[1],
	}
	var WorstCase = map[string]float64{
		"A": interestFunds.FundA[2],
		"B": interestFunds.FundB[2],
		"C": interestFunds.FundC[2],
		"D": interestFunds.FundD[2],
		"E": interestFunds.FundE[2],
	}

	return BestCase, AvrgCase, WorstCase
}

func estimation(ctx handlerContext) func(w http.ResponseWriter, r *http.Request) {
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

		// Get Best case
		balance, hist := p.Balance(BestCaseInterest)
		retirement := p.Retirement(BestCaseInterest)
		BestCase := estimationByCaseResponse{
			Retirement: retirement,
			Balance: balanceHistResponse{
				Amount:    balance,
				Histogram: hist,
			},
		}

		// Get Avrg case
		balance, hist = p.Balance(AvrgCaseInterest)
		retirement = p.Retirement(AvrgCaseInterest)
		AvrgCase := estimationByCaseResponse{
			Retirement: retirement,
			Balance: balanceHistResponse{
				Amount:    balance,
				Histogram: hist,
			},
		}
		// Get Best case
		balance, hist = p.Balance(WorstCaseInterest)
		retirement = p.Retirement(WorstCaseInterest)
		WorstCase := estimationByCaseResponse{
			Retirement: retirement,
			Balance: balanceHistResponse{
				Amount:    balance,
				Histogram: hist,
			},
		}

		//prepare response
		er := estimationResponse{
			BestCase:  BestCase,
			AvrgCase:  AvrgCase,
			WorstCase: WorstCase,
		}

		res := Response{
			Data: er,
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
