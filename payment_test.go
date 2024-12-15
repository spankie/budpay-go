package budpay

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSinglePayment(t *testing.T) {
	requestResponseJSON := `{"success":true,"message":"Successful Message","data":        {
		"reference": "trf_j51m4695fk57nf",
		"currency": "NGN",
		"amount": "7000",
		"fee": "10",
		"bank_code": "000013",
		"bank_name": "Budpay Bank",
		"account_number": "0000101001",
		"account_name": "OYENIYI TOLULOPE OYEBIYI",
		"narration": "school fees",
		"domain": "test",
		"status": "pending",
		"updated_at": "2022-03-30T00:03:12.000000Z",
		"created_at": "2022-03-30T00:03:12.000000Z"
	}}`

	hclient := &http.Client{Transport: RoundTripFunc(func(req *http.Request) *http.Response {
		require.Equal(t, testBaseURL+"v1/bank_transfer", req.URL.String())
		require.Equal(t, http.MethodPost, req.Method)
		require.Equal(t, "application/json", req.Header.Get("Accept"))
		require.Equal(t, "application/json", req.Header.Get("Content-Type"))
		require.Equal(t, "Bearer testApiKey", req.Header.Get(AuthorizationHeaderKey))

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(requestResponseJSON)),
		}
	})}

	client := BudPayClient{HTTPClient: hclient, apiKey: "testApiKey", encryptionkey: []byte("encryptKey"), BaseURL: testBaseURL}

	singlePaymentRequest := &SinglePaymentRequest{
		Currency:      "NGN",
		Amount:        "7000",
		BankCode:      "000013",
		BankName:      "Budpay Bank",
		AccountNumber: "0000101001",
		Narration:     "school fees",
	}
	paymentResponse, err := client.SinglePayment(singlePaymentRequest)
	require.NoError(t, err)
	require.NotNil(t, paymentResponse)
}

func TestBulkPayment(t *testing.T) {
	requestResponseJSON := `{"status":true,"message":"Successful Message","data":[        {
		"reference": "trf_j51m4695fk57nf",
		"currency": "NGN",
		"amount": "7000",
		"fee": "10",
		"bank_code": "000013",
		"bank_name": "Budpay Bank",
		"account_number": "0000101001",
		"account_name": "OYENIYI TOLULOPE OYEBIYI",
		"narration": "school fees",
		"domain": "test",
		"status": "pending",
		"updated_at": "2022-03-30T00:03:12.000000Z",
		"created_at": "2022-03-30T00:03:12.000000Z"
	}]}`

	hclient := &http.Client{Transport: RoundTripFunc(func(req *http.Request) *http.Response {
		require.Equal(t, testBaseURL+"v2/bulk_bank_transfer", req.URL.String())
		require.Equal(t, http.MethodPost, req.Method)
		require.Equal(t, "application/json", req.Header.Get("Accept"))
		require.Equal(t, "application/json", req.Header.Get("Content-Type"))
		require.Equal(t, "Bearer testApiKey", req.Header.Get(AuthorizationHeaderKey))

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(requestResponseJSON)),
		}
	})}

	client := BudPayClient{HTTPClient: hclient, apiKey: "testApiKey", encryptionkey: []byte("encryptKey"), BaseURL: testBaseURL}

	bulkPaymentRequest := &BulkPaymentRequest{
		Currency: "NGN",
		Transfers: []Payment{
			{
				Amount:        "7000",
				BankCode:      "000013",
				BankName:      "Budpay Bank",
				AccountNumber: "0000101001",
				Narration:     "school fees",
			},
		},
	}
	bulkPaymentResponse, err := client.BulkPayment(bulkPaymentRequest)
	require.NoError(t, err)
	require.NotNil(t, bulkPaymentResponse)
}

func TestFetchTransferFee(t *testing.T) {
	requestResponseJSON := `{"success": true,"message": "Transfer Fee Fetched","fee": "10"}`

	hclient := &http.Client{Transport: RoundTripFunc(func(req *http.Request) *http.Response {
		require.Equal(t, testBaseURL+"v2/payout_fee", req.URL.String())
		require.Equal(t, http.MethodPost, req.Method)
		require.Equal(t, "application/json", req.Header.Get("Accept"))
		require.Equal(t, "application/json", req.Header.Get("Content-Type"))
		require.Equal(t, "Bearer testApiKey", req.Header.Get(AuthorizationHeaderKey))

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(requestResponseJSON)),
		}
	})}

	var client = NewBudPayClient(testBaseURL, "testApiKey", "encryptKey")
	client.HTTPClient = hclient

	transferFeeRequest := &TransferFeeRequest{
		Currency: "NGN",
		Amount:   "1000",
	}
	fetchTransferResponse, err := client.FetchTransferFee(transferFeeRequest)
	require.NoError(t, err)
	require.NotNil(t, fetchTransferResponse)
}

func TestCreatePaymentLink(t *testing.T) {
	responseJson := `{
    "status": true,
    "message": "Payment Link created successfully",
    "data": {
        "ref_id": "29494829",
        "payment_link": "https://www.budpay.com/paymentlink/30838291l2dxf6fs1d7m6137c64n"
    }
}`

	clientTest := &http.Client{Transport: RoundTripFunc(func(req *http.Request) *http.Response {
		require.Equal(t, testBaseURL+"v2/create_payment_link", req.URL.String())
		require.Equal(t, http.MethodPost, req.Method)
		require.Equal(t, "application/json", req.Header.Get("Accept"))
		require.Equal(t, "application/json", req.Header.Get("Content-Type"))
		require.Equal(t, "Bearer testApiKey", req.Header.Get(AuthorizationHeaderKey))

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(responseJson)),
		}
	})}

	client := BudPayClient{HTTPClient: clientTest, apiKey: "testApiKey", encryptionkey: []byte("encryptKey"), BaseURL: testBaseURL + "v2/"}

	paymentLinkRequest := &PaymentLinkRequest{
		Amount:      "2500",
		Currency:    "NGN",
		Name:        "Name",
		Description: "my description",
		Redirect:    "https://your_redirect_link",
	}
	createPaymentLinkResponse, err := client.CreatePaymentLink(paymentLinkRequest)
	require.NoError(t, err)
	require.NotNil(t, createPaymentLinkResponse)
	require.Contains(t, createPaymentLinkResponse.Data.PaymentLink, "https://www.budpay.com/paymentlink/")
}
