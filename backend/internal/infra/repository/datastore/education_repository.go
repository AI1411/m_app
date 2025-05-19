package datastore

import (
	"context"

	"github.com/AI1411/m_app/internal/domain/model"
	"github.com/AI1411/m_app/internal/infra/db"
)

// EducationRepository は学歴データへのアクセスを提供するインターフェース
type EducationRepository interface {
	ListEducations(ctx context.Context) ([]*model.Education, error)
}

// educationRepository はEducationRepositoryインターフェースの実装
type educationRepository struct {
	sqlHandler *db.SqlHandler
}

// NewEducationRepository は新しいEducationRepositoryを作成します
func NewEducationRepository(sqlHandler *db.SqlHandler) EducationRepository {
	return &educationRepository{
		sqlHandler: sqlHandler,
	}
}

// ListEducations は学歴一覧を取得します
func (r *educationRepository) ListEducations(ctx context.Context) ([]*model.Education, error) {
	var educations []*model.Education
	err := r.sqlHandler.Conn.Order("sort_order").Find(&educations).Error
	if err != nil {
		return nil, err
	}

	return educations, nil
}