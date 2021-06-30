package subscriptions_test

// import (
// 	"fmt"
// 	"os"
// 	"testing"

// 	ruvixapi "github.com/cagodoy/ruvix-api"
// 	"github.com/cagodoy/ruvix-api/internal/util"
// 	"github.com/cagodoy/ruvix-api/pkg/subscriptions"
// )

// func before() (string, string, error) {
// 	host := os.Getenv("HOST")
// 	if host == "" {
// 		err := fmt.Errorf(fmt.Sprintf("Create: missing env variable HOST, failed with %s value", host))
// 		return "", "", err
// 	}

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		err := fmt.Errorf(fmt.Sprintf("Create: missing env variable PORT, failed with %s value", port))
// 		return "", "", err
// 	}

// 	authHost := os.Getenv("AUTH_HOST")
// 	if authHost == "" {
// 		err := fmt.Errorf(fmt.Sprintf("Create: missing env variable AUTH_HOST, failed with %s value", authHost))
// 		return "", "", err
// 	}

// 	authPort := os.Getenv("AUTH_PORT")
// 	if authPort == "" {
// 		err := fmt.Errorf(fmt.Sprintf("Create: missing env variable AUTH_PORT, failed with %s value", authPort))
// 		return "", "", err
// 	}

// 	return fmt.Sprintf("%s:%s", host, port), fmt.Sprintf("%s:%s", authHost, authPort), nil
// }

// func TestCreatePaySubscription(t *testing.T) {
// 	testName := "TestCreatePaySubscription"

// 	addr, authAddr, err := before()
// 	if err != nil {
// 		t.Errorf("%s: before() failed: %v", testName, err.Error())
// 		return
// 	}

// 	user, err := util.FactoryNewAuth(authAddr)
// 	if err != nil {
// 		t.Errorf("%s: util.FactoryNewAuth(authAddr) failed: %v", testName, err.Error())
// 		return
// 	}

// 	opts := ruvixapi.ClientOptions{
// 		Token:       user.Meta.Token,
// 		Environment: "development",
// 	}

// 	before, after, err := subscriptions.FactoryCreatePaySubscription(addr, opts)
// 	if err != nil {
// 		t.Errorf("%s: subscriptions.FactoryCreatePaySubscription(addr, opts) failed: %v", testName, err.Error())
// 		return
// 	}

// 	expected := after.ID
// 	if expected == before.ID {
// 		t.Errorf("%s: before.ID(\"\") failed, expected %v, got %v", testName, expected, before.ID)
// 		return
// 	}

// 	expected = after.Name
// 	if expected != before.Name {
// 		t.Errorf("%s: before.Name(\"\") failed, expected %v, got %v", testName, expected, before.Name)
// 		return
// 	}

// 	expectedFloatNum := after.Price
// 	if expectedFloatNum != before.Price {
// 		t.Errorf("%s: before.Price(\"\") failed, expectedFloatNum %v, got %v", testName, expectedFloatNum, before.Price)
// 		return
// 	}

// 	expectedNum := after.Months
// 	if expectedNum != before.Months {
// 		t.Errorf("%s: before.Months(\"\") failed, expectedNum %v, got %v", testName, expectedNum, before.Months)
// 		return
// 	}

// 	expectedBigNum := after.CreatedAt
// 	if expectedBigNum == before.CreatedAt {
// 		t.Errorf("%s: before.CreatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.CreatedAt)
// 		return
// 	}

// 	expectedBigNum = after.UpdatedAt
// 	if expectedBigNum == before.UpdatedAt {
// 		t.Errorf("%s: before.UpdatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.UpdatedAt)
// 		return
// 	}
// }

// func TestCreateFreeSubscription(t *testing.T) {
// 	testName := "TestCreateFreeSubscription"

// 	addr, authAddr, err := before()
// 	if err != nil {
// 		t.Errorf("%s: before() failed: %v", testName, err.Error())
// 		return
// 	}

// 	user, err := util.FactoryNewAuth(authAddr)
// 	if err != nil {
// 		t.Errorf("%s: util.FactoryNewAuth(authAddr) failed: %v", testName, err.Error())
// 		return
// 	}

// 	opts := ruvixapi.ClientOptions{
// 		Token:       user.Meta.Token,
// 		Environment: "development",
// 	}

