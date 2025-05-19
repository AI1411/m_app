package main

import (
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/AI1411/m_app/gen/grpc/user/v1/userv1connect"
	"github.com/AI1411/m_app/internal/handler"
)

func main() {
	// ハンドラーの初期化
	userHandler := handler.NewUserHandler()

	// ルーティングの設定
	mux := http.NewServeMux()
	path, h := userv1connect.NewUserServiceHandler(userHandler)
	mux.Handle(path, h)

	// サーバーの起動
	addr := "localhost:8080"
	log.Printf("サーバーを起動中: %s", addr)
	err := http.ListenAndServe(
		addr,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		log.Fatalf("サーバー起動エラー: %v", err)
	}
}
