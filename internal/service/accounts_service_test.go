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

func TestCreateAccount(t *testing.T) {
	// Initialize mock datastore
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDS := mock.NewMockDatastore(ctrl)
	logger := logging.Logger()
	srvc := New(mockDS, logger)
	t.Run("Create Account Successfully", func(t *testing.T) {
		documentNumber := int64(1)
		mockDS.EXPECT().NewAccount(documentNumber).Return(int64(12345), nil)
		accountID, err := srvc.CreateAccount(documentNumber)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		assert.Equal(t, int64(12345), accountID, "Account ID should match")
	})

	t.Run("Create Account Failure", func(t *testing.T) {
		documentNumber := int64(1)
		mockDS.EXPECT().NewAccount(documentNumber).Return(int64(0), errors.New("DB error"))
		_, err := srvc.CreateAccount(documentNumber)
		assert.NotNil(t, err, "Expected error but got none")
	})
}

func TestGetAccount(t *testing.T) {
	// Initialize mock datastore
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDS := mock.NewMockDatastore(ctrl)
	logger := logging.Logger()
	srvc := New(mockDS, logger)
	t.Run("Get Account Successfully", func(t *testing.T) {
		accountID := int64(12345)
		expectedAccount := models.Account{AccountID: accountID, DocumentNumber: 1}
		mockDS.EXPECT().GetAccount(accountID).Return(expectedAccount, nil)
		account, err := srvc.GetAccount(accountID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		assert.Equal(t, expectedAccount, account, "Account should match")
	})
	t.Run("Get Account Failure", func(t *testing.T) {
		accountID := int64(12345)
		mockDS.EXPECT().GetAccount(accountID).Return(models.Account{}, errors.New("DB error"))
		_, err := srvc.GetAccount(accountID)
		assert.NotNil(t, err, "Expected error but got none")
	})
}
