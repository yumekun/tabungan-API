-- name: CreateAkun :one
INSERT INTO akun (
    nasabah_id,
    no_rekening,
    saldo
) VALUES (
    $1, $2, 0
) RETURNING *;

-- name: GetAkun :one
SELECT * FROM akun
WHERE no_rekening = $1;
