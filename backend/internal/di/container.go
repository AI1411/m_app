// backend/internal/di/container.go
package di

import (
	"context"
	"errors"
	"net/http"

	"go.uber.org/fx"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/AI1411/m_app/gen/grpc/user/v1/userv1connect"
	"github.com/AI1411/m_app/internal/handler"
	"github.com/AI1411/m_app/internal/infra/db"
	"github.com/AI1411/m_app/internal/infra/repository/datastore"
	"github.com/AI1411/m_app/internal/usecase"
)

// Module はアプリケーションのDIコンテナを定義します
var Module = fx.Options(
	// インフラストラクチャの依存関係
	fx.Provide(
		// データベース接続
		func() (*db.SqlHandler, error) {
			return db.NewSqlHandler(
				"myuser",     // 環境変数から取得するのが理想的
				"mypassword", // 環境変数から取得するのが理想的
				"localhost",  // 環境変数から取得するのが理想的
				"5432",       // 環境変数から取得するのが理想的
				"m_app",      // 環境変数から取得するのが理想的
			)
		},
	),

	// レポジトリの依存関係
	fx.Provide(
		// UserRepository
		func(sqlHandler *db.SqlHandler) datastore.UserRepository {
			return datastore.NewUserRepository(sqlHandler)
		},
	),

	// ユースケースの依存関係
	fx.Provide(
		// UserUseCase
		usecase.NewUserUseCase,
	),

	// ハンドラーの依存関係
	fx.Provide(
		// UserHandler
		handler.NewUserHandler,

		// HTTPサーバーのセットアップ
		func(userHandler handler.UserHandler) *http.ServeMux {
			mux := http.NewServeMux()
			path, h := userv1connect.NewUserServiceHandler(userHandler)
			mux.Handle(path, h)
			return mux
		},

		// HTTPサーバーの設定
		func(mux *http.ServeMux) *http.Server {
			return &http.Server{
				Addr:    "localhost:8080",
				Handler: h2c.NewHandler(mux, &http2.Server{}),
			}
		},
	),

	// アプリケーションのライフサイクル
	fx.Invoke(func(lifecycle fx.Lifecycle, server *http.Server) {
		lifecycle.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				// サーバーの起動
				go func() {
					if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
						panic(err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				// サーバーのグレースフルシャットダウン
				return server.Shutdown(ctx)
			},
		})
	}),
)
