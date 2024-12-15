package budpay

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetBankList(t *testing.T) {
	requestResponseJSON := `{
		"success": true,
		"message": "Bank list retrieved",
		"currency": "NGN", 
		"data":[
			{
				"bank_name": "9PAYMENT SERVICE BANK",
				"bank_code": "120001"
			},
			{
				"bank_name": "AB MICROFINANCE BANK",
				"bank_code": "090270"
			},
			{
				"bank_name": "ABBEY MORTGAGE BANK ",
				"bank_code": "070010"
			},
			{
				"bank_name": "ABUCOOP MICROFINANCE BANK",
				"bank_code": "090424"
			},
			{
				"bank_name": "ACCESS BANK",
				"bank_code": "000014"
			},
			{
				"bank_name": "ZENITH BANK",
				"bank_code": "000015"
			}
		]}`

	baseURL := "notreal.com"
	hclient := &http.Client{Transport: RoundTripFunc(func(req *http.Request) *http.Response {
		require.Equal(t, baseURL+"v2/bank_list/NGN", req.URL.String())
		require.Equal(t, http.MethodGet, req.Method)
		require.Equal(t, "application/json", req.Header.Get("Accept"))
		require.Equal(t, "application/json", req.Header.Get("Content-Type"))
		require.Equal(t, "Bearer testApiKey", req.Header.Get(AuthorizationHeaderKey))

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(requestResponseJSON)),
		}
	})}

	var client = NewBudPayClient(baseURL, "testApiKey", "encryptKey")
	client.HTTPClient = hclient

	bankListResponse, err := client.GetBankList("NGN")
	require.NoError(t, err)
	require.NotNil(t, bankListResponse)
}
