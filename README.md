# Bone

## Docker を使った環境構築
まだ未完成ではありますが、ホームページの表示のみできます。

### 環境構築方法
コンテナを起動させる
```
docker-compose build 
docker-compose up -d
```

Webアプリの実行
```
docker-compose exec web go run main.go
```