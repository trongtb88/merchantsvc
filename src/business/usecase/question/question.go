package question

import (
	"context"
	Question "github.com/trongtb88/merchantsvc/src/business/domain/question"
	"github.com/trongtb88/merchantsvc/src/business/entity"
)

// Usecaseitf uc interface
type Usecaseitf interface {
	GetPredefineQuestionsForAuthenticate(ctx context.Context) ([]entity.QuestionForAuthentication, error)
	GetPredefineQuestionsForBusiness(ctx context.Context, topic string) ([]entity.Question, error)
}

type questionUc struct {
	question Question.DomainItf
}

// Options for uc, that is for config
type Options struct {
}

// InitQuestion
func InitQuestion(
	Question Question.DomainItf,
) Usecaseitf {
	return &questionUc{
		Question,
	}
}
