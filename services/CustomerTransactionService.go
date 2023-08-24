package services

import (
	"banking/interfaces"
	"banking/models"
	"fmt"
	"time"

	// "banking/utils"
	"context"
	// "fmt"

	// "fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type CustomerTransactionService struct {
	CustomerTransactionCollection *mongo.Collection
	ctx                           context.Context
}

func InitCustomerTransactionService(collection *mongo.Collection, ctx context.Context) interfaces.CustomerTransactionInterface {
	return &CustomerTransactionService{collection, ctx}
}
func (a *CustomerTransactionService) CreateCustomerTransaction(customer *models.CustomerTransaction) (*models.CusTranDbResponse, error) {
	res, _ := a.CustomerTransactionCollection.InsertOne(a.ctx, &customer)
	var newUser *models.CusTranDbResponse
	query := bson.M{"_id": res.InsertedID}

	err := a.CustomerTransactionCollection.FindOne(a.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (a *CustomerTransactionService) GetCustomerTransaction() ([]*models.CusTranDbResponse, error) {
	filter := bson.D{}
	options := options.Find()
	res, _ := a.CustomerTransactionCollection.Find(a.ctx, filter, options)
	var customer []*models.CusTranDbResponse
	for res.Next(a.ctx) {
		acc := &models.CusTranDbResponse{}
		err := res.Decode(acc)
		if err != nil {
			return nil, err
		}
		customer = append(customer, acc)
	}
	return customer, nil
}
func (a *CustomerTransactionService) UpdateCustomerTransaction() (*models.CusTranDbResponse, error) {
	filter := bson.D{{Key: "customer_id", Value: "001"}}
	updatingValue := bson.D{{Key: "$set", Value: bson.D{{Key: "transaction_details", Value: "paid $400 in gpay"}}}}
	var updatedAcc *models.CusTranDbResponse
	res, _ := a.CustomerTransactionCollection.UpdateOne(a.ctx, filter, updatingValue)
	fmt.Println(res)
	query := bson.M{"customer_id": "001"}
	err := a.CustomerTransactionCollection.FindOne(a.ctx, query).Decode(&updatedAcc)
	if err != nil {
		return nil, err
	}
	return updatedAcc, nil
}

func (a *CustomerTransactionService) DeleteCustomerTransactionById(id string) (string, error) {
	filter := bson.D{{"customer_id", "001"}}
	_, err := a.CustomerTransactionCollection.DeleteOne(a.ctx, filter)
	if err != nil {
		return "Account Not Deleted", err
	}
	return "Account Successfully Deleted", nil
}

func (a *CustomerTransactionService) GetCustomerTransactionByCustomer() ([]*models.CustomerTransaction, error) {
	filter := bson.M{"customer_id": "002"}
	options := options.Find()
	res, err := a.CustomerTransactionCollection.Find(a.ctx, filter, options)
	if err != nil {
		fmt.Println("!!!!!!!!!!")
		return nil, err
	}
	defer res.Close(a.ctx) // Close the cursor when done

	var cusTra []*models.CustomerTransaction
	for res.Next(a.ctx) {
		acc := &models.CustomerTransaction{}
		err := res.Decode(acc)
		if err != nil {
			return nil, err
		}
		cusTra = append(cusTra, acc)
	}
	return cusTra, nil
}
func (a *CustomerTransactionService) GetTransactionsByDateRange(fromDate, toDate time.Time) ([]*models.CustomerTransaction, error) {
	filter := bson.M{
		"date": bson.M{
			"$gte": fromDate,
			"$lte": toDate,
		},
	}
	options := options.Find()
	res, err := a.CustomerTransactionCollection.Find(a.ctx, filter, options)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer res.Close(a.ctx)

	var transactions []*models.CustomerTransaction
	for res.Next(a.ctx) {
		transaction := &models.CustomerTransaction{}
		err := res.Decode(transaction)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
