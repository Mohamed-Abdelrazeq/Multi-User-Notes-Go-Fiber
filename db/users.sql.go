// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (email, password) VALUES ($1, $2) RETURNING (id, email)
`

type CreateUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.Password)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, password, created_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}
