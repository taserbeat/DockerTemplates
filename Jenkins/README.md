# Jenkins

Jenkins の Docker コンテナを起動する。

# 手順

1. `docker-compose.yml`から`master`サービスのみを立ち上げる

```bash
docker-compose up master
```

2. [http://localhost:8080](http://localhost:8080)にアクセスし、Jenkins の初期設定を済ます

※最初の Unlock はコンソールに表示されるパスワードを入力する

3. SSH の鍵を生成する

```bash
docker exec -it jenkins_master ssh-keygen -t rsa -C ""
```

デフォルトであれば、[jenkins_home/.ssh](jenkins_home/.ssh)に SSH 鍵が生成される。

4. `master`の公開鍵を登録する

[./docker-compose.yml](./docker-compose.yml)の`slave01`サービスの環境変数に`master`の SSH 公開鍵を設定する。  
`JENKINS_SLAVE_SSH_PUBKEY`に`id_rsa.pub`の内容を貼り付ける。

例:

```
JENKINS_SLAVE_SSH_PUBKEY=ssh-rsa AAA...
```

5. compose up したターミナルを終了する

Ctrl + c で`compose up`の処理を強制終了する。

6. すべてのサービスを立ち上げる

```bash
docker-compose up -d
```

※ この時点で`jenkins_master`と`jenkins_slave01`コンテナが起動している状態である。

```bash
docker-compose ps

    Name                    Command               State                 Ports
--------------------------------------------------------------------------------------------
jenkins_master    /sbin/tini -- /usr/local/b ...   Up      50000/tcp, 0.0.0.0:8080->8080/tcp
jenkins_slave01   setup-sshd                       Up      22/tcp
```

7. ジョブ実行用の slave ノードを追加する

Jenkins の Web アプリからジョブ実行用の slave ノードを追加する設定を行う。

`Jenkinsの管理` -> `ノードの選択` -> `新規ノード作成` ->　以下の項目と値を入力する。

| 項目            | 値      |
| --------------- | ------- |
| ノード名        | slave01 |
| Permanent Agent | ✅      |

-> 以下の項目のみ値を変更する。

| 項目                           | 変更後の値                                         |
| ------------------------------ | -------------------------------------------------- |
| リモート FS ルート             | /home/jenkins                                      |
| 起動方法                       | SSH 経由で Unix マシンのスレーブエージェントを起動 |
| ホスト                         | slave01                                            |
| 認証情報 -> 追加               | Jenkins                                            |
| 認証情報の追加 -> 種類         | SSH ユーザー名と秘密鍵                             |
| ユーザー名                     | jenkins                                            |
| 秘密鍵                         | (jenkins_home/.ssh/id_rsa の内容を貼り付ける)      |
| Host Key Verification Strategy | Non verifying Verification Strategy                |
