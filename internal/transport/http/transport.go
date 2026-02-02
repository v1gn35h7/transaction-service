package http

import (
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
	r.Handle("/transactions/{accountID}", makeListTransactionsTransport(e.listTransactionsEndpoint)).Methods("GET").Name("list_transactions")
	r.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)

	return r
}
