package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// **************************************************************************************************************************
// Request utilities
// **************************************************************************************************************************
func decodeCreateTransactionRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var req CreateTransactionRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func encodeCreateTransactionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
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
