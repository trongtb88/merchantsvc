package answer

import (
	"context"
	Answer "github.com/trongtb88/merchantsvc/src/business/domain/answer"
	"github.com/trongtb88/merchantsvc/src/business/entity"
)

// Usecaseitf uc interface
type Usecaseitf interface {
	SubmitQuestionForAnswer(ctx context.Context, question entity.Question) (entity.Answer, error)
}

type answerUc struct {
	answer Answer.DomainItf
}

// Options for uc, that is for config
type Options struct {
}

// InitAnswer
func InitAnswer(
	Answer Answer.DomainItf,
) Usecaseitf {
	return &answerUc{
		Answer,
	}
}
