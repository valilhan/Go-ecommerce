package main

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	middleware "github.com/valilhan/go-ecommerce/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	routes := gin.New()
	routes.Use(gin.Logger())
	routes.Use(middleware.Authentication())

	routes.POST("/addtocart", app.AddToCart())
	routes.POST("/removeitem", app.RemoveItem())
	routes.GET("/cartcheckout", app.BuyFromCart())
	routes.GET("/instantbuy", app.InstantBuy())

	log.Fatal(routes.Run(":" + port))

}
