-- name: GetPregnanciesByCowId :many
SELECT * FROM pregnancies
where cowId =$1 ;

-- name: UpsertPregnancy :exec
INSERT INTO pregnancies(cowID,detectedAt,firstDay, lastDay) VALUES (@cowID, @detectedAt, @firstDay, @lastDay);


