package account

import (
	"context"
	"github.com/trongtb88/merchantsvc/src/business/entity"
	"gorm.io/gorm"
)

// DomainItf domain interface for merchant account
type DomainItf interface {
	CreateAccount(ctx context.Context, param entity.CreateAccountParam) (entity.MerchantAccount, error)
	GetAccountsByParam(ctx context.Context, param entity.GetAccountParam) ([]entity.MerchantAccount, entity.Pagination,  error)
	UpdateAccountBy(ctx context.Context, param entity.MerchantAccount) (entity.MerchantAccount, error)
}

type account struct {
	sql         *gorm.DB
}

// InitAccountDomain domain init
func InitAccountDomain(
	sql *gorm.DB,
	) DomainItf {
	a := &account{
		sql:         sql,
	}
	return a
}
