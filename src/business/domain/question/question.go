package question

import (
	"context"
	"github.com/trongtb88/merchantsvc/src/business/entity"
	"gorm.io/gorm"
)

// DomainItf domain interface for question
type DomainItf interface {
	GetPredefineQuestionsForAuthenticate(ctx context.Context) ([]entity.QuestionForAuthentication, error)
	GetPredefineQuestionsForBusiness(ctx context.Context, topic string) ([]entity.Question, error)
}

type question struct {
	sql *gorm.DB
}

// InitQuestionDomain domain init
func InitQuestionDomain(
	sql *gorm.DB,
) DomainItf {
	a := &question{
		sql: sql,
	}
	return a
}
