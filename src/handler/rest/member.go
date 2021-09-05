package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/trongtb88/merchantsvc/src/business/entity"
	"github.com/trongtb88/merchantsvc/src/common"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var decoderMember = schema.NewDecoder()

// Create Merchant Member
// @Summary Create Merchant Member
// @Description Create Merchant Member
// @Tags MerchantMember
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
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "InvalidJsonBodyRequest",
			Message: err.Error(),
		})
		return
	}

	if err := json.Unmarshal(body, &param); err != nil {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "InvalidJsonBodyRequest",
			Message: err.Error(),
		})
		return
	}

	if (param.MerchantId == 0) ||
		len(param.Email) == 0 ||
		len(param.Name) == 0 {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "InvalidParameter",
			Message: "One of required field is empty",
		})
		return
	}

	acc, _, err := rst.uc.Account.GetAccountsByParam(r.Context(), entity.GetAccountParam{
		Id: param.MerchantId,
	})

	if len(acc) == 0 {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "NotFoundMerchantId",
			Message: "Not Found Merchant Id, please choose other",
		})
		return
	}
	member, _, err := rst.uc.Member.GetMembersByParam(r.Context(), entity.GetMemberParam{
		Email: param.Email,
	})

	if err != nil {
		rst.httpRespError(w, r, http.StatusInternalServerError, entity.ErrorMessage{
			Code:    "CreateMemberError",
			Message: err.Error(),
		})
		return
	}

	if len(member) > 0 {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "DuplicateMemberEmail",
			Message: "Please choose other member email",
		})
		return
	}

	if !common.IsEmailValid(param.Email) {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "InvalidEmail",
			Message: "Email is not valid",
		})
		return
	}

	m, err := rst.uc.Member.CreateMember(r.Context(), param)
	if err != nil {
		rst.httpRespError(w, r, http.StatusInternalServerError, entity.ErrorMessage{
			Code:    "CreateMemberError",
			Message: err.Error(),
		})
		return
	}

	rst.httpRespSuccess(w, r, http.StatusCreated, m, nil)
}

// Get Merchant Members By Parameters
// @Summary Get Merchant Members By Parameters
// @Description Get Merchant Members By Parameters
// @Tags MerchantMember
// @Accept json
// @Produce json
// @Param id query integer false "Merchant Member ID"
// @Param merchantId query integer false "Merchant Account ID"
// @Param email query string false "Merchant Member Email"
// @Param name query string false "Merchant Account Name"
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
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "GetMemberError",
			Message: err.Error(),
		})
		return
	}

	var param entity.GetMemberParam
	log.Println(r.Form)

	err = decoderMember.Decode(&param, r.Form)
	if err != nil {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "GetMemberError",
			Message: err.Error(),
		})
		return
	}

	members, pagination, err := rst.uc.Member.GetMembersByParam(r.Context(), param)

	if err != nil {
		rst.httpRespError(w, r, http.StatusInternalServerError, entity.ErrorMessage{
			Code:    "GetMemberError",
			Message: err.Error(),
		})
		return
	}

	rst.httpRespSuccess(w, r, http.StatusOK, members, &pagination)
}

