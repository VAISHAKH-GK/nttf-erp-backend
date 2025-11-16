-- name: InsertAccountType :one
INSERT INTO account_types(name, is_active) VALUES($1, $2) ON CONFLICT DO NOTHING RETURNING id; 

-- name: InsertUserAccountType :exec
INSERT INTO user_account_types(user_id, account_type_id) VALUES($1, $2) ON CONFLICT DO NOTHING;
