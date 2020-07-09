package profile_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/jmlopezz/uluru-api"
	"github.com/jmlopezz/uluru-api/internal/util"
	"github.com/jmlopezz/uluru-api/pkg/profile"
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

func TestGetProfile(t *testing.T) {
	testName := "TestGetProfile"

	addr, authAddr, err := before()
	if err != nil {
		t.Errorf("%s: before() failed: %v", testName, err)
		return
	}

	user, err := util.FactoryNewAuth(authAddr)
	if err != nil {
		t.Errorf("%s: util.FactoryNewAuth(authAddre) failed: %v", testName, err.Error())
		return
	}

	opts := uluru.ClientOptions{
		Token:       user.Meta.Token,
		Environment: "development",
	}

	before, after, err := profile.FactoryGet(addr, opts)
	if err != nil {
		t.Errorf("%s: util.FactoryProfileGet(authAddre) failed: %v", testName, err.Error())
		return
	}

	expected := after.ID
	if expected == before.ID {
		t.Errorf("%s: before.ID(\"\") failed, expected %v, got %v", testName, expected, before.ID)
		return
	}

	expected = after.UserID
	if expected == before.UserID {
		t.Errorf("%s: before.UserID(\"\") failed, expected %v, got %v", testName, expected, before.UserID)
		return
	}

	expected = after.Fingerprint
	if expected != before.Fingerprint {
		t.Errorf("%s: before.Fingerprint(\"\") failed, expected %v, got %v", testName, expected, before.Fingerprint)
		return
	}

	expectedNum := after.Age
	if expectedNum != before.Age {
		t.Errorf("%s: before.Age(\"\") failed, expected %v, got %v", testName, expectedNum, before.Age)
		return
	}

	expectedNum = after.Birth
	if expectedNum != before.Birth {
		t.Errorf("%s: before.Birth(\"\") failed, expected %v, got %v", testName, expectedNum, before.Birth)
		return
	}

	expected = after.MaritalStatus
	if expected != before.MaritalStatus {
		t.Errorf("%s: before.MaritalStatus(\"\") failed, expected %v, got %v", testName, expected, before.MaritalStatus)
		return
	}

	expectedNum = after.Childs
	if expectedNum != before.Childs {
		t.Errorf("%s: before.Childs(\"\") failed, expected %v, got %v", testName, expectedNum, before.Childs)
		return
	}
	expected = after.Gender
	if expected != before.Gender {
		t.Errorf("%s: before.Gender(\"\") failed, expected %v, got %v", testName, expected, before.Gender)
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

func TestUpdateProfile(t *testing.T) {
	testName := "TestUpdateProfile"

	addr, authAddr, err := before()
	if err != nil {
		t.Errorf("%s: before() failed: %v", testName, err)
		return
	}

	user, err := util.FactoryNewAuth(authAddr)
	if err != nil {
		t.Errorf("%s: util.FactoryNewAuth(authAddre) failed: %v", testName, err.Error())
		return
	}

	opts := uluru.ClientOptions{
		Token:       user.Meta.Token,
		Environment: "development",
	}

	before, after, err := profile.FactoryUpdate(addr, opts)
	if err != nil {
		t.Errorf("%s: util.FactoryProfileUpdate(authAddre) failed: %v", testName, err.Error())
		return
	}

	expected := after.ID
	if expected != before.ID {
		t.Errorf("%s: before.ID(\"\") failed, expected %v, got %v", testName, expected, before.ID)
		return
	}

	expected = after.UserID
	if expected != before.UserID {
		t.Errorf("%s: before.UserID(\"\") failed, expected %v, got %v", testName, expected, before.UserID)
		return
	}

	expected = after.Fingerprint
	if expected != before.Fingerprint {
		t.Errorf("%s: before.Fingerprint(\"\") failed, expected %v, got %v", testName, expected, before.Fingerprint)
		return
	}

	expectedNum := after.Age
	if expectedNum == before.Age {
		t.Errorf("%s: before.Age(\"\") failed, expectedNum %v, got %v", testName, expectedNum, before.Age)
		return
	}

	expectedNum = after.Birth
	if expectedNum == before.Birth {
		t.Errorf("%s: before.Birth(\"\") failed, expectedNum %v, got %v", testName, expectedNum, before.Birth)
		return
	}

	expectedNum = after.Childs
	if expectedNum == before.Childs {
		t.Errorf("%s: before.Childs(\"\") failed, expectedNum %v, got %v", testName, expectedNum, before.Childs)
		return
	}

	expected = after.MaritalStatus
	if expected == before.MaritalStatus {
		t.Errorf("%s: before.MaritalStatus(\"\") failed, expected %v, got %v", testName, expected, before.MaritalStatus)
		return
	}

	expected = after.Gender
	if expected == before.Gender {
		t.Errorf("%s: before.Gender(\"\") failed, expected %v, got %v", testName, expected, before.Gender)
		return
	}

	expectedBigNum := after.CreatedAt
	if expectedBigNum != before.CreatedAt {
		t.Errorf("%s: before.CreatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.CreatedAt)
		return
	}

	expectedBigNum = after.UpdatedAt
	if expectedBigNum <= before.UpdatedAt {
		t.Errorf("%s: before.UpdatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.UpdatedAt)
		return
	}
}

func TestCreateProfile(t *testing.T) {
	testName := "TestCreateProfile"

	addr, authAddr, err := before()
	if err != nil {
		t.Errorf("%s: before() failed: %v", testName, err)
		return
	}

	user, err := util.FactoryNewAuth(authAddr)
	if err != nil {
		t.Errorf("%s: util.FactoryNewAuth(authAddre) failed: %v", testName, err.Error())
		return
	}

	opts := uluru.ClientOptions{
		Token:       user.Meta.Token,
		Environment: "development",
	}

	before, after, err := profile.FactoryGet(addr, opts)
	if err != nil {
		t.Errorf("%s: util.FactoryProfileGet(authAddre) failed: %v", testName, err.Error())
		return
	}

	expected := after.ID
	if expected == before.ID {
		t.Errorf("%s: before.ID(\"\") failed, expected %v, got %v", testName, expected, before.ID)
		return
	}

	expected = after.UserID
	if expected == before.UserID {
		t.Errorf("%s: before.UserID(\"\") failed, expected %v, got %v", testName, expected, before.UserID)
		return
	}

	expected = after.Fingerprint
	if expected != before.Fingerprint {
		t.Errorf("%s: before.Fingerprint(\"\") failed, expected %v, got %v", testName, expected, before.Fingerprint)
		return
	}

	expectedNum := after.Age
	if expectedNum != after.Age {
		t.Errorf("%s: before.Age(\"\") failed, expectedNum %v, got %v", testName, expectedNum, after.Age)
		return
	}

	expectedNum = after.Birth
	if expectedNum != after.Birth {
		t.Errorf("%s: before.Birth(\"\") failed, expectedNum %v, got %v", testName, expectedNum, after.Birth)
		return
	}

	expectedNum = after.Childs
	if expectedNum != after.Childs {
		t.Errorf("%s: before.Childs(\"\") failed, expectedNum %v, got %v", testName, expectedNum, after.Childs)
		return
	}

	expected = after.Gender
	if expected != after.Gender {
		t.Errorf("%s: before.Gender(\"\") failed, expected %v, got %v", testName, expected, after.Gender)
		return
	}

	expected = after.MaritalStatus
	if expected != after.MaritalStatus {
		t.Errorf("%s: before.Childs(\"\") failed, expected %v, got %v", testName, expected, after.Childs)
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
