package member


import (
	"context"
	Account "github.com/trongtb88/merchantsvc/src/business/domain/account"

	Member "github.com/trongtb88/merchantsvc/src/business/domain/member"
	"github.com/trongtb88/merchantsvc/src/business/entity"
)

// Usecaseitf uc insterface
type Usecaseitf interface {
	CreateMember(ctx context.Context, param entity.CreateMemberParam) (entity.MerchantMember, error)
	UpdateMember(ctx context.Context, param entity.MerchantMember) (entity.MerchantMember, error)
	GetMembersByParam(ctx context.Context, para entity.GetMemberParam) ([]entity.MerchantMemberData, entity.Pagination, error)
}

type member struct {
	member Member.DomainItf
	account Account.DomainItf
}

// Options for uc, that is for config
type Options struct {
}

// InitAccount
func InitMember(
	Member Member.DomainItf,
	Account Account.DomainItf,
) Usecaseitf {
	return &member{
		Member,
		Account,
	}
}

