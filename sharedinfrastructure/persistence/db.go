package persistence

import (
	"context"
	"fmt"
	"log"
	"transaction/domain/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repositories struct {
	Transaction repository.TransactionRepository
	dbService   *mongo.Collection
}

func ConnectionDB(dbhost, dbname string) (Repositories, error) {

	clientOptions := options.Client().ApplyURI(dbhost) // Connect to MongoDB

	db, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err)
		return Repositories{}, err
	}

	// Check the connection
	err = db.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
		return Repositories{}, err
	}
	fmt.Println("Connected to MongoDB!")

	conn := db.Database(dbname)

	transactionCollection := conn.Collection("transaction")

	return Repositories{
		Transaction: NewTransactionInfra(transactionCollection),
		dbService:   transactionCollection,
	}, nil
}
