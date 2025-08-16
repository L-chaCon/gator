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


-- name: GetFeed :one
SELECT * FROM feeds
WHERE url = $1;


-- name: CreateFeedFollow :one
WITH insert_feed_follow AS (
  INSERT INTO feed_follows (id, user_id, feed_id)
  VALUES (
    $1,
    $2,
    $3
  )
  RETURNING *
)
SELECT insert_feed_follow.*, users.name AS user_name, feeds.name AS feed_name
FROM insert_feed_follow
INNER JOIN users 
  ON insert_feed_follow.user_id = users.id
INNER JOIN feeds 
  ON insert_feed_follow.feed_id = feeds.id;


-- name: GetFeedFollowsForUser :many
SELECT feeds.name, feeds.url FROM feed_follows
LEFT JOIN feeds
  ON feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1;
