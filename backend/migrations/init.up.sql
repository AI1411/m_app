DROP TABLE IF EXISTS tweets;
DROP TABLE IF EXISTS likes;
DROP TABLE IF EXISTS community_members;
DROP TABLE IF EXISTS communities;
DROP TABLE IF EXISTS user_interests;
DROP TABLE IF EXISTS interests;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS educations;
DROP TABLE IF EXISTS prefectures;
DROP TABLE IF EXISTS regions;
DROP TABLE IF EXISTS blocklists;

-- 1. 地域（リージョン）テーブルの作成
CREATE TABLE regions
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(20)              NOT NULL UNIQUE,
    name_en    VARCHAR(30)              NOT NULL UNIQUE,
    sort_order INTEGER                  NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- テーブルコメント
COMMENT ON TABLE regions IS '地域区分マスタ';
COMMENT ON COLUMN regions.id IS '地域ID';
COMMENT ON COLUMN regions.name IS '地域名（日本語）';
COMMENT ON COLUMN regions.name_en IS '地域名（英語）';
COMMENT ON COLUMN regions.sort_order IS '表示順序';
COMMENT ON COLUMN regions.created_at IS '作成日時';
COMMENT ON COLUMN regions.updated_at IS '更新日時';
COMMENT ON COLUMN regions.deleted_at IS '削除日時（論理削除用）';

-- 地域の初期データ
INSERT INTO regions (name, name_en, sort_order)
VALUES ('北海道', 'Hokkaido', 1),
       ('東北', 'Tohoku', 2),
       ('関東', 'Kanto', 3),
       ('中部', 'Chubu', 4),
       ('関西', 'Kansai', 5),
       ('中国', 'Chugoku', 6),
       ('四国', 'Shikoku', 7),
       ('九州', 'Kyushu', 8),
       ('沖縄', 'Okinawa', 9);

-- 2. 都道府県テーブル（リージョンテーブルへの外部キー付き）
CREATE TABLE prefectures
(
    id         SERIAL PRIMARY KEY,
    code       SMALLINT                                           NOT NULL, -- 都道府県コード（JIS X 0401に準拠: 01-47）
    name       VARCHAR(4)                                         NOT NULL, -- 都道府県名（例：東京都、大阪府）
    name_en    VARCHAR(30)                                        NOT NULL, -- 英語名（例：Tokyo, Osaka）
    region_id  INTEGER REFERENCES regions (id)                    NOT NULL, -- 地域区分ID
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT uk_prefecture_code UNIQUE (code)
);

-- インデックス
CREATE INDEX IF NOT EXISTS idx_prefectures_region_id ON prefectures (region_id);
CREATE INDEX IF NOT EXISTS idx_prefectures_deleted_at ON prefectures (deleted_at);

-- コメント
COMMENT ON TABLE prefectures IS '都道府県マスタ';
COMMENT ON COLUMN prefectures.id IS 'プライマリーキー';
COMMENT ON COLUMN prefectures.code IS '都道府県コード（JIS X 0401準拠）';
COMMENT ON COLUMN prefectures.name IS '都道府県名';
COMMENT ON COLUMN prefectures.name_en IS '都道府県名（英語）';
COMMENT ON COLUMN prefectures.region_id IS '地域区分ID';
COMMENT ON COLUMN prefectures.created_at IS '作成日時';
COMMENT ON COLUMN prefectures.updated_at IS '更新日時';
COMMENT ON COLUMN prefectures.deleted_at IS '削除日時（論理削除用）';

