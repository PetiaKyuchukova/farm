-- name: GetPregnanciesByCowId :many
SELECT * FROM pregnancies
where cowId =$1 ;

