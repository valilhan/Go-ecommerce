package models

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Address struct {
	Id      *string `json:"Id"`
	House   *string `json:"House"`
	Street  *string `json:"Street"`
	City    *string `json:"City"`
	Pincode *string `json:"Pincode"`
}

// Create a custom BookModel type which wraps the sql.DB connection pool.
type AddressModel struct {
	DB *sqlx.DB
}

// init is invoked before main()
func (env *AddressModel) DeleteAddressById(ctx context.Context, id string) error {
	return nil
}
