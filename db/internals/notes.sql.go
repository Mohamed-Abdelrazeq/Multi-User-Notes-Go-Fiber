// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: notes.sql

package internals

import (
	"context"
)

const addToFavourites = `-- name: AddToFavourites :one
UPDATE notes
set is_favourite = true
WHERE  user_id = $1 AND id = $2
RETURNING id, title, content, user_id, is_favourite, created_at, image_url
`

type AddToFavouritesParams struct {
	UserID int32 `json:"user_id"`
	ID     int32 `json:"id"`
}

func (q *Queries) AddToFavourites(ctx context.Context, arg AddToFavouritesParams) (Note, error) {
	row := q.db.QueryRowContext(ctx, addToFavourites, arg.UserID, arg.ID)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.UserID,
		&i.IsFavourite,
		&i.CreatedAt,
		&i.ImageUrl,
	)
	return i, err
}

const createNote = `-- name: CreateNote :one
INSERT INTO notes (user_id, title, content, is_favourite, image_url) VALUES ($1, $2, $3, $4, $5) RETURNING id, title, content, user_id, is_favourite, created_at, image_url
`

type CreateNoteParams struct {
	UserID      int32  `json:"user_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	IsFavourite bool   `json:"is_favourite"`
	ImageUrl    string `json:"image_url"`
}

func (q *Queries) CreateNote(ctx context.Context, arg CreateNoteParams) (Note, error) {
	row := q.db.QueryRowContext(ctx, createNote,
		arg.UserID,
		arg.Title,
		arg.Content,
		arg.IsFavourite,
		arg.ImageUrl,
	)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.UserID,
		&i.IsFavourite,
		&i.CreatedAt,
		&i.ImageUrl,
	)
	return i, err
}

const deleteNote = `-- name: DeleteNote :exec
DELETE FROM notes WHERE user_id = $1 AND id = $2
`

type DeleteNoteParams struct {
	UserID int32 `json:"user_id"`
	ID     int32 `json:"id"`
}

func (q *Queries) DeleteNote(ctx context.Context, arg DeleteNoteParams) error {
	_, err := q.db.ExecContext(ctx, deleteNote, arg.UserID, arg.ID)
	return err
}

const getAllNotes = `-- name: GetAllNotes :many
SELECT id, title, content, user_id, is_favourite, created_at, image_url FROM notes WHERE user_id = $1
`

func (q *Queries) GetAllNotes(ctx context.Context, userID int32) ([]Note, error) {
	rows, err := q.db.QueryContext(ctx, getAllNotes, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Note
	for rows.Next() {
		var i Note
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.UserID,
			&i.IsFavourite,
			&i.CreatedAt,
			&i.ImageUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNoteById = `-- name: GetNoteById :one
SELECT id, title, content, user_id, is_favourite, created_at, image_url FROM notes WHERE user_id = $1 AND id = $2
`

type GetNoteByIdParams struct {
	UserID int32 `json:"user_id"`
	ID     int32 `json:"id"`
}

func (q *Queries) GetNoteById(ctx context.Context, arg GetNoteByIdParams) (Note, error) {
	row := q.db.QueryRowContext(ctx, getNoteById, arg.UserID, arg.ID)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.UserID,
		&i.IsFavourite,
		&i.CreatedAt,
		&i.ImageUrl,
	)
	return i, err
}

const listFavourites = `-- name: ListFavourites :many
SELECT id, title, content, user_id, is_favourite, created_at, image_url FROM notes WHERE user_id = $1 AND is_favourite = true
`

func (q *Queries) ListFavourites(ctx context.Context, userID int32) ([]Note, error) {
	rows, err := q.db.QueryContext(ctx, listFavourites, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Note
	for rows.Next() {
		var i Note
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.UserID,
			&i.IsFavourite,
			&i.CreatedAt,
			&i.ImageUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeFromFavourites = `-- name: RemoveFromFavourites :one
UPDATE notes
set is_favourite = false
WHERE  user_id = $1 AND id = $2
RETURNING id, title, content, user_id, is_favourite, created_at, image_url
`

type RemoveFromFavouritesParams struct {
	UserID int32 `json:"user_id"`
	ID     int32 `json:"id"`
}

func (q *Queries) RemoveFromFavourites(ctx context.Context, arg RemoveFromFavouritesParams) (Note, error) {
	row := q.db.QueryRowContext(ctx, removeFromFavourites, arg.UserID, arg.ID)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.UserID,
		&i.IsFavourite,
		&i.CreatedAt,
		&i.ImageUrl,
	)
	return i, err
}

const updateNote = `-- name: UpdateNote :one
UPDATE notes 
SET title = $3, content = $4
WHERE  user_id = $1 AND id = $2
RETURNING id, title, content, user_id, is_favourite, created_at, image_url
`

type UpdateNoteParams struct {
	UserID  int32  `json:"user_id"`
	ID      int32  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (q *Queries) UpdateNote(ctx context.Context, arg UpdateNoteParams) (Note, error) {
	row := q.db.QueryRowContext(ctx, updateNote,
		arg.UserID,
		arg.ID,
		arg.Title,
		arg.Content,
	)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.UserID,
		&i.IsFavourite,
		&i.CreatedAt,
		&i.ImageUrl,
	)
	return i, err
}
