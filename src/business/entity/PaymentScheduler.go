package entity

import (
	"github.com/google/uuid"
)

type PaymentScheduler struct {
	LoanID          uuid.UUID
	PayDate         string
	PrincipalAmount float64
	InterestAmount  float64
	Curreny         string
}
