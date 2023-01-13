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
	user controllers.EnvUser
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app := &Env{
		user: controllers.EnvUser{User: models.UserModel{database.NewDatabasePool()}},
		
    }
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router, app.user)
	router.Use(middleware.Authentication())
	router.POST("/addtocart", app.AddToCart())
	router.POST("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(routes.Run(":" + port))

}
