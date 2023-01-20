-- ユーザーのテーブルを作成する
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  uid VARCHAR(255) NOT NULL,
  username VARCHAR(255) NOT NULL,
  mail VARCHAR(255) NOT NULL,
  prefecture_id INTEGER REFERENCES PREFECTURES(id),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ユーザーのデータを作成する
INSERT INTO users (uid, username, mail, prefecture_id) VALUES ('1000000001', 'yamada', 'yamada@example.com', 1);
INSERT INTO users (uid, username, mail, prefecture_id) VALUES ('1000000002', 'suzuki', 'suzuki@example.com', 2);
INSERT INTO users (uid, username, mail, prefecture_id) VALUES ('1000000003', 'tanaka', 'tanaka@example.com', 3);
INSERT INTO users (uid, username, mail, prefecture_id) VALUES ('1000000004', 'sato', 'sato@example.com', 4);
INSERT INTO users (uid, username, mail, prefecture_id) VALUES ('1000000005', 'ito', 'ito@example.com', 5);
