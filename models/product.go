package models

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	Id     *string `json:"Id"`
	Name   *string `json:"Name"`
	Price  *uint64 `json:"Price"`
	Rating *uint8  `json:"Rating"`
	Image  *string `json:"Image"`
}

// Create a custom type which wraps the sql.DB connection pool.
type ProductModel struct {
	DB *sqlx.DB
}

func (env *ProductModel) SelectAllProducts(ctx context.Context, productlist *[]Product) error {
	query := "SELECT * FROM Products;"
	rows, err := env.DB.QueryxContext(ctx, query)
	if err != nil {
		return errors.New("Query err, can not select products")
	}
	for rows.Next() {
		var product Product
		err := rows.StructScan(&product)
		if err != nil {
			return errors.New("error in structscan ")
		}
		*productlist = append(*productlist, product)
	}
	return nil
}
