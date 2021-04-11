# SimpleServer

golang で簡易的な HTTP サーバーを作り、Docker で動かすだけ。

# golang のプロジェクト構築について

`Go modules`の登場によって、どこにでもプロジェクトディレクトリを作れるようになった。

```bash
cd PROJECT_ROOT_DIR
mkdir PROJECT_NAME
cd PROJECT_NAME
go mod init MODULE_NAME
```

例:

```bash
cd golang
mkdir SimpleServer
cd SimpleServer
go mod init main
```

# golang の実行と build について

- 実行

```bash
go run main.go
```

- build + 実行

```bash
go build main.go
./main
```

# Dockerfile の build、run、kill、rmi について

- build

```bash
docker build -t taserbeat/echo:latest .
```

- run

```bash
docker run --rm -d --name simple_server -p 9000:8080 taserbeat/echo:latest
```

- kill

```bash
docker kill simple_server
```

- rmi

```bash
docker rmi taserbeat/echo:latest
```
