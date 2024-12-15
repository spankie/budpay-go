package budpay

import (
	"fmt"
)

type encryptResponse struct {
	Card string `json:"card"`
}

func (budpayClient *BudPayClient) encryptCardDetails(cardDetails CardEncryptReq) (string, error) {
	var encRes string

	if err := budpayClient.Post("s2s/test/encryption", cardDetails, &encRes); err != nil {
		return "", fmt.Errorf("error occured while encrypting card details %v", err)
	}
	//log.Printf("Encrypted card details: %v", encRes)
	return encRes, nil
}
