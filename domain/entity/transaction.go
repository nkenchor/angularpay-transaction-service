package entity

import "time"

type amountStruct struct {
	Currency string `bson:"currency" json:"currency"`
	Value    string `bson:"value" json:"value"`
}
type statusStruct string
type transactionTypeStruct string

type investorAccountStruct struct {
	UserReference    string `bson:"user_reference" json:"user_reference"`
	AccountReference string `bson:"account_reference" json:"account_reference"`
	AccountName      string `bson:"account_name" json:"account_name"`
	AccountNumber    string `bson:"account_number" json:"account_number"`
	Iban             string `bson:"iban" json:"iban"`
	BankReference    string `bson:"bank_reference" json:"bank_reference"`
}

type investeeAccountStruct struct {
	UserReference    string `bson:"user_reference" json:"user_reference"`
	AccountReference string `bson:"account_reference" json:"account_reference"`
	AccountName      string `bson:"account_name" json:"account_name"`
	AccountNumber    string `bson:"account_number" json:"account_number"`
	Iban             string `bson:"iban" json:"iban"`
	BankReference    string `bson:"bank_reference" json:"bank_reference"`
}

type TransactionResponse struct {
	Reference string `bson:"reference" json:"reference"`
	Status    string `bson:"status" json:"status"`
}

type TransactionStruct struct {
	Reference                       string                `bson:"reference" json:"reference"`
	ServiceReference                string                `bson:"service-reference" json:"service_reference"`
	ServiceRequestReference         string                `bson:"service_request_reference" json:"service_request_reference"`
	ServiceRequestInvestorReference string                `bson:"service_request_reference" json:"service_request_investor_reference"`
	TransactionDateTime             time.Time                `bson:"transaction_datetime"`
	TransactionType                 transactionTypeStruct `bson:"transaction_type" json:"transaction_type"`
	Amount                          amountStruct          `bson:"amount" json:"amount"`
	Remarks                         string                `bson:"remarks" json:"remarks"`
	InvestorAcount                  investorAccountStruct `bson:"investor_account" json:"investor_account"`
	InvesteeAccount                 investeeAccountStruct `bson:"investee_account" json:"investee_account"`
	Status                          statusStruct          `bson:"status" json:"status"`
}
