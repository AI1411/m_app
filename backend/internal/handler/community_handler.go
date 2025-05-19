package handler

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/timestamppb"

	communityv1 "github.com/AI1411/m_app/gen/community/v1"
	"github.com/AI1411/m_app/internal/domain/model"
	"github.com/AI1411/m_app/internal/infra/logger"
	"github.com/AI1411/m_app/internal/usecase"
)

// CommunityHandler はコミュニティ関連APIのハンドラーインターフェース
type CommunityHandler interface {
	GetCommunity(context.Context, *connect.Request[communityv1.GetCommunityRequest]) (*connect.Response[communityv1.GetCommunityResponse], error)
	SearchCommunities(context.Context, *connect.Request[communityv1.SearchCommunitiesRequest]) (*connect.Response[communityv1.SearchCommunitiesResponse], error)
	CreateCommunity(context.Context, *connect.Request[communityv1.CreateCommunityRequest]) (*connect.Response[communityv1.CreateCommunityResponse], error)
	UpdateCommunity(context.Context, *connect.Request[communityv1.UpdateCommunityRequest]) (*connect.Response[communityv1.UpdateCommunityResponse], error)
	DeleteCommunity(context.Context, *connect.Request[communityv1.DeleteCommunityRequest]) (*connect.Response[communityv1.DeleteCommunityResponse], error)
}

// communityHandler はCommunityHandlerインターフェースの実装
type communityHandler struct {
	communityUseCase usecase.CommunityUseCase
	logger           *logger.Logger
}

// NewCommunityHandler は新しいCommunityHandlerインスタンスを作成します
func NewCommunityHandler(communityUseCase usecase.CommunityUseCase, logger *logger.Logger) CommunityHandler {
	return &communityHandler{
		communityUseCase: communityUseCase,
		logger:           logger,
	}
}

