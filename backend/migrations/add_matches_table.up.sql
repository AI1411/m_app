-- マッチングテーブル（お互いにいいねを送り合ったユーザーのペア）
CREATE TABLE matches
(
    id         UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    user_id_1  UUID                     NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    user_id_2  UUID                     NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    matched_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT uk_matches_users UNIQUE (user_id_1, user_id_2)
);

-- テーブルコメント
COMMENT ON TABLE matches IS 'お互いにいいねを送り合い、マッチングが成立したユーザーペアの情報を格納するテーブル';

-- カラムコメント
COMMENT ON COLUMN matches.id IS 'マッチングの一意識別子（UUIDv4）';
COMMENT ON COLUMN matches.user_id_1 IS 'マッチングしたユーザー1のID';
COMMENT ON COLUMN matches.user_id_2 IS 'マッチングしたユーザー2のID';
COMMENT ON COLUMN matches.matched_at IS 'マッチング成立日時';
COMMENT ON COLUMN matches.created_at IS 'レコード作成日時';
COMMENT ON COLUMN matches.updated_at IS 'レコード更新日時';
COMMENT ON COLUMN matches.deleted_at IS '論理削除日時（NULLは有効なレコードを示す）';

-- インデックス
CREATE INDEX idx_matches_user_id_1 ON matches (user_id_1);
CREATE INDEX idx_matches_user_id_2 ON matches (user_id_2);
CREATE INDEX idx_matches_matched_at ON matches (matched_at);
CREATE INDEX idx_matches_deleted_at ON matches (deleted_at) WHERE deleted_at IS NULL;