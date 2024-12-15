package budpay

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBankTransferCheckout(t *testing.T) {
	requestResponseJSON := `{
		"status": true,
		"message": "Account generated successfully",
		"data": {
			"account_name": "Business Name / Firstname lastname",
			"account_number": "1014692362",
			"bank_name": "BudPay Bank"
		}
	}`

	baseURL := "notreal.com"
	hclient := &http.Client{Transport: RoundTripFunc(func(req *http.Request) *http.Response {
		require.Equal(t, baseURL+"s2s/banktransfer/initialize", req.URL.String())
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

	bankTransferRequest := &BankTransferRequest{
		Email:     "email@email.com",
		Amount:    "100",
		Currency:  "NGN",
		Reference: "1253627873656276350",
		Name:      "Business Name / Firstname lastname",
	}
	customerResponse, err := client.BankTransferCheckout(bankTransferRequest)
	require.NoError(t, err)
	require.NotNil(t, customerResponse)
}
