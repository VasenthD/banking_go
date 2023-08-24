package routes

import (
	"banking/controllers"

	"github.com/gin-gonic/gin"
)

func CustomerTransactionRoutes(router *gin.Engine, controller controllers.CustomerTransactionController) {
	router.POST("/api/banking/CustomerTransaction/create", controller.CreateCustomerTransaction)
	router.GET("/api/banking/CustomerTransaction/get", controller.GetCustomerTransaction)
	router.POST("/api/banking/CustomerTransaction/update", controller.UpdateCustomerTransaction)
	router.DELETE("/api/banking/CustomerTransaction/delete/:id", controller.DeleteCustomerTransactionById)

	router.GET("/api/banking/CustomerTransaction/getcustans", controller.GetCustomerTransactionByCustomer)
	router.POST("/transactions/by-date-range", controller.GetTransactionsByDateRange)
}
