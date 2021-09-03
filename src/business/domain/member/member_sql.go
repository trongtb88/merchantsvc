package member

import (
	"context"
	"github.com/trongtb88/merchantsvc/src/business/entity"
	"github.com/trongtb88/merchantsvc/src/common"
	"log"
	"math"
	"time"
)

func (m *member) SQLCreateMember(ctx context.Context, param entity.CreateMemberParam) (entity.MerchantMember, error) {
	member := entity.MerchantMember{
		MerchantId:        param.MerchantId,
		MemberName:        param.Name,
		MemberEmail:       param.Email,
		MemberStatus:      1,
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
	}

	tx := m.sql.Create(&member)
	if tx.Error != nil {
		return member, tx.Error
	}
	return member, nil
}

func (m *member) SQLUpdateMember(ctx context.Context, acc entity.MerchantMember) (entity.MerchantMember, error) {
	acc.UpdatedAt = time.Now()
	tx := m.sql.Save(&acc)
	if tx.Error != nil {
		return acc, tx.Error
	}
	return acc, nil
}

func (m *member) SQLGetMembersByParam(ctx context.Context, param entity.GetMemberParam) ([]entity.MerchantMember, entity.Pagination, error) {

	var results [] entity.MerchantMember
	var totalRows int64

	member := entity.MerchantMember{
	}

	if param.Id > 0 {
		member.Id = param.Id
	}
	if param.MerchantId > 0 {
		member.MerchantId = param.MerchantId
	}
	if len(param.Email) > 0 {
		member.MemberEmail = param.Email
	}

	member.MemberStatus = common.StatusActive

	if param.Limit <= 0 {
		param.Limit = 10
	}

	if param.Page <= 0 {
		param.Page = 1
	}

	tx := m.sql.Model(&entity.MerchantMember{}).Debug().Where(&member).Limit(param.Limit).Offset((param.Page - 1) * param.Limit).Find(&results)
	m.sql.Model(&entity.MerchantMember{}).Debug().Where(&member).Count(&totalRows)

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

