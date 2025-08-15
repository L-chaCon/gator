-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetFeeds :many
SELECT
  feeds.id,
  feeds.name,
  feeds.url,
  users.name AS user_name,
  feeds.created_at,
  feeds.updated_at
FROM feeds
LEFT JOIN users
  ON users.id = feeds.user_id;
