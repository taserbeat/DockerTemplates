# simple-service

Pod の集合(主に ReplicaSet)に対する経路やサービスディスカバリを提供する Service のサンプルである。

[simple-replicaset-with-label.yml](./simple-replicaset-with-label.yml)に ReplicaSet が 2 つ定義されている。  
release というラベルが付けられ、それぞれラベルの値は `spring` と `summer` になっている。

# デプロイ

デプロイしてみると、release ラベルに`spring` と `summer`を持つ Pod が作成されていることがわかる。

```bash
kubectl apply -f simple-replicaset-with-label.yml

kubectl get pod -l app=echo -l release=spring
kubectl get pod -l app=echo -l release=summer
```

# Service の例

release=summer のラベルを持つ Pod だけにアクセスできるような Service を作ってみる。

[simple-service.yml](./simple-service.yml)の Service で、  
Selector で指定したラベルと一致した Pod はその Service の対象となり、Service を経由してトラフィックが流れるようになる。

![Serviceからトラフィックが流れる例.png](./images/Serviceからトラフィックが流れる例.png)

```bash
kubectl apply -f simple-service.yml

kubectl get svc echo
```

# release=summer の Pod だけにトラフィックが流れるかの確認

Service は Kubernetes クラスタ内からしかアクセスできない。  
そのため、Kubernetes クラスタ内に一時的なデバッグコンテナをデプロイして確認する。
デバッグコンテナに入ったら`http://echo`に対して curl する。

```bash
# デバッグ用コンテナを実行する
kubectl run -i --rm --tty debug --image=gihyodocker/fundamental:0.1.0 --restart=Never -- bash -il

# curlを何回か実行する
curl http://echo
curl http://echo
curl http://echo
```

release=summer の Pod を 1 つ選び、ログを確認すると「received request」が表示される。
しかし、release=spring の Pod にはログが出力されないことが確認できる。

```bash
# Podを確認する
kubectl get pod

NAME                READY   STATUS    RESTARTS   AGE
echo-spring-cvmlx   2/2     Running   0          21m
echo-summer-vt2tj   2/2     Running   0          21m

# summerのPodでログを確認
kubectl logs -f echo-summer-vt2tj -c echo

2021/06/20 12:41:56 start server
2021/06/20 13:00:45 received request
2021/06/20 13:01:04 received request
2021/06/20 13:01:05 received request
2021/06/20 13:01:06 received request


# springのPodでログを確認 (received requestが表示されない)
kubectl logs -f echo-spring-cvmlx -c echo

2021/06/20 12:41:58 start server
```

# Service の名前解決について

Kubernetes クラスタ内の DNS では、Service を`Service名.Namespace名.svc.local`で名前解決できる。
例えば、echo は Namespace は`default`に配置されているので

```bash
curl http://echo.default.svc.local
```

でアクセスできる。

また、`svc.local`は省略可能で、異なる Namespace の Service 名前解決の最短は次のようになる。

```bash
curl http://echo.default
```

Namespace が同一であれば Service 名だけで名前解決できる。

```bash
curl http://echo
```

# Servic の種類について

作成できる Service には様々な種類があり、yml ファイルで指定できる。

## ClusterIP Service

デフォルトで作成される Service である。  
ClusterIP は Kubernetes クラスタ上の内部 IP アドレスとして Service を公開できる。  
よって、Pod から別の Pod へのアクセスで ClusterIP Service を介することができる。
ただし、外からリーチできない。

## NodePort Service

NodePort はクラスタ外からアクセスできる Service である。  
NodePort も ClusterIP を作る点においては ClusterIP Service と同様であるが、  
各ノード上から Service ポートへ接続するためのグローバルなポート開けるという違いがある。

```yml
apiVersion: v1
kind: Service
metadata:
  name: echo
spec:
  type: NodePort
  selector:
    app: echo
  ports:
    - name: http
      port: 80
```

NodePortService を作成した場合は、以下に 80:31058/TCP と表示されているようにノードの 31058 ポートから Service へアクセスできる。

```bash
kubectl get svc echo

NAME TYPE       CLUSTER-IP     EXTERNAL-IP  PORT(S)       AGE
echo NodePort   10.97.191.150  <none>       80:31058/TCP  32m
```

このとき、ローカルから次のようにアクセスできる。

```bash
curl http://127.0.0.1:31058
```

## LoadBalancer Service

クラウドプラットフォームで提供されているロードバランサーと連携するもので、ローカルの Kubernetes 環境では使用できない。

## ExternalName Service

ExternalName Service は selector も port 定義も持たない特殊な Service である。  
Kubernetes クラスタ内から外部のホストを解決するためのエイリアスを提供する。

例えば、次のような Service を作成すると gihyo.jp を `gihyo` で名前解決できるようになります。

```yml
apiVersion: v1
kind: Service
metadata:
  name: gihyo
spec:
  type: ExternalName
  externalName: gihyo.jp
```
