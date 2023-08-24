package routes

import (
	"banking/controllers"

	"github.com/gin-gonic/gin"
)

func  CustomerRoutes(router *gin.Engine,controller controllers.CustomerController){
	router.POST("/api/banking/Customer/create",controller.CreateCustomer)
	router.GET("/api/banking/Customer/get",controller.GetCustomer)
	router.POST("/api/banking/Customer/update",controller.UpdateCustomer)
	router.DELETE("/api/banking/Customer/delete/:id",controller.DeleteCustomerById)
}