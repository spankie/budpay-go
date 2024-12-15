package budpay

import "time"

type BudpayTime struct {
	time.Time
}

// TODO: add test for this and check if there is a RFC time format for it.
func (t *BudpayTime) UnmarshalJSON(b []byte) error {
	if b[len(b)-2] == 'Z' {
		b = append(b[:len(b)-2], b[len(b)-1:]...)
	}
	tt, err := time.Parse(`"2006-01-02T15:04:05"`, string(b))
	if err != nil {
		return err
	}
	*t = BudpayTime{tt}
	return nil
}

func (t *BudpayTime) MarshalJSON() ([]byte, error) {
	return []byte(t.Format(`"2006-01-02T15:04:05"`)), nil
}

type BudPayWebHook struct {
	Notify     string      `json:"notify" binding:"required"`
	NotifyType string      `json:"notifyType" binding:"required"`
	Data       interface{} `json:"data" binding:"required"`
}

type IncomingPaymentWebhookPayload struct {
	Notify          string          `json:"notify"`
	NotifyType      string          `json:"notifyType"`
	Data            IncomingPayment `json:"data"`
	TransferDetails TransferDetails `json:"transferDetails"`
}

type IncomingPayment struct {
	ID              int            `json:"id"`
	BusinessId      int            `json:"business_id"`
	Currency        string         `json:"currency"`
	Amount          string         `json:"amount"`
	Reference       string         `json:"reference"`
	IpAddress       interface{}    `json:"ip_address"`
	Channel         string         `json:"channel"`
	Type            string         `json:"type"`
	Domain          string         `json:"domain"`
	Fees            string         `json:"fees"`
	CustomerId      int            `json:"customer_id"`
	Plan            interface{}    `json:"plan"`
	RequestedAmount string         `json:"requested_amount"`
	Status          string         `json:"status"`
	CardAttempt     int            `json:"card_attempt"`
	Message         string         `json:"message"`
	Gateway         string         `json:"gateway"`
	CreatedAt       BudpayTime     `json:"created_at"`
	PaidAt          string         `json:"paid_at"`
	Customer        CustomerDetail `json:"customer"`
}

type CustomerDetail struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Domain       string `json:"domain"`
	CustomerCode string `json:"customer_code"`
	Metadata     string `json:"metadata"`
	Status       string `json:"status"`
}

type TransferDetails struct {
	PaymentReference        string `json:"paymentReference"`
	SessionID               string `json:"sessionid"`
	Craccount               string `json:"craccount"`
	Narration               string `json:"narration"`
	Amount                  string `json:"amount"`
	Originatoraccountnumber string `json:"originatoraccountnumber"`
	Originatorname          string `json:"originatorname"`
	Bankname                string `json:"bankname"`
	Bankcode                string `json:"bankcode"`
	Craccountname           string `json:"craccountname"`
}

// TODO: should be renamed to payoutwebhookpayload
type OutgoingWebhookPayload struct {
	Notify     string       `json:"notify"`
	NotifyType string       `json:"notifyType"`
	Data       OutgoingData `json:"data"`
}

type OutgoingData struct {
	Id            int    `json:"id"`
	Reference     string `json:"reference"`
	Sessionid     string `json:"sessionid"`
	Currency      string `json:"currency"`
	Amount        string `json:"amount"`
	Fee           string `json:"fee"`
	BankCode      string `json:"bank_code"`
	BankName      string `json:"bank_name"`
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
	CountryCode   string `json:"countryCode"`
	PaymentMode   string `json:"paymentMode"`
	Narration     string `json:"narration"`
	Sender        string `json:"sender"`
	Domain        string `json:"domain"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
