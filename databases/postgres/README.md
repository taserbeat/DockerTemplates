# PostgreSQL

PostgreSQL の Docker コンテナを起動する compose を紹介する。

# 構成

postgres コンテナと pgweb (postgres の Web GUI ツール)コンテナを起動する。  
postres には[./initdb](./initdb/)に登録された.sql ファイルが番号順に実行される。

- postgres

| 内容                                   | 値         |
| -------------------------------------- | ---------- |
| ポートフォワード (ホスト側:コンテナ側) | 15432:5432 |
| postgres ユーザー名                    | root       |
| postgres ユーザーのパスワード          | postgres   |

- pgweb

| 内容                                   | 値                    |
| -------------------------------------- | --------------------- |
| ポートフォワード (ホスト側:コンテナ側) | 8081:8081             |
| アクセス URL                           | http://localhost:8081 |

# 実行方法

- コンテナを起動

```bash
docker-compose up -d --build
```

- postgres コンテナにログイン

```bash
docker-compose exec postgres bash

# postgresコンテナからpostgresエンジンにログイン
psql
```

- ホストから psql クライアントで postgres にログイン

```bash
psql -h localhost -p 15432 -U root -d test_db
# Password for user root: postgres
```

- コンテナを初期化する

```bash
docker-compose down --volumes && docker-compose up -d
```

- コンテナを終了する (ネットワークやイメージも削除する)

```bash
docker-compose down --rmi all --volumes
```
