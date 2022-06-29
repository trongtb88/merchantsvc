package entity

import (
	"github.com/google/uuid"
)

type PendingTransaction struct {
	TransactionId uuid.UUID
	WithdrawDate  string
	Status        string
	BankAccount   string
	BankName      string
	PendingReason string
}
