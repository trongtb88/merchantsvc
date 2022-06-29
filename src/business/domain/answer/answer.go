package answer

import (
	"context"
	"github.com/trongtb88/merchantsvc/src/business/entity"
	"gorm.io/gorm"
)

type DomainItf interface {
	// DomainItf domain interface for answer
	SubmitQuestionForAnswer(ctx context.Context, question entity.Question) (entity.Answer, error)
}

type answer struct {
	sql *gorm.DB
}

// InitAnswerDomain domain init
func InitAnswerDomain(
	sql *gorm.DB,
) DomainItf {
	a := &answer{
		sql: sql,
	}
	return a
}
