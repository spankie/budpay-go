package budpay

import "fmt"

type Payment struct {
	AccountNumber string `json:"account_number" binding:"required"`
	Amount        string `json:"amount" binding:"required"`
	BankCode      string `json:"bank_code" binding:"required"`
	BankName      string `json:"bank_name" binding:"required"`
	Narration     string `json:"narration" binding:"required"`
	Reference     string `json:"reference,omitempty"`
}

type PaymentResponse struct {
	Reference     string `json:"reference"`
	Currency      string `json:"currency"`
	Amount        string `json:"amount"`
	Fee           string `json:"fee"`
	BankCode      string `json:"bank_code"`
	BankName      string `json:"bank_name"`
	AccountNumber string `json:"account_number"`
	Narration     string `json:"narration"`
	Domain        string `json:"domain"`
	Status        string `json:"status"`
	UpdatedAt     string `json:"updated_at"`
	CreatedAt     string `json:"created_at"`
	PaymentMode   string `json:"paymentMode"`
	ServiceCode   string `json:"serviceCode"`
	// AccountName   string `json:"account_name"`
}

type SinglePaymentRequest struct {
	AccountNumber string `json:"account_number" binding:"required"`
	Amount        string `json:"amount" binding:"required"`
	BankCode      string `json:"bank_code" binding:"required"`
	BankName      string `json:"bank_name" binding:"required"`
	Narration     string `json:"narration"`
	Reference     string `json:"reference,omitempty"`
	Currency      string `json:"currency" binding:"required"`
	PaymentMode   string `json:"paymentMode,omitempty"`
	ServiceCode   string `json:"serviceCode,omitempty"`
	AccountName   string `json:"account_name,omitempty"`
}

type SinglePaymentResponse struct {
	ResponseSuccess
	Data PaymentResponse `json:"data"`
}

type BulkPaymentRequest struct {
	Currency  string    `json:"currency"`
	Transfers []Payment `json:"transfers"`
}

type BulkPaymentResponse struct {
	ResponseSuccess
	Data []PaymentResponse `json:"data"`
}

type TransferFeeRequest struct {
	Currency string `json:"currency" binding:"required"`
	Amount   string `json:"amount" binding:"required"`
}

type TransferFeeResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Fee     string `json:"fee"`
}

type PaymentLinkRequest struct {
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Redirect    string `json:"redirect"`
}

type PaymentLinkResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    PaymentLinkData `json:"data"`
}

type PaymentLinkData struct {
	RefId       string `json:"ref_id"`
	PaymentLink string `json:"payment_link"`
}
type PaymentRequest struct {
	Recipient   string `json:"recipient" binding:"required"`
	Amount      string `json:"amount" binding:"required"`
	Currency    string `json:"currency" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CreatePaymentResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func (budpayClient *BudPayClient) SinglePayment(reqBody *SinglePaymentRequest) (*SinglePaymentResponse, error) {
	var response SinglePaymentResponse
	if err := budpayClient.Post("v1/bank_transfer", reqBody, &response); err != nil {
		return nil, fmt.Errorf("error while making single payment: %v", err)
	}
	return &response, nil
}

func (budpayClient *BudPayClient) BulkPayment(reqBody *BulkPaymentRequest) (*BulkPaymentResponse, error) {
	var response BulkPaymentResponse
	if err := budpayClient.Post("v2/bulk_bank_transfer", reqBody, &response); err != nil {
		return nil, fmt.Errorf("error while making bulk payment: %v", err)
	}
	return &response, nil
}

func (budpayClient *BudPayClient) FetchTransferFee(reqBody *TransferFeeRequest) (*TransferFeeResponse, error) {
	var response TransferFeeResponse
	if err := budpayClient.Post("v2/payout_fee", reqBody, &response); err != nil {
		return nil, fmt.Errorf("error while fetching transfer fee: %v", err)
	}
	return &response, nil
}

func (budpayClient *BudPayClient) CreatePaymentRequest(reqBody *PaymentRequest) (*CreatePaymentResponse, error) {
	var response CreatePaymentResponse
	if err := budpayClient.Post("request_payment", reqBody, &response); err != nil {
		return nil, fmt.Errorf("error while creating request payment: %v", err)
	}
	return &response, nil
}

func (budpayClient *BudPayClient) CreatePaymentLink(reqBody *PaymentLinkRequest) (*PaymentLinkResponse, error) {
	var response PaymentLinkResponse
	if err := budpayClient.Post("create_payment_link", reqBody, &response); err != nil {
		return nil, fmt.Errorf("error occured while creating payment link %v", err)
	}
	return &response, nil
}
