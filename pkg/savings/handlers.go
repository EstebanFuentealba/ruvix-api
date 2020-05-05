package savings

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmlopezz/uluru-api"
)

func listInstitutions(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[SavingInstitutions][List][Request] empty = %v", ""))

		institutions, err := ctx.InstitutionStore.List()
		if err != nil {
			fmt.Println(fmt.Sprintf("[SavingInstitutions][List][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := uluru.Response{
			Data: institutions,
		}

		fmt.Println(fmt.Sprintf("[SavingInstitutions][List][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[SavingInstitutions][List][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

type createInstitutionRequest struct {
	Institution *Institution `json:"institution"`
}

func createInstitution(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[SavingInstitutions][Create][Init]"))

		payload := &createInstitutionRequest{}

		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			fmt.Println(fmt.Sprintf("[SavingInstitutions][Create][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Institution == nil {
			err := "undefined institution"
			fmt.Println(fmt.Sprintf("[SavingInstitutions][Create][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Institution.Name == "" || payload.Institution.Slug == "" {
			err := "undefined name or slug"
			fmt.Println(fmt.Sprintf("[SavingInstitutions][Create][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		fmt.Println(fmt.Sprintf("[SavingInstitutions][Create][Request] payload = %v", payload))

		out, err := ctx.InstitutionStore.Create(payload.Institution)
		if err != nil {
			fmt.Println(fmt.Sprintf("[SavingInstitutions][Create][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := uluru.Response{
			Data: out,
		}

		fmt.Println(fmt.Sprintf("[SavingInstitutions][Create][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[SavingInstitutions][Create][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}
