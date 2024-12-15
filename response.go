package budpay

type ResponseStatus struct {
	Status  bool          `json:"status"`
	Message string        `json:"message"`
}

type ResponseSuccess struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
}