// 	before, after, err := subscriptions.FactoryCreateFreeSubscription(addr, opts)
// 	if err != nil {
// 		t.Errorf("%s: subscriptions.FactoryCreateFreeSubscription(addr, opts) failed: %v", testName, err.Error())
// 		return
// 	}

// 	expected := after.ID
// 	if expected == before.ID {
// 		t.Errorf("%s: before.ID(\"\") failed, expected %v, got %v", testName, expected, before.ID)
// 		return
// 	}

// 	expected = after.Name
// 	if expected != before.Name {
// 		t.Errorf("%s: before.Name(\"\") failed, expected %v, got %v", testName, expected, before.Name)
// 		return
// 	}

// 	expectedFloatNum := after.Price
// 	if expectedFloatNum != before.Price {
// 		t.Errorf("%s: before.Price(\"\") failed, expectedFloatNum %v, got %v", testName, expectedFloatNum, before.Price)
// 		return
// 	}

// 	expectedNum := after.Months
// 	if expectedNum != before.Months {
// 		t.Errorf("%s: before.Months(\"\") failed, expectedNum %v, got %v", testName, expectedNum, before.Months)
// 		return
// 	}

// 	expectedBigNum := after.CreatedAt
// 	if expectedBigNum == before.CreatedAt {
// 		t.Errorf("%s: before.CreatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.CreatedAt)
// 		return
// 	}

// 	expectedBigNum = after.UpdatedAt
// 	if expectedBigNum == before.UpdatedAt {
// 		t.Errorf("%s: before.UpdatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.UpdatedAt)
// 		return
// 	}
// }

// func TestListSubscriptions(t *testing.T) {
// 	testName := "TestListSubscriptions"

// 	addr, authAddr, err := before()
// 	if err != nil {
// 		t.Errorf("%s: before() failed: %v", testName, err.Error())
// 		return
// 	}

// 	user, err := util.FactoryNewAuth(authAddr)
// 	if err != nil {
// 		t.Errorf("%s: util.FactoryNewAuth(authAddre) failed: %v", testName, err.Error())
// 		return
// 	}

// 	opts := ruvixapi.ClientOptions{
// 		Token:       user.Meta.Token,
// 		Environment: "development",
// 	}

// 	before, subscriptions, err := subscriptions.FactoryListSubscriptions(addr, opts)
// 	if err != nil {
// 		t.Errorf("%s: subscriptions.FactoryListSubscriptions(addr, opts) failed: %s", testName, err.Error())
// 		return
// 	}

// 	var index int
// 	for i := 0; i < len(subscriptions); i++ {
// 		if subscriptions[i].ID == before.ID {
// 			index = i
// 			break
// 		}
// 	}

// 	after := subscriptions[index]

// 	expected := after.ID
// 	if expected != before.ID {
// 		t.Errorf("%s: after.ID(\"\") failed, expected %v, got %v", testName, expected, before.ID)
// 		return
// 	}

// 	expected = after.Name
// 	if expected != before.Name {
// 		t.Errorf("%s: before.Name(\"\") failed, expected %v, got %v", testName, expected, before.Name)
// 		return
// 	}

// 	expectedFloatNum := after.Price
// 	if expectedFloatNum != before.Price {
// 		t.Errorf("%s: before.Price(\"\") failed, expectedFloatNum %v, got %v", testName, expectedFloatNum, before.Price)
// 		return
// 	}

// 	expectedNum := after.Months
// 	if expectedNum != before.Months {
// 		t.Errorf("%s: before.Months(\"\") failed, expectedNum %v, got %v", testName, expectedNum, before.Months)
// 		return
// 	}

// 	expectedBigNum := before.CreatedAt
// 	if expectedBigNum == 0 {
// 		t.Errorf("%s: after.CreatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.CreatedAt)
// 		return
// 	}

// 	expectedBigNum = before.UpdatedAt
// 	if expectedBigNum == 0 {
// 		t.Errorf("%s: after.UpdatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.UpdatedAt)
// 		return
// 	}
// }

// func TestListProviders(t *testing.T) {
// 	testName := "TestListProviders"

// 	addr, authAddr, err := before()
// 	if err != nil {
// 		t.Errorf("%s: before() failed: %v", testName, err.Error())
// 		return
// 	}

