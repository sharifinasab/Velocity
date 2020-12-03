package dto

// DepositRequest structure for transaction to add fund
type DepositRequest struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Amount     string `json:"load_amount"`
	Time       string `json:"time"`
}
