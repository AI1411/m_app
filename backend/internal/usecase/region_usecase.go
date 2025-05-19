package usecase

import (
	"context"

	"github.com/AI1411/m_app/internal/domain/model"
	"github.com/AI1411/m_app/internal/infra/repository/datastore"
)

// RegionUseCase はリージョン関連のビジネスロジックを提供するインターフェース
type RegionUseCase interface {
	ListRegions(ctx context.Context) ([]*model.Region, error)
	ListRegionsWithPrefectures(ctx context.Context) ([]*model.Region, []*model.Prefecture, error)
}

// regionUseCase はRegionUseCaseインターフェースの実装
type regionUseCase struct {
	regionRepo datastore.RegionRepository
}

// NewRegionUseCase は新しいRegionUseCaseを作成します
func NewRegionUseCase(regionRepo datastore.RegionRepository) RegionUseCase {
	return &regionUseCase{
		regionRepo: regionRepo,
	}
}

// ListRegions はリージョン一覧を取得します
func (u *regionUseCase) ListRegions(ctx context.Context) ([]*model.Region, error) {
	regions, err := u.regionRepo.ListRegions(ctx)
	if err != nil {
		return nil, err
	}
	return regions, nil
}

// ListRegionsWithPrefectures はリージョン一覧と関連する都道府県を取得します
func (u *regionUseCase) ListRegionsWithPrefectures(ctx context.Context) ([]*model.Region, []*model.Prefecture, error) {
	regions, prefectures, err := u.regionRepo.ListRegionsWithPrefectures(ctx)
	if err != nil {
		return nil, nil, err
	}
	return regions, prefectures, nil
}