// 	user, err := util.FactoryNewAuth(authAddr)
// 	if err != nil {
// 		t.Errorf("%s: util.FactoryNewAuth(authAddre) failed: %v", testName, err.Error())
// 		return
// 	}

// 	opts := ruvixapi.ClientOptions{
// 		Token:       user.Meta.Token,
// 		Environment: "development",
// 	}

// 	providers, err := subscriptions.FactoryListProviders(addr, opts)
// 	if err != nil {
// 		t.Errorf("%s: subscriptions.FactoryListProviders(addr, opts) failed: %s", testName, err.Error())
// 		return
// 	}

// 	for _, v := range providers {
// 		if !subscriptions.ValidProvider(v.ID) {
// 			t.Errorf("%s: ID(\"\") failed, %v is invalid", testName, v.ID)
// 			return
// 		}
// 	}
// }

// func TestPaySubscribe(t *testing.T) {
// 	testName := "TestPaySubscribe"

// 	addr, authAddr, err := before()
// 	if err != nil {
// 		t.Errorf("%s: before() failed: %v", testName, err.Error())
// 		return
// 	}

// 	user, err := util.FactoryNewAuth(authAddr)
// 	if err != nil {
// 		t.Errorf("%s: util.FactoryNewAuth(authAddr) failed: %v", testName, err.Error())
// 		return
// 	}

// 	opts := ruvixapi.ClientOptions{
// 		Token:       user.Meta.Token,
// 		Environment: "development",
// 	}

// 	subscription, transaction, err := subscriptions.FactoryPaySubscribe(addr, opts)
// 	if err != nil {
// 		t.Errorf("%s: subscriptions.FactoryPaySubscribe(addr, opts) failed: %v", testName, err.Error())
// 		return
// 	}

// 	expected := transaction.ID
// 	if expected == "" {
// 		t.Errorf("%s: ID(\"\") failed, expected %v, got %v", testName, expected, "")
// 		return
// 	}

// 	expected = transaction.UserID
// 	if expected != user.Data.ID {
// 		t.Errorf("%s: user.Data.ID(\"\") failed, expected %v, got %v", testName, expected, user.Data.ID)
// 		return
// 	}

// 	expected = transaction.SubscriptionID
// 	if expected != subscription.ID {
// 		t.Errorf("%s: subscription.ID(\"\") failed, expected %v, got %v", testName, expected, subscription.ID)
// 		return
// 	}

// 	expected = transaction.ProviderID
// 	if !subscriptions.ValidProvider(expected) {
// 		t.Errorf("%s: subscriptions.ValidProvider()(\"\") failed, expected %v, got %v", testName, expected, subscriptions.ValidProvider(expected))
// 		return
// 	}

// 	expectedBigNum := transaction.DueDate
// 	if expectedBigNum == 0 {
// 		t.Errorf("%s: DueDate(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, 0)
// 		return
// 	}

// 	expectedBigNum = transaction.RemindedAt
// 	if expectedBigNum == 0 {
// 		t.Errorf("%s: RemindedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, 0)
// 		return
// 	}

// 	expected = transaction.Status
// 	if expected != subscriptions.StatusTransactionPending {
// 		t.Errorf("%s: subscriptions.StatusTransactionPending(\"\") failed, expected %v, got %v", testName, expected, subscriptions.StatusTransactionPending)
// 		return
// 	}

// 	expected = transaction.PaymentToken
// 	if expected == "" {
// 		t.Errorf("%s: PaymentToken(\"\") failed, expected %v, got %v", testName, expected, "")
// 		return
// 	}

// 	expected = transaction.OrderNumber
// 	if expected == "" {
// 		t.Errorf("%s: OrderNumber(\"\") failed, expected %v, got %v", testName, expected, "")
// 		return
// 	}

// 	expectedBigNum = transaction.CreatedAt
// 	if expectedBigNum == 0 {
// 		t.Errorf("%s: CreatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, 0)
// 		return
// 	}

// 	expectedBigNum = transaction.UpdatedAt
// 	if expectedBigNum == 0 {
// 		t.Errorf("%s: UpdatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, 0)
// 		return
// 	}
// }

// func TestFreeSubscribe(t *testing.T) {
// 	testName := "TestFreeSubscribe"

// 	addr, authAddr, err := before()
// 	if err != nil {
// 		t.Errorf("%s: before() failed: %v", testName, err.Error())
// 		return
// 	}

