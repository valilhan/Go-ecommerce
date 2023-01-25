package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/valilhan/go-ecommerce/controllers"
	"github.com/valilhan/go-ecommerce/database"
	"github.com/valilhan/go-ecommerce/middleware"
	"github.com/valilhan/go-ecommerce/models"
	"github.com/valilhan/go-ecommerce/routes"
)

type Env struct {
	user    controllers.EnvUser
	EnvCart controllers.EnvCart
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	//For each model we create different POOL connection
	app := &Env{
		user: controllers.EnvUser{User: models.UserModel{database.NewDatabasePool()}},
		EnvCart: controllers.EnvCart{UserModel: models.UserModel{database.NewDatabasePool()},
			ProductModel: models.ProductModel{database.NewDatabasePool()},
		},
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router, app.user, app.EnvCart)
	router.Use(middleware.Authentication())
	router.POST("/addtocart", app.EnvCart.AddProduct())
	router.POST("/removeitem", app.EnvCart.RemoveItem())
	router.GET("/cartcheckout", app.EnvCart.BuyFromCart())
	router.GET("/instantbuy", app.EnvCart.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
