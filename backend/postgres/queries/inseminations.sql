-- name: GetInseminationsByCowId :many
SELECT * FROM inseminations
where cowId =$1 ;

-- name: UpsertInsemination :exec
INSERT INTO inseminations(cowID,date,breed, isArtificial) VALUES (@cowID, @date, @breed,@isArtificial);


