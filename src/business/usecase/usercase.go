package usecase

import (
	"github.com/trongtb88/merchantsvc/src/business/domain"
	"github.com/trongtb88/merchantsvc/src/business/usecase/account"
	"github.com/trongtb88/merchantsvc/src/business/usecase/member"
)

type Usecase struct {
	Account account.Usecaseitf
	Member member.Usecaseitf
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
	}
}
