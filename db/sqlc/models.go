// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"database/sql"
)

type News struct {
	ID            int32          `json:"id"`
	Title         string         `json:"title"`
	Content       sql.NullString `json:"content"`
	DatePublished sql.NullTime   `json:"date_published"`
}

type ReadingList struct {
	ID     int32         `json:"id"`
	UserID sql.NullInt32 `json:"user_id"`
	NewsID sql.NullInt32 `json:"news_id"`
}

type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}