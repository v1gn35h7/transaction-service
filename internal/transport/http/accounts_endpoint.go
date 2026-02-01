package http

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/v1gn35h7/transaction-service/internal/service"
)

// Create Account Endpoint Spec
type CreateAccountRequest struct {
	DocumentNumber int64 `json:"document_number"`
}

type CreateAccountResponse struct {
	AccountID int64 `json:"account_id"`
}

func MakeCreateAccountEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(CreateAccountRequest)
		accountID, err := srvc.CreateAccount(req.DocumentNumber)

		if err != nil {
			return CreateAccountResponse{AccountID: 0}, err
		}

		return CreateAccountResponse{AccountID: accountID}, nil
	}
}

// Get Account Endpoint Spec
type GetAccountRequest struct {
	AccountID int64 `json:"account_id"`
}

type GetAccountResponse struct {
	AccountID      int64 `json:"account_id"`
	DocumentNumber int64 `json:"document_number"`
}

func MakeGetAccountEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAccountRequest)
		account, err := srvc.GetAccount(req.AccountID)
		if err != nil {
			return GetAccountResponse{}, err
		}

		return GetAccountResponse{
			AccountID:      account.AccountID,
			DocumentNumber: account.DocumentNumber,
		}, nil
	}
}
