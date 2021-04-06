# imagemagick

画像処理ツールで生成した画像ファイルを永続化する。

# コマンド

```bash
cd imagemagick

docker run --rm  -v ${PWD}/workspace:/workspace --name magick gihyodocker/imagemagick:latest convert -size 100x100 xc:#000000 /workspace/gihyo.jpg
```
