# simple-cronjob

Job は一度切りの Pod の実行であるが、CronJob リソースを利用するとスケジューリングして定期的に Pod を実行できる。  
Cron や systemdtimer などで定期実行していたジョブの実行に最適である。

CronJob によって独自に Cron でイベントを発行するようなアプリケーションを用意する必要がなくなる。  
通常の Cron はサーバの CronTab で管理するが、CronJob はマニフェストファイルで定義できる。  
スケジューリング定義のレビューを GitHub の PullRequest で運用できるなど、構成のコード管理という点でも有利である。

# デプロイ

```bash
cd kubernetes-advance/simple-cronjob

kubectl apply -f simple-cronjob.yaml
```

Job は指定した Cron スケジュール(毎分ごと)に実行されていることがわかる。

```bash
kubectl get job -l app=pingpong
```

```bash
kubectl logs -l app=pingpong
```

# リソースの削除

```bash
kubectl delete -f simple-cronjob.yaml
```
