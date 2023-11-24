-- name: UpdateSaldo :one
UPDATE akun
SET saldo = $2
WHERE no_rekening = $1
RETURNING *;
-- name: GetAkun :one
SELECT * FROM akun
WHERE no_rekening = $1;