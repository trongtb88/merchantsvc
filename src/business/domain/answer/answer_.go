package answer

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/trongtb88/merchantsvc/src/business/entity"
)

const (
	GetDetailLoansByMemberID    = "GetDetailLoansByMemberID"
	GetNextPaymentSchedulerInfo = "GetNextPaymentSchedulerInfo"
	GetRemainingPrincipalAmount = "GetRemainingTotalAmount"
)

func (a answer) SubmitQuestionForAnswer(ctx context.Context, question entity.Question) (entity.Answer, error) {
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
		paymentScheduler, err := a.getPaymentSchedulerInfo(question.Parameters)
		if err != nil {
			return answer, err
		}
		answer.Content = paymentScheduler
		answer.Id = uuid.New()
		return answer, nil
	case GetRemainingPrincipalAmount:
		totalRemainingPrincipal, err := a.getTotalRemainingPrincipal(question.Parameters)
		if err != nil {
			return answer, err
		}
		answer.Content = totalRemainingPrincipal
		answer.Id = uuid.New()
		return answer, nil
	}

	return answer, nil
}

func (a answer) getDetailLoanByMemberID(parameters []entity.Parameter) ([]entity.LoanDetail, error) {
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
		LoanCode:     "SBMB-190200010",
		LoanStatus:   "Approved",
		RejectReason: "",
		RejectTime:   nil,
		ApprovedTime: &currentDate,
	}
	loan2 := entity.LoanDetail{
		LoanID:       uuid.New(),
		LoanCode:     "SBMB-19020012",
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

func (a answer) getPaymentSchedulerInfo(parameters []entity.Parameter) ([]entity.PaymentScheduler, error) {
	loan1 := entity.PaymentScheduler{
		LoanID:          uuid.New(),
		PayDate:         "2022-07-15",
		PrincipalAmount: 50000.00,
		InterestAmount:  1000.00,
		Curreny:         "SGD",
	}
	loan2 := entity.PaymentScheduler{
		LoanID:          uuid.New(),
		PayDate:         "2022-07-22",
		PrincipalAmount: 100000.00,
		InterestAmount:  2000.00,
		Curreny:         "SGD",
	}
	var result []entity.PaymentScheduler
	result = append(result, loan1)
	result = append(result, loan2)
	return result, nil
}

func (a answer) getTotalRemainingPrincipal(parameters []entity.Parameter) ([]entity.PrincipalAmount, error) {
	loan1 := entity.PrincipalAmount{
		LoanID:               uuid.New(),
		TotalRemainingAmount: "500.00",
		Curreny:              "SGD",
	}
	loan2 := entity.PrincipalAmount{
		LoanID:               uuid.New(),
		TotalRemainingAmount: "1200.00",
		Curreny:              "SGD",
	}
	var result []entity.PrincipalAmount
	result = append(result, loan1)
	result = append(result, loan2)
	return result, nil
}
