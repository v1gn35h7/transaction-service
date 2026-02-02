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
func decodeCreateAccountRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var req CreateAccountRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		return nil, err
	}

	if req.DocumentNumber == 0 {
		return nil, errors.New("document_number is required")
	}

	if req.DocumentNumber < 0 {
		return nil, errors.New("document_number must be positive")
	}

	if req.DocumentNumber > 99999999999 {
		return nil, errors.New("document_number is too large")
	}

	if req.DocumentNumber < 1000000000 {
		return nil, errors.New("document_number is too small")
	}

	return req, nil
}

func encodeCreateAccountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeGetAccountRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var req GetAccountRequest

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

func encodeGetAccountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// ********************************************************************************************************************
// Scaffloding endpoints to transport
// *******************************************************************************************************************
func makeCreateAccountTransport(endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(
		endpoint,
		decodeCreateAccountRequest,
		encodeCreateAccountResponse,
		httpSrvOptions...,
	)
}

func makeGetAccountTransport(endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(
		endpoint,
		decodeGetAccountRequest,
		encodeGetAccountResponse,
		httpSrvOptions...,
	)
}
