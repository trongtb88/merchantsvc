package entity

import (
	"github.com/google/uuid"
)

type PendingTransaction struct {
	TransactionId uuid.UUID
	WithdrawDate  string
	Status        string
	BankAccount   int64
	BankName      string
	PendingReason string
}
