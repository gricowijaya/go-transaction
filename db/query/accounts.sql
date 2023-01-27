-- name: CreateAccount :one
INSERT INTO accounts (
  owner,
  balance,
  currency
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts WHERE accounts.id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts WHERE accounts.id = $1 LIMIT 1 FOR NO KEY UPDATE ;

-- name: AddAccountBalance :one
UPDATE accounts SET balance = balance + sqlc.arg(amount) WHERE id = sqlc.arg(id) RETURNING *;

-- name: ListAccounts :many
SELECT * FROM accounts order by id LIMIT $1 OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts SET balance = $2 WHERE accounts.id = $1 RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE accounts.id = $1;
