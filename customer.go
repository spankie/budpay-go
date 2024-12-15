package budpay

import "fmt"

type CustomerRequest struct {
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Phone     string `json:"phone"`
	MetaData  string `json:"metadata"`
}

type CustomerResponse struct {
	Status  bool                 `json:"status"`
	Message string               `json:"message"`
	Data    CustomerResponseInfo `json:"data"`
}

type CustomerResponseInfo struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Domain       string `json:"domain"`
	CustomerCode string `json:"customer_code"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type Customer struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	CustomerCode string `json:"customer_code"`
	Phone        string `json:"phone"`
	Domain       string `json:"domain"`
	Metadata     string `json:"metadata"`
	Status       string `json:"status"`
}

func (budpayClient *BudPayClient) CreateCustomer(reqBody *CustomerRequest) (*CustomerResponse, error) {
	var response CustomerResponse
	if err := budpayClient.Post("v2/customer", reqBody, &response); err != nil {
		return nil, fmt.Errorf("error while creating customer: %v", err)
	}
	return &response, nil
}
