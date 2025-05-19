package datastore

import (
	"context"

	"github.com/AI1411/m_app/internal/domain/model"
	"github.com/AI1411/m_app/internal/infra/db"
)

// PrefectureRepository は都道府県データへのアクセスを提供するインターフェース
type PrefectureRepository interface {
	ListPrefectures(ctx context.Context, regionID *int32) ([]*model.Prefecture, error)
	GetPrefecture(ctx context.Context, id int32) (*model.Prefecture, error)
}

// prefectureRepository はPrefectureRepositoryインターフェースの実装
type prefectureRepository struct {
	sqlHandler *db.SqlHandler
}

// NewPrefectureRepository は新しいPrefectureRepositoryを作成します
func NewPrefectureRepository(sqlHandler *db.SqlHandler) PrefectureRepository {
	return &prefectureRepository{
		sqlHandler: sqlHandler,
	}
}

// ListPrefectures は都道府県一覧を取得します
func (r *prefectureRepository) ListPrefectures(ctx context.Context, regionID *int32) ([]*model.Prefecture, error) {
	query := r.sqlHandler.Conn.Model(&model.Prefecture{})

	if regionID != nil {
		query = query.Where("region_id = ?", *regionID)
	}

	var prefectures []*model.Prefecture
	err := query.Order("code").Find(&prefectures).Error
	if err != nil {
		return nil, err
	}

	return prefectures, nil
}

// GetPrefecture は指定されたIDの都道府県を取得します
func (r *prefectureRepository) GetPrefecture(ctx context.Context, id int32) (*model.Prefecture, error) {
	var prefecture model.Prefecture
	err := r.sqlHandler.Conn.Model(&model.Prefecture{}).
		Where("id = ?", id).
		First(&prefecture).Error
	if err != nil {
		return nil, err
	}

	// 関連する地域情報を取得
	var region model.Region
	err = r.sqlHandler.Conn.Model(&model.Region{}).
		Where("id = ?", prefecture.RegionID).
		First(&region).Error
	if err != nil {
		return nil, err
	}

	prefecture.Region = region

	return &prefecture, nil
}
