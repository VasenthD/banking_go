package routes

import (
	"banking/controllers"

	"github.com/gin-gonic/gin"
)

func  BankRoutes(router *gin.Engine,controller controllers.BankController){
	router.POST("/api/banking/Bank/create",controller.CreateBank)
	router.GET("/api/banking/Bank/get",controller.GetBank)
	router.POST("/api/banking/Bank/update",controller.UpdateBank)
	router.DELETE("/api/banking/Bank/delete/:id",controller.DeleteBankById)

	router.GET("/api/banking/Bank/getbankdetails",controller.FindBankById)
	router.GET("/api/banking/Bank/getallcusofbank",controller.GetAllCusOfbank)
}