-- name: GetAccountByEmail :one
SELECT id, password_hash, is_locked
FROM accounts
WHERE email = @email;
