package usecase

import (
	"context"

	"github.com/AI1411/m_app/internal/domain/model"
	"github.com/AI1411/m_app/internal/infra/repository/datastore"
)

// UserUseCase はユーザー関連のビジネスロジックを提供するインターフェース
type UserUseCase interface {
	GetUser(ctx context.Context, id string) (*model.User, error)
	SearchUsers(ctx context.Context) ([]*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
}

// userUseCase はUserUseCaseインターフェースの実装
type userUseCase struct {
	userRepo datastore.UserRepository
}

// NewUserUseCase は新しいUserUseCaseを作成します
func NewUserUseCase(userRepo datastore.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

// GetUser は指定されたIDのユーザーを取得します
func (u *userUseCase) GetUser(ctx context.Context, id string) (*model.User, error) {
	user, err := u.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// SearchUsers はユーザーを検索します
func (u *userUseCase) SearchUsers(ctx context.Context) ([]*model.User, error) {
	users, err := u.userRepo.SearchUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userUseCase) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	createdUser, err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
