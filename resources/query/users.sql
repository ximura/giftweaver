-- name: CreateUser :one
INSERT INTO users (telegram_user_id, username)
VALUES ($1, $2)
ON CONFLICT (telegram_user_id)
DO UPDATE SET username = EXCLUDED.username
RETURNING *;

-- name: GetUserByTelegramID :one
SELECT *
FROM users
WHERE telegram_user_id = $1;
