-- name: UpsertTasks :exec
INSERT INTO tasks(cowID,date,type, text,done) VALUES (@cowID, @date, @type,@text, @done);

-- name: DeleteTask :exec
DELETE FROM tasks
where cowID =$1 ;

-- name: GetAllTasks :many
SELECT * FROM tasks
ORDER BY cowID ASC, date ASC;

-- name: UpdateTaskStatus :exec
UPDATE tasks
SET done = $1
where cowID = $2 and date = $3;

-- name: GetTasksByDate :many
SELECT * FROM tasks
where date =$1 ;