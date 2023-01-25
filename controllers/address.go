package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AddAdress() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func EditHomeAddress() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func EditWorkAddress() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (env *EnvCart) DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")
		if user_id == "" {
			log.Println("user id is empty")
			c.Header("Content-Type", "Application")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid index for id"})
			c.Abort()
		}
		var ctx, cancel = context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		err := env.AddressModel.DeleteAddressById(ctx, user_id)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err.Error())
		}
		c.IndentedJSON(http.StatusAccepted, "Successfully Deleted")

	}
}
