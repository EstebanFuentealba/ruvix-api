package profile

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/jmlopezz/uluru-api"
)

type createProfilePayload struct {
	Profile *Profile `json:"profile"`
}

func createProfile(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[Profile][Create][Init]"))

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Println(fmt.Sprintf("[Profile][Create][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		payload := &createProfilePayload{}
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			fmt.Println(fmt.Sprintf("[Profile][Create][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Profile == nil {
			err := "payload.Profile is undefined"
			fmt.Println(fmt.Sprintf("[Profile][Create][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Profile.Age == 0 && payload.Profile.Birth == 0 && payload.Profile.MaritalStatus == "" && payload.Profile.Gender == "" {
			err := "not age, birth or marital_status to change"
			fmt.Println(fmt.Sprintf("[Profile][Create][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		// TODO(ca): validate if gender is present in enum list

		payload.Profile.UserID = userID

		fmt.Println(fmt.Sprintf("[Profile][Create][Request] payload = %v", payload.Profile))

		profile, err := ctx.Store.Create(payload.Profile)
		if err != nil {
			fmt.Println(fmt.Sprintf("[Profile][Create][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := uluru.Response{
			Data: profile,
		}

		fmt.Println(fmt.Sprintf("[Profile][Create][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[Profile][Create][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func getProfile(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[Profile][Get][Init]"))

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Println(fmt.Sprintf("[Profile][Get][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		profile, err := ctx.Store.Get(&Query{
			UserID: userID,
		})
		if err != nil {
			fmt.Println(fmt.Sprintf("[Profile][Get][Errors] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := uluru.Response{
			Data: profile,
		}

		fmt.Println(fmt.Sprintf("[Profile][Get][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[Profile][Get][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

type updateProfilePayload struct {
	Profile *Profile `json:"profile"`
}

func updateProfile(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[Profile][Update][Init]"))

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Println(fmt.Sprintf("[Profile][Update][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		payload := &updateProfilePayload{}
		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			fmt.Println(fmt.Sprintf("[Profile][Update][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		fmt.Println(fmt.Sprintf("[Profile][Update][Request] payload = %v", payload))

		profile, err := ctx.Store.Get(&Query{
			UserID: userID,
		})
		if err != nil {
			fmt.Println(fmt.Sprintf("[Profile][Update][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if profile.UserID != userID {
			fmt.Println(fmt.Sprintf("[Profile][Update][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusForbidden)
			return
		}

		if payload.Profile == nil {
			err := "payload.Profile is undefined"
			fmt.Println(fmt.Sprintf("[Profile][Update][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Profile.Age == 0 && payload.Profile.Birth == 0 && payload.Profile.Childs == 0 && payload.Profile.MaritalStatus == "" && payload.Profile.Gender == "" {
			fmt.Println(fmt.Sprintf("[Profile][Update][Error] %v", "not age, birth, childs or marital_status to change"))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		// TODO(ca): validate if gender is present in enum list

		if payload.Profile.Age != 0 {
			profile.Age = payload.Profile.Age
		}
		if payload.Profile.Birth != 0 {
			profile.Birth = payload.Profile.Birth
		}
		if payload.Profile.Childs != 0 {
			profile.Childs = payload.Profile.Childs
		}
		if payload.Profile.MaritalStatus != "" {
			profile.MaritalStatus = payload.Profile.MaritalStatus
		}
		if payload.Profile.Gender != "" {
			profile.Gender = payload.Profile.Gender
		}

		p, err := ctx.Store.Update(profile)
		if err != nil {
			fmt.Println(fmt.Sprintf("[Profile][Put][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := uluru.Response{
			Data: p,
		}

		fmt.Println(fmt.Sprintf("[Profile][Put][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[Profile][Put][Error] %v", err))
			b, _ := json.Marshal(uluru.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}
