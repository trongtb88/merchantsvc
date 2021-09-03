package domain

import (
	"github.com/trongtb88/merchantsvc/src/business/domain/account"
	"github.com/trongtb88/merchantsvc/src/business/domain/member"
	"gorm.io/gorm"
)

type Domain struct {
	Account       account.DomainItf
	Member        member.DomainItf
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
	}
}
