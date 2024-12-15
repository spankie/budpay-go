package budpay

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"time"
)

type BudPayClient struct {
	BaseURL       string
	apiKey        string
	encryptionkey []byte
	HTTPClient    *http.Client
	logger        *slog.Logger
}

const (
	CURRENCY_CODE_NGN                string = "NGN"
	ACCOUNT_VERIFICATION_UNAVAILABLE string = "account_verification unavailable"
)

func NewBudPayClient(budPayBaseURL, budPayApiKey, budPayEncryptionKey string) *BudPayClient {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	client := &BudPayClient{
		BaseURL:       budPayBaseURL,
		apiKey:        budPayApiKey,
		encryptionkey: []byte(budPayEncryptionKey),
		HTTPClient: &http.Client{
			Timeout: 1 * time.Minute,
		},
		logger: logger,
	}

	return client
}

func (bc *BudPayClient) SetupProxy(budpayProxy string) error {
	if budpayProxy != "" {
		bc.logger.Debug("using proxy for budpay", "proxy", budpayProxy)
		proxyUrl, err := url.Parse(fmt.Sprintf("http://%s", budpayProxy))
		if err != nil {
			bc.logger.Error("error creating proxy url for budpay", "error", err)
			return err
		}
		bc.HTTPClient.Transport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	} else {
		return fmt.Errorf("proxy cannot be empty")
	}

	return nil
}
