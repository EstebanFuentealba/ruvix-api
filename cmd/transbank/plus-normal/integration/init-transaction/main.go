package main

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/cagodoy/ruvix-api/pkg/transbank"
	"github.com/cagodoy/ruvix-api/pkg/transbank/webpay"
)

func main() {
	amount := float64(1000)
	sessionID := "mi-id-de-sesion"
	buyOrder := strconv.Itoa(rand.Intn(99999))
	returnURL := "https://callback/resultado/de/transaccion"
	finalURL := "https://callback/final/post/comprobante/webpay"

	service := webpay.NewIntegrationPlusNormal()
	transaction, err := service.InitTransaction(transbank.InitTransaction{
		Amount:    amount,
		SessionID: sessionID,
		BuyOrder:  buyOrder,
		ReturnURL: returnURL,
		FinalURL:  finalURL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("URL", transaction.URL)
	log.Println("Token", transaction.Token)
}
