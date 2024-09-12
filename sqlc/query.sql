-- name: GetUser :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
    id, name, password, user_role_id
) VALUES (
             $1, $2, $3, $4
         )
    RETURNING *;

-- name: DeleteUser :exec

DELETE FROM users
WHERE id = $1;

