FROM golang:1.11.5-alpine3.9
  
# Gopkg.toml 等を事前にコピーして dep ensure が実行できるようにする
COPY . /go/src/go_sample/

# dep ensure を行うプロジェクトルートに移動する
WORKDIR /go/src/go_sample/

# 必要なパッケージをイメージにインストールする
RUN apk update \
  && apk add --no-cache \
  git \
  bash \
  vim \
  gcc \
  && go get -u github.com/codegangsta/gin \
  && go get -u github.com/pilu/fresh \
  && go get -u github.com/golang/dep/cmd/dep \
  && dep ensure

# コンテナ実行時のデフォルトを設定する
# ライブリロードを実行する
CMD ["fresh"]
