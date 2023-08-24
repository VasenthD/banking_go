package controllers

import (
	"banking/interfaces"
	"banking/models"
	// "fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type LoanController struct{
	LoanService interfaces.LoanInterface
}

func  InitLoanController(loanService interfaces.LoanInterface)LoanController{
	return LoanController{loanService}
}

func (a *LoanController)CreateLoan(ctx *gin.Context){
	var profile *models.Loan
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newProfile, err := a.LoanService.CreateLoan(profile)

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

func (a *LoanController)GetLoan(ctx *gin.Context){
	accounts,_:=a.LoanService.GetLoan()
	ctx.JSON(http.StatusOK,accounts)
}
func (a *LoanController)UpdateLoan(ctx *gin.Context){
	updatedAcc,_:=a.LoanService.UpdateLoan()
	ctx.JSON(http.StatusOK,updatedAcc)
}
func(a *LoanController)DeleteLoanById(ctx *gin.Context){
	deleteAcc,_:=a.LoanService.DeleteLoanById(ctx.Param("id"))
	ctx.JSON(http.StatusOK,deleteAcc)
}
