package datastore

import (
	"context"

	"github.com/AI1411/m_app/internal/domain/model"
	"github.com/AI1411/m_app/internal/infra/db"
)

// InterestRepository は興味・関心データへのアクセスを提供するインターフェース
type InterestRepository interface {
	ListInterests(ctx context.Context) ([]*model.Interest, error)
}

// interestRepository はInterestRepositoryインターフェースの実装
type interestRepository struct {
	sqlHandler *db.SqlHandler
}

// NewInterestRepository は新しいInterestRepositoryを作成します
func NewInterestRepository(sqlHandler *db.SqlHandler) InterestRepository {
	return &interestRepository{
		sqlHandler: sqlHandler,
	}
}

// ListInterests は興味・関心一覧を取得します
func (r *interestRepository) ListInterests(ctx context.Context) ([]*model.Interest, error) {
	var interests []*model.Interest
	err := r.sqlHandler.Conn.Order("sort_order").Find(&interests).Error
	if err != nil {
		return nil, err
	}

	return interests, nil
}
