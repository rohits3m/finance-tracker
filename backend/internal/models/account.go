package models

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Account struct {
	Id          int64     `json:"id"`
	UserId      int64     `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"created_on"`
	UpdatedOn   time.Time `json:"updated_on"`
}

type CreateAccount struct {
	UserId      int64  `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (data *CreateAccount) Validate() error {
	if data.UserId <= 0 {
		return errors.New("user_id is required")
	}
	if data.Title == "" {
		return errors.New("title is required")
	}
	if data.Description == "" {
		return errors.New("description is required")
	}
	return nil
}

type AccountModel struct {
	Db *pgxpool.Pool
}

// Get all accounts
func (model *AccountModel) Get() ([]Account, error) {
	accounts := []Account{}
	sql := "SELECT id, user_id, title, description, created_on, updated_on FROM accounts"

	rows, err := model.Db.Query(context.Background(), sql)
	if err != nil {
		return accounts, err
	}

	for rows.Next() {
		account, err := pgx.RowToStructByName[Account](rows)
		if err != nil {
			continue
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

// Creates a new account and returns it's `id`
func (model *AccountModel) Create(data CreateAccount) (int64, error) {
	if err := data.Validate(); err != nil {
		return -1, err
	}

	sql := "INSERT INTO accounts(user_id, title, description) VALUES($1, $2, $3) RETURNING id"
	args := []any{data.UserId, data.Title, data.Description}

	var insertId int64
	if err := model.Db.QueryRow(context.Background(), sql, args...).Scan(&insertId); err != nil {
		return -1, err
	}

	return insertId, nil
}
