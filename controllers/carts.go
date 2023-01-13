package controllers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/valilhan/go-ecommerce/database"
	"github.com/valilhan/go-ecommerce/models"
)

type EnvCart struct {
	UserModel    models.UserModel
	ProductModel models.ProductModel
}

func (env *EnvCart) AddProduct() gin.HandlerFunc {
	//http://localhost:8000/addtocart?id=xxxproduct_idxxx&userID=xxxxxxuser_idxxxxxx
	return func(c *gin.Context) {
		productId := c.Query("id")
		if productId == "" {
			c.AbortWithError(http.StatusBadRequest, errors.New("Product id is empty"))
			return
		}
		userID := c.Query("userID")
		if userID == "" {
			c.AbortWithError(http.StatusBadRequest, errors.New("User id is empty"))
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		err := database.AddProductToCart(ctx, env.UserModel, env.ProductModel, productId, userID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Successfully add to the cart")
	}
}

func (env *EnvCart) RemoveItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (env *EnvCart) GetItemFromCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (env *EnvCart) BuyFromCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (env *EnvCart) InstantBuy() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
