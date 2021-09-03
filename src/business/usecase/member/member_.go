package member

import (
	"context"
	"github.com/trongtb88/merchantsvc/src/business/entity"
	"github.com/trongtb88/merchantsvc/src/common"
	"log"
)

func (m *member) CreateMember(ctx context.Context, param entity.CreateMemberParam) (entity.MerchantMember, error) {
	member, err := m.member.CreateMember(ctx, param)
	member.MemberStatusInStr = common.StatusActiveStr
	return  member, err
}

func (m *member) UpdateMember(ctx context.Context, param entity.MerchantMember) (entity.MerchantMember, error) {
	return m.member.UpdateMember(ctx, param)
}

func (m *member) GetMembersByParam(ctx context.Context, param entity.GetMemberParam) ([]entity.MerchantMemberData, entity.Pagination, error) {
	var (
		membersData [] entity.MerchantMemberData
		err error
	)
	// Can do some business logic in usecase
	members, p, err := m.member.GetMembersByParam(ctx, param)
	if err == nil {
		for idx, _ := range members {
			memberData := entity.MerchantMemberData{
				Id:               members[idx].Id,
				MemberName:      members[idx].MemberName,
				MemberEmail:     members[idx].MemberEmail,
				CreatedAt:       members[idx].CreatedAt,
				UpdatedAt:       members[idx].UpdatedAt,
			}

			if members[idx].MemberStatus == common.StatusActive {
				memberData.MemberStatus = "Active"
			} else {
				memberData.MemberStatus = "InActive"
			}

			acc, _, err := m.account.GetAccountsByParam(ctx, entity.GetAccountParam{
				Id:  members[idx].MerchantId,
			})

			if err != nil {
				log.Println("Error when get merchant account", err)
			} else {
				memberData.MerchantAccount = acc[0]
			}
			membersData = append(membersData, memberData)
		}
	}
	return membersData, p, err
}
