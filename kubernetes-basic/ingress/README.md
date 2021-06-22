# Ingress

NodePort は Kubernetes クラスタ外に Service を公開するすることが可能だが、それは L4 層レベルのルーティングとなる。
HTTP/HTTPS のようにパスベースで転送先を切り替えるという L7 層のルーティングは Ingress を使用することで実現できる。

# デプロイ

素の状態のローカル Kubernetes 環境では Ingress を使った Service を公開することはできない。
クラスタ外から HTTP リクエストを Service にルーティングするための`nginx_ingress_controller`をデプロイする。

```bash
kubectl apply -f mandatory.yaml
kubectl apply -f cloud-generic.yaml

kubectl -n ingress-nginx get service,pod
```

Ingress を通して Service にアクセスするために Service を作成する。
`spec.type`が未指定なので ClusterIP Service が作成される。

```bash
kubectl apply -f simple-service-for-ingress.yml

kubectl get svc echo
```

Ingress を定義したマニフェストファイルを反映する。

```bash
kubectl apply -f simple-ingress.yml

kubectl get ingress
```

---

[動作未検証]

ローカルから HTTP リクエストを投げるとバックエンドに存在する echo Service からレスポンスが返るようになる。
**次の curl を実行すると 503 エラーとなってしまう**

```bash
curl http://localhost -H 'Host: ch05.gihyo.local'
```

---

# リソースの削除

```bash
kubectl delete -f mandatory.yaml
kubectl delete -f cloud-generic.yaml
kubectl delete -f simple-service-for-ingress.yml
kubectl delete -f simple-ingress.yml
```

# 参考文献

[Ingress](https://kubernetes.io/ja/docs/concepts/services-networking/ingress/)  
[kubernetes/ngress-nginx](https://github.com/kubernetes/ingress-nginx/blob/nginx-0.30.0/docs/deploy/index.md)
