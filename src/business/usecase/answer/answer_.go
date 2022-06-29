package answer

import (
	"context"
	"github.com/google/uuid"
	"github.com/trongtb88/merchantsvc/src/business/entity"
	"time"
)

const (
	GetDetailLoansByMemberID    = "GetDetailLoansByMemberID"
	GetNextPaymentSchedulerInfo = "GetNextPaymentSchedulerInfo"
	GetTotalRemainingPrincipal  = "GetTotalRemainingPrincipal"
)

func (a answerUc) SubmitQuestionForAnswer(ctx context.Context, question entity.Question) (entity.Answer, error) {
	var answer entity.Answer
	switch question.Function {
	case GetDetailLoansByMemberID:
		loans, err := a.getDetailLoanByMemberID(question.Parameters)
		if err != nil {
			return answer, err
		}
		answer.Content = loans
		answer.Id = uuid.New()
		return answer, nil
	case GetNextPaymentSchedulerInfo:
		return answer, nil
	}
	return answer, nil

}

func (a answerUc) getDetailLoanByMemberID(parameters []entity.Parameter) ([]entity.LoanDetail, error) {
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
