-- ユーザーのテーブルを作成する
CREATE TABLE recreations (
  id SERIAL PRIMARY KEY,
  user_id UUID NOT NULL,
  uuid UUID NOT NULL UNIQUE,
  genre JSONB,
  title VARCHAR(255),
  content TEXT,
  target_number INT,
  required_time INT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ユーザーのデータを作成する
INSERT INTO recreations (user_id, uuid, genre, title, content, target_number, required_time) VALUES ('f8ee08c0-cb6b-4e9c-9744-4c02fcc88fd6', 'c5c29ef8-dd34-45f8-8013-e80840356ea4', '[1, 2]', 'Example Title 1', 'Content Example 1', 10, 30);
INSERT INTO recreations (user_id, uuid, genre, title, content, target_number, required_time) VALUES ('afea09a5-bbe3-45c1-aa76-2d51944b8683', '0c2643c1-4b20-40c6-9163-1d8567954334', '[3, 4]', 'Example Title 2', 'Content Example 2', 20, 40);
INSERT INTO recreations (user_id, uuid, genre, title, content, target_number, required_time) VALUES ('0f8a75a9-3079-4417-bec2-6105a67042f4', 'df3ed291-32b3-44d3-86c2-e01537e066c6', '[5, 6]', 'Example Title 3', 'Content Example 3', 30, 50);
