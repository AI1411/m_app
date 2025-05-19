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
	GetRegionByID(ctx context.Context, id int32) (*model.Region, error)
	CreateRegion(ctx context.Context, region *model.Region) (*model.Region, error)
	UpdateRegion(ctx context.Context, region *model.Region) (*model.Region, error)
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

// GetRegionByID はリージョンIDからリージョンを取得します
func (u *regionUseCase) GetRegionByID(ctx context.Context, id int32) (*model.Region, error) {
	region, err := u.regionRepo.GetRegionByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return region, nil
}

// CreateRegion は新しいリージョンを作成します
func (u *regionUseCase) CreateRegion(ctx context.Context, region *model.Region) (*model.Region, error) {
	createdRegion, err := u.regionRepo.CreateRegion(ctx, region)
	if err != nil {
		return nil, err
	}
	return createdRegion, nil
}

// UpdateRegion はリージョンを更新します
func (u *regionUseCase) UpdateRegion(ctx context.Context, region *model.Region) (*model.Region, error) {
	updatedRegion, err := u.regionRepo.UpdateRegion(ctx, region)
	if err != nil {
		return nil, err
	}
	return updatedRegion, nil
}
