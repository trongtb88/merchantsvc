package entity

import "github.com/google/uuid"

type QuestionForAuthentication struct {
	Id           uuid.UUID
	Content      string
	QuestionType string
	IsRequired   bool
	SupportOTP   bool
	Field        string
}

type Question struct {
	Id           uuid.UUID
	Content      string
	QuestionType string
	Topic        string
	Function     string
	Parameters   []Parameter
}

type Parameter struct {
	Name       string
	Value      string
	IsRequired bool
}
