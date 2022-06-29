package question

import (
	"context"
	"github.com/trongtb88/merchantsvc/src/business/entity"
)

func (q questionUc) GetPredefineQuestionsForAuthenticate(ctx context.Context) ([]entity.QuestionForAuthentication, error) {
	return q.question.GetPredefineQuestionsForAuthenticate(ctx)
}

func (q questionUc) GetPredefineQuestionsForBusiness(ctx context.Context, topic string) ([]entity.Question, error) {
	return q.question.GetPredefineQuestionsForBusiness(ctx, topic)
}
