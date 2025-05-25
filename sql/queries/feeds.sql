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
INNER JOIN users
ON feeds.user_id = users.id;

-- name: GetFeedID :one
SELECT id FROM feeds WHERE url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = $1, updated_at = $1
WHERE id = $2;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at NULLS FIRST
LIMIT 1;