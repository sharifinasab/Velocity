package dto

// DepositResponse is an ack for adding fund
type DepositResponse struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Accepted   bool   `json:"accepted"`
}
