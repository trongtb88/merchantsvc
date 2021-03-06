package answer

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/trongtb88/merchantsvc/src/business/entity"
)

const (
	GetDetailLoansByMemberID            = "GetDetailLoansByMemberID"
	GetNextPaymentSchedulerInfo         = "GetNextPaymentSchedulerInfo"
	GetTotalRemainingPrincipal          = "GetTotalRemainingPrincipal"
	ReportFakeAccount                   = "ReportFakeAccount"
	GetPendingTransactionForBankAccount = "GetPendingTransactionForBankAccount"
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
		paymentScheduler, err := a.getPaymentSchedulerInfo(question.Parameters)
		if err != nil {
			return answer, err
		}
		answer.Content = paymentScheduler
		answer.Id = uuid.New()
		return answer, nil
	case GetTotalRemainingPrincipal:
		totalRemainingPrincipal, err := a.getTotalRemainingPrincipal(question.Parameters)
		if err != nil {
			return answer, err
		}
		answer.Content = totalRemainingPrincipal
		answer.Id = uuid.New()
		return answer, nil
	case ReportFakeAccount:
		answer.Content = "Please fill bank account detail report base on template \n Input bank account number : \n Input bank account name : \n Input bank name : "
		answer.Id = uuid.New()
		return answer, nil
	case GetPendingTransactionForBankAccount:
		var pendingTransactions []entity.PendingTransaction
		pendingTransaction1 := entity.PendingTransaction{
			TransactionId: uuid.New(),
			WithdrawDate:  "2022-06-15 14:45:12",
			Status:        "pending",
			BankAccount:   "10000100005678",
			BankName:      "OCB",
			PendingReason: "Not enough money to draw",
		}

		pendingTransaction2 := entity.PendingTransaction{
			TransactionId: uuid.New(),
			WithdrawDate:  "2022-06-20 14:45:12",
			Status:        "pending",
			BankAccount:   "10000100005678",
			BankName:      "OCB",
			PendingReason: "Can not withdrawl because you already paid all for last payment schedule",
		}

		pendingTransactions = append(pendingTransactions, pendingTransaction1)
		pendingTransactions = append(pendingTransactions, pendingTransaction2)

		answer.Content = pendingTransactions
		answer.Id = uuid.New()
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
		LoanCode:     "SBMB-190200010",
		LoanStatus:   "Approved",
		RejectReason: "",
		RejectTime:   nil,
		ApprovedTime: &currentDate,
	}
	loan2 := entity.LoanDetail{
		LoanID:       uuid.New(),
		LoanCode:     "SBMB-190200012",
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

func (a answerUc) getPaymentSchedulerInfo(parameters []entity.Parameter) ([]entity.PaymentScheduler, error) {
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

func (a answerUc) getTotalRemainingPrincipal(parameters []entity.Parameter) ([]entity.PrincipalAmount, error) {
	loan1 := entity.PrincipalAmount{
		LoanID:               uuid.New(),
		TotalRemainingAmount: "1000.00",
		Curreny:              "SGD",
	}
	loan2 := entity.PrincipalAmount{
		LoanID:               uuid.New(),
		TotalRemainingAmount: "2000.00",
		Curreny:              "SGD",
	}
	var result []entity.PrincipalAmount
	result = append(result, loan1)
	result = append(result, loan2)
	return result, nil
}