// 	user, err := util.FactoryNewAuth(authAddr)
// 	if err != nil {
// 		t.Errorf("%s: util.FactoryNewAuth(authAddr) failed: %v", testName, err.Error())
// 		return
// 	}

// 	opts := ruvixapi.ClientOptions{
// 		Token:       user.Meta.Token,
// 		Environment: "development",
// 	}

// 	subscription, transaction, err := subscriptions.FactoryFreeSubscribe(addr, opts)
// 	if err != nil {
// 		t.Errorf("%s: subscriptions.FactoryFreeSubscribe(addr, opts) failed: %v", testName, err.Error())
// 		return
// 	}

// 	expected := transaction.ID
// 	if expected == "" {
// 		t.Errorf("%s: ID(\"\") failed, expected %v, got %v", testName, expected, "")
// 		return
// 	}

// 	expected = transaction.UserID
// 	if expected != user.Data.ID {
// 		t.Errorf("%s: user.Data.ID(\"\") failed, expected %v, got %v", testName, expected, user.Data.ID)
// 		return
// 	}

// 	expected = transaction.SubscriptionID
// 	if expected != subscription.ID {
// 		t.Errorf("%s: subscription.ID(\"\") failed, expected %v, got %v", testName, expected, subscription.ID)
// 		return
// 	}

// 	expected = transaction.ProviderID
// 	if !subscriptions.ValidProvider(expected) {
// 		t.Errorf("%s: subscriptions.ValidProvider()(\"\") failed, expected %v, got %v", testName, expected, subscriptions.ValidProvider(expected))
// 		return
// 	}

// 	expectedBigNum := transaction.DueDate
// 	if expectedBigNum != 0 {
// 		t.Errorf("%s: DueDate(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, 0)
// 		return
// 	}

// 	expectedBigNum = transaction.RemindedAt
// 	if expectedBigNum != 0 {
// 		t.Errorf("%s: RemindedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, 0)
// 		return
// 	}

// 	expected = transaction.Status
// 	if expected != subscriptions.StatusTransactionCompleted {
// 		t.Errorf("%s: subscriptions.StatusTransactionCompleted(\"\") failed, expected %v, got %v", testName, expected, subscriptions.StatusTransactionCompleted)
// 		return
// 	}

// 	expected = transaction.PaymentToken
// 	if expected != "" {
// 		t.Errorf("%s: PaymentToken(\"\") failed, expected %v, got %v", testName, expected, "")
// 		return
// 	}

// 	expected = transaction.OrderNumber
// 	if expected == "" {
// 		t.Errorf("%s: OrderNumber(\"\") failed, expected %v, got %v", testName, expected, "")
// 		return
// 	}

// 	expectedBigNum = transaction.CreatedAt
// 	if expectedBigNum == 0 {
// 		t.Errorf("%s: CreatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, 0)
// 		return
// 	}

// 	expectedBigNum = transaction.UpdatedAt
// 	if expectedBigNum == 0 {
// 		t.Errorf("%s: UpdatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, 0)
// 		return
// 	}
// }

// func TestPayUnsubscribe(t *testing.T) {
// 	testName := "TestPayUnsubscribe"

// 	addr, authAddr, err := before()
// 	if err != nil {
// 		t.Errorf("%s: before() failed: %v", testName, err.Error())
// 		return
// 	}

// 	user, err := util.FactoryNewAuth(authAddr)
// 	if err != nil {
// 		t.Errorf("%s: util.FactoryNewAuth(authAddr) failed: %v", testName, err.Error())
// 		return
// 	}

// 	opts := ruvixapi.ClientOptions{
// 		Token:       user.Meta.Token,
// 		Environment: "development",
// 	}

// 	before, after, err := subscriptions.FactoryPayUnsubscribe(addr, opts)
// 	if err != nil {
// 		t.Errorf("%s: subscriptions.FactoryPayUnsubscribe(addr, opts) failed: %v", testName, err.Error())
// 		return
// 	}

// 	expected := after.ID
// 	if expected != before.ID {
// 		t.Errorf("%s: before.ID(\"\") failed, expected %v, got %v", testName, expected, before.ID)
// 		return
// 	}

// 	expected = after.UserID
// 	if expected != user.Data.ID {
// 		t.Errorf("%s: user.Data.ID(\"\") failed, expected %v, got %v", testName, expected, user.Data.ID)
// 		return
// 	}

