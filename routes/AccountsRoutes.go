package routes

import (
	"banking/controllers"

	"github.com/gin-gonic/gin"
)

func  AccountsRoutes(router *gin.Engine,controller controllers.AccountsController){
	router.POST("/api/banking/Accounts/create",controller.CreateAccount)
	router.GET("/api/banking/Accounts/get",controller.GetAccount)
	router.POST("/api/banking/Accounts/update",controller.UpdateAccount)
	router.DELETE("/api/banking/Accounts/delete/:id",controller.DeleteAccountById)
}