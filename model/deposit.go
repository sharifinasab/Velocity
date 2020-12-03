package model

import (
	"time"

	"KOHO/dto"
)

//Deposit is a parsed version of DepositRequest
type Deposit struct {
	DepositRequest dto.DepositRequest
	ParsedAmount   float64
	ParsedTime     time.Time
}
