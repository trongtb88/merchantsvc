package account

import (
	"context"
	"github.com/trongtb88/merchantsvc/src/business/entity"
	"github.com/trongtb88/merchantsvc/src/common"
)


func (a *account) CreateAccount(ctx context.Context, param entity.CreateAccountParam) (entity.MerchantAccount, error) {
	acc, err := a.account.CreateAccount(ctx, param)
	acc.MerchantStatusInStr = common.StatusActiveStr
	return  acc, err
}

func (a *account) UpdateAccount(ctx context.Context, param entity.MerchantAccount) (entity.MerchantAccount, error) {
	return a.account.UpdateAccountBy(ctx, param)
}

func (a *account) GetAccountsByParam(ctx context.Context, param entity.GetAccountParam) ([]entity.MerchantAccount, entity.Pagination, error) {
	// Can do some business logic in usecase
	accs, p, err := a.account.GetAccountsByParam(ctx, param)
	if err == nil {
		for idx, _ := range accs {
			if accs[idx].MerchantStatus == common.StatusActive {
				accs[idx].MerchantStatusInStr = "Active"
			} else {
				accs[idx].MerchantStatusInStr = "InActive"
			}
		}
	}
	return accs, p, err
}