// 	expected = after.UserID
// 	if expected != before.UserID {
// 		t.Errorf("%s: before.UserID(\"\") failed, expected %v, got %v", testName, expected, before.UserID)
// 		return
// 	}

// 	expected = after.SubscriptionID
// 	if expected != before.SubscriptionID {
// 		t.Errorf("%s: before.SubscriptionID(\"\") failed, expected %v, got %v", testName, expected, before.SubscriptionID)
// 		return
// 	}

// 	expected = after.ProviderID
// 	if expected != before.ProviderID {
// 		t.Errorf("%s: before.ProviderID(\"\") failed, expected %v, got %v", testName, expected, before.ProviderID)
// 		return
// 	}

// 	expectedBigNum := after.DueDate
// 	if expectedBigNum != before.DueDate {
// 		t.Errorf("%s: before.DueDate(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.DueDate)
// 		return
// 	}

// 	expectedBigNum = after.RemindedAt
// 	if expectedBigNum != before.RemindedAt {
// 		t.Errorf("%s: before.RemindedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.RemindedAt)
// 		return
// 	}

// 	expected = after.Status
// 	if expected != subscriptions.StatusTransactionCanceled {
// 		t.Errorf("%s: subscriptions.StatusTransactionCanceled(\"\") failed, expected %v, got %v", testName, expected, subscriptions.StatusTransactionCanceled)
// 		return
// 	}

// 	expected = after.PaymentToken
// 	if expected != before.PaymentToken {
// 		t.Errorf("%s: before.PaymentToken(\"\") failed, expected %v, got %v", testName, expected, before.PaymentToken)
// 		return
// 	}

// 	expected = after.OrderNumber
// 	if expected != before.OrderNumber {
// 		t.Errorf("%s: before.OrderNumber(\"\") failed, expected %v, got %v", testName, expected, before.OrderNumber)
// 		return
// 	}

// 	expectedBigNum = after.CreatedAt
// 	if expectedBigNum != before.CreatedAt {
// 		t.Errorf("%s: before.CreatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.CreatedAt)
// 		return
// 	}

// 	expectedBigNum = after.UpdatedAt
// 	if expectedBigNum < before.UpdatedAt {
// 		t.Errorf("%s: before.UpdatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.UpdatedAt)
// 		return
// 	}
// }

// func TestFreeUnsubscribe(t *testing.T) {
// 	testName := "TestFreeUnsubscribe"

// 	addr, authAddr, err := before()
// 	if err != nil {
// 		t.Errorf("%s: before() failed: %v", testName, err.Error())
// 		return
// 	}

// 	user, err := util.FactoryNewAuth(authAddr)
// 	if err != nil {
// 		t.Errorf("%s: util.FactoryNewAuth(authAddr) failed: %v", testName, err.Error())
// 		return
// 	}

// 	opts := ruvixapi.ClientOptions{
// 		Token:       user.Meta.Token,
// 		Environment: "development",
// 	}

// 	_, _, err = subscriptions.FactoryFreeUnsubscribe(addr, opts)
// 	if err != nil && err.Error() != "transaction already completed or cancelled" {
// 		t.Errorf("%s: subscriptions.FactoryFreeUnsubscribe(addr, opts) failed: %v", testName, err.Error())
// 		return
// 	}
// }

// func TestPayRefresh(t *testing.T) {
// 	testName := "TestPayRefresh"

// 	addr, authAddr, err := before()
// 	if err != nil {
// 		t.Errorf("%s: before() failed: %v", testName, err.Error())
// 		return
// 	}

// 	user, err := util.FactoryNewAuth(authAddr)
// 	if err != nil {
// 		t.Errorf("%s: util.FactoryNewAuth(authAddr) failed: %v", testName, err.Error())
// 		return
// 	}

// 	opts := ruvixapi.ClientOptions{
// 		Token:       user.Meta.Token,
// 		Environment: "development",
// 	}

// 	before, after, err := subscriptions.FactoryPayRefresh(addr, opts)
// 	if err != nil {
// 		t.Errorf("%s: subscriptions.FactoryPayRefresh(addr, opts) failed: %v", testName, err.Error())
// 		return
// 	}

// 	expected := after.ID
// 	if expected != before.ID {
// 		t.Errorf("%s: before.ID(\"\") failed, expected %v, got %v", testName, expected, before.ID)
// 		return
// 	}

