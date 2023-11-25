-- name: CreateMutasi :one
INSERT INTO mutasi (
    kode_transaksi,
    nominal,
    no_rekening
) VALUES (
    $1, $2, $3
) RETURNING *;