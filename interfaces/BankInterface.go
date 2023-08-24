package interfaces

import "banking/models"

type BankInterface interface{
	CreateBank(account *models.Bank)(*models.BankDbResponse,error)
	GetBank()([]*models.BankDbResponse,error)
	UpdateBank()(*models.BankDbResponse,error)
	DeleteBankById(id string)(string,error)
	FindBankById()([]*models.Bank,error)
	GetAllCusOfbank()([]*models.Customer,error)
}	
