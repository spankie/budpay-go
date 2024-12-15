package budpay

import "fmt"

type PayoutVerificationResponse struct {
	Success bool                           `json:"success"`
	Message string                         `json:"message"`
	Data    PayoutVerificationResponseData `json:"data"`
}

type PayoutVerificationResponseData struct {
	ID            int64  `json:"id"`
	Reference     string `json:"reference"`
	SessionID     string `json:"sessionid"`
	Currency      string `json:"currency"`
	Amount        string `json:"amount"`
	Fee           string `json:"fee"`
	BankCode      string `json:"bank_code"`
	BankName      string `json:"bank_name"`
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
	Narration     string `json:"narration"`
	Domain        string `json:"domain"`
	Status        string `json:"status"`
	SubAccount    string `json:"subaccount"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

func (budpayClient *BudPayClient) VerifyPayout(reference string) (*PayoutVerificationResponse, error) {
	var response PayoutVerificationResponse
	if err := budpayClient.Get(fmt.Sprintf("v2/payout/:%s", reference), &response); err != nil {
		return nil, fmt.Errorf("error occured while verifying transactions %v", err)
	}
	return &response, nil
}

// TODO: list all transfer payouts

type IncomingPaymentVerificationResponse struct {
	Status          bool                                        `json:"status"`
	Message         string                                      `json:"message"`
	Data            IncomingPaymentVerificationResponseData     `json:"data"`
	Fees            IncomingPaymentVerificationResponseFees     `json:"fees"`
	Customer        IncomingPaymentVerificationResponseCustomer `json:"customer"`
	Plan            IncomingPaymentVerificationResponsePlan     `json:"plan"`
	RequestedAmount string                                      `json:"requested_amount"`
}

type IncomingPaymentVerificationResponseData struct {
	Amount            string                    `json:"amount"`
	Currency          string                    `json:"currency"`
	Status            string                    `json:"status"`
	TransactionDate   string                    `json:"transaction_date"`
	Reference         string                    `json:"reference"`
	Domain            string                    `json:"domain"`
	GatewayResponse   string                    `json:"gateway_response"`
	Channel           string                    `json:"channel"`
	IPAddress         string                    `json:"ip_address"`
	Log               IncomingPaymentWebhookLog `json:"log"`
	TransactionFees   string                    `json:"transaction_fees"`
	TransactionFee    string                    `json:"transaction_fee"`
	TransactionCharge string                    `json:"transaction_charge"`
}

type IncomingPaymentWebhookLog struct {
	TimeSpent      int64  `json:"time_spent"`
	Attempts       int64  `json:"attempts"`
	Authentication string `json:"authentication"`
	Errors         int64  `json:"errors"`
	Success        bool   `json:"success"`
	Channel        string `json:"channel"`
	History        []struct {
		Type    string `json:"type"`
		Message string `json:"message"`
		Time    int64  `json:"time"`
	} `json:"history"`
}

type IncomingPaymentVerificationResponseFees struct {
}

type IncomingPaymentVerificationResponseCustomer struct {
	ID           int64  `json:"id"`
	CustomerCode string `json:"customer_code"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
}

type IncomingPaymentVerificationResponsePlan struct {
}

func (budpayClient *BudPayClient) VerifyIncomingPayment(reference string) (*IncomingPaymentVerificationResponse, error) {
	var response IncomingPaymentVerificationResponse
	if err := budpayClient.Get(fmt.Sprintf("transaction/verify/:%s", reference), &response); err != nil {
		return nil, fmt.Errorf("error occured while verifying transactions %v", err)
	}
	return &response, nil
}
