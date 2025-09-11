-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: InsertUser :exec
INSERT INTO users(email, username, password) VALUES($1, $2, $3) ON CONFLICT DO NOTHING;
