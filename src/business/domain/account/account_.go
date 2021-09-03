package account

import (
	"context"
	"github.com/trongtb88/merchantsvc/src/business/entity"
)

func (a *account) CreateAccount(ctx context.Context, param entity.CreateAccountParam) (entity.MerchantAccount, error) {
	return a.SQLCreateAccount(ctx, param)
}

func (a *account) UpdateAccountBy(ctx context.Context, param entity.MerchantAccount) (entity.MerchantAccount, error) {
	return a.SQLUpdateAccount(ctx, param)
}

func (a *account) GetAccountsByParam(ctx context.Context, param entity.GetAccountParam) ([]entity.MerchantAccount, entity.Pagination, error) {
	return a.SQLGetAccountsByParam(ctx, param)
}
