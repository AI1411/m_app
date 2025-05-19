package usecase

import (
	"context"

	"github.com/AI1411/m_app/internal/domain/model"
	"github.com/AI1411/m_app/internal/infra/repository/datastore"
)

// PrefectureUseCase は都道府県関連のビジネスロジックを提供するインターフェース
type PrefectureUseCase interface {
	ListPrefectures(ctx context.Context, regionID *int32) ([]*model.Prefecture, error)
	GetPrefecture(ctx context.Context, id int32) (*model.Prefecture, error)
}

// prefectureUseCase はPrefectureUseCaseインターフェースの実装
type prefectureUseCase struct {
	prefectureRepo datastore.PrefectureRepository
}

// NewPrefectureUseCase は新しいPrefectureUseCaseを作成します
func NewPrefectureUseCase(prefectureRepo datastore.PrefectureRepository) PrefectureUseCase {
	return &prefectureUseCase{
		prefectureRepo: prefectureRepo,
	}
}

// ListPrefectures は都道府県一覧を取得します
func (u *prefectureUseCase) ListPrefectures(ctx context.Context, regionID *int32) ([]*model.Prefecture, error) {
	prefectures, err := u.prefectureRepo.ListPrefectures(ctx, regionID)
	if err != nil {
		return nil, err
	}
	return prefectures, nil
}

// GetPrefecture は指定されたIDの都道府県を取得します
func (u *prefectureUseCase) GetPrefecture(ctx context.Context, id int32) (*model.Prefecture, error) {
	prefecture, err := u.prefectureRepo.GetPrefecture(ctx, id)
	if err != nil {
		return nil, err
	}
	return prefecture, nil
}