-- 初期データ挿入（地域IDをリージョンテーブルから参照）
INSERT INTO prefectures (code, name, name_en, region_id)
VALUES (1, '北海道', 'Hokkaido', (SELECT id FROM regions WHERE name = '北海道')),
       (2, '青森県', 'Aomori', (SELECT id FROM regions WHERE name = '東北')),
       (3, '岩手県', 'Iwate', (SELECT id FROM regions WHERE name = '東北')),
       (4, '宮城県', 'Miyagi', (SELECT id FROM regions WHERE name = '東北')),
       (5, '秋田県', 'Akita', (SELECT id FROM regions WHERE name = '東北')),
       (6, '山形県', 'Yamagata', (SELECT id FROM regions WHERE name = '東北')),
       (7, '福島県', 'Fukushima', (SELECT id FROM regions WHERE name = '東北')),
       (8, '茨城県', 'Ibaraki', (SELECT id FROM regions WHERE name = '関東')),
       (9, '栃木県', 'Tochigi', (SELECT id FROM regions WHERE name = '関東')),
       (10, '群馬県', 'Gunma', (SELECT id FROM regions WHERE name = '関東')),
       (11, '埼玉県', 'Saitama', (SELECT id FROM regions WHERE name = '関東')),
       (12, '千葉県', 'Chiba', (SELECT id FROM regions WHERE name = '関東')),
       (13, '東京都', 'Tokyo', (SELECT id FROM regions WHERE name = '関東')),
       (14, '神奈川県', 'Kanagawa', (SELECT id FROM regions WHERE name = '関東')),
       (15, '新潟県', 'Niigata', (SELECT id FROM regions WHERE name = '中部')),
       (16, '富山県', 'Toyama', (SELECT id FROM regions WHERE name = '中部')),
       (17, '石川県', 'Ishikawa', (SELECT id FROM regions WHERE name = '中部')),
       (18, '福井県', 'Fukui', (SELECT id FROM regions WHERE name = '中部')),
       (19, '山梨県', 'Yamanashi', (SELECT id FROM regions WHERE name = '中部')),
       (20, '長野県', 'Nagano', (SELECT id FROM regions WHERE name = '中部')),
       (21, '岐阜県', 'Gifu', (SELECT id FROM regions WHERE name = '中部')),
       (22, '静岡県', 'Shizuoka', (SELECT id FROM regions WHERE name = '中部')),
       (23, '愛知県', 'Aichi', (SELECT id FROM regions WHERE name = '中部')),
       (24, '三重県', 'Mie', (SELECT id FROM regions WHERE name = '関西')),
       (25, '滋賀県', 'Shiga', (SELECT id FROM regions WHERE name = '関西')),
       (26, '京都府', 'Kyoto', (SELECT id FROM regions WHERE name = '関西')),
       (27, '大阪府', 'Osaka', (SELECT id FROM regions WHERE name = '関西')),
       (28, '兵庫県', 'Hyogo', (SELECT id FROM regions WHERE name = '関西')),
       (29, '奈良県', 'Nara', (SELECT id FROM regions WHERE name = '関西')),
       (30, '和歌山県', 'Wakayama', (SELECT id FROM regions WHERE name = '関西')),
       (31, '鳥取県', 'Tottori', (SELECT id FROM regions WHERE name = '中国')),
       (32, '島根県', 'Shimane', (SELECT id FROM regions WHERE name = '中国')),
       (33, '岡山県', 'Okayama', (SELECT id FROM regions WHERE name = '中国')),
       (34, '広島県', 'Hiroshima', (SELECT id FROM regions WHERE name = '中国')),
       (35, '山口県', 'Yamaguchi', (SELECT id FROM regions WHERE name = '中国')),
       (36, '徳島県', 'Tokushima', (SELECT id FROM regions WHERE name = '四国')),
       (37, '香川県', 'Kagawa', (SELECT id FROM regions WHERE name = '四国')),
       (38, '愛媛県', 'Ehime', (SELECT id FROM regions WHERE name = '四国')),
       (39, '高知県', 'Kochi', (SELECT id FROM regions WHERE name = '四国')),
       (40, '福岡県', 'Fukuoka', (SELECT id FROM regions WHERE name = '九州')),
       (41, '佐賀県', 'Saga', (SELECT id FROM regions WHERE name = '九州')),
       (42, '長崎県', 'Nagasaki', (SELECT id FROM regions WHERE name = '九州')),
       (43, '熊本県', 'Kumamoto', (SELECT id FROM regions WHERE name = '九州')),
       (44, '大分県', 'Oita', (SELECT id FROM regions WHERE name = '九州')),
       (45, '宮崎県', 'Miyazaki', (SELECT id FROM regions WHERE name = '九州')),
       (46, '鹿児島県', 'Kagoshima', (SELECT id FROM regions WHERE name = '九州')),
       (47, '沖縄県', 'Okinawa', (SELECT id FROM regions WHERE name = '沖縄'));

