package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/valilhan/go-ecommerce/controllers"
)

func UserRoutes(incomingRequest *gin.Engine, user controllers.EnvUser, app controllers.EnvCart) {
	incomingRequest.POST("/users/signup", user.Signup())
	incomingRequest.POST("/users/login", user.Login())
	incomingRequest.POST("/admin/addproduct", user.ProductViewerAdmin())
	incomingRequest.GET("/users/productview", app.SearchProducts())
	incomingRequest.GET("/users/search", app.SearchProductsByQuery())
}
