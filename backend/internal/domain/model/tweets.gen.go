// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTweet = "tweets"

// Tweet mapped from table <tweets>
type Tweet struct {
	ID        string         `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid();comment:つぶやきの一意識別子（UUIDv4）" json:"id"`                                                                // つぶやきの一意識別子（UUIDv4）
	UserID    string         `gorm:"column:user_id;type:uuid;not null;index:idx_tweets_user_id,priority:1;comment:つぶやきを投稿したユーザーのID" json:"user_id"`                                                // つぶやきを投稿したユーザーのID
	Content   string         `gorm:"column:content;type:text;not null;comment:つぶやきの内容" json:"content"`                                                                                             // つぶやきの内容
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp with time zone;not null;index:idx_tweets_created_at,priority:1;default:CURRENT_TIMESTAMP;comment:つぶやき作成日時" json:"created_at"` // つぶやき作成日時
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP;comment:レコード更新日時" json:"updated_at"`                                        // レコード更新日時
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;index:idx_tweets_deleted_at,priority:1;comment:論理削除日時（NULLは有効なレコードを示す）" json:"deleted_at"`                     // 論理削除日時（NULLは有効なレコードを示す）
	User      User           `json:"user"`
}

// TableName Tweet's table name
func (*Tweet) TableName() string {
	return TableNameTweet
}
