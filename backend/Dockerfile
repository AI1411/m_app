FROM golang:1.24-alpine AS builder

# 作業ディレクトリの設定
WORKDIR /app

# Goモジュールの依存関係をコピーしてダウンロード
COPY go.mod go.sum ./
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# メインパッケージを探して実行可能ファイルをビルド
# cmd/mainまたはcmd/appなどのディレクトリにあると仮定
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o /app/main ./cmd/server

# ---- 実行ステージ ----
FROM gcr.io/distroless/static-debian11 AS runner

# ビルドステージからコンパイル済みバイナリをコピー
COPY --from=builder /app/main /app/main

# 設定ファイルなどの必要なファイルをコピー
COPY --from=builder /app/.env* /app/

# ポートの公開
EXPOSE 8080

# 作業ディレクトリの設定
WORKDIR /app

# アプリケーションの実行
ENTRYPOINT ["/app/main"]