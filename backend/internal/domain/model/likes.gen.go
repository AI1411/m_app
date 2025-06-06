// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameLike = "likes"

// Like mapped from table <likes>
type Like struct {
	ID        int32          `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:いいねID" json:"id"`                                                                   // いいねID
	UserID    *string        `gorm:"column:user_id;type:uuid;index:idx_likes_user_id,priority:1;comment:いいねをしたユーザーID" json:"user_id"`                                                // いいねをしたユーザーID
	TargetID  string         `gorm:"column:target_id;type:uuid;not null;index:idx_likes_target_id,priority:1;comment:いいねの対象となるエンティティのID" json:"target_id"`                           // いいねの対象となるエンティティのID
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp with time zone;not null;index:idx_likes_created_at,priority:1;default:now();comment:いいね作成日時" json:"created_at"` // いいね作成日時
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp with time zone;not null;default:now();comment:レコード更新日時" json:"updated_at"`                                      // レコード更新日時
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;index:idx_likes_deleted_at,priority:1;comment:論理削除日時（NULLは有効なレコードを示す）" json:"deleted_at"`        // 論理削除日時（NULLは有効なレコードを示す）
	User      User           `json:"user"`
	LikedUser User           `gorm:"foreignKey:LikedUserID;references:ID" json:"liked_user"`
}

// TableName Like's table name
func (*Like) TableName() string {
	return TableNameLike
}
