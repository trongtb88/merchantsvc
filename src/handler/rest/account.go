package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/trongtb88/merchantsvc/src/business/entity"
	"github.com/trongtb88/merchantsvc/src/common"
	"io/ioutil"
	"net/http"
	"strconv"
)
var decoder = schema.NewDecoder()


// Create Merchant Account godoc
// @Summary Create Merchant Account
// @Description Create Merchant Account
// @Tags MerchantAccount
// @Accept json
// @Produce json
// @Param data body entity.CreateAccountParam true "Body Request"
// @Success 201 {object} rest.ResponseCreateAccount
// @Failure 400 {object} rest.HTTPErrResp
// @Failure 401 {object} rest.HTTPErrResp
// @Failure 500 {object} rest.HTTPErrResp
// @Router /v1/accounts [post]
func (rst *rest) CreateMerchantAccount(w http.ResponseWriter, r *http.Request) {
	var param entity.CreateAccountParam
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rst.httpRespError(w, r, http.StatusBadRequest, ErrorMessage{
			Code:     "InvalidJsonBodyRequest",
			Message: err.Error(),
		})
		return
	}

	if err := json.Unmarshal(body, &param); err != nil {
		rst.httpRespError(w, r, http.StatusBadRequest, ErrorMessage{
			Code:     "InvalidJsonBodyRequest",
			Message: err.Error(),
		})
		return
	}

	accounts, _,  err := rst.uc.Account.GetAccountsByParam(r.Context(), entity.GetAccountParam{
		Code:  param.Code,
	})

	if err != nil {
		rst.httpRespError(w, r, http.StatusInternalServerError, ErrorMessage{
			Code:     "CreateAccountError",
			Message: err.Error(),
		})
		return
	}

	if len(accounts) > 0 {
		rst.httpRespError(w, r,http.StatusBadRequest, ErrorMessage{
			Code:     "DuplicateMerchantCode",
			Message: "Please choose other merchant code",
		})
		return
	}

	account, err := rst.uc.Account.CreateAccount(r.Context(), param)
	if err != nil {
		rst.httpRespError(w, r, http.StatusInternalServerError, ErrorMessage{
			Code:     "CreateAccountError",
			Message: err.Error(),
		})
		return
	}

	rst.httpRespSuccess(w, r, http.StatusCreated, account, nil)
}

// Update Merchant Account By Id godoc
// @Summary Update Merchant Account By Id
// @Description Update Merchant Account By Id
// @Tags MerchantAccount
// @Accept json
// @Produce json
// @Param data body entity.UpdateMerchantAccount true "Body Request"
// @Success 200 {object} rest.ResponseCreateAccount
// @Failure 400 {object} rest.HTTPErrResp
// @Failure 401 {object} rest.HTTPErrResp
// @Failure 500 {object} rest.HTTPErrResp
// @Router /v1/accounts [put]
func (rst *rest) UpdateMerchantAccounts(w http.ResponseWriter, r *http.Request) {
	var param entity.UpdateMerchantAccount
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rst.httpRespError(w, r, http.StatusBadRequest, ErrorMessage{
			Code:     "InvalidJsonBodyRequest",
			Message: err.Error(),
		})
		return
	}

	if err := json.Unmarshal(body, &param); err != nil {
		rst.httpRespError(w, r, http.StatusBadRequest, ErrorMessage{
			Code:     "InvalidJsonBodyRequest",
			Message: err.Error(),
		})
		return
	}

	if param.Id <= 0 {
		rst.httpRespError(w, r, http.StatusBadRequest, ErrorMessage{
			Code:     "InvalidMerchantId",
			Message: "Invalid MerchantId",
		})
		return
	}

	if len(param.MerchantCode) <= 0 {
		rst.httpRespError(w, r, http.StatusBadRequest, ErrorMessage{
			Code:     "InvalidMerchantCode",
			Message: "Invalid Merchant Code",
		})
		return
	}

	accountsById, _,  err := rst.uc.Account.GetAccountsByParam(r.Context(), entity.GetAccountParam{
		Id:  param.Id,
	})

	accountsByCode, _,  err := rst.uc.Account.GetAccountsByParam(r.Context(), entity.GetAccountParam{
		Code:  param.MerchantCode,
	})

	if len(accountsById) == 0 {
		rst.httpRespError(w, r, http.StatusNotFound, ErrorMessage{
			Code:     "NotFoundMerchantId",
			Message: "Not Found Merchant Id",
		})
		return
	}

	if len(accountsByCode) > 0 && accountsByCode[0].Id != param.Id {
		rst.httpRespError(w, r, http.StatusNotFound, ErrorMessage{
			Code:     "ExistMerchantCode",
			Message: "This merchant code is belong to other merchant, please choose other value",
		})
		return
	}

	acc := accountsById[0]
	acc.MerchantCode = param.MerchantCode
	acc.MerchantName = param.MerchantName

	accounts, err := rst.uc.Account.UpdateAccount(r.Context(), acc)

	if err != nil {
		rst.httpRespError(w, r, http.StatusInternalServerError, ErrorMessage{
			Code:     "GetAccountsError",
			Message: err.Error(),
		})
		return
	}

	rst.httpRespSuccess(w, r, http.StatusOK, accounts, nil)
}

