package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/context"
)

func listUsers(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[HTTP][Users][List][Init]")

		listedUsers, err := ctx.UserService.List()
		if err != nil {
			fmt.Printf("[HTTP][Users][List][Error] %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := Response{
			Data: listedUsers,
		}

		fmt.Printf("[HTTP][Users][List][Response] %v\n", res)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Printf("[HTTP][Users][List][Error] %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func getUser(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[Users][Get][Init]")

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Printf("[HTTP][Get][Error] %v\n", err)
			b, _ := json.Marshal(Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		user, err := ctx.UserService.GetByID(userID)
		if err != nil {
			fmt.Printf("[HTTP][Users][Get][Error] %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := Response{
			Data: user,
		}

		fmt.Printf("[HTTP][Users][Get][Response] %v\n", res)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Printf("[HTTP][Users][Get][Error] %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func createUser(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[HTTP][Users][Create][Init]")

		payload := &struct {
			User *User `json:"user"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			fmt.Printf("[HTTP][Users][Create][Error] %v\n", err)
			b, _ := json.Marshal(Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.User == nil {
			err := "undefined user"
			fmt.Printf("[HTTP][Users][Create][Error] %v\n", err)
			b, _ := json.Marshal(Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.User.Email == "" || payload.User.Name == "" || payload.User.Password == "" {
			err := "undefined email, name or password"
			fmt.Printf("[HTTP][Users][Create][Error] %v\n", err)
			b, _ := json.Marshal(Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		fmt.Printf("[HTTP][Users][Create][Request] payload = %v\n", payload)

		err := ctx.UserService.Create(payload.User)
		if err != nil {
			fmt.Printf("[HTTP][Users][Create][Error] %v\n", err)
			b, _ := json.Marshal(Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := Response{
			Data: payload.User,
		}

		fmt.Printf("[HTTP][Users][Create][Response] %v\n", res)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Printf("[HTTP][Users][Create][Error] %v\n", err)
			b, _ := json.Marshal(Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func updateUser(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[HTTP][Users][Update][Init]")

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Printf("[HTTP][Update][Error] %v\n", err)
			b, _ := json.Marshal(Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		payload := &struct {
			User *User `json:"user"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			fmt.Printf("[HTTP][Users][Update][Error] %v\n", err)
			b, _ := json.Marshal(Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.User == nil {
			err := "undefined user"
			fmt.Printf("[HTTP][Users][Update][Error] %v\n", err)
			b, _ := json.Marshal(Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		fmt.Printf("[HTTP][Users][Update][Request] payload = %v\n", payload)

		payload.User.ID = userID

		err := ctx.UserService.Update(payload.User)
		if err != nil {
			fmt.Printf("[HTTP][Users][Update][Error] %v\n", err)
			b, _ := json.Marshal(Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := Response{
			Data: payload.User,
		}

		fmt.Printf("[HTTP][Users][Update][Response] %v\n", res)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Printf("[HTTP][Users][Update][Error] %v\n", err)
			b, _ := json.Marshal(Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func deleteUser(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[HTTP][Users][Delete][Init]")

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Printf("[HTTP][Delete][Error] %v\n", err)
			b, _ := json.Marshal(Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		err := ctx.UserService.Delete(&User{
			ID: userID,
		})
		if err != nil {
			fmt.Printf("[HTTP][Users][Delete][Error] %v\n", err)
			b, _ := json.Marshal(Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := Response{}

		fmt.Printf("[HTTP][Users][Delete][Response] %v\n", res)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Printf("[HTTP][Users][Delete][Error] %v\n", err)
			b, _ := json.Marshal(Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}
