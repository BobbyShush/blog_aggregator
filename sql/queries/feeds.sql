-- name: CreateFeed :one
INSERT into feeds(id, created_at, updated_at, name, url, user_id)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.id, feeds.name, feeds.url, users.name AS creator_name
FROM feeds
LEFT JOIN users
ON feeds.user_id = users.id;