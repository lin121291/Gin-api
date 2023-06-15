// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: create.sql

package db

import (
	"context"
	"database/sql"
)

const appendList = `-- name: AppendList :one
INSERT INTO reading_list (user_id, news_id) VALUES ($1,$2) RETURNING id, user_id, news_id
`

type AppendListParams struct {
	UserID sql.NullInt32 `json:"user_id"`
	NewsID sql.NullInt32 `json:"news_id"`
}

func (q *Queries) AppendList(ctx context.Context, arg AppendListParams) (ReadingList, error) {
	row := q.db.QueryRowContext(ctx, appendList, arg.UserID, arg.NewsID)
	var i ReadingList
	err := row.Scan(&i.ID, &i.UserID, &i.NewsID)
	return i, err
}

const createNews = `-- name: CreateNews :one
INSERT INTO news (title, content, date_published) VALUES ($1,$2,$3) RETURNING id, title, content, date_published
`

type CreateNewsParams struct {
	Title         string         `json:"title"`
	Content       sql.NullString `json:"content"`
	DatePublished sql.NullTime   `json:"date_published"`
}

func (q *Queries) CreateNews(ctx context.Context, arg CreateNewsParams) (News, error) {
	row := q.db.QueryRowContext(ctx, createNews, arg.Title, arg.Content, arg.DatePublished)
	var i News
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.DatePublished,
	)
	return i, err
}

const createUsers = `-- name: CreateUsers :one
INSERT INTO users (username, password) VALUES ($1,$2) RETURNING id, username, password
`

type CreateUsersParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (q *Queries) CreateUsers(ctx context.Context, arg CreateUsersParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUsers, arg.Username, arg.Password)
	var i User
	err := row.Scan(&i.ID, &i.Username, &i.Password)
	return i, err
}

const getTopNews = `-- name: GetTopNews :many
SELECT id, title, content, date_published FROM news Limit 10
`

func (q *Queries) GetTopNews(ctx context.Context) ([]News, error) {
	rows, err := q.db.QueryContext(ctx, getTopNews)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []News
	for rows.Next() {
		var i News
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.DatePublished,
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
