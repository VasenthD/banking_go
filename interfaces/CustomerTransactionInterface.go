package interfaces

import (
	"banking/models"
	"time"
)

type CustomerTransactionInterface interface {
	CreateCustomerTransaction(custTrans *models.CustomerTransaction) (*models.CusTranDbResponse, error)
	GetCustomerTransaction() ([]*models.CusTranDbResponse, error)
	UpdateCustomerTransaction() (*models.CusTranDbResponse, error)
	DeleteCustomerTransactionById(id string) (string, error)
	GetCustomerTransactionByCustomer() ([]*models.CustomerTransaction, error)
	GetTransactionsByDateRange(fromDate, toDate time.Time) ([]*models.CustomerTransaction, error)
}
