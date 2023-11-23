-- name: GetMutasi :many
SELECT * FROM mutasi 
WHERE no_rekening = $1
ORDER BY waktu;