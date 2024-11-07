FROM golang:1.23-alpine

# MySQL クライアントをインストール
RUN apk add --no-cache mysql-client

WORKDIR /app

RUN go mod download

# コンテナ起動時に実行されるコマンド
CMD ["go", "run", "cmd/server/main.go"]
