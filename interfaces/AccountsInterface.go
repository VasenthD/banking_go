package interfaces

import "banking/models"

type AccountsInterface interface{
	CreateAccount(account *models.Accounts)(*models.AccountsDbResponse,error)
	GetAccount()([]*models.AccountsDbResponse,error)
	UpdateAccount()(*models.AccountsDbResponse,error)
	DeleteAccountById(id string)(string,error)
}	