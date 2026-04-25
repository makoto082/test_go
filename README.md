
ginとPostgreSQLで適当なチュートリアルを試した。

テーブル用
テーブル作成
CREATE TABLE albums (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    artist TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL
);

データ作成
INSERT INTO albums (id, title, artist, price) VALUES
('1', 'Blue Train', 'John Coltrane', 56.99),
('2', 'Jeru', 'Gerry Mulligan', 17.99),
('3', 'Sarah Vaughan and Clifford Brown', 'Sarah Vaughan', 39.99);

一覧用curl
curl http://localhost:8080/albums

データ追加用curl
curl http://localhost:8080/albums \
  --include \
  --header "Content-Type: application/json" \
  --request "POST" \
  --data '{"id":"4","title":"New Album","artist":"Me","price":12.34}'

