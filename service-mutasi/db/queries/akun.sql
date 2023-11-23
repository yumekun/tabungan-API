-- name: UpdateSaldo :one
UPDATE akun
SET saldo = $2
WHERE no_rekening = $1
RETURNING *;