package handler

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"gorm.io/gorm"

	prefecturev1 "github.com/AI1411/m_app/gen/prefecture/v1"
	regionv1 "github.com/AI1411/m_app/gen/region/v1"
	"github.com/AI1411/m_app/internal/infra/db"
	"github.com/AI1411/m_app/internal/infra/logger"
)

// Prefecture は都道府県モデル
type Prefecture struct {
	ID        int32          `gorm:"column:id;primaryKey;autoIncrement:true"`
	Code      int16          `gorm:"column:code;not null"`
	Name      string         `gorm:"column:name;not null"`
	NameEn    string         `gorm:"column:name_en;not null"`
	RegionID  int32          `gorm:"column:region_id;not null"`
	CreatedAt time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

// TableName はテーブル名を返します
func (Prefecture) TableName() string {
	return "prefectures"
}

// PrefectureHandler は都道府県関連APIのハンドラーインターフェース
type PrefectureHandler interface {
	ListPrefectures(context.Context, *connect.Request[prefecturev1.ListPrefecturesRequest]) (*connect.Response[prefecturev1.ListPrefecturesResponse], error)
	GetPrefecture(context.Context, *connect.Request[prefecturev1.GetPrefectureRequest]) (*connect.Response[prefecturev1.GetPrefectureResponse], error)
}

// prefectureHandler はPrefectureHandlerインターフェースの実装
type prefectureHandler struct {
	sqlHandler *db.SqlHandler
	logger     *logger.Logger
}

// NewPrefectureHandler は新しいPrefectureHandlerインスタンスを作成します
func NewPrefectureHandler(sqlHandler *db.SqlHandler, logger *logger.Logger) PrefectureHandler {
	return &prefectureHandler{
		sqlHandler: sqlHandler,
		logger:     logger,
	}
}

// ListPrefectures は都道府県一覧取得APIを処理します
func (h *prefectureHandler) ListPrefectures(
	ctx context.Context,
	req *connect.Request[prefecturev1.ListPrefecturesRequest],
) (*connect.Response[prefecturev1.ListPrefecturesResponse], error) {
	h.logger.Info("ListPrefectures called", "headers", req.Header())

	query := h.sqlHandler.Conn.Model(&Prefecture{})

	if req.Msg.RegionId != nil {
		query = query.Where("region_id = ?", *req.Msg.RegionId)
	}

	var prefectures []*Prefecture
	err := query.Order("code").Find(&prefectures).Error
	if err != nil {
		h.logger.LogError(err, "failed to list prefectures")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	var prefectureProtos []*prefecturev1.Prefecture
	for _, p := range prefectures {
		prefectureProtos = append(prefectureProtos, &prefecturev1.Prefecture{
			Id:       p.ID,
			Code:     int32(p.Code),
			Name:     p.Name,
			NameEn:   p.NameEn,
			RegionId: p.RegionID,
			Region: &regionv1.Region{
				Id: p.RegionID,
				// Note: We're not loading the full region data here
				// In a real implementation, you might want to join with the regions table
			},
		})
	}

	h.logger.Info("prefectures retrieved successfully", "count", len(prefectureProtos))
	return connect.NewResponse(&prefecturev1.ListPrefecturesResponse{
		Prefectures: prefectureProtos,
	}), nil
}

// GetPrefecture は都道府県詳細取得APIを処理します
func (h *prefectureHandler) GetPrefecture(
	ctx context.Context,
	req *connect.Request[prefecturev1.GetPrefectureRequest],
) (*connect.Response[prefecturev1.GetPrefectureResponse], error) {
	h.logger.Info("GetPrefecture called", "headers", req.Header(), "prefecture_id", req.Msg.Id)

	var prefecture Prefecture
	err := h.sqlHandler.Conn.Model(&Prefecture{}).
		Where("id = ?", req.Msg.Id).
		First(&prefecture).Error
	if err != nil {
		h.logger.LogError(err, "failed to get prefecture", "prefecture_id", req.Msg.Id)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// 関連する地域情報を取得
	var region struct {
		ID        int32  `gorm:"column:id"`
		Name      string `gorm:"column:name"`
		NameEn    string `gorm:"column:name_en"`
		SortOrder int32  `gorm:"column:sort_order"`
	}
	err = h.sqlHandler.Conn.Table("regions").
		Select("id, name, name_en, sort_order").
		Where("id = ?", prefecture.RegionID).
		First(&region).Error
	if err != nil {
		h.logger.LogError(err, "failed to get region for prefecture", "region_id", prefecture.RegionID)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	prefectureProto := &prefecturev1.Prefecture{
		Id:       prefecture.ID,
		Code:     int32(prefecture.Code),
		Name:     prefecture.Name,
		NameEn:   prefecture.NameEn,
		RegionId: prefecture.RegionID,
		Region: &regionv1.Region{
			Id:        region.ID,
			Name:      region.Name,
			NameEn:    region.NameEn,
			SortOrder: region.SortOrder,
		},
	}

	h.logger.Info("prefecture retrieved successfully", "prefecture_id", prefecture.ID)
	return connect.NewResponse(&prefecturev1.GetPrefectureResponse{
		Prefecture: prefectureProto,
	}), nil
}
