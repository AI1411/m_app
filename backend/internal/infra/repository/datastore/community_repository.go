package datastore

import (
	"context"

	"github.com/AI1411/m_app/internal/domain/model"
	"github.com/AI1411/m_app/internal/infra/db"
)

// CommunityRepository はコミュニティデータへのアクセスを提供するインターフェース
type CommunityRepository interface {
	GetCommunityByID(ctx context.Context, id string) (*model.Community, error)
	SearchCommunities(ctx context.Context, params map[string]interface{}, page, pageSize int) ([]*model.Community, int64, error)
	CreateCommunity(ctx context.Context, community *model.Community) (*model.Community, error)
	UpdateCommunity(ctx context.Context, community *model.Community) error
	DeleteCommunity(ctx context.Context, id string) error
}

// communityRepository はCommunityRepositoryインターフェースの実装
type communityRepository struct {
	sqlHandler *db.SqlHandler
}

// NewCommunityRepository は新しいCommunityRepositoryを作成します
func NewCommunityRepository(sqlHandler *db.SqlHandler) CommunityRepository {
	return &communityRepository{
		sqlHandler: sqlHandler,
	}
}

// GetCommunityByID はIDでコミュニティを取得します
func (r *communityRepository) GetCommunityByID(ctx context.Context, id string) (*model.Community, error) {
	var community model.Community
	result := r.sqlHandler.Conn.
		Where("id = ?", id).
		Preload("CommunityMembers").
		Preload("User").
		First(&community)
	if result.Error != nil {
		return nil, result.Error
	}
	return &community, nil
}

// SearchCommunities はコミュニティを検索します
func (r *communityRepository) SearchCommunities(ctx context.Context, params map[string]interface{}, page, pageSize int) ([]*model.Community, int64, error) {
	var communities []*model.Community
	var totalCount int64

	query := r.sqlHandler.Conn.Model(&model.Community{})

	// 検索条件の適用
	if name, ok := params["name"].(string); ok && name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if isPrivate, ok := params["is_private"].(bool); ok {
		query = query.Where("is_private = ?", isPrivate)
	}

	// 総件数の取得
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// ページネーション
	offset := (page - 1) * pageSize
	result := query.
		Offset(offset).
		Limit(pageSize).
		Find(&communities)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return communities, totalCount, nil
}

// CreateCommunity は新しいコミュニティを作成します
func (r *communityRepository) CreateCommunity(ctx context.Context, community *model.Community) (*model.Community, error) {
	result := r.sqlHandler.Conn.Create(community)
	if result.Error != nil {
		return nil, result.Error
	}
	return community, nil
}

// UpdateCommunity はコミュニティを更新します
func (r *communityRepository) UpdateCommunity(ctx context.Context, community *model.Community) error {
	result := r.sqlHandler.Conn.Model(&model.Community{}).
		Where("id = ?", community.ID).
		Updates(community)
	return result.Error
}

// DeleteCommunity はコミュニティを削除します
func (r *communityRepository) DeleteCommunity(ctx context.Context, id string) error {
	result := r.sqlHandler.Conn.Delete(&model.Community{}, "id = ?", id)
	return result.Error
}
