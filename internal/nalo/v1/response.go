package nalo

type SMSResponse struct {
	Destination string
	MessageId   string
}

type CreditBalanceResponse struct {
	Balance  float64 `json:"balance"`
	TotalSMS int64   `json:"total_sms"`
}
