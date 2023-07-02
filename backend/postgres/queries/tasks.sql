-- name: UpsertTasks :exec
INSERT INTO tasks(cowID,date,type, text) VALUES (@cowID, @date, @type,@text);

-- name: DeleteTask :exec
DELETE FROM tasks
where cowID =$1 ;

-- name: GetAllTasks :many
SELECT * FROM tasks
ORDER BY cowID ASC, date ASC;

-- name: GetTaskByCowId :many
SELECT * FROM tasks
where cowID =$1 ;

-- name: GetTasksByDate :many
SELECT * FROM tasks
where date =$1 ;