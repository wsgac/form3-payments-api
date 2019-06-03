package payments

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// func TestMain(m *testing.M) {
// 	// Setup
// 	config.LoadConfig()
// 	config.Database = "api_test"
// 	config.APIPort = 9000
// 	Start()
// 	// Run tests
// 	os.Exit(m.Run())
// 	// Teardown
// 	db.DropDatabase()
// }

// func TestAddPayment(t *testing.T) {
// 	fmt.Println("Running test: TestAddPayment")
// 	// client := &http.Client{}
// 	addRequestBody := `{
// 		"type": "Payment",
// 		"version": 0,
// 		"organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
// 		"attributes": {
// 		"amount": "10000.21",
// 		"beneficiary_party": {
// 			"account_name": "W Owens",
// 			"account_number": "31926819",
// 			"account_number_code": "BBAN",
// 			"account_type": 0,
// 			"address": "1 The Beneficiary Localtown SE2",
// 			"bank_id": "403000",
// 			"bank_id_code": "GBDSC",
// 			"name": "Wilfred Jeremiah Owens"
// 		},
// 		"charges_information": {
// 			"bearer_code": "SHAR",
// 			"sender_charges": [
// 			{
// 				"amount": "5.00",
// 				"currency": "GBP"
// 			},
// 			{
// 				"amount": "10.00",
// 				"currency": "USD"
// 			}
// 			],
// 			"receiver_charges_amount": "1.00",
// 			"receiver_charges_currency": "USD"
// 		},
// 		"currency": "GBP",
// 		"debtor_party": {
// 			"account_name": "EJ Brown Black",
// 			"account_number": "GB29XABC10161234567801",
// 			"account_number_code": "IBAN",
// 			"address": "10 Debtor Crescent Sourcetown NE1",
// 			"bank_id": "203301",
// 			"bank_id_code": "GBDSC",
// 			"name": "Emelia Jane Brown"
// 		},
// 		"end_to_end_reference": "Wil piano Jan",
// 		"fx": {
// 			"contract_reference": "FX123",
// 			"exchange_rate": "2.00000",
// 			"original_amount": "200.42",
// 			"original_currency": "USD"
// 		},
// 		"numeric_reference": "1002001",
// 		"payment_id": "123456789012345678",
// 		"payment_purpose": "Paying for goods/services",
// 		"payment_scheme": "FPS",
// 		"payment_type": "Credit",
// 		"processing_date": "2017-01-18",
// 		"reference": "Payment for Em's piano lessons",
// 		"scheme_payment_sub_type": "InternetBanking",
// 		"scheme_payment_type": "ImmediatePayment",
// 		"sponsor_party": {
// 			"account_number": "56781234",
// 			"bank_id": "123123",
// 			"bank_id_code": "GBDSC"
// 		}
// 		}
// 	}`
// 	req, _ := http.NewRequest("POST",
// 		"localhost"+string(config.APIPort)+"/payments",
// 		bytes.NewBuffer([]byte(addRequestBody)),
// 	)
// 	rr := httptest.NewRecorder()
// 	router.HandleFunc("/payments", AddPayment)
// 	router.ServeHTTP(rr, req)

// 	// response, err := client.Do(req)
// 	// if err != nil {
// 	// 	t.Fatalf("Request failed with error: %v", err)
// 	// }
// 	// if response.StatusCode != http.StatusCreated {
// 	// 	t.Fatalf("Did not receive Created status. Got %d instead", response.StatusCode)
// 	// }
// 	if rr.Code != http.StatusCreated {
// 		t.Errorf("Incorrect response status. Expecting: %d Got: %d", http.StatusCreated, rr.Code)
// 	}
// }

var addRequestBody = `{
	"type": "Payment",
	"version": 0,
	"organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
	"attributes": {
	"amount": "10000.21",
	"beneficiary_party": {
		"account_name": "W Owens",
		"account_number": "31926819",
		"account_number_code": "BBAN",
		"account_type": 0,
		"address": "1 The Beneficiary Localtown SE2",
		"bank_id": "403000",
		"bank_id_code": "GBDSC",
		"name": "Wilfred Jeremiah Owens"
	},
	"charges_information": {
		"bearer_code": "SHAR",
		"sender_charges": [
		{
			"amount": "5.00",
			"currency": "GBP"
		},
		{
			"amount": "10.00",
			"currency": "USD"
		}
		],
		"receiver_charges_amount": "1.00",
		"receiver_charges_currency": "USD"
	},
	"currency": "GBP",
	"debtor_party": {
		"account_name": "EJ Brown Black",
		"account_number": "GB29XABC10161234567801",
		"account_number_code": "IBAN",
		"address": "10 Debtor Crescent Sourcetown NE1",
		"bank_id": "203301",
		"bank_id_code": "GBDSC",
		"name": "Emelia Jane Brown"
	},
	"end_to_end_reference": "Wil piano Jan",
	"fx": {
		"contract_reference": "FX123",
		"exchange_rate": "2.00000",
		"original_amount": "200.42",
		"original_currency": "USD"
	},
	"numeric_reference": "1002001",
	"payment_id": "123456789012345678",
	"payment_purpose": "Paying for goods/services",
	"payment_scheme": "FPS",
	"payment_type": "Credit",
	"processing_date": "2017-01-18",
	"reference": "Payment for Em's piano lessons",
	"scheme_payment_sub_type": "InternetBanking",
	"scheme_payment_type": "ImmediatePayment",
	"sponsor_party": {
		"account_number": "56781234",
		"bank_id": "123123",
		"bank_id_code": "GBDSC"
	}
	}
}`

