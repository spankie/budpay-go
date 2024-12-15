package budpay

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testBaseURL = "notreal.com"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func TestGet(t *testing.T) {
	requestResponseJSON := `{"status":true,"message":"Successful Message","data":["randomData"]}`

	hclient := &http.Client{Transport: RoundTripFunc(func(req *http.Request) *http.Response {
		require.Equal(t, testBaseURL+"v2/anyRoute", req.URL.String())
		require.Equal(t, http.MethodGet, req.Method)
		require.Equal(t, "application/json", req.Header.Get("Accept"))
		require.Equal(t, "application/json", req.Header.Get("Content-Type"))
		require.Equal(t, "Bearer testApiKey", req.Header.Get(AuthorizationHeaderKey))

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(requestResponseJSON)),
		}
	})}

	client := BudPayClient{HTTPClient: hclient, apiKey: "testApiKey", encryptionkey: []byte("encryptKey"), BaseURL: testBaseURL + "v2/"}

	expectedResponse := &struct {
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{}
	err := client.Get("anyRoute", expectedResponse)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestPost(t *testing.T) {
	requestResponseJSON := `{"status":true,"message":"Successful Message","data":["randomData"]}`

	hclient := &http.Client{Transport: RoundTripFunc(func(req *http.Request) *http.Response {
		require.Equal(t, testBaseURL+"v2/anyRoute", req.URL.String())
		require.Equal(t, http.MethodPost, req.Method)
		require.Equal(t, "application/json", req.Header.Get("Accept"))
		require.Equal(t, "application/json", req.Header.Get("Content-Type"))
		require.Equal(t, "Bearer testApiKey", req.Header.Get(AuthorizationHeaderKey))

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(requestResponseJSON)),
		}
	})}

	client := BudPayClient{HTTPClient: hclient, apiKey: "testApiKey", encryptionkey: []byte("encryptKey"), BaseURL: testBaseURL + "v2/"}

	expectedResponse := &struct {
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{}
	err := client.Post("anyRoute", "randomBody", expectedResponse)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestDelete(t *testing.T) {
	requestResponseJSON := `{"status":true,"message":"Successful Message","data":["randomData"]}`

	hclient := &http.Client{Transport: RoundTripFunc(func(req *http.Request) *http.Response {
		require.Equal(t, testBaseURL+"v2/anyRoute", req.URL.String())
		require.Equal(t, http.MethodDelete, req.Method)
		require.Equal(t, "application/json", req.Header.Get("Accept"))
		require.Equal(t, "application/json", req.Header.Get("Content-Type"))
		require.Equal(t, "Bearer testApiKey", req.Header.Get(AuthorizationHeaderKey))

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(requestResponseJSON)),
		}
	})}

	client := BudPayClient{HTTPClient: hclient, apiKey: "testApiKey", encryptionkey: []byte("encryptKey"), BaseURL: testBaseURL + "v2/"}

	expectedResponse := &struct {
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{}
	err := client.Delete("anyRoute", expectedResponse)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGenerateRequest(t *testing.T) {
	client := BudPayClient{apiKey: "testApiKey", encryptionkey: []byte("encryptKey"), BaseURL: testBaseURL + "v2/"}

	testCases := []struct {
		method   string
		endpoint string
		body     interface{}
	}{
		{
			method:   http.MethodPost,
			endpoint: "randomEndpoint",
			body:     &struct{}{},
		},
		{
			method:   http.MethodGet,
			endpoint: "randomEndpoint",
			body:     nil,
		},
	}
	for _, tc := range testCases {
		req, err := client.generateRequest(tc.method, tc.endpoint, tc.body)

		require.NoError(t, err)
		require.NotNil(t, req)
		require.Equal(t, testBaseURL+"v2/randomEndpoint", req.URL.String())
		require.Equal(t, tc.method, req.Method)
		require.Equal(t, "application/json", req.Header.Get("Accept"))
		require.Equal(t, "application/json", req.Header.Get("Content-Type"))
		require.Equal(t, "Bearer testApiKey", req.Header.Get(AuthorizationHeaderKey))
	}
}
