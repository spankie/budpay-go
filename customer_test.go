package budpay

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateCustomer(t *testing.T) {
	requestResponseJSON := `{"status":true,"message":"Successful Message","data":{
		"email": "email@email.com",
		"domain": "test",
		"customer": "CUS_cutomercode",
		"id": 6,
		"updated_at": "2022-03-30T00:03:12.000000Z",
		"created_at": "2022-03-30T00:03:12.000000Z"
	}}`

	baseURL := "notreal.com"
	hclient := &http.Client{Transport: RoundTripFunc(func(req *http.Request) *http.Response {
		require.Equal(t, baseURL+"v2/customer", req.URL.String())
		require.Equal(t, http.MethodPost, req.Method)
		require.Equal(t, "application/json", req.Header.Get("Accept"))
		require.Equal(t, "application/json", req.Header.Get("Content-Type"))
		require.Equal(t, "Bearer testApiKey", req.Header.Get(AuthorizationHeaderKey))

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(requestResponseJSON)),
		}
	})}

	client := BudPayClient{HTTPClient: hclient, apiKey: "testApiKey", encryptionkey: []byte("encryptKey"), BaseURL: baseURL}
	client.HTTPClient = hclient

	customerRequest := &CustomerRequest{
		Email:     "email@email.com",
		FirstName: "firstname",
		LastName:  "lastname",
		Phone:     "+2340000000000",
	}
	customerResponse, err := client.CreateCustomer(customerRequest)
	require.NoError(t, err)
	require.NotNil(t, customerResponse)
}

func TestCreatePaymentRequest(t *testing.T) {
	responseJson := `{
    "status": true,
    "message": "Payment request processed successfully"
}`

	baseURL := "notreal.com"
	testClient := &http.Client{Transport: RoundTripFunc(func(req *http.Request) *http.Response {
		require.Equal(t, baseURL+"v2/request_payment", req.URL.String())
		require.Equal(t, http.MethodPost, req.Method)
		require.Equal(t, "application/json", req.Header.Get("Accept"))
		require.Equal(t, "application/json", req.Header.Get("Content-Type"))
		require.Equal(t, "Bearer testApiKey", req.Header.Get(AuthorizationHeaderKey))

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(responseJson)),
		}
	})}

	client := BudPayClient{HTTPClient: testClient, apiKey: "testApiKey", encryptionkey: []byte("encryptKey"), BaseURL: baseURL + "v2/"}
	client.HTTPClient = testClient

	createPaymentRequest := &PaymentRequest{
		Recipient:   "toluxsys@yahoo.ca,07036218209,sam@bud.africa,08161112404",
		Amount:      "1000",
		Currency:    "NGN",
		Description: "Payment for goods",
	}
	customerResponse, err := client.CreatePaymentRequest(createPaymentRequest)
	require.NoError(t, err)
	require.NotNil(t, customerResponse)
}
