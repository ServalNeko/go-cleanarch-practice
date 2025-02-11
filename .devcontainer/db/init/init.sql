
-- データベース作成
CREATE DATABASE sns;

-- 作成したデータベースに接続
\c sns

-- Usersテーブル
CREATE TABLE users (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    type INT NOT NULL
);

-- Circlesテーブル
CREATE TABLE circles (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    owner_id VARCHAR(255) NULL REFERENCES users(id) ON DELETE SET NULL
);

-- UserCirclesテーブル
CREATE TABLE user_circles (
    user_id VARCHAR(255) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    circle_id VARCHAR(255) NOT NULL REFERENCES circles(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, circle_id)
);

-- インデックス作成 (circles.owner_id)
CREATE INDEX ix_circles_owner_id ON circles (owner_id);

-- インデックス作成 (user_circles.circle_id)
CREATE INDEX ix_user_circles_circle_id ON user_circles (circle_id);
