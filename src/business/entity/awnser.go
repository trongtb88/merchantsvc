package entity

import (
	"github.com/google/uuid"
)

type Answer struct {
	Id      uuid.UUID
	Content interface{}
}
