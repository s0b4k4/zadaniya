-- Task 2: sqlc queries
-- name: GetUserDocuments :many
SELECT * FROM documents WHERE user_id = $1;

-- name: CreateUser :one
INSERT INTO users (name) VALUES ($1) RETURNING *;
