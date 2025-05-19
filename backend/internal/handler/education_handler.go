package handler

import (
	"context"

	"connectrpc.com/connect"

	educationv1 "github.com/AI1411/m_app/gen/education/v1"
	"github.com/AI1411/m_app/internal/infra/logger"
	"github.com/AI1411/m_app/internal/usecase"
)

// EducationHandler は学歴関連APIのハンドラーインターフェース
type EducationHandler interface {
	ListEducations(context.Context, *connect.Request[educationv1.ListEducationsRequest]) (*connect.Response[educationv1.ListEducationsResponse], error)
}

// educationHandler はEducationHandlerインターフェースの実装
type educationHandler struct {
	educationUseCase usecase.EducationUseCase
	logger           *logger.Logger
}

// NewEducationHandler は新しいEducationHandlerインスタンスを作成します
func NewEducationHandler(educationUseCase usecase.EducationUseCase, logger *logger.Logger) EducationHandler {
	return &educationHandler{
		educationUseCase: educationUseCase,
		logger:           logger,
	}
}

func (h *educationHandler) ListEducations(
	ctx context.Context,
	req *connect.Request[educationv1.ListEducationsRequest],
) (*connect.Response[educationv1.ListEducationsResponse], error) {
	h.logger.Info("ListEducations called", "headers", req.Header())

	// 学歴一覧を取得
	educations, err := h.educationUseCase.ListEducations(ctx)
	if err != nil {
		h.logger.LogError(err, "failed to list educations")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// レスポンスの構築
	var educationProtos []*educationv1.Education
	for _, e := range educations {
		educationProto := &educationv1.Education{
			Id:        e.ID,
			Name:      e.Name,
			SortOrder: e.SortOrder,
		}

		educationProtos = append(educationProtos, educationProto)
	}

	h.logger.Info("educations retrieved successfully", "count", len(educationProtos))
	return connect.NewResponse(&educationv1.ListEducationsResponse{
		Educations: educationProtos,
	}), nil
}
