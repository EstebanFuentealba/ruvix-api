package subscriptions_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/jmlopezz/uluru-api"
	"github.com/jmlopezz/uluru-api/internal/util"
	"github.com/jmlopezz/uluru-api/pkg/subscriptions"
)

func before() (string, string, error) {
	host := os.Getenv("HOST")
	if host == "" {
		err := fmt.Errorf(fmt.Sprintf("Create: missing env variable HOST, failed with %s value", host))
		return "", "", err
	}

	port := os.Getenv("PORT")
	if port == "" {
		err := fmt.Errorf(fmt.Sprintf("Create: missing env variable PORT, failed with %s value", port))
		return "", "", err
	}

	authHost := os.Getenv("AUTH_HOST")
	if authHost == "" {
		err := fmt.Errorf(fmt.Sprintf("Create: missing env variable AUTH_HOST, failed with %s value", authHost))
		return "", "", err
	}

	authPort := os.Getenv("AUTH_PORT")
	if authPort == "" {
		err := fmt.Errorf(fmt.Sprintf("Create: missing env variable AUTH_PORT, failed with %s value", authPort))
		return "", "", err
	}

	return fmt.Sprintf("%s:%s", host, port), fmt.Sprintf("%s:%s", authHost, authPort), nil
}

func TestCreateSubscription(t *testing.T) {
	testName := "TestCreateSubscription"

	addr, authAddr, err := before()
	if err != nil {
		t.Errorf("%s: before() failed: %v", testName, err.Error())
		return
	}

	user, err := util.FactoryNewAuth(authAddr)
	if err != nil {
		t.Errorf("%s: util.FactoryNewAuth(authAddr) failed: %v", testName, err.Error())
		return
	}

	opts := uluru.ClientOptions{
		Token:       user.Meta.Token,
		Environment: "development",
	}

	before, after, err := subscriptions.FactoryCreateSubscription(addr, opts)
	if err != nil {
		t.Errorf("%s: subscriptions.FactoryCreateSubscription(addr, opts) failed: %v", testName, err.Error())
		return
	}

	expected := after.ID
	if expected == before.ID {
		t.Errorf("%s: before.ID(\"\") failed, expected %v, got %v", testName, expected, before.ID)
		return
	}

	expected = after.Name
	if expected != before.Name {
		t.Errorf("%s: before.Name(\"\") failed, expected %v, got %v", testName, expected, before.Name)
		return
	}

	expectedBigNum := after.CreatedAt
	if expectedBigNum == before.CreatedAt {
		t.Errorf("%s: before.CreatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.CreatedAt)
		return
	}

	expectedBigNum = after.UpdatedAt
	if expectedBigNum == before.UpdatedAt {
		t.Errorf("%s: before.UpdatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.UpdatedAt)
		return
	}
}
