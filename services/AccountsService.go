package services

import (
	"banking/interfaces"
	"banking/models"
	"context"
	"fmt"

	// "fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AccountsService struct {
	AccountCollection *mongo.Collection
	ctx               context.Context
}

func InitAccountsService(collection *mongo.Collection, ctx context.Context) interfaces.AccountsInterface {
	return &AccountsService{collection, ctx}
}
func (a *AccountsService) CreateAccount(account *models.Accounts) (*models.AccountsDbResponse, error) {
	res, _ := a.AccountCollection.InsertOne(a.ctx, &account)
	var newUser *models.AccountsDbResponse
	query := bson.M{"_id": res.InsertedID}

	err := a.AccountCollection.FindOne(a.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (a *AccountsService) GetAccount() ([]*models.AccountsDbResponse, error) {
	filter := bson.D{}
	options := options.Find()
	res, _ := a.AccountCollection.Find(a.ctx, filter, options)
	var accounts []*models.AccountsDbResponse
	for res.Next(a.ctx) {
		acc := &models.AccountsDbResponse{}
		err := res.Decode(acc)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, acc)
	}
	return accounts, nil
}
func (a *AccountsService) UpdateAccount() (*models.AccountsDbResponse, error) {
	filter := bson.D{{"bank_id", "SBI001"}}
	updatingValue := bson.D{{"$set", bson.D{{"name", "akil adharsan"}}}}
	var updatedAcc *models.AccountsDbResponse
	res, _ := a.AccountCollection.UpdateOne(a.ctx, filter, updatingValue)
	fmt.Println(res)
	query := bson.M{"bank_id": "SBI001"}
	err := a.AccountCollection.FindOne(a.ctx, query).Decode(&updatedAcc)
	if err != nil {
		return nil, err
	}
	return updatedAcc, nil
}

func (a *AccountsService) DeleteAccountById(id string) (string, error) {
	filter := bson.D{{"customer_id", "001"}}
	_, err := a.AccountCollection.DeleteOne(a.ctx, filter)
	if err != nil {
		return "Account Not Deleted", err
	}
	return "Account Successfully Deleted", nil
}
