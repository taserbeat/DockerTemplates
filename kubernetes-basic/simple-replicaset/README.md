# simple-replicaset

複数の Pod を 1 つのマニフェストファイルで構成するサンプル

# デプロイ手順

```bash
cd kubernetes-basic/simple-replicaset
kubectl apply -f simple-replicaset.yml

kubectl get pod
```

# 削除手順

```bash
kubectl delete -f simple-replicaset.yml
```
