-- name: GetPregnanciesByCowId :many
SELECT * FROM pregnancies
where cowId =$1 order by detectedAt DESC;

-- name: UpsertPregnancy :exec
INSERT INTO pregnancies(cowID,detectedAt,firstDay, lastDay) VALUES (@cowID, @detectedAt, @firstDay, @lastDay)
    ON CONFLICT(cowID,detectedAt)
    DO UPDATE SET
    cowID = @cowID,
    detectedAt = @detectedAt,
    firstDay = @firstDay,
    lastDay = @lastDay;

-- name: DeletePregnancy :exec
DELETE FROM pregnancies
where cowID =$1;