var getResponseBody = `{
	"type": "Payment",
	"id": "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
	"version": 0,
	"organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
	"attributes": {
	"amount": "10000.21",
	"beneficiary_party": {
		"account_name": "W Owens",
		"account_number": "31926819",
		"account_number_code": "BBAN",
		"account_type": 0,
		"address": "1 The Beneficiary Localtown SE2",
		"bank_id": "403000",
		"bank_id_code": "GBDSC",
		"name": "Wilfred Jeremiah Owens"
	},
	"charges_information": {
		"bearer_code": "SHAR",
		"sender_charges": [
		{
			"amount": "5.00",
			"currency": "GBP"
		},
		{
			"amount": "10.00",
			"currency": "USD"
		}
		],
		"receiver_charges_amount": "1.00",
		"receiver_charges_currency": "USD"
	},
	"currency": "GBP",
	"debtor_party": {
		"account_name": "EJ Brown Black",
		"account_number": "GB29XABC10161234567801",
		"account_number_code": "IBAN",
		"address": "10 Debtor Crescent Sourcetown NE1",
		"bank_id": "203301",
		"bank_id_code": "GBDSC",
		"name": "Emelia Jane Brown"
	},
	"end_to_end_reference": "Wil piano Jan",
	"fx": {
		"contract_reference": "FX123",
		"exchange_rate": "2.00000",
		"original_amount": "200.42",
		"original_currency": "USD"
	},
	"numeric_reference": "1002001",
	"payment_id": "123456789012345678",
	"payment_purpose": "Paying for goods/services",
	"payment_scheme": "FPS",
	"payment_type": "Credit",
	"processing_date": "2017-01-18",
	"reference": "Payment for Em's piano lessons",
	"scheme_payment_sub_type": "InternetBanking",
	"scheme_payment_type": "ImmediatePayment",
	"sponsor_party": {
		"account_number": "56781234",
		"bank_id": "123123",
		"bank_id_code": "GBDSC"
	}
	}
}`

func MockAddPayment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	JSONResponse(w, http.StatusOK, map[string]string{
		"result": "success",
	})
}

func TestAddPayment(t *testing.T) {
	req := httptest.NewRequest("POST", "/payments",
		bytes.NewBuffer([]byte(addRequestBody)))
	rec := httptest.NewRecorder()
	MockAddPayment(rec, req)
	resp := rec.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bodyDict map[string]string
	json.Unmarshal(body, &bodyDict)
	t.Logf("Body: %#v", bodyDict)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Incorrect request status. Expecting: %d Got: %d", http.StatusOK, resp.StatusCode)
	}
	if bodyDict["result"] != "success" {
		t.Errorf("Incorrect response body. Expecting: \"success\" Got: %s", bodyDict["result"])
	}
}

func MockListPayments(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var payment interface{}
	json.Unmarshal([]byte(getResponseBody), &payment)
	JSONResponse(w, http.StatusOK, map[string][]interface{}{
		"data": []interface{}{payment},
	})
}
func TestListPayments(t *testing.T) {
	req := httptest.NewRequest("GET", "/payments", nil)
	rec := httptest.NewRecorder()
	MockListPayments(rec, req)
	resp := rec.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bodyStruct map[string][]interface{}
	json.Unmarshal(body, &bodyStruct)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Incorrect request status. Expecting: %d Got: %d", http.StatusOK, resp.StatusCode)
	}
	if len(bodyStruct["data"]) != 1 {
		t.Errorf("Incorrect payment cound. Expecting: %d Got: %d", 1, len(bodyStruct["data"]))
	}
}

func MockGetPayment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var payment interface{}
	json.Unmarshal([]byte(getResponseBody), &payment)
	JSONResponse(w, http.StatusOK, payment)
}

func TestGetPayment(t *testing.T) {
	id := "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43"
	req := httptest.NewRequest("GET", "/payments/"+id, nil)
	rec := httptest.NewRecorder()
	MockGetPayment(rec, req)
	resp := rec.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bodyStruct map[string]interface{}
	json.Unmarshal(body, &bodyStruct)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Incorrect request status. Expecting: %d Got: %d", http.StatusOK, resp.StatusCode)
	}
	if bodyStruct["id"] != id {
		t.Errorf("Incorrect payment ID. Expecting: %s Got: %s", id, bodyStruct["id"])
	}
}

func MockUpdatePayment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	JSONResponse(w, http.StatusOK, map[string]string{
		"result": "success",
	})
}

func TestUpdatePayment(t *testing.T) {
	req := httptest.NewRequest("PUT", "/payments",
		bytes.NewBuffer([]byte(addRequestBody)))
	rec := httptest.NewRecorder()
	MockUpdatePayment(rec, req)
	resp := rec.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bodyDict map[string]string
	json.Unmarshal(body, &bodyDict)
	t.Logf("Body: %#v", bodyDict)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Incorrect request status. Expecting: %d Got: %d", http.StatusOK, resp.StatusCode)
	}
	if bodyDict["result"] != "success" {
		t.Errorf("Incorrect response body. Expecting: \"success\" Got: %s", bodyDict["result"])
	}
}

func MockDeletePayment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	JSONResponse(w, http.StatusOK, map[string]string{
		"result": "success",
	})
}
func TestDeletePayment(t *testing.T) {
	id := "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43"
	req := httptest.NewRequest("DELETE", "/payments/"+id, nil)
	rec := httptest.NewRecorder()
	MockDeletePayment(rec, req)
	resp := rec.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Incorrect request status. Expecting: %d Got: %d", http.StatusOK, resp.StatusCode)
	}
}
