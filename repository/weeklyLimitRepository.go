package repository

import (
	"log"
	"time"

	"KOHO/util"
)

// WeeklyLimit contains configurable limits to avoid hardcoding the limitations
type WeeklyLimit struct {
	remainingWeeklyDeposit float64
	effectiveDate          time.Time
}

// NewWeeklyLimitRule initialize a rule for an account
func NewWeeklyLimitRule(t time.Time) *WeeklyLimit {
	return &WeeklyLimit{
		remainingWeeklyDeposit: 20000,
		effectiveDate:          util.GetWeekStartTime(t),
	}
}

// Validate checks the weekly limits
// Times are in UTC
func (wl *WeeklyLimit) Validate(amount float64) bool {
	if wl.remainingWeeklyDeposit-amount < 0 {
		log.Println("The transaction amount ", amount,
			" violates remaining daily limit ", wl.remainingWeeklyDeposit)

		return false
	}

	return true
}

// UpdateQuota decreases the remaining weekly deposit room after a successful transaction
func (wl *WeeklyLimit) UpdateQuota(amount float64) {
	wl.remainingWeeklyDeposit -= amount
}
