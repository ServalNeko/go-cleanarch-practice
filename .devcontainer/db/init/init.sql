
-- データベース作成
CREATE DATABASE sns;

-- 作成したデータベースに接続
\c sns

-- Usersテーブル
CREATE TABLE Users (
    Id VARCHAR(255) NOT NULL PRIMARY KEY,
    Name VARCHAR(20) NOT NULL,
    Type INT NOT NULL
);

-- Circlesテーブル
CREATE TABLE Circles (
    Id VARCHAR(255) NOT NULL PRIMARY KEY,
    Name VARCHAR(20) NOT NULL,
    OwnerId VARCHAR(255) NULL REFERENCES Users(Id) ON DELETE SET NULL
);

-- UserCirclesテーブル
CREATE TABLE UserCircles (
    UserId VARCHAR(255) NOT NULL REFERENCES Users(Id) ON DELETE CASCADE,
    CircleId VARCHAR(255) NOT NULL REFERENCES Circles(Id) ON DELETE CASCADE,
    PRIMARY KEY (UserId, CircleId)
);

-- インデックス作成 (Circles.OwnerId)
CREATE INDEX IX_Circles_OwnerId ON Circles (OwnerId);

-- インデックス作成 (UserCircles.CircleId)
CREATE INDEX IX_UserCircles_CircleId ON UserCircles (CircleId);
