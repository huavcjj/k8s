# ベースイメージとして最新のGoを使用
FROM golang:latest

# 環境変数の設定
ENV PORT=8082
ENV PROJECT=search

# 作業ディレクトリの作成
WORKDIR /app

# アプリケーションコードのコピー
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# アプリケーションをビルド
RUN go build -o myapp .

# ポートの公開
EXPOSE ${PORT}

# アプリケーションの実行
CMD ["./myapp"]