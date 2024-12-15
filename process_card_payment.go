package budpay

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func getTransactionReference() string {
	ref := make([]byte, aes.BlockSize)
	_, err := rand.Read(ref)
	if err != nil {
		fmt.Printf("error generating transaction reference: %v", err)
		return ""
	}
	return hex.EncodeToString(ref)
}

type cardData struct {
	Number      string `json:"number"`
	ExpiryMonth string `json:"expiryMonth"`
	ExpiryYear  string `json:"expiryYear"`
	CVV         string `json:"cvv"`
	Pin         string `json:"pin,omitempty"`
}

type CardEncryptReq struct {
	Data      cardData `json:"data"`
	Reference string   `json:"reference"`
}

type PaymentInfo struct {
	Amount    string `json:"amount"`
	Card      string `json:"card"`
	Currency  string `json:"currency"`
	Email     string `json:"email"`
	Pin       string `json:"pin"`
	Reference string `json:"reference"`
}

type CardPaymentRequest struct {
	Amount      string `json:"amount"`
	CardNumber  string `json:"card_number"`
	ExpiryMonth string `json:"expiry_month"`
	ExpiryYear  string `json:"expiry_year"`
	CVV         string `json:"cvv"`
	Pin         string `json:"pin,omitempty"`
	Currency    string `json:"currency"`
	Email       string `json:"email"`
}

type CardPaymentResponse struct {
	Message   string `json:"message"`
	Status    bool   `json:"status"`
	Reference string `json:"reference"`
	Link      string `json:"link"`
	PaymentID string `json:"paymentid"`
}

func (BudPayClient *BudPayClient) ProcessCardPayment(req *CardPaymentRequest, ref string) (*CardPaymentResponse, error) {
	ref = getTransactionReference()
	cardEncryptReq := CardEncryptReq{
		Data: cardData{
			Number:      req.CardNumber,
			ExpiryMonth: req.ExpiryMonth,
			ExpiryYear:  req.ExpiryYear,
			CVV:         req.CVV,
			Pin:         req.Pin,
		},
		Reference: ref,
	}
	encryptedData, err := BudPayClient.encryptCardDetails(cardEncryptReq)
	if err != nil {
		return nil, err
	}
	paymentInfo := PaymentInfo{
		Amount:    req.Amount,
		Card:      encryptedData,
		Currency:  req.Currency,
		Email:     req.Email,
		Pin:       req.Pin,
		Reference: ref,
	}

	res, err := BudPayClient.initiateCardPayment(&paymentInfo)
	if err != nil {
		return nil, err
	}

	message, ok := res["message"].(string)
	if !ok {
		return nil, fmt.Errorf("error while processing card payment: %v", res)
	}

	cardPaymentResponse := &CardPaymentResponse{
		Message:   message,
		Status:    res["status"].(bool), // FIXME: handle type assertion error
		Reference: ref,
	}

	// TODO: add an example
	if strings.Contains(message, "OTP") {
		data, ok := res["data"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("error while processing card payment: %v", res)
		}
		// reference, paymentid, link
		paymentID, ok := data["paymentid"].(string)
		if !ok {
			return nil, fmt.Errorf("error while processing card payment: %v", res)
		}
		link, ok := data["_links"].(map[string]interface{})["url"].(string)
		if !ok {
			return nil, fmt.Errorf("error while processing card payment: %v", res)
		}
		cardPaymentResponse.PaymentID = paymentID
		cardPaymentResponse.Link = link
	}

	return cardPaymentResponse, nil
}

// TODO: response should not be a map
func (budpayClient *BudPayClient) initiateCardPayment(req *PaymentInfo) (map[string]interface{}, error) {
	var response map[string]interface{}

	if err := budpayClient.Post("s2s/transaction/initialize", req, &response); err != nil {
		return nil, fmt.Errorf("error while initiating card payment: %v", err)
	}
	return response, nil
}

type OtpPaymentRequest struct {
	Otp   string `json:"otp"`
	PayID string `json:"payid"`
	Ref   string `json:"ref"`
}

func (budpayClient *BudPayClient) SendOTP(link string, otpPaymentRequest OtpPaymentRequest) (*CardPaymentResponse, error) {
	reqBytes, err := json.Marshal(otpPaymentRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://payments.budpay.com/api/i5678930tyuhjns-cardotp", bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set(AuthorizationHeaderKey, fmt.Sprintf("Bearer %s", budpayClient.apiKey))
	res, err := budpayClient.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error while sending OTP: %v", res.Status)
	}

	var response CardPaymentResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
