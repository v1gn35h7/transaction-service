package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// **************************************************************************************************************************
// Request utilities
// **************************************************************************************************************************
func decodeCreateTransactionRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var req CreateTransactionRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		return nil, err
	}

	if req.Amount == 0 {
		return nil, errors.New("amount is required")
	}

	if req.AccountID <= 0 {
		return nil, errors.New("account_id must be positive")
	}

	if req.AccountID > 9999999999 {
		return nil, errors.New("account_id is too large")
	}

	if req.OperationTypeID <= 0 {
		return nil, errors.New("operation_type_id is required")
	}

	if req.OperationTypeID > 4 {
		return nil, errors.New("invalid operation_type_id")
	}

	return req, nil
}

func encodeCreateTransactionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeListTransactionsRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var req ListTransactionsRequest

	vars := mux.Vars(request)
	id, ok := vars["accountID"]
	if !ok {
		return nil, errors.New("Bad route")
	}
	accountID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	req.AccountID = int64(accountID)

	return req, nil
}

// ********************************************************************************************************************
// Scaffloding endpoints to transport
// *******************************************************************************************************************
func makeCreateTransactionTransport(endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(
		endpoint,
		decodeCreateTransactionRequest,
		encodeCreateTransactionResponse,
		httpSrvOptions...,
	)
}

func makeListTransactionsTransport(endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(
		endpoint,
		decodeListTransactionsRequest,
		encodeCreateTransactionResponse,
		httpSrvOptions...,
	)
}