// Update Merchant Member By Id
// @Summary Update Merchant Member By Id
// @Description Update Merchant Member By Id
// @Tags MerchantMember
// @Accept json
// @Produce json
// @Param data body entity.UpdateMerchantMember true "Body Request"
// @Success 200 {object} rest.ResponseCreateMember
// @Failure 400 {object} rest.HTTPErrResp
// @Failure 401 {object} rest.HTTPErrResp
// @Failure 500 {object} rest.HTTPErrResp
// @Router /v1/accounts/members [put]
func (rst *rest) UpdateMerchantMember(w http.ResponseWriter, r *http.Request) {
	var (
		param entity.UpdateMerchantMember
		m     entity.MerchantMember
	)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "InvalidJsonBodyRequest",
			Message: err.Error(),
		})
		return
	}

	if err := json.Unmarshal(body, &param); err != nil {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "InvalidJsonBodyRequest",
			Message: err.Error(),
		})
		return
	}

	if param.Id <= 0 {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "InvalidMemberId",
			Message: "Invalid Member Id",
		})
		return
	}

	memberById, _, err := rst.uc.Member.GetMembersByParam(r.Context(), entity.GetMemberParam{
		Id: param.Id,
	})

	if len(memberById) == 0 {
		rst.httpRespError(w, r, http.StatusNotFound, entity.ErrorMessage{
			Code:    "NotFoundMemberId",
			Message: "We not found member Id",
		})
		return
	}
	m = entity.MerchantMember{
		Id:           param.Id,
		MerchantId:   memberById[0].MerchantAccount.Id,
		MemberEmail:  memberById[0].MemberEmail,
		MemberName:   memberById[0].MemberName,
		MemberStatus: common.FillStatus(memberById[0].MemberStatus),
		UpdatedAt:    time.Time{},
		CreatedAt:    memberById[0].CreatedAt,
	}

	if len(param.Email) > 0 {
		if !common.IsEmailValid(param.Email) {
			rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
				Code:    "InvalidEmailFormat",
				Message: "Invalid Email Format",
			})
			return
		}

		memberByEmail, _, err := rst.uc.Member.GetMembersByParam(r.Context(), entity.GetMemberParam{
			Email: strings.TrimSpace(param.Email),
		})

		if err != nil {
			rst.httpRespError(w, r, http.StatusNotFound, entity.ErrorMessage{
				Code:    "ErrorGetMemberByEmail",
				Message: err.Error(),
			})
			return
		}

		if (len(memberByEmail) > 0) && memberByEmail[0].Id != param.Id {
			rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
				Code:    "InvalidUpdateEmail",
				Message: "Email already used by other member",
			})
			return
		}
		m.MemberEmail = param.Email
	}

	if len(param.Name) > 0 {
		m.MemberName = param.Name
	}

	members, err := rst.uc.Member.UpdateMember(r.Context(), m)
	members.MemberStatusInStr = common.FillStatusInStr(m.MemberStatus)

	if err != nil {
		rst.httpRespError(w, r, http.StatusInternalServerError, entity.ErrorMessage{
			Code:    "UpdateMemberError",
			Message: err.Error(),
		})
		return
	}

	rst.httpRespSuccess(w, r, http.StatusOK, members, nil)
}

// Delete Merchant Member By Id
// @Summary Delete Merchant Member By Id
// @Description Delete Merchant Member By Id
// @Tags MerchantMember
// @Accept json
// @Produce json
// @Param member_id path integer true "Merchant Member ID"
// @Success 200 {object} rest.ResponseCreateAccount
// @Failure 400 {object} rest.HTTPErrResp
// @Failure 401 {object} rest.HTTPErrResp
// @Failure 500 {object} rest.HTTPErrResp
// @Router /v1/accounts/members/{member_id} [delete]
func (rst *rest) DeleteMerchantMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var memberId int64
	var err error
	idParam, _ := vars["member_id"]

	if len(idParam) == 0 {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "InvalidMerchantMemberId",
			Message: "Invalid Merchant Member Id",
		})
		return
	}

	if memberId, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "InvalidMerchantMemberId",
			Message: "Member Id must be number",
		})
		return
	}

	membersById, _, err := rst.uc.Member.GetMembersByParam(r.Context(), entity.GetMemberParam{
		Id: memberId,
	})

	if len(membersById) == 0 {
		rst.httpRespError(w, r, http.StatusNotFound, entity.ErrorMessage{
			Code:    "NotFoundMember",
			Message: "Not Found Merchant Member",
		})
		return
	}

	m := entity.MerchantMember{
		Id:                memberId,
		MerchantId:        membersById[0].MerchantAccount.Id,
		MemberName:        membersById[0].MemberName,
		MemberEmail:       membersById[0].MemberEmail,
		MemberStatus:      common.StatusInActive,
		MemberStatusInStr: common.FillStatusInStr(common.StatusInActive),
		CreatedAt:         membersById[0].CreatedAt,
		UpdatedAt:         time.Time{},
	}

	_, err = rst.uc.Member.UpdateMember(r.Context(), m)
	if err != nil {
		rst.httpRespError(w, r, http.StatusInternalServerError, entity.ErrorMessage{
			Code:    "DeleteAccountsError",
			Message: err.Error(),
		})
		return
	}
	rst.httpRespSuccess(w, r, http.StatusOK, m, nil)
}
