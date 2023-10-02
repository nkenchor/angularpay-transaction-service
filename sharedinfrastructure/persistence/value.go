package persistence

import (
	"log"
	"time"
	"transaction/domain/entity"

	"github.com/google/uuid"
)

func Create(c entity.TransactionStruct) (entity.TransactionStruct, error) {
	transactionType, err := c.TransactionType.TransactionString()
	if err != nil {
		log.Println("400", "data validation failed", "entity.TransactionStruct", err)
		return entity.TransactionStruct{}, err
	}
	log.Println(transactionType)

	statusCode, err := c.Status.StatusString()
	if err != nil {
		log.Println("400", "data validation failed", "entity.TransactionStruct", err)
		return entity.TransactionStruct{}, err
	}
	log.Println(statusCode)
	c.TransactionDateTime = time.Now()
	c.Reference = uuid.New().String()
	c.ServiceReference = uuid.New().String()
	c.ServiceRequestInvestorReference = uuid.New().String()
	c.ServiceRequestReference = uuid.New().String()
	c.InvesteeAccount.AccountReference = uuid.New().String()
	c.InvesteeAccount.UserReference = uuid.New().String()
	c.InvesteeAccount.BankReference = uuid.New().String()
	c.InvestorAcount.AccountReference = uuid.New().String()
	c.InvestorAcount.BankReference = uuid.New().String()
	c.InvestorAcount.UserReference = uuid.New().String()

	return c, nil
}
