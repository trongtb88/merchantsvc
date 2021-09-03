package account

import (
	"context"
	Account "github.com/trongtb88/merchantsvc/src/business/domain/account"
	"github.com/trongtb88/merchantsvc/src/business/entity"
)

// Usecaseitf uc insterface
type Usecaseitf interface {
	CreateAccount(ctx context.Context, param entity.CreateAccountParam) (entity.MerchantAccount, error)
	UpdateAccount(ctx context.Context, param entity.MerchantAccount) (entity.MerchantAccount, error)
	GetAccountsByParam(ctx context.Context, para entity.GetAccountParam) ([]entity.MerchantAccount, entity.Pagination, error)
}

type account struct {
	account Account.DomainItf
}

// Options for uc, that is for config
type Options struct {
}

// InitAccount
func InitAccount(
	Account Account.DomainItf,
) Usecaseitf {
	return &account{
		Account,
	}
}
