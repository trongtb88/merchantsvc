package rest

import "github.com/trongtb88/merchantsvc/src/business/entity"

type ResponseSuccessNonPagination struct {
	Meta entity.Meta `json:"metadata"`
	Data interface{} `json:"data"`
}

type ResponseSuccessPagination struct {
	Meta       entity.Meta       `json:"metadata"`
	Data       interface{}       `json:"data"`
	Pagination entity.Pagination `json:"pagination"`
}

type ResponseCreateAccount struct {
	Meta entity.Meta            `json:"metadata"`
	Data entity.MerchantAccount `json:"data"`
}

type ResponseCreateMember struct {
	Meta entity.Meta           `json:"metadata"`
	Data entity.MerchantMember `json:"data"`
}

type ResponseGetAccounts struct {
	Meta       entity.Meta              `json:"metadata"`
	Data       []entity.MerchantAccount `json:"data"`
	Pagination entity.Pagination        `json:"pagination"`
}

type ResponseGetMembers struct {
	Meta       entity.Meta                 `json:"metadata"`
	Data       []entity.MerchantMemberData `json:"data"`
	Pagination entity.Pagination           `json:"pagination"`
}

type HTTPErrResp struct {
	Meta entity.Meta `json:"metadata"`
}
type HTTPEmptyResp struct {
	Meta entity.Meta `json:"metadata"`
}
