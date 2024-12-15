package budpay

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	AuthorizationHeaderKey = "Authorization"
)

func (bp *BudPayClient) Get(endPointURL string, expectedResponse interface{}) error {
	return bp.sendRequest(http.MethodGet, endPointURL, nil, expectedResponse)
}

func (bp *BudPayClient) Post(endPointURL string, reqBody, expectedResponse interface{}) error {
	return bp.sendRequest(http.MethodPost, endPointURL, reqBody, expectedResponse)
}

func (bp *BudPayClient) Delete(endPointURL string, expectedResponse interface{}) error {
	return bp.sendRequest(http.MethodDelete, endPointURL, nil, expectedResponse)
}

func (bp *BudPayClient) generateRequest(method, endpoint string, body interface{}) (*http.Request, error) {
	var bodyReader io.Reader
	var hashedBody string
	if body != nil {
		bodyByte, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		// TODO: monnitor the performance of this
		mac := hmac.New(sha512.New, bp.encryptionkey)
		mac.Write(bodyByte)
		hashedBody = hex.EncodeToString(mac.Sum(nil))

		bodyReader = bytes.NewReader(bodyByte)
	}
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", bp.BaseURL, endpoint), bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set(AuthorizationHeaderKey, fmt.Sprintf("Bearer %s", bp.apiKey))
	req.Header.Set("Encryption", string(hashedBody))

	return req, nil
}

func (bp *BudPayClient) sendRequest(method, endpoint string, body, output interface{}) error {
	req, err := bp.generateRequest(method, endpoint, body)
	if err != nil {
		return err
	}

	res, err := bp.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request %+v", err)
	}
	defer res.Body.Close()
	if _, ok := output.(*string); ok {
		buf := new(bytes.Buffer)
		_, err2 := buf.ReadFrom(res.Body)
		if err2 != nil {
			return err2
		}
		// budpay can sometimes send a string
		*output.(*string) = buf.String()
	} else {
		if err := json.NewDecoder(res.Body).Decode(output); err != nil {
			return fmt.Errorf("error marshalling client response: %s", err)
		}
	}

	if res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("client response with status code: %v message: %v", res.StatusCode, output)
	}

	return nil
}
