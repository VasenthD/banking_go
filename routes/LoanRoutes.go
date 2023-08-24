package routes

import (
	"banking/controllers"

	"github.com/gin-gonic/gin"
)

func  LoanRoutes(router *gin.Engine,controller controllers.LoanController){
	router.POST("/api/banking/Loan/create",controller.CreateLoan)
	router.GET("/api/banking/Loan/get",controller.GetLoan)
	router.POST("/api/banking/Loan/update",controller.UpdateLoan)
	router.DELETE("/api/banking/Loan/delete/:id",controller.DeleteLoanById)
}