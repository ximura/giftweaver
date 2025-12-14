-- name: UpsertWish :one
INSERT INTO wishes (room_id, user_id, text)
VALUES ($1, $2, $3)
ON CONFLICT (room_id, user_id)
DO UPDATE SET text = EXCLUDED.text
RETURNING *;

-- name: GetWish :one
SELECT *
FROM wishes
WHERE room_id = $1 AND user_id = $2;

-- name: ListWishesByRoom :many
SELECT *
FROM wishes
WHERE room_id = $1;

-- name: CountMissingWishes :one
SELECT COUNT(*)
FROM room_members rm
LEFT JOIN wishes w
  ON w.room_id = rm.room_id
 AND w.user_id = rm.user_id
WHERE rm.room_id = $1
  AND w.id IS NULL;