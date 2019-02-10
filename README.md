# Go Lang SampleDockerAPI Ripositori

## ディレクトリ構成

```
root
  ┣━ docker-compose.yml
  ┣━ .dockerignore
  ┣━ sample
  ┃    ┣━ main.go
  ┃    ┣━ Dockerfile
  ┃    ┣━ controllers (ディレクトリ)
  ┃    ┣━ services    (ディレクトリ)
  ┃    ┣━ models      (ディレクトリ)
  ┃    ┣━ Gopkg.toml  (Gemfileと同じ)
  ┃    ┗━ Gopkg.lock  (Gemfile.lockと同じ)
  ┣━ mysql
  ┃    ┗━ data (Docker内のデータをマウント)
  ┗━ redis
       ┗━ data (Docker内のデータをマウント)
```
---- 

## パッケージ管理 depの使い方

```
$ dep ensure        # bundle installと同じ
$ dep               # helpコマンドと同じ意味
$ dep status        # パッケージ一覧
```
---- 

## 環境

- ミドルウェア
  - Redis (高速化キャッシュ)
  - MySQL (RDB)
- フレームワーク
  - Gin


## Dockerの使い方

```
$ docker-compose build
$ docker-compose up
$ docker-compose exec go_sample bash
```
