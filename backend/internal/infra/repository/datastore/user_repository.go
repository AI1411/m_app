package datastore

import (
	"context"

	"github.com/AI1411/m_app/internal/domain/model"
	"github.com/AI1411/m_app/internal/infra/db"
)

// UserRepository はユーザーデータへのアクセスを提供するインターフェース
type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	SearchUsers(ctx context.Context) ([]*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
}

// userRepository はUserRepositoryインターフェースの実装
type userRepository struct {
	sqlHandler *db.SqlHandler
}

// NewUserRepository は新しいUserRepositoryを作成します
func NewUserRepository(sqlHandler *db.SqlHandler) UserRepository {
	return &userRepository{
		sqlHandler: sqlHandler,
	}
}

// GetUserByID はIDでユーザーを取得します
func (r *userRepository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	result := r.sqlHandler.Conn.WithContext(ctx).Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// SearchUsers はユーザーを検索します
func (r *userRepository) SearchUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	result := r.sqlHandler.Conn.WithContext(ctx).Preload("UserInterests").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// CreateUser は新しいユーザーを作成します
func (r *userRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	result := r.sqlHandler.Conn.WithContext(ctx).Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
