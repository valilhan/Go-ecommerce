package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/valilhan/go-ecommerce/models"
	"golang.org/x/crypto/bcrypt"
)

type EnvUser struct {
	User models.UserModel
}

var validate *validator.Validate

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(password string, passwordDb string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordDb), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (env *EnvUser) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
		defer cancel()
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		validate = validator.New()
		if err := validate.Struct(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		users, err := env.User.FindAllUserByEmail(ctx, *user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if len(users) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user with email already exists"})
			return
		}
		users, err = env.User.FindAllUserByPhone(ctx, *user.Phone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if len(users) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user with phone already exists"})
		}
		password, err := HashPassword(*user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "can not hash the password"})
			return
		}
		user.Password = &password
		user.CreatedAt, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		user.UpdatedAt, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		token, refreshToken := generate.TokenGenerator(*user.Email, *user.Phone, *user.FirstName, *user.LastName)
		user.Token = &token
		user.RefreshToken = &refreshToken
		userId, err := env.User.InsertUser(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		user.UserId = userId
		c.JSON(http.StatusCreated, gin.H{"Successfully Signed Up!!": user})
	}
}

func (env *EnvUser) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		//Check exists this user in Database
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		users, err := env.User.FindAllUserByEmail(ctx, *user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		//Check password user and database's user password
		passwordDb := users[0].Password
		err = VerifyPassword(*user.Password, *passwordDb)

		if len(users) != 1 || err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email or password is incorrect"})
		}

		token, refreshToken := generate.TokenGenerator(*user.Email, *user.Phone, *user.FirstName, *user.LastName)
		err = env.User.UpdateToken(ctx, token, refreshToken, users[0].UserId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusFound, users[0])
		
	}
}

func (env *EnvCart) SearchProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var productList []models.Product
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		err := env.ProductModel.SelectAllProducts(ctx, &productList)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusFound, productList)
	}
}

func (env *EnvUser) SearchProductsByQuery() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (env *EnvUser) ProductViewerAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
