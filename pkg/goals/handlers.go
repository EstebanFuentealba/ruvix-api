package goals

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/jmlopezz/uluru-api"
)

func listGoals(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[Goals][List][Request] empty = %v", ""))

		goals, err := ctx.GoalStore.ListGoals()
		if err != nil {
			fmt.Println(fmt.Sprintf("[Goals][List][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := uluru.Response{
			Data: goals,
		}

		fmt.Println(fmt.Sprintf("[Goals][List][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[Goals][List][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

type createGoalRequest struct {
	Goal *Goal `json:"goal"`
}

func createGoal(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[Goals][Create][Init]"))

		payload := &createGoalRequest{}

		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			fmt.Println(fmt.Sprintf("[Goals][Create][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Goal == nil {
			err := "undefined goal"
			fmt.Println(fmt.Sprintf("[Goals][Create][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Goal.Name == "" {
			err := "undefined name"
			fmt.Println(fmt.Sprintf("[Goals][Create][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		fmt.Println(fmt.Sprintf("[Goals][Create][Request] payload = %v", payload))

		g, err := ctx.GoalStore.CreateGoal(payload.Goal)
		if err != nil {
			fmt.Println(fmt.Sprintf("[Goals][Create][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := uluru.Response{
			Data: g,
		}

		fmt.Println(fmt.Sprintf("[Goals][Create][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[Goals][Create][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func getLastRetirement(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[Profile][GetLast][Init]"))

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Println(fmt.Sprintf("[Profile][GetLast][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		fmt.Println(fmt.Sprintf("[Retirement][GetLast][Request] empty = %v", ""))

		retirement, err := ctx.GoalStore.GetLastRetirementGoal(&RetirementGoalQuery{
			UserID: userID,
		})
		if err != nil {
			fmt.Println(fmt.Sprintf("[Retirement][GetLast][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := uluru.Response{
			Data: retirement,
		}

		fmt.Println(fmt.Sprintf("[Retirement][GetLast][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[Retirement][GetLast][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

type createRetirementRequest struct {
	Retirement *RetirementGoal `json:"retirement"`
}

func createRetirementGoal(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[Profile][CreateRetirement][Init]")

		payload := &createRetirementRequest{}

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Println(fmt.Sprintf("[Profile][CreateRetirement][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			fmt.Println(fmt.Sprintf("[Retirement][CreateRetirement][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Retirement == nil {
			err := "undefined retirement"
			fmt.Println(fmt.Sprintf("[Retirement][CreateRetirement][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Retirement.GoalID == "" {
			err := "undefined GoalID"
			fmt.Println(fmt.Sprintf("[Retirement][CreateRetirement][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err})
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
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := uluru.Response{
			Data: retirement,
		}

		fmt.Println(fmt.Sprintf("[Retirement][CreateRetirement][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[Retirement][CreateRetirement][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}
