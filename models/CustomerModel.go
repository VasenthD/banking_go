package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	Name       string `json:"name" bson:"name" binding:"required"`
	CustomerId string `json:"customer_id" bson:"customer_id" binding:"required"`
	BankID     string `json:"bank_id" bson:"bank_id" binding:"required"`
	Password   string `json:"password" bson:"password" binding:"required"`
}

type CustomerDbResponse struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Name       string             `json:"name" bson:"name" binding:"required"`
	CustomerId string             `json:"customer_id" bson:"customer_id" binding:"required"`
	BankID     string             `json:"bank_id" bson:"bank_id" binding:"required"`
	Password   string             `json:"password" bson:"password" binding:"required"`
}
