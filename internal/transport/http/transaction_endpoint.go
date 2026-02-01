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
