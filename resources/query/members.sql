-- name: JoinRoom :exec
INSERT INTO room_members (room_id, user_id)
VALUES ($1, $2)
ON CONFLICT DO NOTHING;

-- name: LeaveRoom :exec
DELETE FROM room_members
WHERE room_id = $1 AND user_id = $2;

-- name: ListRoomMembers :many
SELECT u.*
FROM users u
JOIN room_members rm ON rm.user_id = u.id
WHERE rm.room_id = $1;

-- name: IsUserInRoom :one
SELECT EXISTS (
    SELECT 1
    FROM room_members
    WHERE room_id = $1 AND user_id = $2
);
