package controllers

import (
	"banking/interfaces"
	"banking/models"
	// "fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type BankController struct{
	BankService interfaces.BankInterface
}

func  InitBankController(loanService interfaces.BankInterface)BankController{
	return BankController{loanService}
}

func (a *BankController)CreateBank(ctx *gin.Context){
	var profile *models.Bank
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newProfile, err := a.BankService.CreateBank(profile)

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

func (a *BankController)GetBank(ctx *gin.Context){
	accounts,_:=a.BankService.GetBank()
	ctx.JSON(http.StatusOK,accounts)
}
func (a *BankController)UpdateBank(ctx *gin.Context){
	updatedAcc,_:=a.BankService.UpdateBank()
	ctx.JSON(http.StatusOK,updatedAcc)
}
func(a *BankController)DeleteBankById(ctx *gin.Context){
	deleteAcc,_:=a.BankService.DeleteBankById(ctx.Param("id"))
	ctx.JSON(http.StatusOK,deleteAcc)
}

func (a *BankController)FindBankById(ctx *gin.Context){
	accounts,_:=a.BankService.FindBankById()
	ctx.JSON(http.StatusOK,accounts)
}

func (a *BankController)GetAllCusOfbank(ctx *gin.Context){
	accounts,_:=a.BankService.GetAllCusOfbank()
	ctx.JSON(http.StatusOK,accounts)
}
