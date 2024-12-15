package budpay

import "fmt"

type BudpayBanks struct {
	BankName string `json:"bank_name"`
	BankCode string `json:"bank_code"`
}

type BankListResponse struct {
	ResponseSuccess
	Currency string        `json:"currency"`
	Data     []BudpayBanks `json:"data"`
}

func (budpayClient *BudPayClient) GetBankList(currency string) (*BankListResponse, error) {
	var response BankListResponse
	if err := budpayClient.Get(fmt.Sprintf("v2/bank_list/%s", currency), &response); err != nil {
		return nil, fmt.Errorf("error occured while retrieving bank list %v", err)
	}
	return &response, nil
}