-- 3. 学歴マスターテーブルの作成
CREATE TABLE educations
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(100)             NOT NULL UNIQUE,
    sort_order INTEGER                  NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- テーブルコメント
COMMENT ON TABLE educations IS '学歴を格納するマスターテーブル';

-- カラムコメント
COMMENT ON COLUMN educations.id IS '学歴ID';
COMMENT ON COLUMN educations.name IS '学歴名（高卒、専門卒、大卒など）';
COMMENT ON COLUMN educations.sort_order IS '表示順序';
COMMENT ON COLUMN educations.created_at IS 'レコード作成日時';
COMMENT ON COLUMN educations.updated_at IS 'レコード更新日時';
COMMENT ON COLUMN educations.deleted_at IS '論理削除日時（NULLは有効なレコードを示す）';

-- 初期データの挿入
INSERT INTO educations (name, sort_order)
VALUES ('中卒', 1),
       ('高卒', 2),
       ('専門卒', 3),
       ('短大卒', 4),
       ('大卒', 5),
       ('大学院卒', 6);

-- 4. カテゴリマスターテーブルの作成
CREATE TABLE categories
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(50)              NOT NULL UNIQUE,
    display_name VARCHAR(100)             NOT NULL,
    description  TEXT,
    icon_url     VARCHAR(255),
    color_code   VARCHAR(20),
    sort_order   INTEGER                  NOT NULL DEFAULT 0,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at   TIMESTAMP WITH TIME ZONE
);

-- テーブルコメント
COMMENT ON TABLE categories IS '趣味のカテゴリを格納するマスターテーブル';

-- カラムコメント
COMMENT ON COLUMN categories.id IS 'カテゴリID';
COMMENT ON COLUMN categories.name IS 'カテゴリの識別子（英語、小文字、アンダースコア区切り）';
COMMENT ON COLUMN categories.display_name IS 'カテゴリの表示名（ユーザーに表示される名前）';
COMMENT ON COLUMN categories.description IS 'カテゴリの説明';
COMMENT ON COLUMN categories.icon_url IS 'カテゴリを表すアイコンのURL';
COMMENT ON COLUMN categories.color_code IS 'カテゴリの色コード（UI表示用）';
COMMENT ON COLUMN categories.sort_order IS '表示順序';
COMMENT ON COLUMN categories.created_at IS 'レコード作成日時';
COMMENT ON COLUMN categories.updated_at IS 'レコード更新日時';
COMMENT ON COLUMN categories.deleted_at IS '論理削除日時（NULLは有効なレコードを示す）';

-- カテゴリの初期データ
INSERT INTO categories (name, display_name, description, color_code, sort_order)
VALUES ('sports', 'スポーツ', '各種スポーツやフィットネス活動', '#4CAF50', 10),
       ('outdoor', 'アウトドア', '屋外での活動やアドベンチャー', '#8BC34A', 20),
       ('music', '音楽', '音楽の鑑賞や演奏に関する趣味', '#2196F3', 30),
       ('entertainment', 'エンタメ', '映画、アニメ、ゲームなどのエンターテイメント', '#9C27B0', 40),
       ('food_drink', '料理・食', '料理や飲食に関する趣味', '#FF9800', 50),
       ('travel', '旅行', '国内外の旅行に関する趣味', '#03A9F4', 60),
       ('art', '芸術', '美術、写真、手芸などの創作活動', '#E91E63', 70),
       ('lifestyle', 'ライフスタイル', 'ファッションや生活に関する趣味', '#607D8B', 80),
       ('learning', '学び', '自己啓発や知識習得に関する趣味', '#795548', 90),
       ('crafts', '趣味・クラフト', '創作やものづくりの趣味', '#FF5722', 100),
       ('pets_animals', 'ペット・動物', 'ペットや動物に関する趣味', '#CDDC39', 110),
       ('wellness', '健康・癒し', '健康増進やリラクゼーションに関する趣味', '#009688', 120),
       ('technology', 'テクノロジー', 'テクノロジーやガジェットに関する趣味', '#3F51B5', 130),
       ('vehicles', '乗り物', '車やバイク、乗り物に関する趣味', '#F44336', 140),
       ('other', 'その他', 'その他のカテゴリに分類されない趣味', '#9E9E9E', 999);

