package models

import "github.com/jmoiron/sqlx"

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
