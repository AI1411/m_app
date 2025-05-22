package handler

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/timestamppb"

	notificationv1 "github.com/AI1411/m_app/gen/notification/v1"
	"github.com/AI1411/m_app/internal/domain/model"
	"github.com/AI1411/m_app/internal/infra/logger"
	"github.com/AI1411/m_app/internal/usecase"
)

// NotificationHandler は通知関連APIのハンドラーインターフェース
type NotificationHandler interface {
	GetNotifications(context.Context, *connect.Request[notificationv1.GetNotificationsRequest]) (*connect.Response[notificationv1.GetNotificationsResponse], error)
	MarkNotificationAsRead(context.Context, *connect.Request[notificationv1.MarkNotificationAsReadRequest]) (*connect.Response[notificationv1.MarkNotificationAsReadResponse], error)
	MarkAllNotificationsAsRead(context.Context, *connect.Request[notificationv1.MarkAllNotificationsAsReadRequest]) (*connect.Response[notificationv1.MarkAllNotificationsAsReadResponse], error)
	CreateNotification(context.Context, *connect.Request[notificationv1.CreateNotificationRequest]) (*connect.Response[notificationv1.CreateNotificationResponse], error)
	DeleteNotification(context.Context, *connect.Request[notificationv1.DeleteNotificationRequest]) (*connect.Response[notificationv1.DeleteNotificationResponse], error)
}

// notificationHandler はNotificationHandlerインターフェースの実装
type notificationHandler struct {
	notificationUseCase usecase.NotificationUseCase
	logger              *logger.Logger
}

// NewNotificationHandler は新しいNotificationHandlerインスタンスを作成します
func NewNotificationHandler(notificationUseCase usecase.NotificationUseCase, logger *logger.Logger) NotificationHandler {
	return &notificationHandler{
		notificationUseCase: notificationUseCase,
		logger:              logger,
	}
}

