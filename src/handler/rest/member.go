package rest

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"github.com/trongtb88/merchantsvc/src/business/entity"
	"github.com/trongtb88/merchantsvc/src/common"
	"io/ioutil"
	"log"
	"net/http"
)

var decoderMember = schema.NewDecoder()


// Create Merchant Account godoc
// @Summary Create Merchant Account
// @Description Create Merchant Account
// @Tags MerchantAccount
// @Accept json
// @Produce json
// @Param data body entity.CreateMemberParam true "Body Request"
// @Success 201 {object} rest.ResponseCreateMember
// @Failure 400 {object} rest.HTTPErrResp
// @Failure 401 {object} rest.HTTPErrResp
// @Failure 500 {object} rest.HTTPErrResp
// @Router /v1/accounts/members [post]
func (rst *rest) CreateMerchantMember(w http.ResponseWriter, r *http.Request) {
	var param entity.CreateMemberParam
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

	if (param.MerchantId == 0) ||
		len(param.Email) == 0 ||
			len(param.Name) == 0 {
		rst.httpRespError(w, r, http.StatusBadRequest, ErrorMessage{
			Code:     "InvalidParameter",
			Message: "One of required field is empty",
		})
		return
	}

	acc, _,  err := rst.uc.Account.GetAccountsByParam(r.Context(), entity.GetAccountParam{
		Id:  param.MerchantId,
	})

	if len(acc) == 0 {
		rst.httpRespError(w, r, http.StatusBadRequest, ErrorMessage{
			Code:     "NotFoundMerchantId",
			Message: "Not Found Merchant Id, please choose other",
		})
		return
	}
	member, _,  err := rst.uc.Member.GetMembersByParam(r.Context(), entity.GetMemberParam{
		Email:  param.Email,
	})

	if err != nil {
		rst.httpRespError(w, r, http.StatusInternalServerError, ErrorMessage{
			Code:     "CreateMemberError",
			Message: err.Error(),
		})
		return
	}

	if len(member) > 0 {
		rst.httpRespError(w, r,http.StatusBadRequest, ErrorMessage{
			Code:     "DuplicateMemberEmail",
			Message: "Please choose other member email",
		})
		return
	}

	if !common.IsEmailValid(param.Email) {
		rst.httpRespError(w, r,http.StatusBadRequest, ErrorMessage{
			Code:     "InvalidEmail",
			Message: "Email is not valid",
		})
		return
	}

	m, err := rst.uc.Member.CreateMember(r.Context(), param)
	if err != nil {
		rst.httpRespError(w, r, http.StatusInternalServerError, ErrorMessage{
			Code:     "CreateMemberError",
			Message: err.Error(),
		})
		return
	}

	rst.httpRespSuccess(w, r, http.StatusCreated, m, nil)
}

// Get Merchant Members By Parameters
// @Summary Get Merchant Members By Parameters
// @Description Get Merchant Members By Parameters
// @Tags MerchantMembers
// @Accept json
// @Produce json
// @Param id query integer false "Merchant Member ID"
// @Param merchant_id query integer false "Merchant Account ID"
// @Param member_email query string false "Merchant Member Email"
// @Param member_name query string false "Merchant Account Name"
// @Param page query int false " "
// @Param limit query int false " "
// @Success 200 {object} rest.ResponseGetMembers
// @Failure 400 {object} rest.HTTPErrResp
// @Failure 401 {object} rest.HTTPErrResp
// @Failure 500 {object} rest.HTTPErrResp
// @Router /v1/accounts/members [get]
func (rst *rest) GetMerchantMembers(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Print("GO HERE")
		rst.httpRespError(w, r, http.StatusBadRequest, ErrorMessage{
			Code:     "GetMemberError",
			Message: err.Error(),
		})
		return
	}

	var param entity.GetMemberParam
	log.Println(r.Form)

	err = decoderMember.Decode(&param, r.Form)
	if err != nil {
		log.Print("GO HERE2")
		rst.httpRespError(w, r, http.StatusBadRequest, ErrorMessage{
			Code:     "GetMemberError",
			Message: err.Error(),
		})
		return
	}

	members, pagination, err := rst.uc.Member.GetMembersByParam(r.Context(), param)


	if err != nil {
		rst.httpRespError(w, r, http.StatusInternalServerError, ErrorMessage{
			Code:     "GetMemberError",
			Message: err.Error(),
		})
		return
	}

	rst.httpRespSuccess(w, r, http.StatusOK, members, &pagination)
}
