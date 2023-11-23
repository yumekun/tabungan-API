  -- name: CreateNasabah :one
INSERT INTO nasabah (
    nama,
    nik,
    no_hp
) VALUES (
    $1, $2, $3
) RETURNING *;