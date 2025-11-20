-- name: InsertSession :one
INSERT INTO sessions(user_id, user_agent, ip_address) VALUES($1, $2, $3) ON CONFLICT DO NOTHING RETURNING id;

-- name: InsertRefreshToken :one
INSERT INTO refresh_tokens(session_id, token) VALUES($1, $2) ON CONFLICT DO NOTHING RETURNING id;

-- name: GetRefreshTokenWithSession :one
SELECT rt.token, rt.expires_at as token_expires_at, rt.is_revoked, rt.session_id, s.expires_at as session_expires_at, s.user_id  FROM refresh_tokens as rt INNER JOIN sessions as s ON rt.session_id = s.id WHERE token = $1 ;

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens SET is_revoked = true WHERE token = $1;
