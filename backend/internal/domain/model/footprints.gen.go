// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameFootprint = "footprints"

// Footprint mapped from table <footprints>
type Footprint struct {
	ID            string         `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid();comment:足あとの一意識別子（UUIDv4）" json:"id"`                                                                       // 足あとの一意識別子（UUIDv4）
	VisitorUserID string         `gorm:"column:visitor_user_id;type:uuid;not null;index:idx_footprints_visitor_user_id,priority:1;comment:プロフィールを閲覧したユーザーのID" json:"visitor_user_id"`                        // プロフィールを閲覧したユーザーのID
	VisitedUserID string         `gorm:"column:visited_user_id;type:uuid;not null;index:idx_footprints_visited_user_id,priority:1;comment:プロフィールを閲覧されたユーザーのID" json:"visited_user_id"`                       // プロフィールを閲覧されたユーザーのID
	VisitedAt     time.Time      `gorm:"column:visited_at;type:timestamp with time zone;not null;index:idx_footprints_visited_at,priority:1;default:CURRENT_TIMESTAMP;comment:プロフィール閲覧日時" json:"visited_at"` // プロフィール閲覧日時
	CreatedAt     time.Time      `gorm:"column:created_at;type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP;comment:レコード作成日時" json:"created_at"`                                              // レコード作成日時
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP;comment:レコード更新日時" json:"updated_at"`                                              // レコード更新日時
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;index:idx_footprints_deleted_at,priority:1;comment:論理削除日時（NULLは有効なレコードを示す）" json:"deleted_at"`                       // 論理削除日時（NULLは有効なレコードを示す）
	User          User           `json:"user"`
	VisitedUser   User           `gorm:"foreignKey:VisitedUserID;references:ID" json:"visited_user"`
}

// TableName Footprint's table name
func (*Footprint) TableName() string {
	return TableNameFootprint
}
