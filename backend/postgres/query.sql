
-- name: UpsertCow :exec
INSERT INTO cows (id, birthdate,colour,motherid) VALUES (@id,@birthdate, @colour, @motherid)
ON CONFLICT(id,birthDate)
    DO UPDATE SET
    colour = @colour,
           motherid = @motherid;

-- name: DeleteCow :exec
DELETE FROM cows
where id =$1 ;

-- name: GetAllCows :many
SELECT * FROM cows
ORDER BY id ASC, birthdate ASC;

-- name: GetCowById :one
SELECT * FROM cows
where id =$1 ;