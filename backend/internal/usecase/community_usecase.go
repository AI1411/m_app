package usecase

import (
	"context"

	"github.com/AI1411/m_app/internal/domain/model"
	"github.com/AI1411/m_app/internal/infra/repository/datastore"
)

// CommunityUseCase はコミュニティ関連のビジネスロジックを提供するインターフェース
type CommunityUseCase interface {
	GetCommunity(ctx context.Context, id string) (*model.Community, error)
	SearchCommunities(ctx context.Context, params map[string]interface{}, page, pageSize int) ([]*model.Community, int64, error)
	CreateCommunity(ctx context.Context, community *model.Community) (*model.Community, error)
	UpdateCommunity(ctx context.Context, community *model.Community) error
	DeleteCommunity(ctx context.Context, id string) error
}

// communityUseCase はCommunityUseCaseインターフェースの実装
type communityUseCase struct {
	communityRepo datastore.CommunityRepository
}

// NewCommunityUseCase は新しいCommunityUseCaseを作成します
func NewCommunityUseCase(communityRepo datastore.CommunityRepository) CommunityUseCase {
	return &communityUseCase{
		communityRepo: communityRepo,
	}
}

// GetCommunity は指定されたIDのコミュニティを取得します
func (u *communityUseCase) GetCommunity(ctx context.Context, id string) (*model.Community, error) {
	community, err := u.communityRepo.GetCommunityByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return community, nil
}

// SearchCommunities はコミュニティを検索します
func (u *communityUseCase) SearchCommunities(ctx context.Context, params map[string]interface{}, page, pageSize int) ([]*model.Community, int64, error) {
	communities, totalCount, err := u.communityRepo.SearchCommunities(ctx, params, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return communities, totalCount, nil
}

// CreateCommunity は新しいコミュニティを作成します
func (u *communityUseCase) CreateCommunity(ctx context.Context, community *model.Community) (*model.Community, error) {
	createdCommunity, err := u.communityRepo.CreateCommunity(ctx, community)
	if err != nil {
		return nil, err
	}
	return createdCommunity, nil
}

// UpdateCommunity はコミュニティを更新します
func (u *communityUseCase) UpdateCommunity(ctx context.Context, community *model.Community) error {
	return u.communityRepo.UpdateCommunity(ctx, community)
}

// DeleteCommunity はコミュニティを削除します
func (u *communityUseCase) DeleteCommunity(ctx context.Context, id string) error {
	return u.communityRepo.DeleteCommunity(ctx, id)
}
