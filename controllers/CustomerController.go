package controllers

import (
	"banking/interfaces"
	"banking/models"
	// "fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type CustomerController struct{
	CustomerService interfaces.CustomerInterface
}

func  InitCustomerController(cusService interfaces.CustomerInterface)CustomerController{
	return CustomerController{cusService}
}

func (a *CustomerController)CreateCustomer(ctx *gin.Context){
	var profile *models.Customer
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newProfile, err := a.CustomerService.CreateCustomer(profile)

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

func (a *CustomerController)GetCustomer(ctx *gin.Context){
	accounts,_:=a.CustomerService.GetCustomer()
	ctx.JSON(http.StatusOK,accounts)
}
func (a *CustomerController)UpdateCustomer(ctx *gin.Context){
	updatedAcc,_:=a.CustomerService.UpdateCustomer()
	ctx.JSON(http.StatusOK,updatedAcc)
}
func(a *CustomerController)DeleteCustomerById(ctx *gin.Context){
	deleteAcc,_:=a.CustomerService.DeleteCustomerById(ctx.Param("id"))
	ctx.JSON(http.StatusOK,deleteAcc)
}
