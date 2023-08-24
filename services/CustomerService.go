package services

import (
	"banking/interfaces"
	"banking/models"
	"banking/utils"
	"context"
	"fmt"

	// "fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type CustomerService struct{
	CustomerCollection *mongo.Collection
	ctx context.Context
}
func InitCustomerService(collection *mongo.Collection,ctx context.Context)interfaces.CustomerInterface{
	return &CustomerService{collection,ctx}
}
func(a *CustomerService) CreateCustomer(customer *models.Customer)(*models.CustomerDbResponse,error){
	hashedPassword, _ := utils.HashPassword(customer.Password)
	customer.Password = string(hashedPassword)
	res,_:=a.CustomerCollection.InsertOne(a.ctx,&customer)
	var newUser *models.CustomerDbResponse
	query := bson.M{"_id": res.InsertedID}

	err := a.CustomerCollection.FindOne(a.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser,nil
}

func(a *CustomerService)GetCustomer()([]*models.CustomerDbResponse,error){
	filter:=bson.D{}
	options:=options.Find()
	res,_:=a.CustomerCollection.Find(a.ctx,filter,options)
	var customer[]*models.CustomerDbResponse
	for res.Next(a.ctx){
		acc:=&models.CustomerDbResponse{}
		err:=res.Decode(acc)
		if err!=nil{
			return nil,err
		}
		customer=append(customer, acc)
	}
	return customer,nil
}
func (a *CustomerService)UpdateCustomer()(*models.CustomerDbResponse,error){
	filter:=bson.D{{"bank_id","SBI001"}}
	updatingValue:=bson.D{{"$set",bson.D{{"name","akil adharsan"}}}}
	var updatedAcc *models.CustomerDbResponse
	res,_:=a.CustomerCollection.UpdateOne(a.ctx,filter,updatingValue)
	fmt.Println(res)
	query := bson.M{"bank_id": "SBI001"}
	err := a.CustomerCollection.FindOne(a.ctx, query).Decode(&updatedAcc)
	if err!=nil{
		return nil,err
	}
	return updatedAcc,nil
}

func(a *CustomerService)DeleteCustomerById(id string)(string,error){
	filter:=bson.D{{"bank_id","SBI001"}}
	_,err:=a.CustomerCollection.DeleteOne(a.ctx,filter)
	if  err!=nil{
		return "Account Not Deleted",err
	}
	return "Account Successfully Deleted",nil
}
