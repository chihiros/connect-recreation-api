-- 都道府県のテーブルマスターを作成する
CREATE TABLE prefecture (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Path: docker/postgres/seeds/0003-create-prefecture-data.sql
-- 都道府県のデータを作成する
INSERT INTO prefecture (name) VALUES ('北海道');
INSERT INTO prefecture (name) VALUES ('青森県');
INSERT INTO prefecture (name) VALUES ('岩手県');
INSERT INTO prefecture (name) VALUES ('宮城県');
INSERT INTO prefecture (name) VALUES ('秋田県');
INSERT INTO prefecture (name) VALUES ('山形県');
INSERT INTO prefecture (name) VALUES ('福島県');
INSERT INTO prefecture (name) VALUES ('茨城県');
INSERT INTO prefecture (name) VALUES ('栃木県');
INSERT INTO prefecture (name) VALUES ('群馬県');
INSERT INTO prefecture (name) VALUES ('埼玉県');
INSERT INTO prefecture (name) VALUES ('千葉県');
INSERT INTO prefecture (name) VALUES ('東京都');
INSERT INTO prefecture (name) VALUES ('神奈川県');
INSERT INTO prefecture (name) VALUES ('新潟県');
INSERT INTO prefecture (name) VALUES ('富山県');
INSERT INTO prefecture (name) VALUES ('石川県');
INSERT INTO prefecture (name) VALUES ('福井県');
INSERT INTO prefecture (name) VALUES ('山梨県');
INSERT INTO prefecture (name) VALUES ('長野県');
INSERT INTO prefecture (name) VALUES ('岐阜県');
INSERT INTO prefecture (name) VALUES ('静岡県');
INSERT INTO prefecture (name) VALUES ('愛知県');
INSERT INTO prefecture (name) VALUES ('三重県');
INSERT INTO prefecture (name) VALUES ('滋賀県');
INSERT INTO prefecture (name) VALUES ('京都府');
INSERT INTO prefecture (name) VALUES ('大阪府');
INSERT INTO prefecture (name) VALUES ('兵庫県');
INSERT INTO prefecture (name) VALUES ('奈良県');
INSERT INTO prefecture (name) VALUES ('和歌山県');
INSERT INTO prefecture (name) VALUES ('鳥取県');
INSERT INTO prefecture (name) VALUES ('島根県');
INSERT INTO prefecture (name) VALUES ('岡山県');
INSERT INTO prefecture (name) VALUES ('広島県');
INSERT INTO prefecture (name) VALUES ('山口県');
INSERT INTO prefecture (name) VALUES ('徳島県');
INSERT INTO prefecture (name) VALUES ('香川県');
INSERT INTO prefecture (name) VALUES ('愛媛県');
INSERT INTO prefecture (name) VALUES ('高知県');
INSERT INTO prefecture (name) VALUES ('福岡県');
INSERT INTO prefecture (name) VALUES ('佐賀県');
INSERT INTO prefecture (name) VALUES ('長崎県');
INSERT INTO prefecture (name) VALUES ('熊本県');
INSERT INTO prefecture (name) VALUES ('大分県');
INSERT INTO prefecture (name) VALUES ('宮崎県');
INSERT INTO prefecture (name) VALUES ('鹿児島県');
INSERT INTO prefecture (name) VALUES ('沖縄県');
