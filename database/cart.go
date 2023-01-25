package database

import (
	"context"

	"github.com/valilhan/go-ecommerce/models"
)

func AddProductToCart(ctx context.Context, user models.UserModel, product models.ProductModel, productQueryId string, userQueryId string) error {
	return nil
}

func RemoveProductFromCart(ctx context.Context, user models.UserModel, product models.ProductModel, productQueryId string, userQueryId string) error {
	return nil
}

func BuyProductInstance(ctx context.Context, user models.UserModel, product models.ProductModel, productQueryId string, userQueryId string) error {
	return nil
}

func BuyFromCart(ctx context.Context, user models.UserModel, userQueryId string) error {
	return nil
}
