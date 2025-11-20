-- name: InsertSession :one
INSERT INTO sessions(user_id, user_agent, ip_address) VALUES($1, $2, $3) ON CONFLICT DO NOTHING RETURNING id;

-- name: InsertRefreshToken :one
INSERT INTO refresh_tokens(session_id, token) VALUES($1, $2) ON CONFLICT DO NOTHING RETURNING id;

-- name: GetRefreshToken :one
SELECT rt.session_id, rt.token, rt.expires_at, rt.is_revoked FROM refresh_tokens as rt INNER JOIN sessions as s ON rt.session_id = s.id WHERE rt.token = $1; 
