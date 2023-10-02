package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"transaction/interfaces"
	"transaction/routes"
	"transaction/sharedinfrastructure/helper"
	"transaction/sharedinfrastructure/persistence"
)

func init() {
	fileName := "log/transaction-service.log"
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(f)
}

func main() {
	address, port, mode, dbhost, dbname := helper.LoadConfig()

	transactions, err := persistence.ConnectionDB(dbhost, dbname)
	if err != nil {
		fmt.Println(err)
	}

	transactionEndPoint := interfaces.NewTransaction(transactions.Transaction)

	fmt.Println("App running on " + address + ":" + port)
	if mode == "dev" {
		r := routes.SetupRouter(port, address, transactionEndPoint)
		http.ListenAndServe(":"+port, r)
	}

}
