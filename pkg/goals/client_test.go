package goals_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/jmlopezz/uluru-api"
	"github.com/jmlopezz/uluru-api/internal/util"
	"github.com/jmlopezz/uluru-api/pkg/goals"
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

func TestCreateGoal(t *testing.T) {
	testName := "TestCreateGoal"

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

	before, after, err := goals.FactoryCreateGoal(addr, opts)
	if err != nil {
		t.Errorf("%s: goals.FactoryCreateGoal(addr, opts) failed: %v", testName, err.Error())
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

func TestListGoals(t *testing.T) {
	testName := "TestListGoals"

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

	before, goals, err := goals.FactoryListGoals(addr, opts)
	if err != nil {
		t.Errorf("%s: goals.FactoryListGoals(addr, opts) failed: %v", testName, err.Error())
		return
	}

	var index int
	for i := 0; i < len(goals); i++ {
		if goals[i].ID == before.ID {
			index = i
			break
		}
	}

	goal := goals[index]

	expected := goal.ID
	if expected != before.ID {
		t.Errorf("%s: goal.ID(\"\") failed, expected %v, got %v", testName, expected, before.ID)
		return
	}

	expected = goal.Name
	if expected != before.Name {
		t.Errorf("%s: goal.Name(\"\") failed, expected %v, got %v", testName, expected, before.Name)
		return
	}

	expectedBigNum := before.CreatedAt
	if expectedBigNum == 0 {
		t.Errorf("%s: before.CreatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.CreatedAt)
		return
	}

	expectedBigNum = before.UpdatedAt
	if expectedBigNum == 0 {
		t.Errorf("%s: before.UpdatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.UpdatedAt)
		return
	}
}

func TestCreateRetirementGoal(t *testing.T) {
	testName := "TestCreateRetirementGoal"

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

	before, after, err := goals.FactoryCreateRetirementGoal(addr, opts)
	if err != nil {
		t.Errorf("%s: goals.FactoryCreateRetirementGoal(addr, opts) failed: %v", testName, err.Error())
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

	expected = after.GoalID
	if expected != before.GoalID {
		t.Errorf("%s: before.GoalID(\"\") failed, expected %v, got %v", testName, expected, before.GoalID)
		return
	}

	expectedFloatNum := after.MonthlySalary
	if expectedFloatNum != before.MonthlySalary {
		t.Errorf("%s: before.MonthlySalary(\"\") failed, expectedFloatNum %v, got %v", testName, expectedFloatNum, before.MonthlySalary)
		return
	}

	expectedFloatNum = after.MonthlyRetirement
	if expectedFloatNum != before.MonthlyRetirement {
		t.Errorf("%s: before.MonthlyRetirement(\"\") failed, expectedFloatNum %v, got %v", testName, expectedFloatNum, before.MonthlyRetirement)
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

func TestGetLastRetirementGoal(t *testing.T) {
	testName := "TestGetLastRetirementGoal"

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

	before, after, err := goals.FactoryGetLastRetirementGoal(addr, opts)
	if err != nil {
		t.Errorf("%s: goals.FactoryGetLastRetirementGoal(addr, opts) failed: %s", testName, err.Error())
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

	expected = after.GoalID
	if expected != before.GoalID {
		t.Errorf("%s: before.GoalID(\"\") failed, expected %v, got %v", testName, expected, before.GoalID)
		return
	}

	expectedFloatNum := after.MonthlySalary
	if expectedFloatNum != before.MonthlySalary {
		t.Errorf("%s: before.MonthlySalary(\"\") failed, expectedFloatNum %v, got %v", testName, expectedFloatNum, before.MonthlySalary)
		return
	}

	expectedFloatNum = after.MonthlyRetirement
	if expectedFloatNum != before.MonthlyRetirement {
		t.Errorf("%s: before.MonthlyRetirement(\"\") failed, expectedFloatNum %v, got %v", testName, expectedFloatNum, before.MonthlyRetirement)
		return
	}

	expectedBigNum := after.CreatedAt
	if expectedBigNum != before.CreatedAt {
		t.Errorf("%s: CreatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.CreatedAt)
		return
	}

	expectedBigNum = after.UpdatedAt
	if expectedBigNum != before.UpdatedAt {
		t.Errorf("%s: UpdatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.UpdatedAt)
		return
	}
}

func TestCreateGuestRetirementGoal(t *testing.T) {
	testName := "TestCreateGuestRetirementGoal"

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

	before, after, err := goals.FactoryCreateGuestRetirementGoal(addr, opts)
	if err != nil {
		t.Errorf("%s: goals.FactoryCreateGuestRetirementGoal(addr, opts) failed: %v", testName, err.Error())
		return
	}

	expected := after.ID
	if expected == before.ID {
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

	expected = after.GoalID
	if expected != before.GoalID {
		t.Errorf("%s: before.GoalID(\"\") failed, expected %v, got %v", testName, expected, before.GoalID)
		return
	}

	expectedFloatNum := after.MonthlySalary
	if expectedFloatNum != before.MonthlySalary {
		t.Errorf("%s: before.MonthlySalary(\"\") failed, expectedFloatNum %v, got %v", testName, expectedFloatNum, before.MonthlySalary)
		return
	}

	expectedFloatNum = after.MonthlyRetirement
	if expectedFloatNum != before.MonthlyRetirement {
		t.Errorf("%s: before.MonthlyRetirement(\"\") failed, expectedFloatNum %v, got %v", testName, expectedFloatNum, before.MonthlyRetirement)
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