-- 5. ユーザーテーブルの作成
CREATE TABLE users
(
    id                UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    email             VARCHAR(255)             NOT NULL UNIQUE,
    password_hash     VARCHAR(255)             NOT NULL,
    name              VARCHAR(100)             NOT NULL,
    nickname          VARCHAR(50),
    birth_date        DATE                     NOT NULL,
    gender            VARCHAR(20)              NOT NULL,
    profile_image_url VARCHAR(255),
    about_me          TEXT,
    job_title         VARCHAR(100),
    company           VARCHAR(100),
    education_id      INTEGER                  REFERENCES educations (id) ON DELETE SET NULL,
    prefecture_id     INTEGER                  REFERENCES prefectures (id) ON DELETE SET NULL,
    looking_for       TEXT,
    last_active       TIMESTAMP WITH TIME ZONE,
    is_verified       BOOLEAN                           DEFAULT FALSE,
    is_premium        BOOLEAN                           DEFAULT FALSE,
    created_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at        TIMESTAMP WITH TIME ZONE
);

-- テーブルコメント
COMMENT ON TABLE users IS 'アプリのユーザー情報を格納するテーブル';

-- カラムコメント
COMMENT ON COLUMN users.id IS 'ユーザーの一意識別子（UUIDv4）';
COMMENT ON COLUMN users.email IS 'ユーザーのメールアドレス（認証と通知に使用）';
COMMENT ON COLUMN users.password_hash IS 'パスワードのハッシュ値（平文保存は不可）';
COMMENT ON COLUMN users.name IS 'ユーザーの本名';
COMMENT ON COLUMN users.nickname IS 'ユーザーのニックネーム（表示名）';
COMMENT ON COLUMN users.birth_date IS '生年月日';
COMMENT ON COLUMN users.gender IS '性別（マッチング条件として使用）';
COMMENT ON COLUMN users.profile_image_url IS 'プロフィール画像のURL';
COMMENT ON COLUMN users.about_me IS '自己紹介文';
COMMENT ON COLUMN users.job_title IS '職業・職種';
COMMENT ON COLUMN users.company IS '会社名・組織名';
COMMENT ON COLUMN users.education_id IS '学歴ID';
COMMENT ON COLUMN users.prefecture_id IS '都道府県ID';
COMMENT ON COLUMN users.looking_for IS '求めている関係性の説明';
COMMENT ON COLUMN users.last_active IS '最終アクティブ日時';
COMMENT ON COLUMN users.is_verified IS 'アカウント認証済みフラグ';
COMMENT ON COLUMN users.is_premium IS 'プレミアムアカウントフラグ';
COMMENT ON COLUMN users.created_at IS 'レコード作成日時';
COMMENT ON COLUMN users.updated_at IS 'レコード更新日時';
COMMENT ON COLUMN users.deleted_at IS '論理削除日時（NULLは有効なレコードを示す）';

-- インデックス
CREATE INDEX IF NOT EXISTS idx_users_education_id ON users (education_id);
CREATE INDEX IF NOT EXISTS idx_users_prefecture_id ON users (prefecture_id);
CREATE INDEX IF NOT EXISTS idx_users_gender_birth_date ON users (gender, birth_date);
CREATE INDEX IF NOT EXISTS idx_users_last_active ON users (last_active);
CREATE INDEX IF NOT EXISTS idx_users_premium_last_active ON users (is_premium, last_active);
CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users (deleted_at) WHERE deleted_at IS NULL;

