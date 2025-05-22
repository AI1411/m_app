package usecase

import (
	"context"

	"github.com/AI1411/m_app/internal/domain/model"
	"github.com/AI1411/m_app/internal/infra/repository/datastore"
)

// NotificationUseCase は通知関連のビジネスロジックを提供するインターフェース
type NotificationUseCase interface {
	GetNotifications(ctx context.Context, userID string, page, pageSize int, unreadOnly bool) ([]*model.Notification, int64, int64, error)
	MarkNotificationAsRead(ctx context.Context, notificationID, userID string) error
	MarkAllNotificationsAsRead(ctx context.Context, userID string) (int64, error)
	CreateNotification(ctx context.Context, notification *model.Notification) (*model.Notification, error)
	DeleteNotification(ctx context.Context, notificationID, userID string) error
}

// notificationUseCase はNotificationUseCaseインターフェースの実装
type notificationUseCase struct {
	notificationRepo datastore.NotificationRepository
}

// NewNotificationUseCase は新しいNotificationUseCaseを作成します
func NewNotificationUseCase(notificationRepo datastore.NotificationRepository) NotificationUseCase {
	return &notificationUseCase{
		notificationRepo: notificationRepo,
	}
}

// GetNotifications はユーザーの通知リストを取得します
func (u *notificationUseCase) GetNotifications(ctx context.Context, userID string, page, pageSize int, unreadOnly bool) ([]*model.Notification, int64, int64, error) {
	return u.notificationRepo.GetNotifications(ctx, userID, page, pageSize, unreadOnly)
}

// MarkNotificationAsRead は通知を既読にします
func (u *notificationUseCase) MarkNotificationAsRead(ctx context.Context, notificationID, userID string) error {
	return u.notificationRepo.MarkNotificationAsRead(ctx, notificationID, userID)
}

// MarkAllNotificationsAsRead はユーザーの全通知を既読にします
func (u *notificationUseCase) MarkAllNotificationsAsRead(ctx context.Context, userID string) (int64, error) {
	return u.notificationRepo.MarkAllNotificationsAsRead(ctx, userID)
}

// CreateNotification は新しい通知を作成します
func (u *notificationUseCase) CreateNotification(ctx context.Context, notification *model.Notification) (*model.Notification, error) {
	return u.notificationRepo.CreateNotification(ctx, notification)
}

// DeleteNotification は通知を削除します
func (u *notificationUseCase) DeleteNotification(ctx context.Context, notificationID, userID string) error {
	return u.notificationRepo.DeleteNotification(ctx, notificationID, userID)
}
