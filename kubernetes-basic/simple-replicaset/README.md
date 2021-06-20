# simple-replicaset

同じ仕様の Pod を 複数構成するサンプル

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
