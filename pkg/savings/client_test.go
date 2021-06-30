package savings_test

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

// func TestCreateInstitution(t *testing.T) {
// 	testName := "TestCreateInstitution"

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

// 	before, after, err := savings.FactoryCreateInstitution(addr, opts)
// 	if err != nil {
// 		t.Errorf("%s: savings.FactoryCreateInstitution(addr, opts) failed: %s", testName, err.Error())
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

// 	expected = after.Slug
// 	if expected != before.Slug {
// 		t.Errorf("%s: before.Slug(\"\") failed, expected %v, got %v", testName, expected, before.Slug)
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

// func TestCreateInstitutionWithAccounts(t *testing.T) {
// 	testName := "TestCreateInstitutionWithAccounts"

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

// 	before, after, err := savings.FactoryCreateInstitutionWithAccounts(addr, opts)
// 	if err != nil {
// 		t.Errorf("%s: savings.FactoryCreateInstitutionWithAccounts(addr, opts) failed: %s", testName, err.Error())
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

// 	expected = after.Slug
// 	if expected != before.Slug {
// 		t.Errorf("%s: before.Slug(\"\") failed, expected %v, got %v", testName, expected, before.Slug)
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

// 	for i := 0; i < len(after.Accounts); i++ {
// 		account := after.Accounts[i]

// 		expected = account.Name
// 		if !strings.Contains(expected, "fake_account_name_") {
// 			t.Errorf("%s: strings.Contains(\"\") failed, expected %v, got %v", testName, expected, !strings.Contains(expected, "fake_account_name_"))
// 			return
// 		}

// 		expected = account.Slug
// 		if !strings.Contains(expected, "fake_account_slug_") {
// 			t.Errorf("%s: strings.Contains(\"\") failed, expected %v, got %v", testName, expected, !strings.Contains(expected, "fake_account_slug_"))
// 			return
// 		}
// 	}
// }

// func TestCreateInstitutionWithAccountsAndInstruments(t *testing.T) {
// 	testName := "TestCreateInstitutionWithAccountsAndInstruments"

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

// 	before, after, err := savings.FactoryCreateInstitutionWithAccountsAndInstruments(addr, opts)
// 	if err != nil {
// 		t.Errorf("%s: savings.FactoryCreateInstitutionWithAccountsAndInstruments(addr, opts) failed: %s", testName, err.Error())
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

// 	expected = after.Slug
// 	if expected != before.Slug {
// 		t.Errorf("%s: before.Slug(\"\") failed, expected %v, got %v", testName, expected, before.Slug)
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

// 	for i := 0; i < len(after.Accounts); i++ {
// 		account := after.Accounts[i]

// 		expected = account.Name
// 		if !strings.Contains(expected, "fake_account_name_") {
// 			t.Errorf("%s: strings.Contains(\"\") failed, expected %v, got %v", testName, expected, !strings.Contains(expected, "fake_account_name_"))
// 			return
// 		}

// 		expected = account.Slug
// 		if !strings.Contains(expected, "fake_account_slug_") {
// 			t.Errorf("%s: strings.Contains(\"\") failed, expected %v, got %v", testName, expected, !strings.Contains(expected, "fake_account_slug_"))
// 			return
// 		}

// 		for j := 0; j < len(account.Instruments); j++ {
// 			instrument := account.Instruments[j]

// 			expected = instrument.Name
// 			if !strings.Contains(expected, "fake_instrument_name_") {
// 				t.Errorf("%s: strings.Contains(\"\") failed, expected %v, got %v", testName, expected, !strings.Contains(expected, "fake_instrument_name_"))
// 				return
// 			}

// 			expected = instrument.Slug
// 			if !strings.Contains(expected, "fake_instrument_slug_") {
// 				t.Errorf("%s: strings.Contains(\"\") failed, expected %v, got %v", testName, expected, !strings.Contains(expected, "fake_instrument_slug_"))
// 				return
// 			}

// 			expectedFloat := instrument.Return1m
// 			if expectedFloat != instrument.Return1m {
// 				t.Errorf("%s: instrument.Return1m(\"\") failed, expectedFloat %v, got %v", testName, expectedFloat, instrument.Return1m)
// 				return
// 			}

// 			expectedFloat = instrument.Return1y
// 			if expectedFloat != instrument.Return1y {
// 				t.Errorf("%s: instrument.Return1y(\"\") failed, expectedFloat %v, got %v", testName, expectedFloat, instrument.Return1y)
// 				return
// 			}

// 			expectedFloat = instrument.Return5y
// 			if expectedFloat != instrument.Return5y {
// 				t.Errorf("%s: instrument.Return5y(\"\") failed, expectedFloat %v, got %v", testName, expectedFloat, instrument.Return5y)
// 				return
// 			}

// 			expectedFloat = instrument.Return10y
// 			if expectedFloat != instrument.Return10y {
// 				t.Errorf("%s: instrument.Return10y(\"\") failed, expectedFloat %v, got %v", testName, expectedFloat, instrument.Return10y)
// 				return
// 			}

// 			expectedFloat = instrument.ProjectedWorstCase
// 			if expectedFloat != instrument.ProjectedWorstCase {
// 				t.Errorf("%s: instrument.ProjectedWorstCase(\"\") failed, expectedFloat %v, got %v", testName, expectedFloat, instrument.ProjectedWorstCase)
// 				return
// 			}

// 			expectedFloat = instrument.ProjectedAvgCase
// 			if expectedFloat != instrument.ProjectedAvgCase {
// 				t.Errorf("%s: instrument.ProjectedAvgCase(\"\") failed, expectedFloat %v, got %v", testName, expectedFloat, instrument.ProjectedAvgCase)
// 				return
// 			}

// 			expectedFloat = instrument.ProjectedBestCase
// 			if expectedFloat != instrument.ProjectedBestCase {
// 				t.Errorf("%s: instrument.ProjectedBestCase(\"\") failed, expectedFloat %v, got %v", testName, expectedFloat, instrument.ProjectedBestCase)
// 				return
// 			}
// 		}
// 	}
// }

// func TestListInstrument(t *testing.T) {
// 	testName := "TestListInstrument"

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

// 	before, institutions, err := savings.FactoryListInstitutions(addr, opts)
// 	if err != nil {
// 		t.Errorf("%s: savings.FactoryListInstitutions(addr, opts) failed: %s", testName, err.Error())
// 		return
// 	}

// 	var index int
// 	for i := 0; i < len(institutions); i++ {
// 		if institutions[i].ID == before.ID {
// 			index = i
// 			break
// 		}
// 	}

// 	after := institutions[index]

// 	expected := after.ID
// 	if expected != before.ID {
// 		t.Errorf("%s: after.ID(\"\") failed, expected %v, got %v", testName, expected, before.ID)
// 		return
// 	}

// 	expected = after.Name
// 	if expected != before.Name {
// 		t.Errorf("%s: after.Name(\"\") failed, expected %v, got %v", testName, expected, before.Name)
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
