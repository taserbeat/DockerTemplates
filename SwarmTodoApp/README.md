# SwarmTodoApp

Docker Swarm を使って実用的な構成の TODO アプリを作る。

# アプリケーションの仕様

- TODO を登録・更新・削除できる
- 登録されている TODO の一覧を表示できる
- ブラウザから利用できる Web アプリケーションとして構築する
- ブラウザ以外のプラットフォームからでも利用できるように、JSON API のエンドポイントも作成する

# アーキテクチャ

![アーキテクチャ](./images/アーキテクチャ.png)

# アーキテクチャの構成要素

| イメージ名 | 用途                                   | Service 名                | Stack 名             |
| ---------- | -------------------------------------- | ------------------------- | -------------------- |
| MySQL      | データストア                           | mysql_master, mysql_slave | MySQL                |
| API        | データストアを操作する API サーバ      | app_api                   | Application          |
| Web        | ビューを生成するアプリケーションサーバ | frontend_web              | Frontend             |
| Nginx      | プロキシサーバ                         | app_nginx, frontend_nginx | Application Frontend |

# 用語

| 名称    | 役割                                                                    | 対応するコマンド |
| ------- | ----------------------------------------------------------------------- | ---------------- |
| Compose | 複数のコンテナを使う Docker アプリケーションの管理 (主にシングルホスト) | `docker-compose` |
| Swarm   | クラスタの構築や管理を担う (主にマルチホスト)                           | `docker swarm`   |
| Service | Swarm で、Service (1 つ以上のコンテナの集まり)を管理する                | `docker service` |
| Stack   | Swarm で、複数の Service をまとめたアプリケーション全体の管理           | `docker stack`   |

# コンテナ配置について

コンテナのオーケストレーションを行う上で複数のホストを用意する必要があるが、導入コストが高い。  
そこで Docker in Docker で「Docker ホストとして機能する Docker コンテナ」を複数立てる。

複数ホストと見なしたコンテナは次の通り。

- registry × 1  
  Docker イメージのレジストリ。  
  あらかじめイメージを push しておき、manager や worker からイメージを pull する。  
  普通は https でアクセスするものだが、http でもアクセスできるような回避策を取る。

- manager × 1  
  Swarm クラスタ全体を制御し、複数実行されている worker へ Service が保持するコンテナを適切に配置する。

- worker × 3  
  複数の worker によってクラスタ全体のアプリケーションを実行する。

# TODO アプリケーション構築の全体像

アプリケーションを構築する流れは次の通り。

1. データストアとなる Master/Slave 構成の MySQL Service の構築

2. MySQL とデータをやり取りするための API を実装

3. Web アプリケーションと API サーバ間にリバースプロキシとなる Nginx を通じてアクセスできるように設定

4. API を利用してサーバサイドレンダリングをする Web アプリケーションを実装

5. フロント側にリバースプロキシ (Nginx) を置く

# 実行

## 事前準備

```bash
# 複数ホストの起動
cd SwarmTodoApp
docker-compose up -d

# managerでswarm initを実行してSwarmモードにする (このときswarm joinのトークンが表示されるので控えておく)
docker exec -it manager docker swarm init

# 各workerを登録する
docker exec -it worker01 docker swarm join --token {SWMTKN-1-...} manager:2377
docker exec -it worker02 docker swarm join --token {SWMTKN-1-...} manager:2377
docker exec -it worker03 docker swarm join --token {SWMTKN-1-...} manager:2377

# ノードを確認する (master × 1 と worker × 3 が確認できる)
docker-compose exec manager docker node ls

# あらかじめ専用のオーバーレイネットワークを構築する
# オーバーレイネットワークを構築することで、Dockerホストを問わずに配置されているコンテナがあたかも同一NW上に存在するように扱える
docker-compose exec manager docker network create --driver=overlay --attachable todoapp

```

## Swarm で MySQL のスタックを構築する

```bash
# ホスト上でMySQLのDockerイメージ(masterとslaveを兼用)をビルドし、registryにプッシュする
cd tododb
docker build -t ch04/tododb:latest .
docker image tag ch04/tododb:latest localhost:5000/ch04/tododb:latest
docker push localhost:5000/ch04/tododb:latest

# Swarm上でMySQLのMaster/Slaveサービスを実行する
docker-compose exec manager docker stack deploy -c /stack/todo-mysql.yml todo_mysql
docker-compose exec manager docker service ls
```

この時点で Swarm クラスターは以下のようになっている。

![MySQLスタックのデプロイ時](./images/mysqlスタックデプロイ時.png)
