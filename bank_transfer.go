package budpay

import "fmt"

type BankTransferRequest struct {
	Email     string `json:"email" binding:"required"`
	Amount    string `json:"amount" binding:"required"`
	Currency  string `json:"currency" binding:"required"`
	Reference string `json:"reference" binding:"required"`
	Name      string `json:"name" binding:"required"`
}

type BankTransferResponse struct {
	ResponseStatus
	Data    AccountDetail `json:"data"`
}

type AccountDetail struct {
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
}

func (budpayClient *BudPayClient) BankTransferCheckout(reqBody *BankTransferRequest) (*BankTransferResponse, error) {
	var response BankTransferResponse
	if err := budpayClient.Post("s2s/banktransfer/initialize", reqBody, &response); err != nil {
		return nil, fmt.Errorf("error while making transfer checkout: %v", err)
	}
	return &response, nil
}
