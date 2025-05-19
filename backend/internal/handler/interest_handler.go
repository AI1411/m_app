package handler

import (
	"context"

	"connectrpc.com/connect"

	interestv1 "github.com/AI1411/m_app/gen/interest/v1"
	"github.com/AI1411/m_app/internal/infra/logger"
	"github.com/AI1411/m_app/internal/usecase"
)

// InterestHandler は興味・関心関連APIのハンドラーインターフェース
type InterestHandler interface {
	ListInterests(context.Context, *connect.Request[interestv1.ListInterestsRequest]) (*connect.Response[interestv1.ListInterestsResponse], error)
}

// interestHandler はInterestHandlerインターフェースの実装
type interestHandler struct {
	interestUseCase usecase.InterestUseCase
	logger          *logger.Logger
}

// NewInterestHandler は新しいInterestHandlerインスタンスを作成します
func NewInterestHandler(interestUseCase usecase.InterestUseCase, logger *logger.Logger) InterestHandler {
	return &interestHandler{
		interestUseCase: interestUseCase,
		logger:          logger,
	}
}

func (h *interestHandler) ListInterests(
	ctx context.Context,
	req *connect.Request[interestv1.ListInterestsRequest],
) (*connect.Response[interestv1.ListInterestsResponse], error) {
	h.logger.Info("ListInterests called", "headers", req.Header())

	// 興味・関心一覧を取得
	interests, err := h.interestUseCase.ListInterests(ctx)
	if err != nil {
		h.logger.LogError(err, "failed to list interests")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// レスポンスの構築
	var interestProtos []*interestv1.Interest
	for _, i := range interests {
		interestProto := &interestv1.Interest{
			Id:          i.ID,
			Name:        i.Name,
			DisplayName: i.DisplayName,
			SortOrder:   i.SortOrder,
		}

		if i.CategoryID != nil {
			interestProto.CategoryId = i.CategoryID
		}

		if i.IconURL != nil {
			interestProto.IconUrl = i.IconURL
		}

		interestProtos = append(interestProtos, interestProto)
	}

	h.logger.Info("interests retrieved successfully", "count", len(interestProtos))
	return connect.NewResponse(&interestv1.ListInterestsResponse{
		Interests: interestProtos,
	}), nil
}
