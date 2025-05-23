package datastore

import (
	"context"

	"gorm.io/gorm"

	"github.com/AI1411/m_app/internal/domain/model"
	"github.com/AI1411/m_app/internal/infra/db"
)

// NotificationRepository は通知データへのアクセスを提供するインターフェース
type NotificationRepository interface {
	GetNotificationByID(ctx context.Context, id string) (*model.Notification, error)
	GetNotifications(ctx context.Context, userID string, page, pageSize int, unreadOnly bool) ([]*model.Notification, int64, int64, error)
	MarkNotificationAsRead(ctx context.Context, notificationID, userID string) error
	MarkAllNotificationsAsRead(ctx context.Context, userID string) (int64, error)
	CreateNotification(ctx context.Context, notification *model.Notification) (*model.Notification, error)
	DeleteNotification(ctx context.Context, notificationID, userID string) error
}

// notificationRepository はNotificationRepositoryインターフェースの実装
type notificationRepository struct {
	sqlHandler *db.SqlHandler
}

// NewNotificationRepository は新しいNotificationRepositoryを作成します
func NewNotificationRepository(sqlHandler *db.SqlHandler) NotificationRepository {
	return &notificationRepository{
		sqlHandler: sqlHandler,
	}
}

// GetNotificationByID はIDで通知を取得します
func (r *notificationRepository) GetNotificationByID(ctx context.Context, id string) (*model.Notification, error) {
	var notification model.Notification
	result := r.sqlHandler.Conn.WithContext(ctx).Where("id = ?", id).First(&notification)
	if result.Error != nil {
		return nil, result.Error
	}
	return &notification, nil
}

// GetNotifications はユーザーの通知リストを取得します
func (r *notificationRepository) GetNotifications(ctx context.Context, userID string, page, pageSize int, unreadOnly bool) ([]*model.Notification, int64, int64, error) {
	var notifications []*model.Notification
	var totalCount int64
	var unreadCount int64

	// クエリの構築
	query := r.sqlHandler.Conn.WithContext(ctx).Model(&model.Notification{}).Where("user_id = ?", userID)

	// 未読のみのフィルタリング
	if unreadOnly {
		query = query.Where("is_read = ?", false)
	}

	// 未読通知の総数を取得
	if err := r.sqlHandler.Conn.WithContext(ctx).Model(&model.Notification{}).
		Where("user_id = ? AND is_read = ?", userID, false).
		Count(&unreadCount).Error; err != nil {
		return nil, 0, 0, err
	}

	// 総件数の取得
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, 0, err
	}

	// ページネーション
	offset := (page - 1) * pageSize
	result := query.Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&notifications)

	if result.Error != nil {
		return nil, 0, 0, result.Error
	}

	return notifications, totalCount, unreadCount, nil
}

// MarkNotificationAsRead は通知を既読にします
func (r *notificationRepository) MarkNotificationAsRead(ctx context.Context, notificationID, userID string) error {
	result := r.sqlHandler.Conn.WithContext(ctx).Model(&model.Notification{}).
		Where("id = ? AND user_id = ?", notificationID, userID).
		Update("is_read", true)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// MarkAllNotificationsAsRead はユーザーの全通知を既読にします
func (r *notificationRepository) MarkAllNotificationsAsRead(ctx context.Context, userID string) (int64, error) {
	result := r.sqlHandler.Conn.WithContext(ctx).Model(&model.Notification{}).
		Where("user_id = ? AND is_read = ?", userID, false).
		Update("is_read", true)

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

// CreateNotification は新しい通知を作成します
func (r *notificationRepository) CreateNotification(ctx context.Context, notification *model.Notification) (*model.Notification, error) {
	result := r.sqlHandler.Conn.WithContext(ctx).Create(notification)
	if result.Error != nil {
		return nil, result.Error
	}
	return notification, nil
}

// DeleteNotification は通知を削除します
func (r *notificationRepository) DeleteNotification(ctx context.Context, notificationID, userID string) error {
	result := r.sqlHandler.Conn.WithContext(ctx).
		Where("id = ? AND user_id = ?", notificationID, userID).
		Delete(&model.Notification{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
