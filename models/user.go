package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	UserId        *string       `json:"UserId"`
	FirstName     *string       `json:"FirstName"`
	LastName      *string       `json:"LastName"`
	Password      *string       `json:"Password"`
	Email         *string       `json:"Email"`
	Phone         *string       `json:"Phone"`
	Token         *string       `json:"Token"`
	RefreshToken  *string       `json:"RefreshToken"`
	CreatedAt     time.Time     `json:"CreatedAt"`
	UpdatedAt     time.Time     `json:"UpdatedAt"`
	UserCart      []ProductUser `json:"UserCart"`
	AddressDetail []Address     `json:"AddressDetail"`
	OrderStatus   []Order       `json:"OrderStatus"`
}

// Create a custom type which wraps the sql.DB connection pool.
type UserModel struct {
	DB *sqlx.DB
}
