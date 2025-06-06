// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameNotification = "notifications"

// Notification mapped from table <notifications>
type Notification struct {
	ID                  string         `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid();comment:通知の一意識別子（UUIDv4）" json:"id"`                                                                                           // 通知の一意識別子（UUIDv4）
	UserID              string         `gorm:"column:user_id;type:uuid;not null;index:idx_notifications_user_id,priority:1;comment:通知の受信者となるユーザーID" json:"user_id"`                                                                   // 通知の受信者となるユーザーID
	Title               string         `gorm:"column:title;type:character varying(255);not null;comment:通知のタイトル" json:"title"`                                                                                                        // 通知のタイトル
	Content             string         `gorm:"column:content;type:text;not null;comment:通知の内容" json:"content"`                                                                                                                        // 通知の内容
	NotificationType    string         `gorm:"column:notification_type;type:character varying(50);not null;index:idx_notifications_notification_type,priority:1;comment:通知のタイプ（system, like, follow, etc.）" json:"notification_type"` // 通知のタイプ（system, like, follow, etc.）
	RelatedResourceID   *string        `gorm:"column:related_resource_id;type:uuid;comment:関連リソースのID（ツイートID、ユーザーIDなど）" json:"related_resource_id"`                                                                                    // 関連リソースのID（ツイートID、ユーザーIDなど）
	RelatedResourceType *string        `gorm:"column:related_resource_type;type:character varying(50);comment:関連リソースのタイプ（tweet, user, etc.）" json:"related_resource_type"`                                                            // 関連リソースのタイプ（tweet, user, etc.）
	IsRead              bool           `gorm:"column:is_read;type:boolean;not null;index:idx_notifications_is_read,priority:1;comment:既読フラグ（true: 既読, false: 未読）" json:"is_read"`                                                     // 既読フラグ（true: 既読, false: 未読）
	CreatedAt           time.Time      `gorm:"column:created_at;type:timestamp with time zone;not null;index:idx_notifications_created_at,priority:1;default:CURRENT_TIMESTAMP;comment:通知作成日時" json:"created_at"`                     // 通知作成日時
	UpdatedAt           time.Time      `gorm:"column:updated_at;type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP;comment:レコード更新日時" json:"updated_at"`                                                                 // レコード更新日時
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;index:idx_notifications_deleted_at,priority:1;comment:論理削除日時（NULLは有効なレコードを示す）" json:"deleted_at"`                                       // 論理削除日時（NULLは有効なレコードを示す）
	Users               []User         `json:"users"`
}

// TableName Notification's table name
func (*Notification) TableName() string {
	return TableNameNotification
}
