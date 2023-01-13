package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/valilhan/go-ecommerce/controllers"
)

func UserRoutes(incomingRequest *gin.Engine, app controllers.EnvUser) {
	incomingRequest.POST("/users/signup", app.Signup())
	incomingRequest.POST("/users/login", app.Login())
	incomingRequest.POST("/admin/addproduct", app.ProductViewerAdmin())
	incomingRequest.GET("/users/productview", app.SearchProducts())
	incomingRequest.GET("/users/search", app.SearchProductsByQuery())
}