// GetCommunity はコミュニティ情報取得APIを処理します
func (h *communityHandler) GetCommunity(
	ctx context.Context,
	req *connect.Request[communityv1.GetCommunityRequest],
) (*connect.Response[communityv1.GetCommunityResponse], error) {
	h.logger.Info("GetCommunity called", "request_id", req.Msg.Id, "headers", req.Header())

	community, err := h.communityUseCase.GetCommunity(ctx, req.Msg.Id)
	if err != nil {
		h.logger.LogError(err, "failed to get community", "request_id", req.Msg.Id)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	h.logger.Info("community retrieved successfully", "community_id", community.ID)
	return connect.NewResponse(&communityv1.GetCommunityResponse{
		Community: convertToCommunityProto(community),
	}), nil
}

// SearchCommunities はコミュニティ検索APIを処理します
func (h *communityHandler) SearchCommunities(
	ctx context.Context,
	req *connect.Request[communityv1.SearchCommunitiesRequest],
) (*connect.Response[communityv1.SearchCommunitiesResponse], error) {
	h.logger.Info("SearchCommunities called", "headers", req.Header())

	// 検索パラメータの構築
	params := make(map[string]interface{})
	if req.Msg.Name != nil {
		params["name"] = *req.Msg.Name
	}
	if req.Msg.IsPrivate != nil {
		params["is_private"] = *req.Msg.IsPrivate
	}

	// ページネーションパラメータの設定
	page := int(req.Msg.Page)
	if page <= 0 {
		page = 1
	}
	pageSize := int(req.Msg.PageSize)
	if pageSize <= 0 {
		pageSize = 10
	}

	communities, totalCount, err := h.communityUseCase.SearchCommunities(ctx, params, page, pageSize)
	if err != nil {
		h.logger.LogError(err, "failed to search communities")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	var communityPreviews []*communityv1.CommunityPreview
	for _, community := range communities {
		communityPreviews = append(communityPreviews, convertToCommunityPreviewProto(community))
	}

	h.logger.Info("communities retrieved successfully", "count", len(communityPreviews))
	res := connect.NewResponse(&communityv1.SearchCommunitiesResponse{
		Communities: communityPreviews,
		TotalCount:  int32(totalCount),
		Page:        req.Msg.Page,
		PageSize:    req.Msg.PageSize,
	})
	return res, nil
}

// CreateCommunity はコミュニティ作成APIを処理します
func (h *communityHandler) CreateCommunity(
	ctx context.Context,
	req *connect.Request[communityv1.CreateCommunityRequest],
) (*connect.Response[communityv1.CreateCommunityResponse], error) {
	h.logger.Info("CreateCommunity called", "headers", req.Header(), "name", req.Msg.Name)

	// リクエストからモデルへの変換
	community := &model.Community{
		Name:      req.Msg.Name,
		IsPrivate: req.Msg.IsPrivate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// オプションフィールドの設定
	if req.Msg.Description != nil {
		community.Description = req.Msg.Description
	}
	if req.Msg.ProfileImageUrl != nil {
		community.ProfileImageURL = req.Msg.ProfileImageUrl
	}
	if req.Msg.CoverImageUrl != nil {
		community.CoverImageURL = req.Msg.CoverImageUrl
	}
	if req.Msg.CreatorId != nil {
		community.CreatorID = req.Msg.CreatorId
	}

	createdCommunity, err := h.communityUseCase.CreateCommunity(ctx, community)
	if err != nil {
		h.logger.LogError(err, "failed to create community", "name", req.Msg.Name)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	h.logger.Info("community created successfully", "id", createdCommunity.ID)
	return connect.NewResponse(&communityv1.CreateCommunityResponse{
		Id: createdCommunity.ID,
	}), nil
}

// UpdateCommunity はコミュニティ更新APIを処理します
func (h *communityHandler) UpdateCommunity(
	ctx context.Context,
	req *connect.Request[communityv1.UpdateCommunityRequest],
) (*connect.Response[communityv1.UpdateCommunityResponse], error) {
	h.logger.Info("UpdateCommunity called", "headers", req.Header(), "id", req.Msg.Id)

	// 既存のコミュニティを取得
	existingCommunity, err := h.communityUseCase.GetCommunity(ctx, req.Msg.Id)
	if err != nil {
		h.logger.LogError(err, "failed to get community for update", "id", req.Msg.Id)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	// 更新フィールドの設定
	if req.Msg.Name != nil {
		existingCommunity.Name = *req.Msg.Name
	}
	if req.Msg.Description != nil {
		existingCommunity.Description = req.Msg.Description
	}
	if req.Msg.ProfileImageUrl != nil {
		existingCommunity.ProfileImageURL = req.Msg.ProfileImageUrl
	}
	if req.Msg.CoverImageUrl != nil {
		existingCommunity.CoverImageURL = req.Msg.CoverImageUrl
	}
	if req.Msg.IsPrivate != nil {
		existingCommunity.IsPrivate = *req.Msg.IsPrivate
	}
	existingCommunity.UpdatedAt = time.Now()

	err = h.communityUseCase.UpdateCommunity(ctx, existingCommunity)
	if err != nil {
		h.logger.LogError(err, "failed to update community", "id", req.Msg.Id)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	h.logger.Info("community updated successfully", "id", req.Msg.Id)
	return connect.NewResponse(&communityv1.UpdateCommunityResponse{
		Success: true,
	}), nil
}

// DeleteCommunity はコミュニティ削除APIを処理します
func (h *communityHandler) DeleteCommunity(
	ctx context.Context,
	req *connect.Request[communityv1.DeleteCommunityRequest],
) (*connect.Response[communityv1.DeleteCommunityResponse], error) {
	h.logger.Info("DeleteCommunity called", "headers", req.Header(), "id", req.Msg.Id)

	err := h.communityUseCase.DeleteCommunity(ctx, req.Msg.Id)
	if err != nil {
		h.logger.LogError(err, "failed to delete community", "id", req.Msg.Id)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	h.logger.Info("community deleted successfully", "id", req.Msg.Id)
	return connect.NewResponse(&communityv1.DeleteCommunityResponse{
		Success: true,
	}), nil
}

// モデルからプロトへの変換ヘルパー関数
func convertToCommunityProto(community *model.Community) *communityv1.Community {
	proto := &communityv1.Community{
		Id:        community.ID,
		Name:      community.Name,
		IsPrivate: community.IsPrivate,
		CreatedAt: timestamppb.New(community.CreatedAt),
		UpdatedAt: timestamppb.New(community.UpdatedAt),
	}

	// オプションフィールドの設定
	if community.Description != nil {
		proto.Description = community.Description
	}
	if community.ProfileImageURL != nil {
		proto.ProfileImageUrl = community.ProfileImageURL
	}
	if community.CoverImageURL != nil {
		proto.CoverImageUrl = community.CoverImageURL
	}
	if community.CreatorID != nil {
		proto.CreatorId = community.CreatorID
	}

	return proto
}

// モデルからプレビュープロトへの変換ヘルパー関数
func convertToCommunityPreviewProto(community *model.Community) *communityv1.CommunityPreview {
	proto := &communityv1.CommunityPreview{
		Id:        community.ID,
		Name:      community.Name,
		IsPrivate: community.IsPrivate,
	}

	// オプションフィールドの設定
	if community.Description != nil {
		proto.Description = community.Description
	}
	if community.ProfileImageURL != nil {
		proto.ProfileImageUrl = community.ProfileImageURL
	}
	// Note: Category and Prefecture are not directly accessible from Community model
	// They would be loaded by the repository if needed

	return proto
}
