package models

import (
	"context"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	UserId       int64     `json:"UserId" db:"userid"`
	FirstName    *string   `json:"FirstName" db:"firstname" validate:"required"`
	LastName     *string   `json:"LastName" db:"lastname" validate:"required"`
	Password     *string   `json:"Password" db:"password" validate:"required"`
	Email        *string   `json:"Email" db:"email" validate:"required"`
	Phone        *string   `json:"Phone" db:"phone" validate:"required"`
	Token        *string   `json:"Token" db:"token"`
	RefreshToken *string   `json:"RefreshToken" db:"refreshtoken"`
	CreatedAt    time.Time `json:"CreatedAt" db:"createdat"`
	UpdatedAt    time.Time `json:"UpdatedAt" db:"updatedat"`
}

// Create a custom type which wraps the sql.DB connection pool.
type UserModel struct {
	DB *sqlx.DB
}

var (
	ErrQueryExec  = errors.New("error with executing query")
	ErrScanRow    = errors.New("can not scan row to model")
	ErrLastInsert = errors.New("can not see last user id")
)

func (user *UserModel) FindAllUserByEmail(ctx context.Context, email string) ([]User, error) {
	rows, err := user.DB.QueryxContext(ctx, "SELECT * FROM USERS WHERE email = $1", email)
	if err != nil {
		return nil, ErrQueryExec
	}
	var users []User
	for rows.Next() {
		var user User
		err := rows.StructScan(&user)
		if err != nil {
			return nil, ErrScanRow
		}
		users = append(users, user)
	}
	return users, nil
}

func (user *UserModel) FindAllUserByPhone(ctx context.Context, phone string) ([]User, error) {
	rows, err := user.DB.QueryxContext(ctx, "SELECT * FROM USERS WHERE phone = $1", phone)
	if err != nil {
		return nil, ErrQueryExec
	}
	var users []User
	for rows.Next() {
		var user User
		err := rows.StructScan(&user)
		if err != nil {
			return nil, ErrScanRow
		}
		users = append(users, user)
	}
	return users, nil
}

func (user *UserModel) InsertUser(ctx context.Context, model User) (int64, error) {
	query := "INSERT INTO people (firstname, lastname, password, email, phone, token, refreshtoken,  createdat, updatedat) VALUES (:firstname, :lastname, :password, :email, :phone, :token, :refreshtoken, :createdat, :updatedat)"
	result, err := user.DB.NamedExecContext(ctx, query, user)
	if err != nil {
		return -1, ErrQueryExec
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return -1, ErrLastInsert
	}
	return userId, nil
}
