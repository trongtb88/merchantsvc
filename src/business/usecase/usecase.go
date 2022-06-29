package usecase

import (
	"github.com/trongtb88/merchantsvc/src/business/domain"
	"github.com/trongtb88/merchantsvc/src/business/usecase/account"
	"github.com/trongtb88/merchantsvc/src/business/usecase/answer"
	"github.com/trongtb88/merchantsvc/src/business/usecase/member"
	"github.com/trongtb88/merchantsvc/src/business/usecase/question"
)

type Usecase struct {
	Account  account.Usecaseitf
	Member   member.Usecaseitf
	Question question.Usecaseitf
	Answer   answer.Usecaseitf
}

// Init all usecase
func Init(
	dom *domain.Domain,
) *Usecase {

	return &Usecase{
		Account: account.InitAccount(
			dom.Account,
		),
		Member: member.InitMember(
			dom.Member,
			dom.Account,
		),
		Question: question.InitQuestion(
			dom.Question,
		),
		Answer: answer.InitAnswer(
			dom.Answer,
		),
	}
}
