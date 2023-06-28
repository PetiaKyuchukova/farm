-- name: GetInseminationsByCowId :many
SELECT * FROM inseminations
where cowId =$1 ;

