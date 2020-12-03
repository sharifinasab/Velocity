package repository

import (
	"log"
	"time"

	"KOHO/util"
)

// DailyLimit contains configurable limits to avoid hardcoding the limitations
type DailyLimit struct {
	remainingDailyDeposit         float64
	remainingDailyTransationCount int
	effectiveDate                 time.Time
}

// NewDailyLimitRule initialize a rule for an account
func NewDailyLimitRule(t time.Time) *DailyLimit {
	return &DailyLimit{
		remainingDailyDeposit:         5000,
		remainingDailyTransationCount: 3,
		effectiveDate:                 util.GetDayStartTime(t),
	}
}

// Validate checks the daily limits
func (dl *DailyLimit) Validate(amount float64) bool {
	if dl.remainingDailyDeposit-amount < 0 || dl.remainingDailyTransationCount-1 < 0 {

		if dl.remainingDailyDeposit-amount < 0 {
			log.Println("The transaction amount ", amount,
				" violates remaining daily limit ", dl.remainingDailyDeposit)
		}

		if dl.remainingDailyTransationCount-1 < 0 {
			log.Println("The transaction count violates remaining daily limit ",
				dl.remainingDailyTransationCount)
		}

		return false
	}

	return true
}

// UpdateQuota decreases the remaining daily deposit room after a successful transaction
func (dl *DailyLimit) UpdateQuota(amount float64) {
	dl.remainingDailyDeposit -= amount
	dl.remainingDailyTransationCount--
}
