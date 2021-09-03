package rest

import "github.com/trongtb88/merchantsvc/src/business/entity"

type ErrorMessage struct {
	Code     string `json:"code"`
	Message string `json:"message"`
}

type Meta struct {
	Path       string `json:"path"`
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
	Error      ErrorMessage `json:"error,omitempty"`
	Timestamp  string `json:"timestamp"`
}

type HTTPErrResp struct {
	Meta Meta `json:"metadata"`
}
type HTTPEmptyResp struct {
	Meta Meta `json:"metadata"`
}

type ResponseSuccessNonPagination struct {
	Meta Meta                   `json:"metadata"`
	Data interface{} `json:"data"`
}

type ResponseSuccessPagination struct {
	Meta Meta                   `json:"metadata"`
	Data interface{}            `json:"data"`
	Pagination entity.Pagination `json:"pagination"`
}

type ResponseCreateAccount struct {
	Meta Meta                   `json:"metadata"`
	Data entity.MerchantAccount `json:"data"`
}

type ResponseCreateMember struct {
	Meta Meta                   `json:"metadata"`
	Data entity.MerchantMember `json:"data"`
}

type ResponseGetAccounts struct {
	Meta Meta                      `json:"metadata"`
	Data []entity.MerchantAccount  `json:"data"`
	Pagination entity.Pagination   `json:"pagination"`
}

type ResponseGetMembers struct {
	Meta Meta                      `json:"metadata"`
	Data []entity.MerchantMemberData  `json:"data"`
	Pagination entity.Pagination   `json:"pagination"`
}
