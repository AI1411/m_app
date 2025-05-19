package usecase

import (
	"context"

	"github.com/AI1411/m_app/internal/domain/model"
	"github.com/AI1411/m_app/internal/infra/repository/datastore"
)

// InterestUseCase は興味・関心関連のビジネスロジックを提供するインターフェース
type InterestUseCase interface {
	ListInterests(ctx context.Context) ([]*model.Interest, error)
}

// interestUseCase はInterestUseCaseインターフェースの実装
type interestUseCase struct {
	interestRepo datastore.InterestRepository
}

// NewInterestUseCase は新しいInterestUseCaseを作成します
func NewInterestUseCase(interestRepo datastore.InterestRepository) InterestUseCase {
	return &interestUseCase{
		interestRepo: interestRepo,
	}
}

// ListInterests は興味・関心一覧を取得します
func (u *interestUseCase) ListInterests(ctx context.Context) ([]*model.Interest, error) {
	interests, err := u.interestRepo.ListInterests(ctx)
	if err != nil {
		return nil, err
	}
	return interests, nil
}