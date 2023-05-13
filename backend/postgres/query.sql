
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

-- name: UpsertNotification :exec
INSERT INTO notifications(cowID,date,type, text) VALUES (@cowID, @date, @type,@text);

-- name: DeleteNotification :exec
DELETE FROM notifications
where id =$1 ;

-- name: GetAllNotification :many
SELECT * FROM notifications
ORDER BY id ASC, type ASC;

-- name: GetNotificationByCowId :one
SELECT * FROM notifications
where cowID =$1 ;

-- name: GetNotificationsByDate :many
SELECT * FROM notifications
where date =$1 ;