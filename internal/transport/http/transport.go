package http

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/v1gn35h7/transaction-service/internal/service"
)

var (
	httpSrvOptions []httptransport.ServerOption
)

func MakeHandlers(srvc service.Service) http.Handler {
	r := mux.NewRouter()
	e := MakeEndpoints(srvc)
	logger := kitlog.NewLogfmtLogger(os.Stderr)
	httpSrvOptions = []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}
	r.Handle("/metrics", promhttp.Handler()).Methods("GET")
	r.Handle("/accounts", makeCreateAccountTransport(e.createAccountEndpoint)).Methods("POST").Name("create_account")
	r.Handle("/transactions", makeCreateTransactionTransport(e.createTransactionEndpoint)).Methods("POST").Name("create_transaction")
	r.Handle("/accounts/{accountID}", makeGetAccountTransport(e.getAccountEndpoint)).Methods("GET").Name("get_account")
	r.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)

	return r
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	// switch err {
	// case ErrNotFound:
	// 	return http.StatusNotFound
	// case ErrAlreadyExists, ErrInconsistentIDs:
	// 	return http.StatusBadRequest
	// default:
	// 	return http.StatusInternalServerError
	// }
	return http.StatusInternalServerError
}
