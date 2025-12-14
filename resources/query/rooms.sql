-- name: CreateRoom :one
INSERT INTO rooms (name, description, created_by)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetRoomByID :one
SELECT *
FROM rooms
WHERE id = $1;

-- name: GetRoomsByCreator :many
SELECT *
FROM rooms
WHERE created_by = $1;

-- name: ListRooms :many
SELECT *
FROM rooms
ORDER BY created_at DESC;
