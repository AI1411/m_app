package handler

import (
	"context"

	"connectrpc.com/connect"

	userv1 "github.com/AI1411/m_app/gen/user/v1"
	"github.com/AI1411/m_app/internal/infra/logger"
	"github.com/AI1411/m_app/internal/usecase"
)

// UserHandler はユーザー関連APIのハンドラーインターフェース
type UserHandler interface {
	GetUser(context.Context, *connect.Request[userv1.GetUserRequest]) (*connect.Response[userv1.GetUserResponse], error)
	SearchUsers(context.Context, *connect.Request[userv1.SearchUsersRequest]) (*connect.Response[userv1.SearchUsersResponse], error)
	CreateUser(context.Context, *connect.Request[userv1.CreateUserRequest]) (*connect.Response[userv1.CreateUserResponse], error)
	UpdateUser(context.Context, *connect.Request[userv1.UpdateUserRequest]) (*connect.Response[userv1.UpdateUserResponse], error)
}

// userHandler はUserHandlerインターフェースの実装
type userHandler struct {
	userUseCase usecase.UserUseCase
	logger      *logger.Logger
}

// NewUserHandler は新しいUserHandlerインスタンスを作成します
func NewUserHandler(userUseCase usecase.UserUseCase, logger *logger.Logger) UserHandler {
	return &userHandler{
		userUseCase: userUseCase,
		logger:      logger,
	}
}

// GetUser はユーザー情報取得APIを処理します
func (h *userHandler) GetUser(
	ctx context.Context,
	req *connect.Request[userv1.GetUserRequest],
) (*connect.Response[userv1.GetUserResponse], error) {
	h.logger.Info("GetUser called", "request_id", req.Msg.Id, "headers", req.Header())

	user, err := h.userUseCase.GetUser(ctx, req.Msg.Id)
	if err != nil {
		h.logger.LogError(err, "failed to get user", "request_id", req.Msg.Id)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	h.logger.Info("user retrieved successfully", "user_id", user.ID)
	return connect.NewResponse(&userv1.GetUserResponse{
		User: &userv1.User{
			Id:   user.ID,
			Name: user.Name,
			// 他のフィールドをmodelからDTOに変換
		},
	}), nil
}

// SearchUsers はユーザー検索APIを処理します
func (h *userHandler) SearchUsers(
	ctx context.Context,
	req *connect.Request[userv1.SearchUsersRequest],
) (*connect.Response[userv1.SearchUsersResponse], error) {
	h.logger.Info("SearchUsers called", "headers", req.Header())

	users, err := h.userUseCase.SearchUsers(ctx)
	if err != nil {
		h.logger.LogError(err, "failed to search users")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	var userPreviews []*userv1.UserPreview
	for _, user := range users {
		userPreviews = append(userPreviews, &userv1.UserPreview{
			Id:   user.ID,
			Name: user.Name,
		})
	}

	h.logger.Info("users retrieved successfully", "count", len(userPreviews))
	res := connect.NewResponse(&userv1.SearchUsersResponse{
		Users: userPreviews,
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

// CreateUser はユーザー作成APIを処理します
func (h *userHandler) CreateUser(
	ctx context.Context,
	req *connect.Request[userv1.CreateUserRequest],
) (*connect.Response[userv1.CreateUserResponse], error) {
	h.logger.Info("CreateUser called", "headers", req.Header(), "name", req.Msg.Name)
	// ここに実装を追加
	h.logger.Info("user created successfully", "id", "new-user-id")
	return connect.NewResponse(&userv1.CreateUserResponse{
		Id: "new-user-id",
	}), nil
}

// UpdateUser はユーザー更新APIを処理します
func (h *userHandler) UpdateUser(
	ctx context.Context,
	req *connect.Request[userv1.UpdateUserRequest],
) (*connect.Response[userv1.UpdateUserResponse], error) {
	h.logger.Info("UpdateUser called", "headers", req.Header(), "id", req.Msg.Id)
	// ここに実装を追加
	h.logger.Info("user updated successfully", "id", req.Msg.Id)
	return connect.NewResponse(&userv1.UpdateUserResponse{
		Success: true,
	}), nil
}
