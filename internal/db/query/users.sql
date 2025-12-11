-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: InsertUser :one
INSERT INTO users(email, name, password) VALUES($1, $2, $3) ON CONFLICT DO NOTHING RETURNING id;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;