-- 6. 趣味テーブルの作成
CREATE TABLE interests
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(100)             NOT NULL UNIQUE,
    display_name VARCHAR(100)             NOT NULL,
    category_id  INTEGER REFERENCES categories (id),
    icon_url     VARCHAR(255),
    sort_order   INTEGER                  NOT NULL DEFAULT 0,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at   TIMESTAMP WITH TIME ZONE
);

-- カテゴリのインデックス
CREATE INDEX IF NOT EXISTS idx_interests_category ON interests (category_id);
CREATE INDEX IF NOT EXISTS idx_interests_name ON interests (name);

-- サンプルの趣味データをカテゴリと紐付けて挿入するクエリ例
INSERT INTO interests (name, display_name, category_id, sort_order)
VALUES
-- スポーツカテゴリ
('running', 'ランニング', (SELECT id FROM categories WHERE name = 'sports'), 10),
('swimming', '水泳', (SELECT id FROM categories WHERE name = 'sports'), 11),
('cycling', 'サイクリング', (SELECT id FROM categories WHERE name = 'sports'), 12),
('yoga', 'ヨガ', (SELECT id FROM categories WHERE name = 'sports'), 15),
('gym', 'ジム・筋トレ', (SELECT id FROM categories WHERE name = 'sports'), 16),
('soccer', 'サッカー', (SELECT id FROM categories WHERE name = 'sports'), 17),
('baseball', '野球', (SELECT id FROM categories WHERE name = 'sports'), 18),
('basketball', 'バスケットボール', (SELECT id FROM categories WHERE name = 'sports'), 19),
('tennis', 'テニス', (SELECT id FROM categories WHERE name = 'sports'), 20),
('golf', 'ゴルフ', (SELECT id FROM categories WHERE name = 'sports'), 21),

-- アウトドアカテゴリ
('hiking', 'ハイキング', (SELECT id FROM categories WHERE name = 'outdoor'), 10),
('camping', 'キャンプ', (SELECT id FROM categories WHERE name = 'outdoor'), 11),
('fishing', '釣り', (SELECT id FROM categories WHERE name = 'outdoor'), 12),
('mountain_climbing', '登山', (SELECT id FROM categories WHERE name = 'outdoor'), 13),
('surfing', 'サーフィン', (SELECT id FROM categories WHERE name = 'outdoor'), 14),

-- 音楽カテゴリ
('music', '音楽鑑賞', (SELECT id FROM categories WHERE name = 'music'), 10),
('playing_music', '楽器演奏', (SELECT id FROM categories WHERE name = 'music'), 11),
('singing', '歌・カラオケ', (SELECT id FROM categories WHERE name = 'music'), 12),
('concerts', 'コンサート・ライブ', (SELECT id FROM categories WHERE name = 'music'), 13),
('classical_music', 'クラシック音楽', (SELECT id FROM categories WHERE name = 'music'), 14),

-- エンタメカテゴリ
('movies', '映画鑑賞', (SELECT id FROM categories WHERE name = 'entertainment'), 10),
('anime', 'アニメ', (SELECT id FROM categories WHERE name = 'entertainment'), 11),
('manga', '漫画', (SELECT id FROM categories WHERE name = 'entertainment'), 12),
('games', 'ゲーム', (SELECT id FROM categories WHERE name = 'entertainment'), 13),
('reading', '読書', (SELECT id FROM categories WHERE name = 'entertainment'), 14),

-- 料理・食カテゴリ
('cooking', '料理', (SELECT id FROM categories WHERE name = 'food_drink'), 10),
('baking', 'お菓子作り', (SELECT id FROM categories WHERE name = 'food_drink'), 11),
('wine', 'ワイン', (SELECT id FROM categories WHERE name = 'food_drink'), 12),
('beer', 'ビール・クラフトビール', (SELECT id FROM categories WHERE name = 'food_drink'), 13),
('coffee', 'コーヒー', (SELECT id FROM categories WHERE name = 'food_drink'), 14),
('dining_out', '外食・グルメ', (SELECT id FROM categories WHERE name = 'food_drink'), 15);

