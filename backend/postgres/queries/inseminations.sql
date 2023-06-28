-- name: GetInseminationsByCowId :many
SELECT * FROM inseminations
where cowId =$1 ;

-- name: UpsertInsemination :exec
INSERT INTO inseminations(cowID,date,breed, isArtificial) VALUES (@cowID, @date, @breed,@isArtificial)
    ON CONFLICT(cowID,date)
    DO UPDATE SET
    cowID = @cowID,
    date = @date,
    breed = @breed,
    isArtificial = @isArtificial;


