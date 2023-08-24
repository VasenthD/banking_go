package services

import (
	"banking/interfaces"
	"banking/models"
	// "banking/utils"
	"context"
	"fmt"

	// "fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LoanService struct{
	LoanCollection *mongo.Collection
	ctx context.Context
}
func InitLoanService(collection *mongo.Collection,ctx context.Context)interfaces.LoanInterface{
	return &LoanService{collection,ctx}
}
func(a *LoanService) CreateLoan(loan *models.Loan)(*models.LoanDbResponse,error){
	res,_:=a.LoanCollection.InsertOne(a.ctx,&loan)
	var newUser *models.LoanDbResponse
	query := bson.M{"_id": res.InsertedID}

	err := a.LoanCollection.FindOne(a.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser,nil
}

func(a *LoanService)GetLoan()([]*models.LoanDbResponse,error){
	filter:=bson.D{}
	options:=options.Find()
	res,_:=a.LoanCollection.Find(a.ctx,filter,options)
	var loan[]*models.LoanDbResponse
	for res.Next(a.ctx){
		acc:=&models.LoanDbResponse{}
		err:=res.Decode(acc)
		if err!=nil{
			return nil,err
		}
		loan=append(loan, acc)
	}
	return loan,nil
}
func (a *LoanService)UpdateLoan()(*models.LoanDbResponse,error){
	filter:=bson.D{{"bank_id","SBI001"}}
	updatingValue:=bson.D{{"$set",bson.D{{"name","akil adharsan"}}}}
	var updatedAcc *models.LoanDbResponse
	res,_:=a.LoanCollection.UpdateOne(a.ctx,filter,updatingValue)
	fmt.Println(res)
	query := bson.M{"bank_id": "SBI001"}
	err := a.LoanCollection.FindOne(a.ctx, query).Decode(&updatedAcc)
	if err!=nil{
		return nil,err
	}
	return updatedAcc,nil
}

func(a *LoanService)DeleteLoanById(id string)(string,error){
	filter:=bson.D{{"bank_id","SBI001"}}
	_,err:=a.LoanCollection.DeleteOne(a.ctx,filter)
	if  err!=nil{
		return "Account Not Deleted",err
	}
	return "Account Successfully Deleted",nil
}
