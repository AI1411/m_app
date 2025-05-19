package usecase

import (
	"context"

	"github.com/AI1411/m_app/internal/domain/model"
	"github.com/AI1411/m_app/internal/infra/repository/datastore"
)

// EducationUseCase は学歴関連のビジネスロジックを提供するインターフェース
type EducationUseCase interface {
	ListEducations(ctx context.Context) ([]*model.Education, error)
}

// educationUseCase はEducationUseCaseインターフェースの実装
type educationUseCase struct {
	educationRepo datastore.EducationRepository
}

// NewEducationUseCase は新しいEducationUseCaseを作成します
func NewEducationUseCase(educationRepo datastore.EducationRepository) EducationUseCase {
	return &educationUseCase{
		educationRepo: educationRepo,
	}
}

// ListEducations は学歴一覧を取得します
func (u *educationUseCase) ListEducations(ctx context.Context) ([]*model.Education, error) {
	educations, err := u.educationRepo.ListEducations(ctx)
	if err != nil {
		return nil, err
	}
	return educations, nil
}