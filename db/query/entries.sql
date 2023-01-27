-- name: CreateEntries :one
INSERT INTO entries (
  account_id,
  amount
) VALUES (
  $1, $2 
) RETURNING *;

-- name: GetEntries :one
SELECT * FROM entries WHERE entries.id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries order by id LIMIT $1 OFFSET $2;

-- name: UpdateEntries :one
UPDATE entries SET amount = $2 WHERE entries.id = $1 RETURNING *;

-- name: DeleteEntries :exec
DELETE FROM entries WHERE entries.id = $1;
