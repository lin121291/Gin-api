-- name: CreateUsers :one
INSERT INTO users (username, password) VALUES ($1,$2) RETURNING *;

-- name: AppendList :one
INSERT INTO reading_list (user_id, news_id) VALUES ($1,$2) RETURNING *;

-- name: CreateNews :one
INSERT INTO news (title, content, date_published) VALUES ($1,$2,$3) RETURNING *;

-- name: GetTopNews :many
SELECT * FROM news Limit 10;

