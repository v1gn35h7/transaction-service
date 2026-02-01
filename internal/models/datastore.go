package models

type Datastore interface {
	AccountsDatastore
	TransactionsDatastore
}
