package entity

import "errors"

const (
	Pending    statusStruct = "pending"
	Successful statusStruct = "successful"
	Failed     statusStruct = "failed"
)
const (
	InvestmentPayment transactionTypeStruct = "investmentpayment"
	RevenuePayment transactionTypeStruct = "revenuepayment"
)

func (c statusStruct) StatusString() (string, error) {
	var response string
	statustype := [...]string{"PENDING", "SUCCESSFUL", "FAILED"}

	x := string(c)
	for _, v := range statustype {
		if v == x {

			response = x

			return response, nil
		}
	}

	response = ""

	return response, errors.New("incorrect status code")
}

func (c transactionTypeStruct) TransactionString() (string, error) {
	var response string
	transactiontype := [...]string{"INVESTMENT_PAYMENT", "REVENUE_PAYMENT"}

	x := string(c)
	for _, v := range transactiontype {
		if v == x {

			response = x

			return response, nil
		}
	}

	response = ""

	return response, errors.New("incorrect transaction type")
}
