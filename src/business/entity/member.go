package entity

import "time"

type MerchantMember struct {
	Id                int64     `gorm:"primaryKey;autoIncrement;not_null" json:"id"`
	MerchantId        int64     `gorm:"type:bigint(20);" json:"merchant_id"`
	MemberName        string    `gorm:"type:varchar(100);" json:"member_name"`
	MemberEmail       string    `gorm:"unique;index;type:varchar(100);" json:"member_email"`
	MemberStatus      int       `gorm:"type:tinyint(1);" json:"-"`
	MemberStatusInStr string    `gorm:"-" json:"member_status"`
	CreatedAt         time.Time `gorm:""DEFAULT:current_timestamp; type:timestamp"" json:"created_at"`
	UpdatedAt         time.Time `gorm:""DEFAULT:current_timestamp; type:timestamp"" json:"updated_at"`
}

type MerchantMemberData struct {
	Id              int64           `json:"id"`
	MerchantAccount MerchantAccount `json:"merchant"`
	MemberName      string          `json:"member_name"`
	MemberEmail     string          `json:"member_email"`
	MemberStatus    string          `json:"member_status"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type CreateMemberParam struct {
	MerchantId int64  `json:"merchant_id"`
	Email   string    `json:"member_email"`
	Name    string    `json:"member_name"`
}

type GetMemberParam struct {
	Id         int64  `json:"id"`
	MerchantId int64  `json:"merchant_id"`
	Email      string `json:"member_email"`
	Name       string `json:"member_name"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
}