package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/valilhan/go-ecommerce/controllers"
)

func UserRoutes(incomingRequest *gin.Engine) {
	incomingRequest.POST("/users/signup", controllers.Signup())
	incomingRequest.POST("/users/login", controllers.Login())
	incomingRequest.POST("/admin/addproduct", controllers.AddProduct())
	incomingRequest.GET("/users/productview", controllers.SearchProducts())
	incomingRequest.GET("/users/search", controllers.SearchProductsByQuery())
}
