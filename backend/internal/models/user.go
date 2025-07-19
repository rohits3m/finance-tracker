package models

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Id         int64  `json:"id"`
	FirebaseId string `json:"firebase_id"`
	FullName   string `json:"full_name"`
	Email      string `json:"email"`
	CreatedOn  string `json:"created_on"`
	UpdatedOn  string `json:"updated_on"`
}

type CreateUser struct {
	FirebaseId string `json:"firebase_id"`
	FullName   string `json:"full_name"`
	Email      string `json:"email"`
}

func (data *CreateUser) Validate() error {
	if data.FirebaseId == "" {
		return errors.New("firebase_id is required")
	}

	if data.FullName == "" {
		return errors.New("full_name is required")
	}

	if data.Email == "" {
		return errors.New("email is required")
	}

	return nil
}

type UserModel struct {
	Db *pgxpool.Pool
}

// Creates a new user and returns the `id`
func (model *UserModel) Create(data CreateUser) (int64, error) {
	if err := data.Validate(); err != nil {
		return -1, err
	}

	sql := "INSERT INTO users(firebase_id, full_name, email) VALUES($1, $2, $3) RETURNING id"
	args := []any{data.FirebaseId, data.FullName, data.Email}

	var insertId int64
	if err := model.Db.QueryRow(context.Background(), sql, args...).Scan(&insertId); err != nil {
		return -1, err
	}

	return insertId, nil
}

// Creates a new user and returns the `id`
func (model *UserModel) GetByFirebaseId(firebaseId string) (User, error) {
	sql := "SELECT id, firebase_id, full_name, email, created_on, updated_on FROM users WHERE firebase_id=$1"
	args := []any{firebaseId}

	var user User
	err := model.Db.QueryRow(context.Background(), sql, args...).Scan(
		&user.Id,
		&user.FirebaseId,
		&user.FullName,
		&user.Email,
		&user.CreatedOn,
		&user.UpdatedOn,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return user, errors.New("user not found")
		}

		return user, err
	}

	return user, nil
}
