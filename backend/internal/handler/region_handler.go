package handler

import (
	"context"

	"connectrpc.com/connect"

	prefecturev1 "github.com/AI1411/m_app/gen/prefecture/v1"
	regionv1 "github.com/AI1411/m_app/gen/region/v1"
	"github.com/AI1411/m_app/internal/infra/logger"
	"github.com/AI1411/m_app/internal/usecase"
)

// RegionHandler はリージョン関連APIのハンドラーインターフェース
type RegionHandler interface {
	ListRegions(context.Context, *connect.Request[regionv1.ListRegionsRequest]) (*connect.Response[regionv1.ListRegionsResponse], error)
}

// regionHandler はRegionHandlerインターフェースの実装
type regionHandler struct {
	regionUseCase usecase.RegionUseCase
	logger        *logger.Logger
}

// NewRegionHandler は新しいRegionHandlerインスタンスを作成します
func NewRegionHandler(regionUseCase usecase.RegionUseCase, logger *logger.Logger) RegionHandler {
	return &regionHandler{
		regionUseCase: regionUseCase,
		logger:        logger,
	}
}

// ListRegions はリージョン一覧取得APIを処理します
func (h *regionHandler) ListRegions(
	ctx context.Context,
	req *connect.Request[regionv1.ListRegionsRequest],
) (*connect.Response[regionv1.ListRegionsResponse], error) {
	h.logger.Info("ListRegions called", "headers", req.Header())

	// リージョンと都道府県を取得
	regions, prefectures, err := h.regionUseCase.ListRegionsWithPrefectures(ctx)
	if err != nil {
		h.logger.LogError(err, "failed to list regions with prefectures")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// 都道府県をリージョンIDでグループ化
	prefecturesByRegionID := make(map[int32][]*prefecturev1.Prefecture)
	for _, p := range prefectures {
		prefecturesByRegionID[p.RegionID] = append(prefecturesByRegionID[p.RegionID], &prefecturev1.Prefecture{
			Id:       p.ID,
			Code:     int32(p.Code),
			Name:     p.Name,
			NameEn:   p.NameEn,
			RegionId: p.RegionID,
		})
	}

	// レスポンスの構築
	var regionProtos []*regionv1.Region
	for _, r := range regions {
		// 標準のRegionオブジェクトを作成
		regionProto := &regionv1.Region{
			Id:        r.ID,
			Name:      r.Name,
			NameEn:    r.NameEn,
			SortOrder: r.SortOrder,
		}

		regionProtos = append(regionProtos, regionProto)
	}

	// 注意: ここでは都道府県情報はレスポンスに含まれていませんが、
	// クライアント側で必要に応じて別途都道府県一覧APIを呼び出して取得することができます。
	// 本来はprotoファイルを更新して再生成する必要があります。

	h.logger.Info("regions retrieved successfully", "count", len(regionProtos))
	return connect.NewResponse(&regionv1.ListRegionsResponse{
		Regions: regionProtos,
	}), nil
}
