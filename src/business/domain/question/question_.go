package question

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/trongtb88/merchantsvc/src/business/entity"
)

func (q question) GetPredefineQuestionsForAuthenticate(ctx context.Context) ([]entity.QuestionForAuthentication, error) {
	data := `[{
		"Id":"94ec5e84-f02b-11ec-8ea0-0242ac120002",
		"Content": "Please input your identification number",
		"QuestionType" : "textbox",
		"IsRequired": true,
		"SupportOTP": false,
		"Field" : "identification_number"
	},
	{
		"Id":"94ec5e84-f02b-11ec-8ea0-0242ac120003",
		"Content": "Please input your full name",
		"QuestionType" : "textbox",
		"IsRequired": true,
		"SupportOTP": false,
		"Field" : "name"
	},
{
		"Id":"94ec5e84-f02b-11ec-8ea0-0242ac120004",
		"Content": "Please input your email",
		"QuestionType" : "textbox",
		"IsRequired": true,
		"SupportOTP": true,
		"Field" : "email"
	},
{
		"Id":"94ec5e84-f02b-11ec-8ea0-0242ac120004",
		"Content": "",
		"QuestionType" : "textbox",
		"IsRequired": true,
		"SupportOTP": true,
		"Field" : "phoneNumber"
	}
]`

	//Id           uuid.UUID
	//Content      string
	//QuestionType string
	//IsRequired   bool
	//SupportOTP   bool

	var m []entity.QuestionForAuthentication
	err := json.Unmarshal([]byte(data), &m)
	return m, err
}

func (q question) GetPredefineQuestionsForBusiness(ctx context.Context, topic string) ([]entity.Question, error) {
	switch topic {
	case "loan":
		return q.GetPredefineQuestionsForLoan(ctx)
	case "bank_account":
		return q.GetPredefineQuestionsForBankAccount(ctx)
	}
	return []entity.Question{}, nil
}

func (q question) GetPredefineQuestionsForLoan(ctx context.Context) ([]entity.Question, error) {
	var questions []entity.Question
	questions = append(questions, entity.Question{
		Id:           uuid.New(),
		Content:      "Give me loan details which I am having ?",
		QuestionType: "quires",
		Topic:        "loan",
		Function:     "GetDetailLoansByMemberID",
		Parameters:   q.getParametersForMemberLoan(ctx),
	})
	questions = append(questions, entity.Question{
		Id:           uuid.New(),
		Content:      "Give me next payment scheduler information ?",
		QuestionType: "quires",
		Topic:        "loan",
		Function:     "GetNextPaymentSchedulerInfo",
		Parameters:   q.getParametersForGetDetailLoan(ctx),
	})
	questions = append(questions, entity.Question{
		Id:           uuid.New(),
		Content:      "Give me the total remaining of principal of loan ?",
		QuestionType: "quires",
		Topic:        "loan",
		Function:     "GetTotalRemainingPrincipal",
		Parameters:   q.getParametersForGetDetailLoan(ctx),
	})
	return questions, nil
}

func (q question) getParametersForMemberLoan(ctx context.Context) []entity.Parameter {
	var parameters []entity.Parameter
	parameters = append(parameters, entity.Parameter{
		Name:       "member_id",
		IsRequired: true,
	})
	return parameters
}

func (q question) getParametersForGetDetailLoan(ctx context.Context) []entity.Parameter {
	var parameters []entity.Parameter
	parameters = append(parameters, entity.Parameter{
		Name:       "loan_id",
		IsRequired: true,
	})
	return parameters
}

// Bank account
func (q question) GetPredefineQuestionsForBankAccount(ctx context.Context) ([]entity.Question, error) {
	var questions []entity.Question
	questions = append(questions, entity.Question{
		Id:           uuid.New(),
		Content:      "I want to report fake bank account",
		QuestionType: "submit",
		Topic:        "bank_account",
		Function:     "ReportFakeAccount",
		Parameters:   q.getParametersForFakeBankAccount(ctx),
	})
	questions = append(questions, entity.Question{
		Id:           uuid.New(),
		Content:      "I want to see my pending withdrawals transaction for bank account",
		QuestionType: "queries",
		Topic:        "bank_account",
		Function:     "GetPendingTransactionForBankAccount",
		Parameters:   q.getParametersForTransactionBankAccount(ctx),
	})
	return questions, nil
}

func (q question) getParametersForFakeBankAccount(ctx context.Context) []entity.Parameter {
	var parameters []entity.Parameter
	parameters = append(parameters, entity.Parameter{
		Name:       "bank_account",
		IsRequired: true,
	})
	return parameters
}

func (q question) getParametersForTransactionBankAccount(ctx context.Context) []entity.Parameter {
	var parameters []entity.Parameter
	parameters = append(parameters, entity.Parameter{
		Name:       "bank_account",
		IsRequired: true,
	})
	return parameters
}
