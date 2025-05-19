package datastore

import (
	"context"

	"github.com/AI1411/m_app/internal/domain/model"
	"github.com/AI1411/m_app/internal/infra/db"
)

// RegionRepository はリージョンデータへのアクセスを提供するインターフェース
type RegionRepository interface {
	ListRegions(ctx context.Context) ([]*model.Region, error)
	ListRegionsWithPrefectures(ctx context.Context) ([]*model.Region, []*model.Prefecture, error)
	GetRegionByID(ctx context.Context, id int32) (*model.Region, error)
	CreateRegion(ctx context.Context, region *model.Region) (*model.Region, error)
	UpdateRegion(ctx context.Context, region *model.Region) (*model.Region, error)
}

// regionRepository はRegionRepositoryインターフェースの実装
type regionRepository struct {
	sqlHandler *db.SqlHandler
}

// NewRegionRepository は新しいRegionRepositoryを作成します
func NewRegionRepository(sqlHandler *db.SqlHandler) RegionRepository {
	return &regionRepository{
		sqlHandler: sqlHandler,
	}
}

// ListRegions はリージョン一覧を取得します
func (r *regionRepository) ListRegions(ctx context.Context) ([]*model.Region, error) {
	var regions []*model.Region
	err := r.sqlHandler.Conn.Order("sort_order").Find(&regions).Error
	if err != nil {
		return nil, err
	}

	return regions, nil
}

// ListRegionsWithPrefectures はリージョン一覧と関連する都道府県を取得します
func (r *regionRepository) ListRegionsWithPrefectures(ctx context.Context) ([]*model.Region, []*model.Prefecture, error) {
	// リージョン一覧を取得
	var regions []*model.Region
	err := r.sqlHandler.Conn.Order("sort_order").Find(&regions).Error
	if err != nil {
		return nil, nil, err
	}

	// 都道府県一覧を取得
	var prefectures []*model.Prefecture
	err = r.sqlHandler.Conn.Order("code").Find(&prefectures).Error
	if err != nil {
		return nil, nil, err
	}

	return regions, prefectures, nil
}

// GetRegionByID はリージョンIDからリージョンを取得します
func (r *regionRepository) GetRegionByID(ctx context.Context, id int32) (*model.Region, error) {
	var region model.Region
	err := r.sqlHandler.Conn.Where("id = ?", id).First(&region).Error
	if err != nil {
		return nil, err
	}

	return &region, nil
}

// CreateRegion は新しいリージョンを作成します
func (r *regionRepository) CreateRegion(ctx context.Context, region *model.Region) (*model.Region, error) {
	err := r.sqlHandler.Conn.Create(region).Error
	if err != nil {
		return nil, err
	}

	return region, nil
}

// UpdateRegion はリージョンを更新します
func (r *regionRepository) UpdateRegion(ctx context.Context, region *model.Region) (*model.Region, error) {
	err := r.sqlHandler.Conn.Save(region).Error
	if err != nil {
		return nil, err
	}

	return region, nil
}
