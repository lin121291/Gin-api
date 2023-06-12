-- name: createUsers :exec
INSERT INTO users (username, password) VALUES ($1,$2);

-- name: checkoutUpdata :many
SELECT * FROM users WHERE username=$1;

-- name: deleteUsers :exec
DELETE FROM users WHERE id = $1

