-- name: CreateAssignment :exec
INSERT INTO assignments (
    room_id,
    giver_user_id,
    receiver_user_id,
    wish_id
)
VALUES ($1, $2, $3, $4);


-- name: ListAssignmentsForRoom :many
SELECT
    a.giver_user_id,
    a.receiver_user_id,
    w.text AS wish_text
FROM assignments a
JOIN wishes w ON w.id = a.wish_id
WHERE a.room_id = $1;

-- name: HasAssignments :one
SELECT EXISTS (
    SELECT 1
    FROM assignments
    WHERE room_id = $1
);