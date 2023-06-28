
-- name: UpsertCow :exec
INSERT INTO cows (id, birthdate,gender, breed,colour,motherId,fatherId,fatherBreed,isPregnant,ovulation) VALUES (@id, @birthdate, @gender, @breed, @colour, @motherId, @fatherId, @fatherBreed, @isPregnant, @ovulation)
    ON CONFLICT(id)
    DO UPDATE SET
    id = @id,
    birthdate = @birthdate,
    gender = @gender,
    breed = @breed,
    colour = @colour,
    motherId = @motherId,
    fatherId = @fatherId,
    fatherBreed = @fatherBreed,
    isPregnant = @isPregnant,
    ovulation = @ovulation;


-- name: DeleteCow :exec
DELETE FROM cows
where id =$1 ;

-- name: GetAllCows :many
SELECT * FROM cows
ORDER BY id ASC, birthdate ASC;

-- name: GetCowById :one
SELECT * FROM cows
where id =$1 ;