// 	expected = after.UserID
// 	if expected != user.Data.ID {
// 		t.Errorf("%s: user.Data.ID(\"\") failed, expected %v, got %v", testName, expected, user.Data.ID)
// 		return
// 	}

// 	expected = after.UserID
// 	if expected != before.UserID {
// 		t.Errorf("%s: before.UserID(\"\") failed, expected %v, got %v", testName, expected, before.UserID)
// 		return
// 	}

// 	expected = after.SubscriptionID
// 	if expected != before.SubscriptionID {
// 		t.Errorf("%s: before.SubscriptionID(\"\") failed, expected %v, got %v", testName, expected, before.SubscriptionID)
// 		return
// 	}

// 	expected = after.ProviderID
// 	if expected != before.ProviderID {
// 		t.Errorf("%s: before.ProviderID(\"\") failed, expected %v, got %v", testName, expected, before.ProviderID)
// 		return
// 	}

// 	expectedBigNum := after.DueDate
// 	if expectedBigNum != before.DueDate {
// 		t.Errorf("%s: before.DueDate(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.DueDate)
// 		return
// 	}

// 	expectedBigNum = after.RemindedAt
// 	if expectedBigNum != before.RemindedAt {
// 		t.Errorf("%s: before.RemindedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.RemindedAt)
// 		return
// 	}

// 	expected = after.Status
// 	if expected != subscriptions.StatusTransactionPending {
// 		t.Errorf("%s: subscriptions.StatusTransactionPending(\"\") failed, expected %v, got %v", testName, expected, subscriptions.StatusTransactionPending)
// 		return
// 	}

// 	expected = after.PaymentToken
// 	if expected == before.PaymentToken {
// 		t.Errorf("%s: before.PaymentToken(\"\") failed, expected %v, got %v", testName, expected, before.PaymentToken)
// 		return
// 	}

// 	expected = after.OrderNumber
// 	if expected != before.OrderNumber {
// 		t.Errorf("%s: before.OrderNumber(\"\") failed, expected %v, got %v", testName, expected, before.OrderNumber)
// 		return
// 	}

// 	expectedBigNum = after.CreatedAt
// 	if expectedBigNum != before.CreatedAt {
// 		t.Errorf("%s: before.CreatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.CreatedAt)
// 		return
// 	}

// 	expectedBigNum = after.UpdatedAt
// 	if expectedBigNum < before.UpdatedAt {
// 		t.Errorf("%s: before.UpdatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.UpdatedAt)
// 		return
// 	}
// }

// func TestFreeRefresh(t *testing.T) {
// 	testName := "TestFreeRefresh"

// 	addr, authAddr, err := before()
// 	if err != nil {
// 		t.Errorf("%s: before() failed: %v", testName, err.Error())
// 		return
// 	}

// 	user, err := util.FactoryNewAuth(authAddr)
// 	if err != nil {
// 		t.Errorf("%s: util.FactoryNewAuth(authAddr) failed: %v", testName, err.Error())
// 		return
// 	}

// 	opts := ruvixapi.ClientOptions{
// 		Token:       user.Meta.Token,
// 		Environment: "development",
// 	}

// 	_, _, err = subscriptions.FactoryFreeRefresh(addr, opts)
// 	if err != nil && err.Error() != "transaction already completed or cancelled" {
// 		t.Errorf("%s: subscriptions.FactoryFreeRefresh(addr, opts) failed: %v", testName, err.Error())
// 		return
// 	}
// }

// // TODO(ca): should implement TestPayVerify testing method with mock on interface

// func TestUnpayVerify(t *testing.T) {
// 	testName := "TestPayVerify"

// 	addr, authAddr, err := before()
// 	if err != nil {
// 		t.Errorf("%s: before() failed: %v", testName, err.Error())
// 		return
// 	}

// 	user, err := util.FactoryNewAuth(authAddr)
// 	if err != nil {
// 		t.Errorf("%s: util.FactoryNewAuth(authAddr) failed: %v", testName, err.Error())
// 		return
// 	}

// 	opts := ruvixapi.ClientOptions{
// 		Token:       user.Meta.Token,
// 		Environment: "development",
// 	}

