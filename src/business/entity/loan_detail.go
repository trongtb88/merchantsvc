package entity

import (
	"github.com/google/uuid"
	"time"
)

type LoanDetail struct {
	LoanID       uuid.UUID
	LoanCode     string
	LoanStatus   string
	RejectReason string
	RejectTime   *time.Time
	ApprovedTime *time.Time
}
