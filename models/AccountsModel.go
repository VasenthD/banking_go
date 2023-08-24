package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Accounts struct{
	AccountId string `json:"account_id" bson:"account_id" binding:"required"`
	CustomerId string `json:"customer_id" bson:"customer_id" binding:"required"`
	AccouontType string `json:"account_type" bson:"account_type" binding:"required"`
}

type AccountsDbResponse struct{
	ID primitive.ObjectID `json:"id" bson:"_id"`
	AccountId string `json:"account_id" bson:"account_id" binding:"required"`
	CustomerId string `json:"customer_id" bson:"customer_id" binding:"required"`
	AccouontType string `json:"account_type" bson:"account_type" binding:"required"`
}