-- 残りの趣味データも同様に挿入（スペースの関係で省略）

-- 7. ユーザーと趣味の関連を管理する中間テーブル
CREATE TABLE user_interests
(
    user_id     UUID REFERENCES users (id) ON DELETE CASCADE,
    interest_id INTEGER REFERENCES interests (id) ON DELETE CASCADE,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, interest_id)
);

-- テーブルコメント
COMMENT ON TABLE user_interests IS 'ユーザーと趣味の関連を管理する中間テーブル';

-- カラムコメント
COMMENT ON COLUMN user_interests.user_id IS 'ユーザーID';
COMMENT ON COLUMN user_interests.interest_id IS '趣味ID';
COMMENT ON COLUMN user_interests.created_at IS '関連付け作成日時';

-- インデックスの作成
CREATE INDEX IF NOT EXISTS idx_user_interests_user_id ON user_interests (user_id);
CREATE INDEX IF NOT EXISTS idx_user_interests_interest_id ON user_interests (interest_id);

-- コミュニティテーブル
CREATE TABLE communities
(
    id                UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    name              VARCHAR(100)             NOT NULL,
    description       TEXT,
    profile_image_url VARCHAR(255),
    cover_image_url   VARCHAR(255),
    is_private        BOOLEAN                  NOT NULL DEFAULT FALSE,
    creator_id        UUID REFERENCES users (id),
    created_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at        TIMESTAMP WITH TIME ZONE
);

-- テーブルコメント
COMMENT ON TABLE communities IS 'コミュニティ情報を格納するテーブル';

-- カラムコメント
COMMENT ON COLUMN communities.id IS 'コミュニティの一意識別子（UUIDv4）';
COMMENT ON COLUMN communities.name IS 'コミュニティ名';
COMMENT ON COLUMN communities.description IS 'コミュニティの説明';
COMMENT ON COLUMN communities.profile_image_url IS 'コミュニティのプロフィール画像URL';
COMMENT ON COLUMN communities.cover_image_url IS 'コミュニティのカバー画像URL';
COMMENT ON COLUMN communities.is_private IS 'プライベートコミュニティフラグ（参加に承認が必要）';
COMMENT ON COLUMN communities.creator_id IS 'コミュニティ作成者のユーザーID';
COMMENT ON COLUMN communities.created_at IS 'レコード作成日時';
COMMENT ON COLUMN communities.updated_at IS 'レコード更新日時';
COMMENT ON COLUMN communities.deleted_at IS '論理削除日時（NULLは有効なレコードを示す）';

-- インデックス
CREATE INDEX IF NOT EXISTS idx_communities_creator_id ON communities (creator_id);
CREATE INDEX IF NOT EXISTS idx_communities_created_at ON communities (created_at);
CREATE INDEX IF NOT EXISTS idx_communities_is_private ON communities (is_private);
CREATE INDEX IF NOT EXISTS idx_communities_deleted_at ON communities (deleted_at) WHERE deleted_at IS NULL;

-- コミュニティメンバーシップテーブル
CREATE TABLE community_members
(
    community_id UUID REFERENCES communities (id) ON DELETE CASCADE,
    user_id      UUID REFERENCES users (id) ON DELETE CASCADE,
    role         VARCHAR(20)              NOT NULL DEFAULT 'member',
    is_approved  BOOLEAN                  NOT NULL DEFAULT FALSE,
    joined_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY (community_id, user_id)
);

-- テーブルコメント
COMMENT ON TABLE community_members IS 'コミュニティのメンバーシップ情報を管理するテーブル';

