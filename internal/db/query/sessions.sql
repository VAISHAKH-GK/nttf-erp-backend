-- name: InsertSession :one
INSERT INTO sessions(user_id, user_agent, ip_address) VALUES($1, $2, $3) ON CONFLICT DO NOTHING RETURNING id;

-- name: InsertRefreshToken :one
INSERT INTO refresh_tokens(session_id, token) VALUES($1, $2) ON CONFLICT DO NOTHING RETURNING id;