// 	_, _, err = subscriptions.FactoryPayVerify(addr, opts)
// 	if err != nil && err.Error() != "Error: code=soap:Server message=<!-- Timeout error(272) -->" {
// 		t.Errorf("%s: subscriptions.FactoryPayVerify(addr, opts) failed: %v", testName, err.Error())
// 		return
// 	}
// }

// func TestFreeVerify(t *testing.T) {
// 	testName := "TestFreeVerify"

// 	addr, authAddr, err := before()
// 	if err != nil {
// 		t.Errorf("%s: before() failed: %v", testName, err.Error())
// 		return
// 	}

// 	user, err := util.FactoryNewAuth(authAddr)
// 	if err != nil {
// 		t.Errorf("%s: util.FactoryNewAuth(authAddr) failed: %v", testName, err.Error())
// 		return
// 	}

// 	opts := ruvixapi.ClientOptions{
// 		Token:       user.Meta.Token,
// 		Environment: "development",
// 	}

// 	_, _, err = subscriptions.FactoryFreeVerify(addr, opts)
// 	if err != nil && err.Error() != "transaction status already complete" {
// 		t.Errorf("%s: subscriptions.FactoryFreeVerify(addr, opts) failed: %v", testName, err.Error())
// 		return
// 	}
// }

// func TestListTransactions(t *testing.T) {
// 	testName := "TestListTransactions"

// 	addr, authAddr, err := before()
// 	if err != nil {
// 		t.Errorf("%s: before() failed: %v", testName, err.Error())
// 		return
// 	}

// 	user, err := util.FactoryNewAuth(authAddr)
// 	if err != nil {
// 		t.Errorf("%s: util.FactoryNewAuth(authAddre) failed: %v", testName, err.Error())
// 		return
// 	}

// 	opts := ruvixapi.ClientOptions{
// 		Token:       user.Meta.Token,
// 		Environment: "development",
// 	}

// 	before, transactions, err := subscriptions.ListTransactions(addr, opts)
// 	if err != nil {
// 		t.Errorf("%s: subscriptions.ListTransactions(addr, opts) failed: %s", testName, err.Error())
// 		return
// 	}

// 	var index int
// 	for i := 0; i < len(transactions); i++ {
// 		if transactions[i].ID == before.ID {
// 			index = i
// 			break
// 		}
// 	}

// 	after := transactions[index]

// 	expected := after.ID
// 	if expected != before.ID {
// 		t.Errorf("%s: before.ID(\"\") failed, expected %v, got %v", testName, expected, before.ID)
// 		return
// 	}

// 	expected = after.UserID
// 	if expected != user.Data.ID {
// 		t.Errorf("%s: user.Data.ID(\"\") failed, expected %v, got %v", testName, expected, user.Data.ID)
// 		return
// 	}

// 	expected = after.UserID
// 	if expected != before.UserID {
// 		t.Errorf("%s: before.UserID(\"\") failed, expected %v, got %v", testName, expected, before.UserID)
// 		return
// 	}

// 	expected = after.SubscriptionID
// 	if expected != before.SubscriptionID {
// 		t.Errorf("%s: before.SubscriptionID(\"\") failed, expected %v, got %v", testName, expected, before.SubscriptionID)
// 		return
// 	}

// 	expected = after.ProviderID
// 	if expected != before.ProviderID {
// 		t.Errorf("%s: before.ProviderID(\"\") failed, expected %v, got %v", testName, expected, before.ProviderID)
// 		return
// 	}

// 	expectedBigNum := after.DueDate
// 	if expectedBigNum != before.DueDate {
// 		t.Errorf("%s: before.DueDate(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.DueDate)
// 		return
// 	}

// 	expectedBigNum = after.RemindedAt
// 	if expectedBigNum != before.RemindedAt {
// 		t.Errorf("%s: before.RemindedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.RemindedAt)
// 		return
// 	}

// 	expected = after.Status
// 	if expected != subscriptions.StatusTransactionPending {
// 		t.Errorf("%s: subscriptions.StatusTransactionPending(\"\") failed, expected %v, got %v", testName, expected, subscriptions.StatusTransactionPending)
// 		return
// 	}

// 	expected = after.PaymentToken
// 	if expected != before.PaymentToken {
// 		t.Errorf("%s: before.PaymentToken(\"\") failed, expected %v, got %v", testName, expected, before.PaymentToken)
// 		return
// 	}

