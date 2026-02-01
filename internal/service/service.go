package service

import (
	"github.com/go-logr/logr"
	"github.com/v1gn35h7/transaction-service/internal/models"
)

type Service interface {
	AccountService
	TransactionService
}

type service struct {
	ds     models.Datastore
	logger logr.Logger
}

func New(ds models.Datastore, logger logr.Logger) service {
	return service{ds: ds, logger: logger}
}
