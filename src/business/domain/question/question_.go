package question

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/trongtb88/merchantsvc/src/business/entity"
	"time"
)

const (
	GetDetailLoansByMemberID    = "GetDetailLoansByMemberID"
	GetNextPaymentSchedulerDate = "GetNextPaymentSchedulerDate"
	GetTotalPaymentScheduler    = "GetTotalPaymentScheduler"
	GetRepaymentScheduler       = "GetRepaymentScheduler"
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
		Content:      "Give me next payment scheduler date  ?",
		QuestionType: "quires",
		Topic:        "loan",
		Function:     "GetNextPaymentSchedulerDate",
		Parameters:   q.getParametersForGetDetailLoan(ctx),
	})
	questions = append(questions, entity.Question{
		Id:           uuid.New(),
		Content:      "Give me amount money of next payment I have to pay ?",
		QuestionType: "quires",
		Topic:        "loan",
		Function:     "GetTotalPaymentScheduler",
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
	questions = append(questions, entity.Question{
		Id:           uuid.New(),
		Content:      "Give me the next payment scheduler of loan ?",
		QuestionType: "quires",
		Topic:        "loan",
		Function:     "GetRepaymentScheduler",
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

func (q question) SubmitQuestionForAnswer(ctx context.Context, question entity.Question) (entity.Answer, error) {
	var answer entity.Answer
	switch question.Function {
	case GetDetailLoansByMemberID:
		loans, err := q.getDetailLoanByMemberID(question.Parameters)
		if err != nil {
			return answer, err
		}
		answer.Content = loans
		answer.Id = uuid.New()
		return answer, nil
	}
	return answer, nil
}

func (q question) getDetailLoanByMemberID(parameters []entity.Parameter) ([]entity.LoanDetail, error) {
	//var memberID uuid.UUID
	//for _, parameter := range parameters {
	//	if parameter.Name == "member_id" {
	//		memberID, _ = uuid.Parse(parameter.Value)
	//		break
	//	}
	//}
	currentDate := time.Now().UTC()
	loan1 := entity.LoanDetail{
		LoanID:       uuid.New(),
		LoanCode:     "SB2000",
		LoanStatus:   "Approved",
		RejectReason: "",
		RejectTime:   nil,
		ApprovedTime: &currentDate,
	}
	loan2 := entity.LoanDetail{
		LoanID:       uuid.New(),
		LoanCode:     "SB2001",
		LoanStatus:   "Rejected",
		RejectReason: "Not qualify our conditions - missing revenue annual",
		RejectTime:   &currentDate,
		ApprovedTime: nil,
	}
	var result []entity.LoanDetail
	result = append(result, loan1)
	result = append(result, loan2)
	return result, nil
}
