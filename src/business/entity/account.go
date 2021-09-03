package entity

import "time"

type MerchantAccount struct {
	Id                  int64     `gorm:"primaryKey;autoIncrement;not_null" json:"id"`
	MerchantCode        string    `gorm:"unique;index;type:varchar(20);" json:"merchant_code"`
	MerchantName        string    `gorm:"type:varchar(200);" json:"merchant_name"`
	MerchantStatus      int       `gorm:"type:tinyint(1);" json:"-"`
	MerchantStatusInStr string    `gorm:"-" json:"merchant_status"`
	CreatedAt           time.Time `gorm:""DEFAULT:current_timestamp; type:timestamp"" json:"created_at"`
	UpdatedAt           time.Time `gorm:""DEFAULT:current_timestamp; type:timestamp"" json:"updated_at"`
}


type UpdateMerchantAccount struct {
	Id                  int64     `json:"id"`
	MerchantCode        string    `json:"merchant_code"`
	MerchantName        string    `json:"merchant_name"`
}


type GetAccountParam struct {
	Id     int64        `json:"id"`
	Code   string       `json:"merchant_code"`
	Name   string       `json:"merchant_name"`
	Page   int          `json:"page"`
	Limit  int          `json:"limit"`
}

type CreateAccountParam struct {
	Code   string       `json:"merchant_code"`
	Name   string       `json:"merchant_name"`
}
