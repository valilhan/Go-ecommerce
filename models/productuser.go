package models

import "github.com/jmoiron/sqlx"

type ProductUser struct {
	Id     *string `json:"Id"`
	Name   *string `json:"Name"`
	Price  uint64  `json:"Price"`
	Rating *uint8  `json:"Rating"`
	Image  *string `json:"Image"`
}

// Create a custom type which wraps the sql.DB connection pool.
type ProductUserModel struct {
	DB *sqlx.DB
}