-- カラムコメント
COMMENT ON COLUMN community_members.community_id IS 'コミュニティID';
COMMENT ON COLUMN community_members.user_id IS 'ユーザーID';
COMMENT ON COLUMN community_members.role IS 'メンバーの役割（member, moderator, admin）';
COMMENT ON COLUMN community_members.is_approved IS '承認状態（プライベートコミュニティの場合）';
COMMENT ON COLUMN community_members.joined_at IS 'コミュニティ参加日時';
COMMENT ON COLUMN community_members.updated_at IS 'レコード更新日時';

-- インデックス
CREATE INDEX IF NOT EXISTS idx_community_members_user_id ON community_members (user_id);
CREATE INDEX IF NOT EXISTS idx_community_members_role ON community_members (role);
CREATE INDEX IF NOT EXISTS idx_community_members_is_approved ON community_members (is_approved);

-- いいね機能のテーブル
CREATE TABLE likes
(
    id         SERIAL PRIMARY KEY,
    user_id    UUID REFERENCES users (id) ON DELETE CASCADE,
    target_id  UUID                     NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT uk_likes_user_target UNIQUE (user_id, target_id)
);

-- テーブルコメント
COMMENT ON TABLE likes IS 'ユーザーのいいね情報を格納するテーブル';

-- カラムコメント
COMMENT ON COLUMN likes.id IS 'いいねID';
COMMENT ON COLUMN likes.user_id IS 'いいねをしたユーザーID';
COMMENT ON COLUMN likes.target_id IS 'いいねの対象となるエンティティのID';
COMMENT ON COLUMN likes.created_at IS 'いいね作成日時';
COMMENT ON COLUMN likes.updated_at IS 'レコード更新日時';
COMMENT ON COLUMN likes.deleted_at IS '論理削除日時（NULLは有効なレコードを示す）';

-- インデックス
CREATE INDEX IF NOT EXISTS idx_likes_user_id ON likes (user_id);
CREATE INDEX IF NOT EXISTS idx_likes_target_id ON likes (target_id);
CREATE INDEX IF NOT EXISTS idx_likes_created_at ON likes (created_at);
CREATE INDEX IF NOT EXISTS idx_likes_deleted_at ON likes (deleted_at) WHERE deleted_at IS NULL;

CREATE TABLE blocklists
(
    id              BIGSERIAL PRIMARY KEY,
    blocker_user_id UUID        NOT NULL,                           -- ブロックしたユーザーのID
    blocked_user_id UUID        NOT NULL,                           -- ブロックされたユーザーのID
    created_at      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, -- ブロック日時

    FOREIGN KEY (blocker_user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (blocked_user_id) REFERENCES users (id) ON DELETE CASCADE,
    UNIQUE (blocker_user_id, blocked_user_id)                       -- 同一の組み合わせでの重複ブロックを防ぐ
);

-- パフォーマンス向上のため、必要に応じてインデックスを作成します
CREATE INDEX idx_blocklists_blocker_user_id ON blocklists (blocker_user_id);
CREATE INDEX idx_blocklists_blocked_user_id ON blocklists (blocked_user_id);

-- つぶやきテーブル
CREATE TABLE tweets
(
    id         UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    user_id    UUID                     NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    content    TEXT                     NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- テーブルコメント
COMMENT ON TABLE tweets IS 'ユーザーのつぶやき情報を格納するテーブル';

-- カラムコメント
COMMENT ON COLUMN tweets.id IS 'つぶやきの一意識別子（UUIDv4）';
COMMENT ON COLUMN tweets.user_id IS 'つぶやきを投稿したユーザーのID';
COMMENT ON COLUMN tweets.content IS 'つぶやきの内容';
COMMENT ON COLUMN tweets.created_at IS 'つぶやき作成日時';
COMMENT ON COLUMN tweets.updated_at IS 'レコード更新日時';
COMMENT ON COLUMN tweets.deleted_at IS '論理削除日時（NULLは有効なレコードを示す）';

-- インデックス
CREATE INDEX idx_tweets_user_id ON tweets (user_id);
CREATE INDEX idx_tweets_created_at ON tweets (created_at);
CREATE INDEX idx_tweets_deleted_at ON tweets (deleted_at) WHERE deleted_at IS NULL;
