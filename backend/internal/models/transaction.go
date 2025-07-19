package models

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Transaction struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	AccountId int64     `json:"account_id"`
	Type      string    `json:"type"`
	Amount    float64   `json:"amount"`
	Remarks   string    `json:"remarks"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

type CreateTransaction struct {
	UserId    int64   `json:"user_id"`
	AccountId int64   `json:"account_id"`
	Type      string  `json:"type"`
	Amount    float64 `json:"amount"`
	Remarks   string  `json:"remarks"`
}

func (data *CreateTransaction) Validate() error {
	if data.UserId <= 0 {
		return errors.New("user_id is required")
	}
	if data.AccountId <= 0 {
		return errors.New("account_id is required")
	}
	if data.Type == "" {
		return errors.New("type is required")
	}
	if data.Amount <= 0.0 {
		return errors.New("amount is required")
	}
	if data.Remarks == "" {
		return errors.New("remarks is required")
	}
	return nil
}

type TransactionModel struct {
	Db *pgxpool.Pool
}

func (model *TransactionModel) Get() ([]Transaction, error) {
	transactions := []Transaction{}
	sql := "SELECT id, user_id, account_id, type, amount, remarks, created_on, updated_on FROM transactions"

	rows, err := model.Db.Query(context.Background(), sql)
	if err != nil {
		return transactions, err
	}

	for rows.Next() {
		transaction, err := pgx.RowToStructByName[Transaction](rows)
		if err != nil {
			continue
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (model *TransactionModel) Create(data CreateTransaction) (int64, error) {
	if err := data.Validate(); err != nil {
		return -1, err
	}

	sql := "INSERT INTO transactions(user_id, account_id, type, amount, remarks) VALUES($1, $2, $3, $4, $5) RETURNING id"
	args := []any{data.UserId, data.AccountId, data.Type, data.Amount, data.Remarks}

	var insertId int64
	if err := model.Db.QueryRow(context.Background(), sql, args...).Scan(&insertId); err != nil {
		return -1, err
	}

	return insertId, nil
}
