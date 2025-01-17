package budpay

import (
	"fmt"
	"strings"
)

type AccountDetails struct {
	BankCode      string `json:"bank_code" binding:"required"`
	AccountNumber string `json:"account_number" binding:"required"`
	Currency      string `json:"currency" binding:"required"`
}

type AccountDetailsResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    string `json:"data"`
	// Data    AccountDetailsResponseData `json:"data"`
}

type AccountDetailsResponseData struct {
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
	BankName      string `json:"bank_name"`
}

func (budpayClient *BudPayClient) VerifyAccountDetails(receiver *AccountDetails) (*AccountDetailsResponse, error) {
	switch receiver.Currency {
	case "NGN":
		var mapResponse map[string]interface{}
		if err := budpayClient.Post("v1/account_name_verify", receiver, &mapResponse); err != nil {
			return nil, fmt.Errorf("error occured while retrieving user account details %v", err)
		}
		response := AccountDetailsResponse{
			Success: mapResponse["success"].(bool),
			Message: mapResponse["message"].(string),
		}
		if name, ok := mapResponse["data"].(string); ok {
			if strings.TrimSpace(name) == "" {
				return nil, fmt.Errorf("could not fetch account details")
			}
			response.Data = name
			return &response, nil
		}
		if data, ok := mapResponse["data"].(map[string]interface{}); ok {
			response.Data = data["account_name"].(string)
			if strings.TrimSpace(response.Data) == "" {
				return nil, fmt.Errorf("could not fetch account details")
			}
			return &response, nil
		}

		return nil, fmt.Errorf("could not fetch account details")
	default:
		return nil, fmt.Errorf("only NGN is supported for account verification")

	}

}
