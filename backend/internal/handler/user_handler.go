package handler

import (
	"context"
	"log"

	"connectrpc.com/connect"

	userv1 "github.com/AI1411/m_app/gen/grpc/user/v1"
)

// UserHandler はユーザー関連APIのハンドラー
type UserHandler struct{}

// NewUserHandler は新しいUserHandlerインスタンスを作成します
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// GetUser はユーザー情報取得APIを処理します
func (h *UserHandler) GetUser(
	ctx context.Context,
	req *connect.Request[userv1.GetUserRequest],
) (*connect.Response[userv1.GetUserResponse], error) {
	log.Println("Request headers: ", req.Header())
	// ここに実装を追加
	return connect.NewResponse(&userv1.GetUserResponse{
		User: &userv1.User{
			Id:   req.Msg.Id,
			Name: "John Doe",
			// 他のフィールドも必要に応じて設定
		},
	}), nil
}

// SearchUsers はユーザー検索APIを処理します
func (h *UserHandler) SearchUsers(
	ctx context.Context,
	req *connect.Request[userv1.SearchUsersRequest],
) (*connect.Response[userv1.SearchUsersResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&userv1.SearchUsersResponse{
		Users: []*userv1.UserPreview{
			{
				Id:   "1",
				Name: "John Doe",
			},
		},
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

// CreateUser はユーザー作成APIを処理します
func (h *UserHandler) CreateUser(
	ctx context.Context,
	req *connect.Request[userv1.CreateUserRequest],
) (*connect.Response[userv1.CreateUserResponse], error) {
	log.Println("Request headers: ", req.Header())
	// ここに実装を追加
	return connect.NewResponse(&userv1.CreateUserResponse{
		Id: "new-user-id",
	}), nil
}

// UpdateUser はユーザー更新APIを処理します
func (h *UserHandler) UpdateUser(
	ctx context.Context,
	req *connect.Request[userv1.UpdateUserRequest],
) (*connect.Response[userv1.UpdateUserResponse], error) {
	log.Println("Request headers: ", req.Header())
	// ここに実装を追加
	return connect.NewResponse(&userv1.UpdateUserResponse{
		Success: true,
	}), nil
}
