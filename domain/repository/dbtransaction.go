package repository

import "transaction/domain/entity"

type TransactionRepository interface {
	CreateTransaction(entity.TransactionStruct) (interface{}, error)
	UpdateTransaction(string, entity.TransactionStruct) (interface{}, error)
	GetTransactionByRef(string, entity.TransactionStruct) (interface{}, error)
	GetTransactionListByCriteria(date, transactionType, ref string, c entity.TransactionStruct) (interface{}, error)
}
