package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerTransaction struct {
	CustomerId         string    `json:"customer_id" bson:"customer_id" binding:"required"`
	TransactionDetails string    `json:"transaction_details" bson:"transaction_details" binding:"required"`
	Date               time.Time `json:"date" bson:"date" binding:"required"`
}

type CusTranDbResponse struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id"`
	CustomerId         string             `json:"customer_id" bson:"customer_id" binding:"required"`
	TransactionDetails string             `json:"transaction_details" bson:"transaction_details" binding:"required"`
	Date               time.Time          `json:"date" bson:"date" binding:"required"`
}
