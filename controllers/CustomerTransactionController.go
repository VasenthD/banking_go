package controllers

import (
	"banking/interfaces"
	"banking/models"
	"time"

	// "fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type CustomerTransactionController struct {
	CustomerTransactionService interfaces.CustomerTransactionInterface
}

func InitCustomerTransactionController(cusTransService interfaces.CustomerTransactionInterface) CustomerTransactionController {
	return CustomerTransactionController{cusTransService}
}

func (a *CustomerTransactionController) CreateCustomerTransaction(ctx *gin.Context) {
	var profile *models.CustomerTransaction
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newProfile, err := a.CustomerTransactionService.CreateCustomerTransaction(profile)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newProfile})
}

func (a *CustomerTransactionController) GetCustomerTransaction(ctx *gin.Context) {
	custTrans, _ := a.CustomerTransactionService.GetCustomerTransaction()
	ctx.JSON(http.StatusOK, custTrans)
}
func (a *CustomerTransactionController) UpdateCustomerTransaction(ctx *gin.Context) {
	updatedcustTrans, _ := a.CustomerTransactionService.UpdateCustomerTransaction()
	ctx.JSON(http.StatusOK, updatedcustTrans)
}
func (a *CustomerTransactionController) DeleteCustomerTransactionById(ctx *gin.Context) {
	deletecustTrans, _ := a.CustomerTransactionService.DeleteCustomerTransactionById(ctx.Param("id"))
	ctx.JSON(http.StatusOK, deletecustTrans)
}
func (a *CustomerTransactionController) GetCustomerTransactionByCustomer(ctx *gin.Context) {
	custTrans, _ := a.CustomerTransactionService.GetCustomerTransactionByCustomer()
	ctx.JSON(http.StatusOK, custTrans)
}

//func (a *CustomerTransactionController) GetTransactionsByDateRange(ctx *gin.Context) {

// FromDateString := time.Date(2020, 9, 1, 0, 0, 0, 0, time.UTC)
// transactions, err := a.CustomerTransactionService.GetTransactionsByDateRange(FromDateString, time.Now())
// if err != nil {
// 	ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
// 	return
// }

// ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": transactions})
func (a *CustomerTransactionController) GetTransactionsByDateRange(ctx *gin.Context) {
	var request struct {
		FromDate string `json:"fromDate"`
		ToDate   string `json:"toDate" `
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
 
	// Parse the fromDate and toDate strings into time.Time objects
	fromDate, err := time.Parse("2006-01-02", request.FromDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid 'fromDate' parameter"})
		return
	}

	toDate, err := time.Parse("2006-01-02", request.ToDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid 'toDate' parameter"})
		return
	}

	transactions, err := a.CustomerTransactionService.GetTransactionsByDateRange(fromDate, toDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": transactions})
}
