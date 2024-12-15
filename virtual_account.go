package budpay

import (
	"fmt"
)

type VirtualAccountRequest struct {
	CustomerCode string `json:"customer"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Phone        string `json:"phone"`
}

type VirtualAccountResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    BankAccount `json:"data"`
}

type BankAccount struct {
	Bank          Bank     `json:"bank"`
	AccountName   string   `json:"account_name"`
	AccountNumber string   `json:"account_number"`
	Currency      string   `json:"currency"`
	Status        string   `json:"status"`
	Reference     string   `json:"reference"`
	ID            int      `json:"id"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     string   `json:"updated_at"`
	Domain        string   `json:"domain"`
	Customer      Customer `json:"customer"`
	Assignment    string   `json:"assignment"`
}

type VirtualAccountList struct {
	ResponseStatus
	Meta MetaData      `json:"meta"`
	Data []BankAccount `json:"data"`
}

type MetaData struct {
	Total int `json:"total"`
}

type Bank struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	BankCode string `json:"bank_code"`
	Prefix   string `json:"prefix"`
}

func (budpayClient *BudPayClient) CreateVirtualAccount(reqBody *VirtualAccountRequest) (*VirtualAccountResponse, error) {
	if reqBody.CustomerCode == "" {
		return nil, fmt.Errorf("customer code is required")
	}

	// TODO: might need to uncomment this later
	// if v, err := budpayClient.findExistingVirtualAccount(reqBody.CustomerCode); err == nil {
	// 	return v, nil
	// }

	var response VirtualAccountResponse
	if err := budpayClient.Post("v2/dedicated_virtual_account", reqBody, &response); err != nil {
		return nil, fmt.Errorf("error occured while creating virtual account: %v", err)
	}
	return &response, nil
}

func (budpayClient *BudPayClient) findExistingVirtualAccount(customerCode string) (*VirtualAccountResponse, error) {
	bankAccounts, err := budpayClient.ListVirtualAccounts()
	if err != nil {
		return nil, err
	}
	for _, account := range *bankAccounts {
		if account.Customer.CustomerCode == customerCode {
			return &VirtualAccountResponse{
				Status:  true,
				Message: "success",
				Data:    account,
			}, nil
		}
	}
	return nil, fmt.Errorf("virtual account not found")
}

func (budpayClient *BudPayClient) ListVirtualAccounts() (*[]BankAccount, error) {
	var response VirtualAccountList
	if err := budpayClient.Get("v2/list_dedicated_accounts", &response); err != nil {
		return nil, fmt.Errorf("error occured while listing virtual accounts: %v", err)
	}
	return &response.Data, nil
}
