package http

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/v1gn35h7/transaction-service/internal/service"
)

type coreEndpoints struct {
	createAccountEndpoint     endpoint.Endpoint
	getAccountEndpoint        endpoint.Endpoint
	createTransactionEndpoint endpoint.Endpoint
}

func MakeEndpoints(srvc service.Service) coreEndpoints {
	return coreEndpoints{
		createAccountEndpoint:     MakeCreateAccountEndpoint(srvc),
		getAccountEndpoint:        MakeGetAccountEndpoint(srvc),
		createTransactionEndpoint: MakeCreateTransactionEndpoint(srvc),
	}
}
