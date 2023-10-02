package application

import (
	"transaction/domain/entity"
	"transaction/domain/repository"
)

type TransactionApp struct {
	theTransaction repository.TransactionRepository
}

var _TransactionApplication = &TransactionApp{}

type TransactionApplication interface {
	CreateTransaction(entity.TransactionStruct) (interface{}, error)
	UpdateTransaction(string, entity.TransactionStruct) (interface{}, error)
	GetTransactionByRef(string, entity.TransactionStruct) (interface{}, error)
	GetTransactionListByCriteria(date, transactionType, ref string, c entity.TransactionStruct) (interface{}, error)
}

func (u *TransactionApp) CreateTransaction(c entity.TransactionStruct) (interface{}, error) {
	return u.theTransaction.CreateTransaction(c)
}

func (u *TransactionApp) UpdateTransaction (ref string, c entity.TransactionStruct) (interface{}, error) {
	return u.theTransaction.UpdateTransaction(ref, c)
}

func (u *TransactionApp) GetTransactionByRef (ref string, c entity.TransactionStruct) (interface{}, error) {
	return u.theTransaction.GetTransactionByRef(ref, c)
}

func (u *TransactionApp) GetTransactionListByCriteria(d, t, ref string, c entity.TransactionStruct) (interface{}, error) {
	return u.theTransaction.GetTransactionListByCriteria(d, t, ref, c)
}