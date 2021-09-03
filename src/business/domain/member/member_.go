package member

import (
	"context"
	"github.com/trongtb88/merchantsvc/src/business/entity"
)

func (m *member) CreateMember(ctx context.Context, param entity.CreateMemberParam) (entity.MerchantMember, error) {
	return m.SQLCreateMember(ctx, param)
}

func (m *member) GetMembersByParam(ctx context.Context, param entity.GetMemberParam) ([]entity.MerchantMember, entity.Pagination, error) {
	return m.SQLGetMembersByParam(ctx, param)
}

func (m *member) UpdateMember(ctx context.Context, param entity.MerchantMember) (entity.MerchantMember, error) {
	return m.SQLUpdateMember(ctx, param)
}

