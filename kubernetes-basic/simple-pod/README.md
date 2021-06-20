# simple-pod

nginx-proxy と echo アプリケーション ([SwarmTutorial](../../SwarmTutorial)と同じ)を Kubernetes にデプロイする

# デプロイ手順

```bash
cd kubernetes-basic/simple-pod
kubectl apply -f simple-pod.yml
```

# Pod の操作

## Pod の一覧取得

```bash
kubectl get pod
```

## kubectl でコンテナにログインする

```bash
kubectl exec -it simple-echo sh -c nginx
```

## Pod 内コンテナの標準出力を表示

```bash
kubectl logs -f simple-echo -c echo
```

## Pod を削除する

kubectl delete で削除できる。

```bash
kubectl delete pod simple-echo
```

マニフェストファイルから削除指定も可能

```bash
kubectl delete -f simple-pod.yml
```
