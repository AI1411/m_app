package handler

import (
	"context"

	"connectrpc.com/connect"

	prefecturev1 "github.com/AI1411/m_app/gen/prefecture/v1"
	regionv1 "github.com/AI1411/m_app/gen/region/v1"
	"github.com/AI1411/m_app/internal/domain/model"
	"github.com/AI1411/m_app/internal/infra/logger"
	"github.com/AI1411/m_app/internal/usecase"
)

// RegionHandler はリージョン関連APIのハンドラーインターフェース
type RegionHandler interface {
	ListRegions(context.Context, *connect.Request[regionv1.ListRegionsRequest]) (*connect.Response[regionv1.ListRegionsResponse], error)
	GetRegion(context.Context, *connect.Request[regionv1.GetRegionRequest]) (*connect.Response[regionv1.GetRegionResponse], error)
	CreateRegion(context.Context, *connect.Request[regionv1.CreateRegionRequest]) (*connect.Response[regionv1.CreateRegionResponse], error)
	UpdateRegion(context.Context, *connect.Request[regionv1.UpdateRegionRequest]) (*connect.Response[regionv1.UpdateRegionResponse], error)
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

	h.logger.Info("regions retrieved successfully", "count", len(regionProtos))
	return connect.NewResponse(&regionv1.ListRegionsResponse{
		Regions: regionProtos,
	}), nil
}

// GetRegion はリージョン取得APIを処理します
func (h *regionHandler) GetRegion(
	ctx context.Context,
	req *connect.Request[regionv1.GetRegionRequest],
) (*connect.Response[regionv1.GetRegionResponse], error) {
	h.logger.Info("GetRegion called", "headers", req.Header())

	// リージョンを取得
	region, err := h.regionUseCase.GetRegionByID(ctx, req.Msg.Id)
	if err != nil {
		h.logger.LogError(err, "failed to get region")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// レスポンスの構築
	regionProto := &regionv1.Region{
		Id:        region.ID,
		Name:      region.Name,
		NameEn:    region.NameEn,
		SortOrder: region.SortOrder,
	}

	h.logger.Info("region retrieved successfully", "id", region.ID)
	return connect.NewResponse(&regionv1.GetRegionResponse{
		Region: regionProto,
	}), nil
}

// CreateRegion はリージョン作成APIを処理します
func (h *regionHandler) CreateRegion(
	ctx context.Context,
	req *connect.Request[regionv1.CreateRegionRequest],
) (*connect.Response[regionv1.CreateRegionResponse], error) {
	h.logger.Info("CreateRegion called", "headers", req.Header())

	// リージョンを作成
	region, err := h.regionUseCase.CreateRegion(ctx, &model.Region{
		Name:      req.Msg.Region.Name,
		NameEn:    req.Msg.Region.NameEn,
		SortOrder: req.Msg.Region.SortOrder,
	})
	if err != nil {
		h.logger.LogError(err, "failed to create region")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// レスポンスの構築
	regionProto := &regionv1.Region{
		Id:        region.ID,
		Name:      region.Name,
		NameEn:    region.NameEn,
		SortOrder: region.SortOrder,
	}

	h.logger.Info("region created successfully", "id", region.ID)
	return connect.NewResponse(&regionv1.CreateRegionResponse{
		Region: regionProto,
	}), nil
}

// UpdateRegion はリージョン更新APIを処理します
func (h *regionHandler) UpdateRegion(
	ctx context.Context,
	req *connect.Request[regionv1.UpdateRegionRequest],
) (*connect.Response[regionv1.UpdateRegionResponse], error) {
	h.logger.Info("UpdateRegion called", "headers", req.Header())

	// リージョンを更新
	region, err := h.regionUseCase.UpdateRegion(ctx, &model.Region{
		ID:        req.Msg.Region.Id,
		Name:      req.Msg.Region.Name,
		NameEn:    req.Msg.Region.NameEn,
		SortOrder: req.Msg.Region.SortOrder,
	})
	if err != nil {
		h.logger.LogError(err, "failed to update region")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// レスポンスの構築
	regionProto := &regionv1.Region{
		Id:        region.ID,
		Name:      region.Name,
		NameEn:    region.NameEn,
		SortOrder: region.SortOrder,
	}

	h.logger.Info("region updated successfully", "id", region.ID)
	return connect.NewResponse(&regionv1.UpdateRegionResponse{
		Region: regionProto,
	}), nil
}
