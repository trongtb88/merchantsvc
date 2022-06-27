package rest

import (
	"encoding/json"
	"github.com/trongtb88/merchantsvc/src/business/entity"
	"io/ioutil"
	"net/http"
)

// GetPredefineQuestionForAuthentication Get Predefine questions for authentication user godoc
// @Summary Get Predefine questions for authentication user
// @Description Get Predefine questions for authentication user
// @Tags CustomerSupport
// @Accept json
// @Produce json
// @Success 200 {object} rest.ResponseGetPredefineQuestionForAuthentication
// @Failure 400 {object} rest.HTTPErrResp
// @Failure 401 {object} rest.HTTPErrResp
// @Failure 500 {object} rest.HTTPErrResp
// @Router /v1/support/pre-define-questions-for-authentication [get]
func (rst *rest) GetPredefineQuestionForAuthentication(w http.ResponseWriter, r *http.Request) {

	questions, err := rst.uc.Question.GetPredefineQuestionsForAuthenticate(r.Context())
	if err != nil {
		rst.httpRespError(w, r, http.StatusInternalServerError, entity.ErrorMessage{
			Code:    "GetPredefineQuestionForAuthentication",
			Message: err.Error(),
		})
		return
	}

	rst.httpRespSuccess(w, r, http.StatusOK, questions, nil)
}

// GetPredefineQuestionForBusiness Get Predefine questions for user godoc
// @Summary Get Predefine questions for user
// @Description Get Predefine questions for user
// @Tags CustomerSupport
// @Accept json
// @Produce json
// @Success 200 {object} rest.ResponseGetPredefineQuestionForBusiness
// @Failure 400 {object} rest.HTTPErrResp
// @Failure 401 {object} rest.HTTPErrResp
// @Failure 500 {object} rest.HTTPErrResp
// @Router /v1/support/pre-define-questions-for-business [get]
func (rst *rest) GetPredefineQuestionForBusiness(w http.ResponseWriter, r *http.Request) {

	questions, err := rst.uc.Question.GetPredefineQuestionsForBusiness(r.Context(), "loan")
	if err != nil {
		rst.httpRespError(w, r, http.StatusInternalServerError, entity.ErrorMessage{
			Code:    "GetPredefineQuestionsForBusiness",
			Message: err.Error(),
		})
		return
	}

	rst.httpRespSuccess(w, r, http.StatusOK, questions, nil)
}

// SubmitQuestionForAnswer POSTSubmitQuestionForAnswer Submit Question for Anwser godoc
// @Summary Submit Question for Anwser
// @Description Submit Question for Anwser
// @Tags CustomerSupport
// @Accept json
// @Produce json
// @Param data body entity.Question true "Body Request"
// @Success 201 {object} rest.ResponseAnswerForQuestion
// @Failure 400 {object} rest.HTTPErrResp
// @Failure 401 {object} rest.HTTPErrResp
// @Failure 500 {object} rest.HTTPErrResp
// @Router /v1/support/submit-questions-for-answer [post]
func (rst *rest) SubmitQuestionForAnswer(w http.ResponseWriter, r *http.Request) {
	var question entity.Question
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "InvalidJsonBodyRequest",
			Message: err.Error(),
		})
		return
	}

	if err := json.Unmarshal(body, &question); err != nil {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "InvalidJsonBodyRequest",
			Message: err.Error(),
		})
		return
	}

	answer, err := rst.uc.Question.SubmitQuestionForAnswer(r.Context(), question)
	if err != nil {
		rst.httpRespError(w, r, http.StatusInternalServerError, entity.ErrorMessage{
			Code:    "GetPredefineQuestionsForBusiness",
			Message: err.Error(),
		})
		return
	}

	rst.httpRespSuccess(w, r, http.StatusOK, answer, nil)
}
