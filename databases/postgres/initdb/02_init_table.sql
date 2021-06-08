set client_encoding = 'UTF8';

-- 01_init_database.sqlで作成したデータベースに接続
\c test_db;

-- 作成済みののテーブルを削除しておく
DROP TABLE IF EXISTS persons;

-- テーブルを作成する
CREATE TABLE persons (
  name VARCHAR(255),
  age INTEGER
);

-- 何かデータを入れておく
INSERT INTO persons (name, age) VALUES ('John', 30)
