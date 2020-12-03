package model

import "time"

// IRule defines an interface for daily / weekly limit rules
type IRule interface {
	Validate(transactionTime time.Time, amount float64) bool
	UpdateQuota(amount float64)
}
