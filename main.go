package main

import (
	"banking/config"
	"banking/constants"
	"banking/controllers"
	"banking/routes"
	"banking/services"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoclient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)

func initRoutes() {
	routes.Default(server)
}


func initApp(mongoClient *mongo.Client) {
	ctx = context.TODO()
	CustomerCollection := mongoClient.Database(constants.Datbasename).Collection("Customer")
	CustomerService := services.InitCustomerService(CustomerCollection, ctx)
	CustomerController := controllers.InitCustomerController(CustomerService)
	routes.CustomerRoutes(server, CustomerController)

	CustomerTransactionCollection := mongoClient.Database(constants.Datbasename).Collection("Customer_Transactions")
	CustomerTransactionService := services.InitCustomerTransactionService(CustomerTransactionCollection, ctx)
	CustomerTransactionController := controllers.InitCustomerTransactionController(CustomerTransactionService)
	routes.CustomerTransactionRoutes(server, CustomerTransactionController)
	
	LoanCollection := mongoClient.Database(constants.Datbasename).Collection("Loan")
	LoanService := services.InitLoanService(LoanCollection, ctx)
	LoanController := controllers.InitLoanController(LoanService)
	routes.LoanRoutes(server, LoanController)

	AccountCollection := mongoClient.Database(constants.Datbasename).Collection("Account")
	AccountService := services.InitAccountsService(AccountCollection, ctx)
	AccountsController := controllers.InitAccountsController(AccountService)
	routes.AccountsRoutes(server, AccountsController)

	BankCollection := mongoClient.Database(constants.Datbasename).Collection("Bank")
	BankService := services.InitBankService(BankCollection,CustomerCollection, ctx)
	BankController := controllers.InitBankController(BankService)
	routes.BankRoutes(server,BankController)

}

func main() {
	server = gin.Default()
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	initRoutes()
	initApp(mongoclient)
	fmt.Println("server running on port", constants.Port)
	log.Fatal(server.Run(constants.Port))
}
//insert transaction every month u make a transaction
//input will be from date and to date
