-- レクリエーションのテーブルを作成する
CREATE TABLE recreations (
  id SERIAL PRIMARY KEY,
  user_id UUID NOT NULL,
  recreation_id UUID NOT NULL UNIQUE,
  genre JSONB,
  title VARCHAR(255),
  content TEXT,
  target_number INT,
  required_time INT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE recreations
ADD CONSTRAINT fk_recreations_user_id
FOREIGN KEY (user_id)
REFERENCES profiles (uuid)
ON DELETE CASCADE;

-- レクリエーションのデータを作成する
INSERT INTO recreations (user_id, recreation_id, genre, title, content, target_number, required_time) VALUES ('c2cc015d-5753-4132-86fe-624787ae49df', 'c5c29ef8-dd34-45f8-8013-e80840356ea4', '[1, 2]', 'Example Title 1', 'Content Example 1', 10, 30);
INSERT INTO recreations (user_id, recreation_id, genre, title, content, target_number, required_time) VALUES ('87cab8ea-2252-4a4d-bc15-3fe6192bbf5d', '0c2643c1-4b20-40c6-9163-1d8567954334', '[3, 4]', 'Example Title 2', 'Content Example 2', 20, 40);
INSERT INTO recreations (user_id, recreation_id, genre, title, content, target_number, required_time) VALUES ('c2cc015d-5753-4132-86fe-624787ae49df', 'df3ed291-32b3-44d3-86c2-e01537e066c6', '[5, 6]', 'Example Title 3', 'Content Example 3', 30, 50);

ALTER TABLE recreations
ADD COLUMN youtube_id text;

-- publish Booleanのカラムを追加する
ALTER TABLE recreations
ADD COLUMN publish BOOLEAN NOT NULL DEFAULT FALSE;

UPDATE
  recreations
SET
  publish = TRUE
WHERE
  recreation_id = 'c5c29ef8-dd34-45f8-8013-e80840356ea4'
  OR recreation_id = '0c2643c1-4b20-40c6-9163-1d8567954334'
  OR recreation_id = 'df3ed291-32b3-44d3-86c2-e01537e066c6'
;
