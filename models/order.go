package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Order struct {
	OrderId        *string       `json:"OrderId"`
	OrderCart      []ProductUser `json:"OrderCart"`
	OrderAt        time.Time     `json:"OrderAt"`
	Price          int64         `json:"Price"`
	Discount       *int64        `json:"Discount"`
	Payment_Method Payment
}

type Payment struct {
	Digital bool `json:"Digital"`
	CDD     bool `json:"CDD"`
}

// Create a custom BookModel type which wraps the sql.DB connection pool.
type OrderModel struct {
	DB *sqlx.DB
}
