# SwarmTutorial

Docker Swarm を使って、複数ホストに複数コンテナを制御するコンテナオーケストレーションツールを学ぶ。

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

# 実行

**Service を作成する**

```bash
# 複数ホストの起動
cd SwarmTutorial
docker-compose up -d

# managerでswarm initを実行してSwarmモードにする (このときswarm joinのトークンが表示されるので控えておく)
docker exec -it manager docker swarm init

# 各workerを登録する
docker exec -it worker01 docker swarm join --token {SWMTKN-1-...} manager:2377
docker exec -it worker02 docker swarm join --token {SWMTKN-1-...} manager:2377
docker exec -it worker03 docker swarm join --token {SWMTKN-1-...} manager:2377

# ノードを確認する (master × 1 と worker × 3 が確認できる)
docker exec -it manager docker node ls

# レジストリにpushするDockerイメージのビルドとタグ付けを行う
docker build -t taserbeat/echo:latest ../SimpleServer
docker tag taserbeat/echo:latest localhost:5000/taserbeat/echo:latest

# registryコンテナにイメージをpushする
docker push localhost:5000/taserbeat/echo:latest

# workerがregistryからDockerイメージをpullできるか確認する
docker exec -it worker01 docker pull registry:5000/taserbeat/echo:latest
docker exec -it worker01 docker images

# "echo"という名前のServiceを作る
docker exec -it manager docker service create --replicas 1 --publish 8000:8080 --name echo registry:5000/taserbeat/echo:latest

# Serviceが作られたことを確認する
docker exec -it manager docker service ls

# Serviceのコンテナの数を増やす (スケールアウトする)
docker exec -it manager docker service scale echo=6

# Serviceがスケールアウトしたことを確認する (REPLICASが増えている)
docker exec -it manager docker service ls

# Swarmクラスタで実行されているコンテナを確認する
# NODEの項目を見ると、Swarmクラスタのノードに分散していることがわかる
docker exec -it manager docker service ps echo

# デプロイしたServiceを削除する
docker exec -it manager docker service rm echo
```
