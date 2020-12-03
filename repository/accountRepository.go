package repository

import (
	"log"
	"time"

	"KOHO/model"
	"KOHO/util"
)

// IAccountRepository represents the interface of the account repository
type IAccountRepository interface {
	Deposit(d *model.Deposit) error
	ValidateDeposit(d *model.Deposit) error
	PrintRemainingLimits()
	SetAccountQuota(t *time.Time)
}

// Account holds the structure of a customer account
type Account struct {
	ID              string
	Balance         float64
	DailyLimitRule  *DailyLimit
	WeeklyLimitRule *WeeklyLimit
}

// NewAccount initialize an account using defaults
func NewAccount(customerID string) *Account {
	return &Account{
		ID:      customerID,
		Balance: 0,
	}
}

// Deposit adds fund to account and updates the limits
func (a *Account) Deposit(d *model.Deposit) error {
	a.DailyLimitRule.UpdateQuota(d.ParsedAmount)
	a.WeeklyLimitRule.UpdateQuota(d.ParsedAmount)

	return nil
}

// ValidateDeposit checks if the rules of transaction are valid
func (a *Account) ValidateDeposit(d *model.Deposit) bool {
	return a.DailyLimitRule.Validate(d.ParsedAmount) &&
		a.WeeklyLimitRule.Validate(d.ParsedAmount)
}

// PrintRemainingDepositLimits prints remaining room to deposite
func (a *Account) PrintRemainingDepositLimits() {
	log.Println("*** Remaining quota for account ", a.ID, " ***")
	log.Println("Remaining daily amount: ", a.DailyLimitRule.remainingDailyDeposit)
	log.Println("Remaining daily count: ", a.DailyLimitRule.remainingDailyTransationCount)
	log.Println("Remaining weekly amount: ", a.WeeklyLimitRule.remainingWeeklyDeposit)
	log.Println("********************")
}

// SetAccountQuota checks if rules are set
func (a *Account) SetAccountQuota(ruleTime time.Time) {
	// New account or Daily rule expired
	if a.DailyLimitRule == nil ||
		!util.SameDay(a.DailyLimitRule.effectiveDate, util.GetDayStartTime(ruleTime)) {
		a.DailyLimitRule = NewDailyLimitRule(ruleTime)
	}

	// New account or Daily rule expired
	if a.WeeklyLimitRule == nil ||
		!util.SameWeek(a.WeeklyLimitRule.effectiveDate, util.GetWeekStartTime(ruleTime)) {
		a.WeeklyLimitRule = NewWeeklyLimitRule(ruleTime)
	}
}
