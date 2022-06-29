package entity

import (
	"github.com/google/uuid"
)

type PrincipalAmount struct {
	LoanID               uuid.UUID
	TotalRemainingAmount string
	Curreny              string
}
