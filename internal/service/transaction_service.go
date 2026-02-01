package service

type TransactionService interface {
	CreateTransaction(account_id int64, operation_type_id int16, amount float64) error
}

func (s service) CreateTransaction(account_id int64, operation_type_id int16, amount float64) error {
	_, err := s.ds.NewTransaction(account_id, operation_type_id, amount)
	if err != nil {
		return err
	}
	return nil
}
