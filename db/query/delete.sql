-- name: DeleteList :exec
DELETE FROM reading_list WHERE id = $1;

-- name: DeleteUsers :exec
DELETE FROM users WHERE id = $1;

-- name: DeleteNews :exec
DELETE FROM news WHERE id = $1;