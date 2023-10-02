package persistence

import (
	"context"
	"errors"
	"log"
	"transaction/domain/entity"
	"transaction/sharedinfrastructure/helper"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionInfra struct {
	collection *mongo.Collection
}

func NewTransactionInfra(collection *mongo.Collection) *TransactionInfra {
	return &TransactionInfra{collection}
}

func (r *TransactionInfra) CreateTransaction(c entity.TransactionStruct) (interface{}, error) {
	errorResponse, validateErr := helper.ValidateStruct(c)
	if validateErr != nil {
		return errorResponse, errors.New("")
	}

	val, err := Create(c)
	if err != nil {
		errorResponse := helper.ReturnedError("validation error", "data validation error")
		return errorResponse, errors.New("")
	}

	_, err = r.collection.InsertOne(context.TODO(), val)
	if err != nil {
		errorResponse := helper.ReturnedError("validation error", "error inserting record into the database")
		return errorResponse, errors.New("")
	}

	ref := c.Reference

	return ref, nil
}

func (r *TransactionInfra) UpdateTransaction(ref string, c entity.TransactionStruct) (interface{}, error) {
	errorResponse, validateErr := helper.ValidateStruct(c)
	if validateErr != nil {
		return errorResponse, errors.New("")
	}

	statusCode, err := c.Status.StatusString()
	if err != nil {
		errorResponse := helper.ReturnedError("validation error", "invalid status")
		return errorResponse, errors.New("")
	}
	log.Println(statusCode)

	filter := bson.M{"reference": ref}
	update := bson.M{"$set": bson.M{
		"status": statusCode,
	}}
	result, err := r.collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		errorResponse := helper.ReturnedError("validation error", "the reference search filter does not match any record in the database")
		return errorResponse, errors.New("")
	}
	log.Println("modified count: ", result.ModifiedCount)
	var rm entity.TransactionResponse
	rm.Reference = c.Reference
	rm.Status = statusCode
	return rm, nil
}

func (r *TransactionInfra) GetTransactionByRef(ref string, c entity.TransactionStruct) (interface{}, error) {
	filter := bson.M{"reference": ref}
	err := r.collection.FindOne(context.TODO(), filter).Decode(&c)
	if err != nil {
		errorResponse := helper.ReturnedError("validation error", "the reference search filter does not match any record in the database")
		return errorResponse, errors.New("")
	}
	return c, nil
}

func (r *TransactionInfra) GetTransactionListByCriteria(d, t, ref string, c entity.TransactionStruct) (interface{}, error) {
	var C []entity.TransactionStruct
	filter := bson.M{"transaction_date": d, "transaction_type": t, "reference": ref}
	cursor, err := r.collection.Find(context.TODO(), filter)
	if err != nil {
		defer cursor.Close(context.Background())
		errorResponse := helper.ReturnedError("validation error", "unable to find all the document in the database")
		return errorResponse, errors.New("")
	}
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&c)
		if err != nil {
			errorResponse := helper.ReturnedError("validation error", "unable to decode the documents from the database")
			return errorResponse, errors.New("")
		}
		C = append(C, c)
	}
	return C, nil
}
