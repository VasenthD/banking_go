package interfaces

import "banking/models"

type CustomerInterface interface{
	CreateCustomer(account *models.Customer)(*models.CustomerDbResponse,error)
	GetCustomer()([]*models.CustomerDbResponse,error)
	UpdateCustomer()(*models.CustomerDbResponse,error)
	DeleteCustomerById(id string)(string,error)
}	
