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
WITH rows AS (
DELETE FROM users
WHERE users.id = $1
    RETURNING *
)
SELECT count(*) FROM rows;





-- name: GetNew :one
SELECT * FROM news
WHERE id = $1
LIMIT 1;

-- name: GetNews :many
SELECT * FROM news
ORDER BY title;

-- name: CreateNew :one
INSERT INTO news (
    id, user_id, title, description, name_image
) VALUES (
             $1, $2, $3, $4, $5
         )
    RETURNING *;

-- name: DeleteNews :exec
WITH rows AS (
DELETE FROM news
WHERE news.id = $1
    RETURNING *
)
SELECT count(*) FROM rows;





-- name: GetTag :one
SELECT * FROM tags
WHERE id = $1
    LIMIT 1;

-- name: GetTags :many
SELECT * FROM tags
ORDER BY name;

-- name: CreateTag :one
INSERT INTO tags (
    id, name
) VALUES (
             $1, $2
         )
    RETURNING *;

-- name: DeleteTags :exec
WITH rows AS (
DELETE FROM tags
WHERE id = $1
    RETURNING *
)
SELECT count(*) FROM rows;





-- name: GetComment :one
SELECT * FROM comments
WHERE id = $1
    LIMIT 1;

-- name: GetComments :many
SELECT * FROM comments
ORDER BY name;

-- name: CreateComment :one
INSERT INTO comments (
    id, user_id, news_id, name, description
) VALUES (
             $1, $2, $3, $4, $5
         )
    RETURNING *;

-- name: DeleteComment :exec
WITH rows AS (
DELETE FROM comments
WHERE id = $1
    RETURNING *
)
SELECT count(*) FROM rows;

