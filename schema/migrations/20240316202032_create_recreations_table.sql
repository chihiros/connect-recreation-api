-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS recreations (
  id SERIAL,
  recreation_id UUID NOT NULL PRIMARY KEY,
  genre JSONB,
  title VARCHAR(255),
  content TEXT,
  target_number INT,
  required_time INT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  youtube_id TEXT,
  publish BOOLEAN NOT NULL DEFAULT FALSE,
  published_at TIMESTAMP DEFAULT NULL,
  drafted_at TIMESTAMP DEFAULT NULL,
  user_id UUID NOT NULL
);

ALTER TABLE recreations
ADD CONSTRAINT fk_recreations_user_id
FOREIGN KEY (user_id)
REFERENCES users (user_id)
ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE recreations;
-- +goose StatementEnd
