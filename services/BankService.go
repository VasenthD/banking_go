package services

import (
	"banking/interfaces"
	"banking/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BankService struct{
	BankCollection *mongo.Collection
	CustomerColl *mongo.Collection
	ctx context.Context
}
func InitBankService(bankcollection *mongo.Collection,cuscollection *mongo.Collection,ctx context.Context)interfaces.BankInterface{
	return &BankService{bankcollection,cuscollection,ctx}
}
func(a *BankService) CreateBank(bank *models.Bank)(*models.BankDbResponse,error){
	res,_:=a.BankCollection.InsertOne(a.ctx,&bank)
	var newUser *models.BankDbResponse
	query := bson.M{"_id": res.InsertedID}

	err := a.BankCollection.FindOne(a.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser,nil
}

func(a *BankService)GetBank()([]*models.BankDbResponse,error){
	filter:=bson.D{}
	options:=options.Find()
	res,_:=a.BankCollection.Find(a.ctx,filter,options)
	var bank[]*models.BankDbResponse
	for res.Next(a.ctx){
		acc:=&models.BankDbResponse{}
		err:=res.Decode(acc)
		if err!=nil{
			return nil,err
		}
		bank=append(bank, acc)
	}
	return bank,nil
}
func (a *BankService)UpdateBank()(*models.BankDbResponse,error){
	filter:=bson.D{{"ifsc_code","123456789"}}
	updatingValue:=bson.D{{"$set",bson.D{{"bank_id","SBI001"}}}}
	var updatedAcc *models.BankDbResponse
	res,_:=a.BankCollection.UpdateOne(a.ctx,filter,updatingValue)
	fmt.Println(res)
	query := bson.M{"ifsc_code": "123456789"}
	err := a.BankCollection.FindOne(a.ctx, query).Decode(&updatedAcc)
	if err!=nil{
		return nil,err
	}
	return updatedAcc,nil
}

func(a *BankService)DeleteBankById(id string)(string,error){
	filter:=bson.D{{"bank_id","SBI001"}}
	_,err:=a.BankCollection.DeleteOne(a.ctx,filter)
	if  err!=nil{
		return "Account Not Deleted",err
	}
	return "Account Successfully Deleted",nil
}

func (a *BankService) FindBankById() ([]*models.Bank, error) {
	filter1 := bson.D{{"customer_id","003"}}
	var cus *models.Customer
	err := a.CustomerColl.FindOne(a.ctx, filter1).Decode(&cus)
	if err != nil {
		fmt.Println("error")
		return nil, err
	}
	filter := bson.M{"bank_id": cus.BankID}
	options := options.Find()
	res, err := a.BankCollection.Find(a.ctx, filter, options)
	if err != nil {
		fmt.Println("!!!!!!!!!!")
		return nil, err
	}
	defer res.Close(a.ctx)
	var banks []*models.Bank
	for res.Next(a.ctx) {
		acc := &models.Bank{}
		err := res.Decode(acc)
		if err != nil {
			return nil, err
		}
		banks = append(banks, acc)
	}
	return banks, nil
}


func (a *BankService) GetAllCusOfbank() ([]*models.Customer, error) {
	filter := bson.M{"bank_id":"SBI001"}
	options := options.Find()
	res, err := a.CustomerColl.Find(a.ctx, filter, options)
	if err != nil {
		fmt.Println("!!!!!!!!!!")
		return nil, err
	}
	defer res.Close(a.ctx)

	var banks []*models.Customer
	for res.Next(a.ctx) {
		acc := &models.Customer{}
		err := res.Decode(acc)
		if err != nil {
			fmt.Println("hello")
			return nil, err
		}
		banks = append(banks, acc)
	}
	return banks, nil
}