// 	expected = after.OrderNumber
// 	if expected != before.OrderNumber {
// 		t.Errorf("%s: before.OrderNumber(\"\") failed, expected %v, got %v", testName, expected, before.OrderNumber)
// 		return
// 	}

// 	expectedBigNum = after.CreatedAt
// 	if expectedBigNum != before.CreatedAt {
// 		t.Errorf("%s: before.CreatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.CreatedAt)
// 		return
// 	}

// 	expectedBigNum = after.UpdatedAt
// 	if expectedBigNum != before.UpdatedAt {
// 		t.Errorf("%s: before.UpdatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.UpdatedAt)
// 		return
// 	}
// }

// func TestLastTransaction(t *testing.T) {
// 	testName := "TestLastTransaction"

// 	addr, authAddr, err := before()
// 	if err != nil {
// 		t.Errorf("%s: before() failed: %v", testName, err.Error())
// 		return
// 	}

// 	user, err := util.FactoryNewAuth(authAddr)
// 	if err != nil {
// 		t.Errorf("%s: util.FactoryNewAuth(authAddre) failed: %v", testName, err.Error())
// 		return
// 	}

// 	opts := ruvixapi.ClientOptions{
// 		Token:       user.Meta.Token,
// 		Environment: "development",
// 	}

// 	before, after, err := subscriptions.LastTransaction(addr, opts)
// 	if err != nil {
// 		t.Errorf("%s: subscriptions.LastTransaction(addr, opts) failed: %s", testName, err.Error())
// 		return
// 	}

// 	expected := after.ID
// 	if expected != before.ID {
// 		t.Errorf("%s: before.ID(\"\") failed, expected %v, got %v", testName, expected, before.ID)
// 		return
// 	}

// 	expected = after.UserID
// 	if expected != user.Data.ID {
// 		t.Errorf("%s: user.Data.ID(\"\") failed, expected %v, got %v", testName, expected, user.Data.ID)
// 		return
// 	}

// 	expected = after.UserID
// 	if expected != before.UserID {
// 		t.Errorf("%s: before.UserID(\"\") failed, expected %v, got %v", testName, expected, before.UserID)
// 		return
// 	}

// 	expected = after.SubscriptionID
// 	if expected != before.SubscriptionID {
// 		t.Errorf("%s: before.SubscriptionID(\"\") failed, expected %v, got %v", testName, expected, before.SubscriptionID)
// 		return
// 	}

// 	expected = after.ProviderID
// 	if expected != before.ProviderID {
// 		t.Errorf("%s: before.ProviderID(\"\") failed, expected %v, got %v", testName, expected, before.ProviderID)
// 		return
// 	}

// 	expectedBigNum := after.DueDate
// 	if expectedBigNum != before.DueDate {
// 		t.Errorf("%s: before.DueDate(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.DueDate)
// 		return
// 	}

// 	expectedBigNum = after.RemindedAt
// 	if expectedBigNum != before.RemindedAt {
// 		t.Errorf("%s: before.RemindedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.RemindedAt)
// 		return
// 	}

// 	expected = after.Status
// 	if expected != subscriptions.StatusTransactionPending {
// 		t.Errorf("%s: subscriptions.StatusTransactionPending(\"\") failed, expected %v, got %v", testName, expected, subscriptions.StatusTransactionPending)
// 		return
// 	}

// 	expected = after.PaymentToken
// 	if expected != before.PaymentToken {
// 		t.Errorf("%s: before.PaymentToken(\"\") failed, expected %v, got %v", testName, expected, before.PaymentToken)
// 		return
// 	}

// 	expected = after.OrderNumber
// 	if expected != before.OrderNumber {
// 		t.Errorf("%s: before.OrderNumber(\"\") failed, expected %v, got %v", testName, expected, before.OrderNumber)
// 		return
// 	}

// 	expectedBigNum = after.CreatedAt
// 	if expectedBigNum != before.CreatedAt {
// 		t.Errorf("%s: before.CreatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.CreatedAt)
// 		return
// 	}

// 	expectedBigNum = after.UpdatedAt
// 	if expectedBigNum != before.UpdatedAt {
// 		t.Errorf("%s: before.UpdatedAt(\"\") failed, expectedBigNum %v, got %v", testName, expectedBigNum, before.UpdatedAt)
// 		return
// 	}
// }
