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
	"github.com/rs/cors"

	a "github.com/microapis/authentication-api"
	authclient "github.com/microapis/authentication-api/client"
	authHTTP "github.com/microapis/authentication-api/http"
	auth "github.com/microapis/authentication-api/run"

	emailclient "github.com/microapis/email-api/client"
	email "github.com/microapis/email-api/run"

	u "github.com/microapis/users-api"
	usersclient "github.com/microapis/users-api/client"
	usersHTTP "github.com/microapis/users-api/http"
	users "github.com/microapis/users-api/run"

	afpsimulatorHTTP "github.com/jmlopezz/afp-simulator/http"

	"github.com/jmlopezz/uluru-api/database"
	authmodel "github.com/jmlopezz/uluru-api/pkg/auth"
	"github.com/jmlopezz/uluru-api/pkg/goals"
	"github.com/jmlopezz/uluru-api/pkg/profile"
	"github.com/jmlopezz/uluru-api/pkg/savings"
	"github.com/jmlopezz/uluru-api/pkg/subscriptions"
	usermodel "github.com/jmlopezz/uluru-api/pkg/users"
)

var (
	httpServer *http.Server

	subscriptionsWorkerPool subscriptions.WorkerPool
)

func main() {
	//
	// INITIALIZE API
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
	uc, err := usersclient.New(usersAddr)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		users.Run(usersAddr, postgresDSN, &u.Events{
			AfterCreate: func() error {
				log.Println("Users: here in AfterCreate event")
				return nil
			},
		})
	}()

	//
	// INITIALIZE EMAIL SERVICE
	//
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		err := errors.New("invalid REDIS_URL env value")
		log.Fatal(err)
	}

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
	ec, err := emailclient.New(emailAddr)
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
	ac, err := authclient.New(authAddr)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		auth.Run(authAddr, postgresDSN, usersAddr, &a.MailingTemplates{
			Signup:          template.SignupTemplate(ec),
			VerifyEmail:     template.VerifyEmailTemplate(ec),
			ForgotPassword:  template.ForgotPasswordTemplate(ec),
			PasswordChanged: template.PasswordChangedTemplate(ec),
		})
	}()

	//
	// HTTP SERVER
	//
	r := mux.NewRouter()
	db, err := database.NewPostgres(postgresDSN)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	db.LogMode(true)

	err = db.Exec("CREATE extension if not exists pgcrypto;").Error
	if err != nil {
		log.Fatalln(err)
	}

	// Auth (microapis/authentication-api)
	authmodel.RunMigrations(db)
	authHTTP.Routes(r, ac)

	// Users (microapis/users-api)
	usermodel.RunMigrations(db)
	usersHTTP.Routes(r, uc)

	// Simulations (jmlopezz/afp-simulator)
	afpsimulatorHTTP.Routes(r)

	// Profile (jmlopezz/uluru-api)
	profile.RunMigrations(db)
	profile.Routes(r, ac, profile.NewStore(db))

	// Goals (jmlopezz/uluru-api)
	goals.RunMigrations(db)
	goals.Routes(r, ac, goals.NewGoalStore(db), goals.NewRetirementGoalStore(db))

	// Savings (jmlopezz/uluru-api)
	savings.RunMigrations(db)
	savings.Routes(r, ac, savings.NewInstitutionStore(db), savings.NewRetirementInstrumentStore(db))

	// Subscriptions (jmlopezz/uluru-api)
	subscriptions.RunMigrations(db)
	ss := subscriptions.NewSubscriptionStore(db)
	subscriptions.Routes(r, ac, ss)
	subscriptionsWorkerPool = subscriptions.NewWorkerPool(ss)
	err = subscriptionsWorkerPool.Start()
	if err != nil {
		log.Fatal(err)
	}

	// Notifications (jmlopezz/uluru-api)

	r.Use(loggingMiddleware)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}).Handler(r)

	log.Println("Starting HTTP service...")
	go func() {
		log.Println(fmt.Sprintf("HTTP service running, Listening on: %v", addr))

		httpServer = &http.Server{Addr: addr, Handler: handler}
		err = httpServer.ListenAndServe()
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

	subscriptionsWorkerPool.Stop()

	if err := httpServer.Shutdown(nil); err != nil {
		log.Fatalln(err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI, r.Method)
		next.ServeHTTP(w, r)
	})
}
