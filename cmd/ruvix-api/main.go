package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/cagodoy/ruvix-api/pkg/auth"
	"github.com/cagodoy/ruvix-api/pkg/goals"
	"github.com/cagodoy/ruvix-api/pkg/profile"
	"github.com/cagodoy/ruvix-api/pkg/savings"
	"github.com/cagodoy/ruvix-api/pkg/subscriptions"
	"github.com/cagodoy/ruvix-api/pkg/users"

	"github.com/cagodoy/ruvix-api/pkg/pigeon/email"
	"github.com/cagodoy/ruvix-api/pkg/pigeon/email/provider/mandrill"
	"github.com/cagodoy/ruvix-api/pkg/pigeon/email/provider/sendgrid"
	"github.com/cagodoy/ruvix-api/pkg/pigeon/email/provider/ses"
	"github.com/cagodoy/ruvix-api/pkg/pigeon/scheduler"

	"github.com/cagodoy/ruvix-api/internal/database"
	"github.com/cagodoy/ruvix-api/internal/template"
)

var (
	httpServer *http.Server

	subscriptionsWorkerPool subscriptions.WorkerPool
)

type Provider interface {
	Approve(*email.Message) error
	Deliver(*email.Message) error
}

var (
	provider Provider
)

func main() {
	//
	// INITIALIZE API
	//
	// env := os.Getenv("ENV")
	// if env == "" {
	// 	log.Fatalln("missing env variable ENV")
	// }
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
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
	// INITIALIZE EMAIL PROVIDER
	//
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		err := errors.New("invalid REDIS_URL env value")
		log.Fatal(err)
	}
	providerSrt := os.Getenv("PROVIDER")
	if providerSrt == "" {
		log.Fatal(errors.New("PROVIDER value not defined"))
	}
	switch providerSrt {
	case mandrill.Name:
		apiKey := os.Getenv("MANDRILL_API_KEY")
		if apiKey == "" {
			log.Fatal(errors.New("MANDRILL_API_KEY value not defined"))
		}
		provider = mandrill.New(apiKey)
	case sendgrid.Name:
		apiKey := os.Getenv("SENDGRID_API_KEY")
		if apiKey == "" {
			log.Fatal(errors.New("SENDGRID_API_KEY value not defined"))
		}
		provider = sendgrid.New(apiKey)
	case ses.Name:
		region := os.Getenv("SES_REGION")
		if region == "" {
			log.Fatal(errors.New("SES_REGION value not defined"))
		}
		key := os.Getenv("SES_KEY")
		if key == "" {
			log.Fatal(errors.New("SES_KEY value not defined"))
		}
		secretKey := os.Getenv("SES_SECRET_KEY")
		if secretKey == "" {
			log.Fatal(errors.New("SES_SECRET_KEY value not defined"))
		}
		provider = ses.New(region, key, secretKey)
	}

	//
	// INITIALIZE EMAIL SERVICE
	//
	pq := scheduler.NewPriorityQueue(redisURL)
	emailService, err := email.NewService(pq, provider)
	if err != nil {
		log.Fatal(err)
	}

	//
	// HTTP SERVER
	//
	router := mux.NewRouter()
	db, err := database.NewPostgres(databaseURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	db.LogMode(true)

	err = db.Exec("CREATE extension if not exists pgcrypto;").Error
	if err != nil {
		log.Fatalln(err)
	}

	// Users
	usersDb, err := users.NewPostgres(databaseURL)
	userService := users.NewService(usersDb, &users.Events{
		AfterCreate: func() error {
			log.Println("Users: here in AfterCreate event")
			return nil
		},
	})
	users.RunMigrations(db)
	users.Routes(router, userService)

	// Auth
	authDb, err := auth.NewPostgres(databaseURL)
	if err != nil {
		log.Fatalln(err)
	}
	authService := auth.NewService(authDb, userService, &auth.MailingTemplates{
		Signup:          template.SignupTemplate(emailService),
		VerifyEmail:     template.VerifyEmailTemplate(emailService),
		ForgotPassword:  template.ForgotPasswordTemplate(emailService),
		PasswordChanged: template.PasswordChangedTemplate(emailService),
	})
	auth.RunMigrations(db)
	auth.Routes(router, authService)

	// Simulations
	// afpsimulatorHTTP.Routes(r)

	// Profile
	profile.RunMigrations(db)
	profile.Routes(router, authService, profile.NewStore(db))

	// Goals
	goals.RunMigrations(db)
	goals.Routes(router, authService, goals.NewGoalStore(db))

	// Savings
	savings.RunMigrations(db)
	savings.Routes(router, authService, savings.NewInstitutionStore(db))

	// Subscriptions
	subscriptions.RunMigrations(db)
	subscriptionStore := subscriptions.NewSubscriptionStore(db)
	subscriptions.Routes(router, authService, subscriptionStore)
	subscriptionsWorkerPool = subscriptions.NewWorkerPool(subscriptionStore)
	err = subscriptionsWorkerPool.Start()
	if err != nil {
		log.Fatal(err)
	}

	// Notifications (cagodoy/uluru-api)

	router.Use(loggingMiddleware)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}).Handler(router)

	go func() {
		log.Println(fmt.Sprintf("[HTTP] Listening on: %v", addr))

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
