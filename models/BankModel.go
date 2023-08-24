package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bank struct{
	BankId string `json:"bank_id" bson:"bank_id" binding:"required"`
	IFSCcode string `json:"ifsc_code" bson:"ifsc_code" binding:"required"`
	BankAddress string `json:"bank_address" bson:"bank_address" binding:"required"`
}

type BankDbResponse struct{
	ID primitive.ObjectID `json:"id" bson:"_id"`
	BankId string `json:"bank_id" bson:"bank_id" binding:"required"`
	IFSCcode string `json:"ifsc_code" bson:"ifsc_code" binding:"required"`
	BankAddress string `json:"bank_address" bson:"bank_address" binding:"required"`
}