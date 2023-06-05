package dto

import (
	"banking/errs"
	"strings"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit atleast 5000.00")
	}
	accType := strings.ToLower(r.AccountType)
	if accType != "saving" && accType != "checking" {
		return errs.NewValidationError("Account type should be checking or saving")
	}
	return nil
}
