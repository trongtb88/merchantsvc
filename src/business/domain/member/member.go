package member

import (
	"context"
	"github.com/trongtb88/merchantsvc/src/business/entity"
	"gorm.io/gorm"
)

// DomainItf domain interface for merchant members of merchants
type DomainItf interface {
	CreateMember(ctx context.Context, param entity.CreateMemberParam) (entity.MerchantMember, error)
	GetMembersByParam(ctx context.Context, param entity.GetMemberParam) ([]entity.MerchantMember, entity.Pagination,  error)
	UpdateMember(ctx context.Context, param entity.MerchantMember) (entity.MerchantMember, error)
}

type member struct {
	sql         *gorm.DB
}

// InitMemberDomain domain init
func InitMemberDomain(
	sql *gorm.DB,
) DomainItf {
	a := &member{
		sql:         sql,
	}
	return a
}
