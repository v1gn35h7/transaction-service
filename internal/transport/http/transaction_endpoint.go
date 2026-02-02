package http

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/v1gn35h7/transaction-service/internal/service"
)

// Create Transaction Endpoints spec
type CreateTransactionRequest struct {
	AccountID       int64   `json:"account_id"`
	OperationTypeID int16   `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

type CreateTransactionResponse struct {
	Success bool `json:"success"`
}

type ListTransactionsRequest struct {
	AccountID int64 `json:"account_id"`
}

type ListTransactionsResponse struct {
	Transactions []interface{} `json:"transactions"`
}

func MakeCreateTransactionEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(CreateTransactionRequest)
		err := srvc.CreateTransaction(req.AccountID, req.OperationTypeID, req.Amount)
		if err != nil {
			return CreateTransactionResponse{Success: false}, err
		}

		return CreateTransactionResponse{Success: true}, nil
	}
}

func MakeListTransactionsEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListTransactionsRequest)
		transactions, err := srvc.ListTransactions(req.AccountID)
		if err != nil {
			return ListTransactionsResponse{Transactions: nil}, err
		}
		var txs []interface{}
		for _, tx := range transactions {
			txs = append(txs, tx)
		}
		return ListTransactionsResponse{Transactions: txs}, nil
	}
}
