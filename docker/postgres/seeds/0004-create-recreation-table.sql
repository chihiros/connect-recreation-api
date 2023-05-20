-- ユーザーのテーブルを作成する
CREATE TABLE recreations (
  id SERIAL PRIMARY KEY,
  user_id UUID NOT NULL,
  uuid UUID NOT NULL UNIQUE,
  genre JSONB,
  title VARCHAR(255),
  target_number INT,
  required_time INT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ユーザーのデータを作成する
INSERT INTO recreations (user_id, uuid, genre, title, target_number, required_time) VALUES ('1000000001', 'yamada', '[1, 2]', 'Example Title 1', 10, 30);
INSERT INTO recreations (user_id, uuid, genre, title, target_number, required_time) VALUES ('1000000002', 'suzuki', '[3, 4]', 'Example Title 2', 20, 40);
INSERT INTO recreations (user_id, uuid, genre, title, target_number, required_time) VALUES ('1000000003', 'tanaka', '[5, 6]', 'Example Title 3', 30, 50);
