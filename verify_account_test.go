package budpay

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifyAccountDetails(t *testing.T) {
	responseJson := `{
    "success": true,
    "message": "Account name retrieved",
    "data": "OYENIYI TOLULOPE OYEBIYI"
}`
	/*
	   {
	   			"account_name": "OYENIYI TOLULOPE OYEBIYI",
	   			"account_number": "0123456789",
	   			"bank_code": "000001",
	   			"bank_name": "Stanbic IBTC"
	   		}
	*/

	clientTest := &http.Client{Transport: RoundTripFunc(func(req *http.Request) *http.Response {
		require.Equal(t, testBaseURL+"v1/account_name_verify", req.URL.String())
		require.Equal(t, http.MethodPost, req.Method)
		require.Equal(t, "application/json", req.Header.Get("Accept"))
		require.Equal(t, "application/json", req.Header.Get("Content-Type"))
		require.Equal(t, "Bearer testApiKey", req.Header.Get(AuthorizationHeaderKey))

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(responseJson)),
		}
	})}

	client := BudPayClient{HTTPClient: clientTest, apiKey: "testApiKey", encryptionkey: []byte("encryptKey"), BaseURL: testBaseURL}

	accountDetails := &AccountDetails{
		BankCode:      "000001",
		AccountNumber: "0123456789",
		Currency:      "NGN",
	}
	response, err := client.VerifyAccountDetails(accountDetails)
	require.NoError(t, err)
	require.NotNil(t, response)
}
