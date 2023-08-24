package controllers

import (
	"banking/interfaces"
	"banking/models"
	// "fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AccountsController struct{
	AccountsService interfaces.AccountsInterface
}

func  InitAccountsController(accountService interfaces.AccountsInterface)AccountsController{
	return AccountsController{accountService}
}

func (a *AccountsController)CreateAccount(ctx *gin.Context){
	var profile *models.Accounts
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newProfile, err := a.AccountsService.CreateAccount(profile)

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

func (a *AccountsController)GetAccount(ctx *gin.Context){
	accounts,_:=a.AccountsService.GetAccount()
	ctx.JSON(http.StatusOK,accounts)
}
func (a *AccountsController)UpdateAccount(ctx *gin.Context){
	updatedAcc,_:=a.AccountsService.UpdateAccount()
	ctx.JSON(http.StatusOK,updatedAcc)
}
func(a *AccountsController)DeleteAccountById(ctx *gin.Context){
	deleteAcc,_:=a.AccountsService.DeleteAccountById(ctx.Param("id"))
	ctx.JSON(http.StatusOK,deleteAcc)
}
