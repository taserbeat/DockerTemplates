# nginx-secret

Kubernetes の Secret リソースを使うことで、TLS/SSL 証明書や秘密鍵、パスワードといった機密情報を Base64Base64 エンコードした状態で利用できる。

ここでは、Nginx の Basic 認証の認証情報を記述したファイルを Secret で管理してみる。

# デプロイ手順

## 1. ユーザー名とパスワードの Base64 暗号化

まず、openssl を利用してユーザー名とパスワードを暗号化し、その結果を Base64 文字列に変換する。

```bash
cd kubernetes-advance/nginx-secret

echo "your_username:$(openssl passwd -quiet -crypt your_password)" | base64
> eW91cl91c2VybmFtZToySVZqVy45TW1POW5BCg==
```

## 2. Secret リソースの作成

Secret リソースを作成する。
このとき、Base64 で暗号化された文字列を yaml ファイルの`data.htpasswd`に設定してから Secret リソースを作成する。

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: nginx-secret
type: Opaque
data:
  .htpasswd: BASE64_STRING
```

```bash
kubectl apply -f nginx-secret.yaml

kubectl get secret nginx-secret
```

Secret が作成されると、kubernetes-dashboard から確認することができる。

```bash
# ダッシュボードのデプロイ
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.3.1/aio/deploy/recommended.yaml

# secret名を取得
# この例では "deployment-controller-token-xxs5d" がsecret名
kubectl -n kube-system get secret | grep deploy
# >> deployment-controller-token-xxs5d kubernetes.io/service-account-token 3 56m

# secret名からトークンを取得
kubectl -n kube-system describe secret deployment-controller-token-xxs5d

# ブラウザで見れるように、プロキシを起動
kubectl proxy

# プロキシ起動後は以下のURLでダッシュボードを閲覧可能
# http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/
```

![Base64暗号化したデータの確認.png](./images/Base64暗号化したデータの確認.png)

## 3. Basic 認証を施した Nginx の構築

この Secret リソースを活用し、Basic 認証を施した Nginx を構築する。  
`basic-auth.yaml`という Service と Deployment が定義されたマニフェストファイルを実行する。

```bash
kubectl apply -f basic-auth.yaml

kubectl get pod -l app=basic-auth
```

## 4. 認証が働いているかの確認

次のように HTTP リクエストを送ると、ステータスコード 401 で認証が施されていることがわかる。

```bash
curl http://localhost:30060
```

次のように、認証情報を付けて HTTP リクエストを送ると、ステータスコード 200 でレスポンスが返る。

```bash
curl -i --user your_username:your_password http://localhost:30060
```

以上、Secret を使うことで機密情報を平文で保存しなくてもよくなることがわかった。  
しかし、完璧な対策にはなっていないので注意が必要。

# リソースの削除

```bash
kubectl delete -f nginx-secret.yaml
kubectl delete -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.3.1/aio/deploy/recommended.yaml
kubectl delete -f basic-auth.yaml
```
