package domain

import (
	"github.com/trongtb88/merchantsvc/src/business/domain/account"
	"github.com/trongtb88/merchantsvc/src/business/domain/member"
	"github.com/trongtb88/merchantsvc/src/business/domain/question"
	"gorm.io/gorm"
)

type Domain struct {
	Account  account.DomainItf
	Member   member.DomainItf
	Question question.DomainItf
}

func Init(
	sql *gorm.DB,
) *Domain {

	return &Domain{
		Account: account.InitAccountDomain(
			sql,
		),
		Member: member.InitMemberDomain(
			sql,
		),
		Question: question.InitQuestionDomain(sql),
	}
}
