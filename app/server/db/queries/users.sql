-- name: GetUser :one
SELECT id, email, name, password_hash, created_at, updated_at
FROM users
WHERE id = $1;
