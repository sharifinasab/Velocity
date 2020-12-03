package service

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"KOHO/dto"
	"KOHO/model"
)

// ParseDepositRequest returns a parsed version of DepositRequest
func ParseDepositRequest(depositRequest string) (*model.Deposit, error) {
	var dr dto.DepositRequest

	if err := json.Unmarshal([]byte(depositRequest), &dr); err != nil {
		log.Println("Error parsing line: ", err)
		return nil, err
	}

	a, err := strconv.ParseFloat(strings.Trim(dr.Amount, "$"), 64)
	if err != nil {
		log.Println("Error parsing amount: ", err)
		return nil, err
	}

	t, err := time.Parse(time.RFC3339, dr.Time)
	if err != nil {
		log.Println("Error parsing time: ", err)
		return nil, err
	}

	return &model.Deposit{
		DepositRequest: dr,
		ParsedAmount:   a,
		ParsedTime:     t,
	}, nil
}

// FormatDepositResponse formats the struct to be written in json
func FormatDepositResponse(id string, customerID string, accepted bool) string {
	response := &dto.DepositResponse{
		ID:         id,
		CustomerID: customerID,
		Accepted:   accepted,
	}

	str, err := json.Marshal(response)

	if err != nil {
		log.Println("Error formatting output: ", err)
		return ""
	}

	return string(str)
}
