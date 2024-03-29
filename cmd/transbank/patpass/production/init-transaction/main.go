package main

import (
	"log"

	"github.com/cagodoy/ruvix-api/pkg/transbank"
	"github.com/cagodoy/ruvix-api/pkg/transbank/webpay"
)

const (
	privateCert       = "private_cert_str"
	publicCert        = "public_cert_str"
	commerceCode      = 0
	commerceEmail     = "commerce_email"
	webpayService     = webpay.ServiceNormal
	webpayEnvironment = webpay.EnvironmentProduction
)

func main() {
	// InitTransanction base params
	amount := float64(10000)
	sessionID := "sesion1234557545"
	buyOrder := "ordenCompra12345678"
	returnURL := "https://callback/resultado/de/transaccion"
	finalURL := "https://callback/final/post/comprobante/webpay"

	// WPMDetail params
	serviceID := "123345567"
	cardHolderID := "12345"
	cardHolderName := "Juan"
	cardHolderLastName1 := "Perez"
	cardHolderLastName2 := "Gonzalez"
	cardHolderMail := "juan.perez@gmail.com"
	cellPhoneNumber := "9912345678"
	expirationDate := "2019-03-20T20:18:20Z"
	commerceMail := "contacto@comercio.cl"
	ufFlag := false

	service, err := webpay.NewPatpass(privateCert, publicCert, commerceCode, commerceEmail, webpayService, webpayEnvironment)
	if err != nil {
		log.Fatalln(err)
	}

	transaction, err := service.InitTransaction(transbank.InitTransaction{
		Amount:    amount,
		SessionID: sessionID,
		BuyOrder:  buyOrder,
		ReturnURL: returnURL,
		FinalURL:  finalURL,
		WPMDetail: &transbank.WPMDetail{
			ServiceID:           serviceID,
			CardHolderID:        cardHolderID,
			CardHolderName:      cardHolderName,
			CardHolderLastName1: cardHolderLastName1,
			CardHolderLastName2: cardHolderLastName2,
			CardHolderMail:      cardHolderMail,
			CellPhoneNumber:     cellPhoneNumber,
			ExpirationDate:      expirationDate,
			CommerceMail:        commerceMail,
			UfFlag:              ufFlag,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("URL", transaction.URL)
	log.Println("Token", transaction.Token)
}
