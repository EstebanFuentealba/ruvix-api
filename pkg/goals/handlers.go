package goals

import (
	"encoding/json"
	"fmt"
	"net/http"

	ruvixapi "github.com/cagodoy/ruvix-api"
	"github.com/gorilla/context"
)

func listGoals(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[Goals][List][Request] empty = %v", ""))

		goals, err := ctx.GoalStore.ListGoals()
		if err != nil {
			fmt.Println(fmt.Sprintf("[Goals][List][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := ruvixapi.Response{
			Data: goals,
		}

		fmt.Println(fmt.Sprintf("[Goals][List][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[Goals][List][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func createGoal(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[Goals][Create][Init]"))

		payload := &struct {
			Goal *Goal `json:"goal"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			fmt.Println(fmt.Sprintf("[Goals][Create][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Goal == nil {
			err := "undefined goal"
			fmt.Println(fmt.Sprintf("[Goals][Create][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Goal.Name == "" {
			err := "undefined name"
			fmt.Println(fmt.Sprintf("[Goals][Create][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		fmt.Println(fmt.Sprintf("[Goals][Create][Request] payload = %v", payload))

		g, err := ctx.GoalStore.CreateGoal(payload.Goal)
		if err != nil {
			fmt.Println(fmt.Sprintf("[Goals][Create][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := ruvixapi.Response{
			Data: g,
		}

		fmt.Println(fmt.Sprintf("[Goals][Create][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[Goals][Create][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func getLastRetirement(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[Goals][GetLastRetirement][Init]"))

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Println(fmt.Sprintf("[Goals][GetLastRetirement][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		fmt.Println(fmt.Sprintf("[Retirement][GetLastRetirement][Request] empty = %v", ""))

		retirement, err := ctx.GoalStore.GetLastRetirementGoal(&RetirementGoalQuery{
			UserID: userID,
		})
		if err != nil {
			fmt.Println(fmt.Sprintf("[Retirement][GetLastRetirement][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := ruvixapi.Response{
			Data: retirement,
		}

		fmt.Println(fmt.Sprintf("[Retirement][GetLastRetirement][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[Retirement][GetLastRetirement][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func createRetirementGoal(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[Goals][CreateRetirement][Init]")

		payload := &struct {
			Retirement *RetirementGoal `json:"retirement"`
		}{}

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Println(fmt.Sprintf("[Goals][CreateRetirement][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			fmt.Println(fmt.Sprintf("[Retirement][CreateRetirement][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Retirement == nil {
			err := "undefined retirement"
			fmt.Println(fmt.Sprintf("[Retirement][CreateRetirement][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Retirement.GoalID == "" {
			err := "undefined GoalID"
			fmt.Println(fmt.Sprintf("[Retirement][CreateRetirement][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		// set user_id in retirement
		payload.Retirement.UserID = userID

		// sets user_id in all retirement instruments
		for i := 0; i < len(payload.Retirement.RetirementInstruments); i++ {
			payload.Retirement.RetirementInstruments[i].UserID = userID
		}

		fmt.Println(fmt.Sprintf("[Retirement][CreateRetirement][Request] payload = %v", payload))

		retirement, err := ctx.GoalStore.CreateRetirementGoal(payload.Retirement)
		if err != nil {
			fmt.Println(fmt.Sprintf("[Retirement][CreateRetirement][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		// calculate estimation
		// pp := &person.PersonParams{}
		// worst, avg, best, err := ctx.Simulator.Estimation(pp)
		// if err != nil {
		// 	fmt.Println(fmt.Sprintf("[Simulator][AFP][Estimation][Error] %v", err))
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// TODO(ca): calculate recommendations

		// prepare meta
		// meta := struct {
		// 	Simulation      interface{} `json:"simulation"`
		// 	Recommendations interface{} `json:"recommendations"`
		// }{
		// 	Simulation: &simulatorHTTP.EstimationResponse{
		// 		Best:  best,
		// 		Avg:   avg,
		// 		Worst: worst,
		// 	},
		// }

		// prepare response
		res := ruvixapi.Response{
			Data: retirement,
			// Meta: meta,
		}

		fmt.Println(fmt.Sprintf("[Retirement][CreateRetirement][Response] %v", res))

		// send response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[Retirement][CreateRetirement][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func createGuestRetirementGoal(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[Goals][CreateGuestRetirement][Init]")

		payload := &struct {
			Retirement *RetirementGoal `json:"retirement"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			fmt.Println(fmt.Sprintf("[Retirement][CreateGuestRetirement][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Retirement == nil {
			err := "undefined retirement"
			fmt.Println(fmt.Sprintf("[Retirement][CreateGuestRetirement][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Retirement.GoalID == "" {
			err := "undefined GoalID"
			fmt.Println(fmt.Sprintf("[Retirement][CreateGuestRetirement][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Retirement.Fingerprint == "" {
			err := "undefined Fingerprint"
			fmt.Println(fmt.Sprintf("[Retirement][CreateGuestRetirement][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		// sets user_id in all retirement instruments
		for i := 0; i < len(payload.Retirement.RetirementInstruments); i++ {
			payload.Retirement.RetirementInstruments[i].Fingerprint = payload.Retirement.Fingerprint
		}

		fmt.Println(fmt.Sprintf("[Retirement][CreateGuestRetirement][Request] payload = %v", payload))

		retirement, err := ctx.GoalStore.CreateRetirementGoal(payload.Retirement)
		if err != nil {
			fmt.Println(fmt.Sprintf("[Retirement][CreateGuestRetirement][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := ruvixapi.Response{
			Data: retirement,
		}

		fmt.Println(fmt.Sprintf("[Retirement][CreateGuestRetirement][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[Retirement][CreateGuestRetirement][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func getLastGuestRetirement(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[Goals][GetLastGuestRetirement][Init]"))

		fingerprint := context.Get(r, "fingerprint").(string)
		if fingerprint == "" {
			err := "fingerprint is not defined"
			fmt.Println(fmt.Sprintf("[Goals][GetLastGuestRetirement][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		fmt.Println(fmt.Sprintf("[Retirement][GetLastGuestRetirement][Request] empty = %v", ""))

		retirement, err := ctx.GoalStore.GetLastGuestRetirementGoal(&RetirementGoalQuery{
			Fingerprint: fingerprint,
		})
		if err != nil {
			fmt.Println(fmt.Sprintf("[Retirement][GetLastGuestRetirement][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := ruvixapi.Response{
			Data: retirement,
		}

		fmt.Println(fmt.Sprintf("[Retirement][GetLastGuestRetirement][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[Retirement][GetLastGuestRetirement][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}