// GetNotifications はユーザーの通知リストを取得します
func (h *notificationHandler) GetNotifications(
	ctx context.Context,
	req *connect.Request[notificationv1.GetNotificationsRequest],
) (*connect.Response[notificationv1.GetNotificationsResponse], error) {
	h.logger.Info("GetNotifications called", "user_id", req.Msg.UserId, "headers", req.Header())

	// デフォルト値の設定
	page := int(req.Msg.Page)
	if page <= 0 {
		page = 1
	}
	pageSize := int(req.Msg.PageSize)
	if pageSize <= 0 {
		pageSize = 10
	}
	unreadOnly := false
	if req.Msg.UnreadOnly != nil {
		unreadOnly = *req.Msg.UnreadOnly
	}

	notifications, totalCount, unreadCount, err := h.notificationUseCase.GetNotifications(
		ctx,
		req.Msg.UserId,
		page,
		pageSize,
		unreadOnly,
	)
	if err != nil {
		h.logger.LogError(err, "failed to get notifications", "user_id", req.Msg.UserId)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// モデルからDTOへの変換
	var notificationDTOs []*notificationv1.Notification
	for _, notification := range notifications {
		notificationDTOs = append(notificationDTOs, convertToNotificationDTO(notification))
	}

	h.logger.Info("notifications retrieved successfully", "user_id", req.Msg.UserId, "count", len(notificationDTOs))
	return connect.NewResponse(&notificationv1.GetNotificationsResponse{
		Notifications: notificationDTOs,
		TotalCount:    int32(totalCount),
		Page:          int32(page),
		PageSize:      int32(pageSize),
		UnreadCount:   int32(unreadCount),
	}), nil
}

// MarkNotificationAsRead は通知を既読にします
func (h *notificationHandler) MarkNotificationAsRead(
	ctx context.Context,
	req *connect.Request[notificationv1.MarkNotificationAsReadRequest],
) (*connect.Response[notificationv1.MarkNotificationAsReadResponse], error) {
	h.logger.Info("MarkNotificationAsRead called", "notification_id", req.Msg.NotificationId, "user_id", req.Msg.UserId, "headers", req.Header())

	err := h.notificationUseCase.MarkNotificationAsRead(ctx, req.Msg.NotificationId, req.Msg.UserId)
	if err != nil {
		h.logger.LogError(err, "failed to mark notification as read", "notification_id", req.Msg.NotificationId, "user_id", req.Msg.UserId)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	h.logger.Info("notification marked as read successfully", "notification_id", req.Msg.NotificationId, "user_id", req.Msg.UserId)
	return connect.NewResponse(&notificationv1.MarkNotificationAsReadResponse{
		Success: true,
	}), nil
}

// MarkAllNotificationsAsRead はユーザーの全通知を既読にします
func (h *notificationHandler) MarkAllNotificationsAsRead(
	ctx context.Context,
	req *connect.Request[notificationv1.MarkAllNotificationsAsReadRequest],
) (*connect.Response[notificationv1.MarkAllNotificationsAsReadResponse], error) {
	h.logger.Info("MarkAllNotificationsAsRead called", "user_id", req.Msg.UserId, "headers", req.Header())

	updatedCount, err := h.notificationUseCase.MarkAllNotificationsAsRead(ctx, req.Msg.UserId)
	if err != nil {
		h.logger.LogError(err, "failed to mark all notifications as read", "user_id", req.Msg.UserId)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	h.logger.Info("all notifications marked as read successfully", "user_id", req.Msg.UserId, "updated_count", updatedCount)
	return connect.NewResponse(&notificationv1.MarkAllNotificationsAsReadResponse{
		UpdatedCount: int32(updatedCount),
	}), nil
}

// CreateNotification は新しい通知を作成します
func (h *notificationHandler) CreateNotification(
	ctx context.Context,
	req *connect.Request[notificationv1.CreateNotificationRequest],
) (*connect.Response[notificationv1.CreateNotificationResponse], error) {
	h.logger.Info("CreateNotification called", "user_id", req.Msg.UserId, "title", req.Msg.Title, "headers", req.Header())

	// DTOからモデルへの変換
	notification := &model.Notification{
		UserID:           req.Msg.UserId,
		Title:            req.Msg.Title,
		Content:          req.Msg.Content,
		NotificationType: req.Msg.NotificationType,
		IsRead:           false,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	// オプショナルフィールドの設定
	if req.Msg.RelatedResourceId != nil {
		notification.RelatedResourceID = req.Msg.RelatedResourceId
	}
	if req.Msg.RelatedResourceType != nil {
		notification.RelatedResourceType = req.Msg.RelatedResourceType
	}

	createdNotification, err := h.notificationUseCase.CreateNotification(ctx, notification)
	if err != nil {
		h.logger.LogError(err, "failed to create notification", "user_id", req.Msg.UserId)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	h.logger.Info("notification created successfully", "notification_id", createdNotification.ID, "user_id", req.Msg.UserId)
	return connect.NewResponse(&notificationv1.CreateNotificationResponse{
		Id: createdNotification.ID,
	}), nil
}

// DeleteNotification は通知を削除します
func (h *notificationHandler) DeleteNotification(
	ctx context.Context,
	req *connect.Request[notificationv1.DeleteNotificationRequest],
) (*connect.Response[notificationv1.DeleteNotificationResponse], error) {
	h.logger.Info("DeleteNotification called", "notification_id", req.Msg.NotificationId, "user_id", req.Msg.UserId, "headers", req.Header())

	err := h.notificationUseCase.DeleteNotification(ctx, req.Msg.NotificationId, req.Msg.UserId)
	if err != nil {
		h.logger.LogError(err, "failed to delete notification", "notification_id", req.Msg.NotificationId, "user_id", req.Msg.UserId)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	h.logger.Info("notification deleted successfully", "notification_id", req.Msg.NotificationId, "user_id", req.Msg.UserId)
	return connect.NewResponse(&notificationv1.DeleteNotificationResponse{
		Success: true,
	}), nil
}

// convertToNotificationDTO はモデルからDTOへの変換を行います
func convertToNotificationDTO(notification *model.Notification) *notificationv1.Notification {
	dto := &notificationv1.Notification{
		Id:               notification.ID,
		UserId:           notification.UserID,
		Title:            notification.Title,
		Content:          notification.Content,
		NotificationType: notification.NotificationType,
		IsRead:           notification.IsRead,
		CreatedAt:        timestamppb.New(notification.CreatedAt),
		UpdatedAt:        timestamppb.New(notification.UpdatedAt),
	}

	// オプショナルフィールドの設定
	if notification.RelatedResourceID != nil {
		dto.RelatedResourceId = notification.RelatedResourceID
	}
	if notification.RelatedResourceType != nil {
		dto.RelatedResourceType = notification.RelatedResourceType
	}

	return dto
}
