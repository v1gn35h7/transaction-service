package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/v1gn35h7/transaction-service/internal/logging"
	"github.com/v1gn35h7/transaction-service/internal/mock"
	"github.com/v1gn35h7/transaction-service/internal/models"
	"go.uber.org/mock/gomock"
)

func TestTransactionService_CreateTransaction(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDS := mock.NewMockDatastore(mockCtrl)
	logger := logging.Logger()
	srvc := New(mockDS, logger)
	t.Run("Create Transaction Successfully", func(t *testing.T) {
		accountID := int64(12345)
		operationTypeID := int16(1)
		amount := 100.50
		mockDS.EXPECT().NewTransaction(accountID, operationTypeID, amount).Return(&models.Transaction{
			AccountID:       accountID,
			TransactionID:   1,
			Amount:          amount,
			OperationTypeID: operationTypeID,
			EventDate:       "2024-01-01T00:00:00Z",
		}, nil)
		err := srvc.CreateTransaction(accountID, operationTypeID, amount)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Create Transaction Failure", func(t *testing.T) {
		accountID := int64(12345)
		operationTypeID := int16(1)
		amount := 100.50
		mockDS.EXPECT().NewTransaction(accountID, operationTypeID, amount).Return(nil, errors.New("DB error"))
		err := srvc.CreateTransaction(accountID, operationTypeID, amount)
		assert.NotNil(t, err, "Expected error but got none")
	})
}
