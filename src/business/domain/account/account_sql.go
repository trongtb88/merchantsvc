package account

import (
	"context"
	"github.com/trongtb88/merchantsvc/src/business/entity"
	"github.com/trongtb88/merchantsvc/src/common"
	"log"
	"math"
	"time"
)

func (a *account) SQLCreateAccount(ctx context.Context, param entity.CreateAccountParam) (entity.MerchantAccount, error) {

	account := entity.MerchantAccount{
		MerchantCode: param.Code,
		MerchantName: param.Name,
		MerchantStatus: common.StatusActive,
	}

	tx := a.sql.Create(&account)
	if tx.Error != nil {
		return account, tx.Error
	}
	return account, nil
}

func (a *account) SQLUpdateAccount(ctx context.Context, acc entity.MerchantAccount) (entity.MerchantAccount, error) {
	acc.UpdatedAt = time.Now()
	tx := a.sql.Save(&acc)
	if tx.Error != nil {
		return acc, tx.Error
	}
	return acc, nil
}

func (a *account) SQLGetAccountsByParam(ctx context.Context, param entity.GetAccountParam) ([]entity.MerchantAccount, entity.Pagination, error) {

	var results [] entity.MerchantAccount
	var totalRows int64

	account := entity.MerchantAccount{
	}

	if param.Id > 0 {
		account.Id = param.Id
	}
	if len(param.Code) > 0 {
		account.MerchantCode = param.Code
	}
	if len(param.Name) > 0 {
		account.MerchantName = param.Name
	}

	if param.Status.Valid {
		account.MerchantStatus = int(param.Status.Int32)
	}

	if param.Limit <= 0 {
		param.Limit = 10
	}

	if param.Page <= 0 {
		param.Page = 1
	}

	tx := a.sql.Model(&entity.MerchantAccount{}).Debug().Where(&account).Limit(param.Limit).Offset((param.Page - 1) * param.Limit).Find(&results)
	a.sql.Model(&entity.MerchantAccount{}).Debug().Where(&account).Count(&totalRows)

	if tx.Error != nil {
		log.Printf("Error on tx", tx)
		return results, entity.Pagination{}, tx.Error
	}

	pagination := entity.Pagination{
		CurrentPage:     int64(param.Page),
		TotalPages:      int64(math.Ceil(float64(totalRows) / float64(param.Limit))),
		TotalElements:   totalRows,
		SortBy:          nil,
		CursorStart:     nil,
		CursorEnd:       nil,
	}
	return results, pagination, nil
}






