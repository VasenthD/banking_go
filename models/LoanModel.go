package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Loan struct{
	LoanAmount string `json:"loan_amount" bson:"loan_amount" binding:"required"`
	TypeOfLoan string `json:"type_of_loan" bson:"type_of_loan" binding:"required"`
}

type LoanDbResponse struct{
	ID primitive.ObjectID `json:"id" bson:"_id"`
	LoanAmount string `json:"loan_amount" bson:"loan_amount" binding:"required"`
	TypeOfLoan string `json:"type_of_loan" bson:"type_of_loan" binding:"required"`
}