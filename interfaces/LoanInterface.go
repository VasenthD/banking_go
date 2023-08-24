package interfaces

import "banking/models"

type LoanInterface interface{
	CreateLoan(account *models.Loan)(*models.LoanDbResponse,error)
	GetLoan()([]*models.LoanDbResponse,error)
	UpdateLoan()(*models.LoanDbResponse,error)
	DeleteLoanById(id string)(string,error)
}	
