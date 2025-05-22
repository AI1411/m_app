-- ユーザー画像テーブルの作成
DROP TABLE IF EXISTS user_images CASCADE;
CREATE TABLE user_images
(
    id          UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    user_id     UUID                     NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    image_url   VARCHAR(255)             NOT NULL,
    description TEXT,
    is_primary  BOOLEAN                  NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP WITH TIME ZONE
);

-- テーブルコメント
COMMENT
ON TABLE user_images IS 'ユーザーの画像を格納するテーブル';

-- カラムコメント
COMMENT
ON COLUMN user_images.id IS '画像の一意識別子（UUIDv4）';
COMMENT
ON COLUMN user_images.user_id IS '画像の所有者であるユーザーのID';
COMMENT
ON COLUMN user_images.image_url IS '画像のURL';
COMMENT
ON COLUMN user_images.description IS '画像の説明';
COMMENT
ON COLUMN user_images.is_primary IS 'プロフィールのメイン画像かどうか（trueの場合はメイン画像）';
COMMENT
ON COLUMN user_images.created_at IS 'レコード作成日時';
COMMENT
ON COLUMN user_images.updated_at IS 'レコード更新日時';
COMMENT
ON COLUMN user_images.deleted_at IS '論理削除日時（NULLは有効なレコードを示す）';

-- インデックス
CREATE INDEX idx_user_images_user_id ON user_images (user_id);
CREATE INDEX idx_user_images_is_primary ON user_images (is_primary);
CREATE INDEX idx_user_images_created_at ON user_images (created_at);
CREATE INDEX idx_user_images_deleted_at ON user_images (deleted_at) WHERE deleted_at IS NULL;