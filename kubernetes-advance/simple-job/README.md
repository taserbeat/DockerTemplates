# simple-job

Job は 1 つ以上の Pod を作成し、指定された数の Pod が正常に完了するまでを管理するリソースである。  
Job による全ての Pod が正常に終了しても、Pod は削除されずに保持されるため、終了後に Pod のログや実行結果を分析できる。  
そのため、Web アプリケーション等の常駐型アプリケーションではなく、**大規模な計算やバッチ指向のアプリケーションに向いている**。

Job は Pod を複数並列で実行することで容易にスケールアウトできる。  
また、Pod として実行されることで Kubernetes の Service と連携した処理を行いやすいという面もある。

# デプロイ

マニフェストファイルをデプロイすると、Pod が実行される。

```bash
cd kubernetes-advance/simple-job

kubectl apply -f simple-job.yaml
```

Pod の実行後、約 20 秒後に全ての処理が完了する。

```bash
kubectl logs -l app=pingpong

[Tue Jun 22 12:38:31 UTC 2021] ping!
[Tue Jun 22 12:38:41 UTC 2021] pong!
[Tue Jun 22 12:38:32 UTC 2021] ping!
[Tue Jun 22 12:38:42 UTC 2021] pong!
[Tue Jun 22 12:38:32 UTC 2021] ping!
[Tue Jun 22 12:38:42 UTC 2021] pong!
```

処理が完了した Pod は`STATUS`が`Completed`になる。

```bash
kubectl get pod -l app=pingpong

NAME             READY   STATUS      RESTARTS   AGE
pingpong-9qj25   0/1     Completed   0          3m15s
pingpong-mltfl   0/1     Completed   0          3m15s
pingpong-sx8k5   0/1     Completed   0          3m15s
```

# リソースの削除

```bash
kubectl delete -f simple-job.yaml
```
