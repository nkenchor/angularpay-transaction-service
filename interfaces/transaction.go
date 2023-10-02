package interfaces

import (
	"encoding/json"
	"errors"
	"net/http"
	"transaction/application"
	"transaction/domain/entity"
	"transaction/sharedinfrastructure/helper"

	"github.com/go-chi/chi"
)

type TransactionInterface struct {
	us application.TransactionApplication
}

func NewTransaction(us application.TransactionApplication) TransactionInterface {
	return TransactionInterface{
		us: us,
	}
}
func (s *TransactionInterface) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction entity.TransactionStruct
	err := decodeJSONBody(w, r, &transaction)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			errorResponse := helper.ReturnedError("invalid struct", err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorResponse)
		} else {
			errorResponse := helper.ReturnedError("invalid struct", err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errorResponse)
		}
		return
	}
	postTransaction, err := s.us.CreateTransaction(transaction)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(postTransaction)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(postTransaction)
}

func (s *TransactionInterface) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction entity.TransactionStruct
	ref := chi.URLParam(r, "reference")
	err := decodeJSONBody(w, r, &transaction)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			errorResponse := helper.ReturnedError("invalid struct", err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorResponse)
		} else {
			errorResponse := helper.ReturnedError("invalid struct", err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errorResponse)
		}
		return
	}
	putTransaction, err := s.us.UpdateTransaction(ref, transaction)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(putTransaction)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(putTransaction)
}

func (s *TransactionInterface) GetTransactionByRef(w http.ResponseWriter, r *http.Request) {
	var transaction entity.TransactionStruct
	ref := chi.URLParam(r, "reference")

	aTransaction, err := s.us.GetTransactionByRef(ref, transaction)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(aTransaction)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(aTransaction)
}

func (s *TransactionInterface) GetTransactionListByCriteria(w http.ResponseWriter, r *http.Request) {
	var transaction entity.TransactionStruct
	date := chi.URLParam(r, "date")
	transactionType := chi.URLParam(r, "type")
	ref := chi.URLParam(r, "reference")

	aTransaction, err := s.us.GetTransactionListByCriteria(date, transactionType, ref, transaction)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(aTransaction)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(aTransaction)
}
