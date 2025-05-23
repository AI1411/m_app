package di

import (
	"context"
	"errors"
	"net/http"

	"go.uber.org/fx"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"connectrpc.com/connect"

	"github.com/AI1411/m_app/gen/community/v1/communityv1connect"
	"github.com/AI1411/m_app/gen/education/v1/educationv1connect"
	"github.com/AI1411/m_app/gen/interest/v1/interestv1connect"
	"github.com/AI1411/m_app/gen/notification/v1/notificationv1connect"
	"github.com/AI1411/m_app/gen/prefecture/v1/prefecturev1connect"
	"github.com/AI1411/m_app/gen/region/v1/regionv1connect"
	"github.com/AI1411/m_app/gen/user/v1/userv1connect"
	"github.com/AI1411/m_app/internal/env"
	"github.com/AI1411/m_app/internal/handler"
	"github.com/AI1411/m_app/internal/infra/db"
	"github.com/AI1411/m_app/internal/infra/logger"
	"github.com/AI1411/m_app/internal/infra/repository/datastore"
	"github.com/AI1411/m_app/internal/usecase"
)

// Module はアプリケーションのDIコンテナを定義します
var Module = fx.Options(
	// インフラストラクチャの依存関係
	fx.Provide(
		// 環境変数の設定
		func(log *logger.Logger) (*env.Values, error) {
			// 環境変数を読み込む
			v, err := env.NewValues()
			if err != nil {
				log.LogError(err, "環境変数の読み込みに失敗しました")
				return nil, err
			}
			return v, nil
		},

		// ロガーの設定
		func() *logger.Logger {
			cfg := logger.DefaultConfig()
			// ログの形式はJSONで
			cfg.JSON = true
			return logger.New(cfg)
		},

		// データベース接続
		func(envValues *env.Values, log *logger.Logger) (*db.SqlHandler, error) { // envValues を引数に追加
			return db.NewSqlHandler(
				envValues.DatabaseUsername,
				envValues.DatabasePassword,
				envValues.DatabaseHost,
				envValues.DatabasePort,
				envValues.DatabaseName,
				log,
			)
		},
	),

	// レポジトリの依存関係
	fx.Provide(
		// UserRepository
		func(sqlHandler *db.SqlHandler) datastore.UserRepository {
			return datastore.NewUserRepository(sqlHandler)
		},

		// PrefectureRepository
		func(sqlHandler *db.SqlHandler) datastore.PrefectureRepository {
			return datastore.NewPrefectureRepository(sqlHandler)
		},

		// RegionRepository
		func(sqlHandler *db.SqlHandler) datastore.RegionRepository {
			return datastore.NewRegionRepository(sqlHandler)
		},

		// InterestRepository
		func(sqlHandler *db.SqlHandler) datastore.InterestRepository {
			return datastore.NewInterestRepository(sqlHandler)
		},

		// EducationRepository
		func(sqlHandler *db.SqlHandler) datastore.EducationRepository {
			return datastore.NewEducationRepository(sqlHandler)
		},

		// CommunityRepository
		func(sqlHandler *db.SqlHandler) datastore.CommunityRepository {
			return datastore.NewCommunityRepository(sqlHandler)
		},

		// NotificationRepository
		func(sqlHandler *db.SqlHandler) datastore.NotificationRepository {
			return datastore.NewNotificationRepository(sqlHandler)
		},
	),

	// ユースケースの依存関係
	fx.Provide(
		// UserUseCase
		usecase.NewUserUseCase,

		// RegionUseCase
		usecase.NewRegionUseCase,

		// InterestUseCase
		usecase.NewInterestUseCase,

		// EducationUseCase
		usecase.NewEducationUseCase,

		// CommunityUseCase
		usecase.NewCommunityUseCase,

		// NotificationUseCase
		usecase.NewNotificationUseCase,
	),

	// ハンドラーの依存関係
	fx.Provide(
		// UserHandler
		handler.NewUserHandler,

		// PrefectureHandler
		handler.NewPrefectureHandler,

		// RegionHandler
		handler.NewRegionHandler,

		// InterestHandler
		handler.NewInterestHandler,

		// EducationHandler
		handler.NewEducationHandler,

		// CommunityHandler
		handler.NewCommunityHandler,

		// NotificationHandler
		handler.NewNotificationHandler,

		// HTTPサーバーのセットアップ
		func(userHandler handler.UserHandler, prefectureHandler handler.PrefectureHandler, regionHandler handler.RegionHandler, interestHandler handler.InterestHandler, educationHandler handler.EducationHandler, communityHandler handler.CommunityHandler, notificationHandler handler.NotificationHandler) *http.ServeMux {
			mux := http.NewServeMux()

			// Connectハンドラーのオプション設定
			// HTTP（JSON）互換性を有効にする
			connectOpts := []connect.HandlerOption{
				connect.WithCompressMinBytes(1024),
			}

			// ユーザーサービスのハンドラー登録
			userPath, userHttpHandler := userv1connect.NewUserServiceHandler(userHandler, connectOpts...)
			mux.Handle(userPath, userHttpHandler)

			// 都道府県サービスのハンドラー登録
			prefecturePath, prefectureHttpHandler := prefecturev1connect.NewPrefectureServiceHandler(prefectureHandler, connectOpts...)
			mux.Handle(prefecturePath, prefectureHttpHandler)

			// リージョンサービスのハンドラー登録
			regionPath, regionHttpHandler := regionv1connect.NewRegionServiceHandler(regionHandler, connectOpts...)
			mux.Handle(regionPath, regionHttpHandler)

			// 興味・関心サービスのハンドラー登録
			interestPath, interestHttpHandler := interestv1connect.NewInterestServiceHandler(interestHandler, connectOpts...)
			mux.Handle(interestPath, interestHttpHandler)

			// 学歴サービスのハンドラー登録
			educationPath, educationHttpHandler := educationv1connect.NewEducationServiceHandler(educationHandler, connectOpts...)
			mux.Handle(educationPath, educationHttpHandler)

			// コミュニティサービスのハンドラー登録
			communityPath, communityHttpHandler := communityv1connect.NewCommunityServiceHandler(communityHandler, connectOpts...)
			mux.Handle(communityPath, communityHttpHandler)

			// 通知サービスのハンドラー登録
			notificationPath, notificationHttpHandler := notificationv1connect.NewNotificationServiceHandler(notificationHandler, connectOpts...)
			mux.Handle(notificationPath, notificationHttpHandler)

			return mux
		},

		// HTTPサーバーの設定
		func(envValues *env.Values, mux *http.ServeMux) *http.Server { // envValues を引数に追加
			return &http.Server{
				Addr:    envValues.ServerPort,
				Handler: h2c.NewHandler(mux, &http2.Server{}),
			}
		},
	),

	// アプリケーションのライフサイクル
	fx.Invoke(func(lifecycle fx.Lifecycle, server *http.Server, log *logger.Logger) {
		lifecycle.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				// サーバーの起動
				log.Info("starting HTTP server", "address", server.Addr)
				go func() {
					if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
						log.LogError(err, "サーバー起動エラー", "address", server.Addr)
						// アプリケーションを終了させるためにシグナルを送る方法として、fx.Shutdownerを利用することも検討してください
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				// サーバーのグレースフルシャットダウン
				log.Info("shutting down HTTP server")
				return server.Shutdown(ctx)
			},
		})
	}),
)
