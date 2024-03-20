-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
	id SERIAL,
	user_id uuid NOT NULL PRIMARY KEY,
	nickname text NOT NULL,
	icon_url text NULL,
	prefecture varchar(2) NULL,
	created_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE users IS 'ユーザー情報';
COMMENT ON COLUMN users.id IS 'ただのシーケンシャルID';
COMMENT ON COLUMN users.user_id IS 'ユーザーID';
COMMENT ON COLUMN users.nickname IS 'ニックネーム';
COMMENT ON COLUMN users.icon_url IS 'アイコンURL';
COMMENT ON COLUMN users.prefecture IS '都道府県番号:JIS X 0401に準拠';
COMMENT ON COLUMN users.created_at IS '作成日時';
COMMENT ON COLUMN users.updated_at IS '更新日時';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