// Delete Merchant Account By Id godoc
// @Summary Delete Merchant Account By Id
// @Description Delete Merchant Account By Id
// @Tags MerchantAccount
// @Accept json
// @Produce json
// @Param account_id path integer true "Merchant Account ID"
// @Success 200 {object} rest.ResponseCreateAccount
// @Failure 400 {object} rest.HTTPErrResp
// @Failure 401 {object} rest.HTTPErrResp
// @Failure 500 {object} rest.HTTPErrResp
// @Router /v1/accounts/{account_id} [delete]
func (rst *rest) DeleteMerchantAccounts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var accountId int64
	var err error
	idParam, _ := vars["account_id"]

	if len(idParam) == 0 {
		rst.httpRespError(w, r, http.StatusBadRequest, ErrorMessage{
			Code:     "InvalidMerchantAccountId",
			Message: "Invalid Merchant Account Id",
		})
		return
	}

	if accountId, err = strconv.ParseInt(idParam,10,64); err != nil {
		rst.httpRespError(w, r, http.StatusBadRequest, ErrorMessage{
			Code:     "InvalidMerchantAccountId",
			Message: "MerchantAccountId must be number",
		})
		return
	}

	accountsById, _,  err := rst.uc.Account.GetAccountsByParam(r.Context(), entity.GetAccountParam{
		Id: accountId ,
	})

	if len(accountsById) == 0 {
		rst.httpRespError(w, r, http.StatusNotFound, ErrorMessage{
			Code:     "NotFoundMerchantId",
			Message: "Not Found Merchant Id",
		})
		return
	}

	acc := accountsById[0]
	acc.MerchantStatus = common.StatusInActive

	_, err = rst.uc.Account.UpdateAccount(r.Context(), acc)
	if err != nil {
		rst.httpRespError(w, r, http.StatusInternalServerError, ErrorMessage{
			Code:     "DeleteAccountsError",
			Message: err.Error(),
		})
		return
	}
	rst.httpRespSuccess(w, r, http.StatusOK, acc, nil)
}

// Get Merchant Accounts
// @Summary Get Merchant Accounts godoc
// @Description Get Merchant Accounts godoc
// @Tags MerchantAccount
// @Accept json
// @Produce json
// @Param id query integer false "Merchant Account ID"
// @Param code query string false "Merchant Account Code"
// @Param name query integer false "Merchant Account Name"
// @Param page query int false " "
// @Param limit query int false " "
// @Success 200 {object} rest.ResponseGetAccounts
// @Failure 400 {object} rest.HTTPErrResp
// @Failure 401 {object} rest.HTTPErrResp
// @Failure 500 {object} rest.HTTPErrResp
// @Router /v1/accounts [get]
func (rst *rest) GetMerchantAccounts(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		rst.httpRespError(w, r, http.StatusBadRequest, ErrorMessage{
			Code:     "GetAccountsError",
			Message: err.Error(),
		})
		return
	}

	var param entity.GetAccountParam

	err = decoder.Decode(&param, r.Form)
	if err != nil {
		rst.httpRespError(w, r, http.StatusBadRequest, ErrorMessage{
			Code:     "GetAccountsError",
			Message: err.Error(),
		})
		return
	}


	accounts, pagination, err := rst.uc.Account.GetAccountsByParam(r.Context(), param)

	if err != nil {
		rst.httpRespError(w, r, http.StatusInternalServerError, ErrorMessage{
			Code:     "GetAccountsError",
			Message: err.Error(),
		})
		return
	}

	rst.httpRespSuccess(w, r, http.StatusOK, accounts, &pagination)
}

