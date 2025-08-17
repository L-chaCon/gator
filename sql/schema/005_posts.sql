-- +goose Up
CREATE TABLE posts (
  id UUID PRIMARY KEY,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  title TEXT,
  url TEXT UNIQUE,
  description TEXT,
  published_at TIMESTAMPTZ,
  feed_id UUID NOT NULL,
  CONSTRAINT fk_feed
    FOREIGN KEY (feed_id)
    REFERENCES feeds(id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;
