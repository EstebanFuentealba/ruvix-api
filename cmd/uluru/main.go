package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/jmlopezz/uluru-api/template"

	authClient "github.com/microapis/auth-api/client"
	authHTTP "github.com/microapis/auth-api/http"
	auth "github.com/microapis/auth-api/run"

	e "github.com/microapis/email-api"
	emailClient "github.com/microapis/email-api/client"
	email "github.com/microapis/email-api/run"

	usersClient "github.com/microapis/users-api/client"
	usersHTTP "github.com/microapis/users-api/http"
	users "github.com/microapis/users-api/run"

	afpsimulatorHTTP "github.com/jmlopezz/afp-simulator/http"
)

func main() {
	//
	// INITIALIZE ULURU
	//
	postgresDSN := os.Getenv("DATABASE_URL")
	if postgresDSN == "" {
		log.Fatalln("missing env variable DATABASE_URL")
	}
	host := os.Getenv("HOST")
	if host == "" {
		log.Fatalln("missing env variable HOST")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("missing env variable PORT")
	}
	addr := fmt.Sprintf("%s:%s", host, port)

	//
	// INITIALIZE USERS SERVICE
	//
	usersHost := os.Getenv("USERS_HOST")
	if usersHost == "" {
		log.Fatalln("missing env variable USERS_HOST")
	}
	usersPort := os.Getenv("USERS_PORT")
	if usersPort == "" {
		log.Fatalln("missing env variable USERS_PORT")
	}
	usersAddr := fmt.Sprintf("%s:%s", usersHost, usersPort)
	uc, err := usersClient.New(usersAddr)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		users.Run(usersAddr, postgresDSN)
	}()

	//
	// INITIALIZE EMAIL SERVICE
	//
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		err := errors.New("invalid REDIS_URL env value")
		log.Fatal(err)
	}

	log.Println("============================")
	log.Println("REDIS_URL", redisURL)
	log.Println("============================")

	providersEnv := os.Getenv("PROVIDERS")
	if providersEnv == "" {
		log.Fatal(errors.New("PROVIDERS value not defined"))
	}
	providers := strings.Split(providersEnv, ",")
	emailHost := os.Getenv("EMAIL_HOST")
	if emailHost == "" {
		log.Fatalln("missing env variable EMAIL_HOST")
	}
	emailPort := os.Getenv("EMAIL_PORT")
	if emailPort == "" {
		log.Fatalln("missing env variable EMAIL_PORT")
	}
	emailAddr := fmt.Sprintf("%s:%s", emailHost, emailPort)
	ec, err := emailClient.New(emailAddr)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		email.Run(emailAddr, redisURL, providers)
	}()

	//
	// INITIALIZE AUTH SERVICE
	//
	authHost := os.Getenv("AUTH_HOST")
	if authHost == "" {
		log.Fatalln("missing env variable AUTH_HOST")
	}
	authPort := os.Getenv("AUTH_PORT")
	if authPort == "" {
		log.Fatalln("missing env variable AUTH_PORT")
	}
	authAddr := fmt.Sprintf("%s:%s", authHost, authPort)
	ac, err := authClient.New(authAddr)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		auth.Run(authAddr, postgresDSN, usersAddr, &e.MailingTemplates{
			Signup:          template.SignupTemplate(ec),
			VerifyEmail:     template.VerifyEmailTemplate(ec),
			ForgotPassword:  template.ForgotPasswordTemplate(ec),
			PasswordChanged: template.PasswordChangedTemplate(ec),
		})
	}()

	//
	// INITIALIZE HTTP SERVER
	//
	r := mux.NewRouter()
	authHTTP.Routes(r, ac)
	usersHTTP.Routes(r, uc)
	afpsimulatorHTTP.Routes(r)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(loggingMiddleware)
	log.Println("Starting HTTP service...")
	go func() {
		log.Println(fmt.Sprintf("HTTP service running, Listening on: %v", addr))
		err = http.ListenAndServe(":5000", r)
		if err != nil {
			log.Fatal(err)
		}
	}()

	//
	// LISTEN SIGNALS
	//
	quit := make(chan struct{})
	listenInterrupt(quit)
	<-quit
	gracefullShutdown()
}

func listenInterrupt(quit chan struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		s := <-c
		log.Println("Signal received", s)
		quit <- struct{}{}
	}()
}

func gracefullShutdown() {
	log.Println("Gracefully shutdown")
	// if err := httpServer.Shutdown(nil); err != nil {
	// 	log.Error(err.Error())
	// }
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI, r.Method)
		next.ServeHTTP(w, r)
	})
}
