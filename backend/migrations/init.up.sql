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
    location          VARCHAR(100),
    job_title         VARCHAR(100),
    company           VARCHAR(100),
    education_id      INT                      REFERENCES educations (id) ON DELETE SET NULL,
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
COMMENT ON COLUMN users.location IS '居住地（都市名など）';
COMMENT ON COLUMN users.job_title IS '職業・職種';
COMMENT ON COLUMN users.company IS '会社名・組織名';
COMMENT ON COLUMN users.education IS '学歴';
COMMENT ON COLUMN users.last_active IS '最終アクティブ日時';
COMMENT ON COLUMN users.is_verified IS 'アカウント認証済みフラグ';
COMMENT ON COLUMN users.is_premium IS 'プレミアムアカウントフラグ';
COMMENT ON COLUMN users.created_at IS 'レコード作成日時';
COMMENT ON COLUMN users.updated_at IS 'レコード更新日時';
COMMENT ON COLUMN users.deleted_at IS '論理削除日時（NULLは有効なレコードを示す）';

-- インデックス（検索パフォーマンス向上用）
-- 性別と年齢範囲での検索用複合インデックス
CREATE INDEX idx_users_gender_birth_date ON users (gender, birth_date);

-- 地域と年齢範囲での検索用複合インデックス
CREATE INDEX idx_users_location_birth_date ON users (location, birth_date);

-- 最終アクティブ日時での検索・ソート用インデックス
CREATE INDEX idx_users_last_active ON users (last_active);

-- プレミアムユーザーと最終アクティブ日時のフィルタリング用複合インデックス
CREATE INDEX idx_users_premium_last_active ON users (is_premium, last_active);

-- 論理削除フィルタリング用インデックス（削除されていないユーザーの検索が多いため）
CREATE INDEX idx_users_deleted_at ON users (deleted_at) WHERE deleted_at IS NULL;

-- GINインデックス（interests配列の検索最適化用）
CREATE INDEX idx_users_interests ON users USING GIN (interests);

-- 学歴マスターテーブルの作成
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

-- education_idカラムの追加
ALTER TABLE users
    ADD COLUMN education_id INTEGER REFERENCES educations (id);

-- インデックスの作成
CREATE INDEX idx_users_education_id ON users (education_id);

-- 趣味の初期データ（カテゴリ別にグループ化）
INSERT INTO interests (name, display_name, category_id, sort_order)
VALUES
-- スポーツ・アウトドア
('running', 'ランニング', 'スポーツ', 10),
('swimming', '水泳', 'スポーツ', 11),
('cycling', 'サイクリング', 'スポーツ', 12),
('hiking', 'ハイキング', 'アウトドア', 13),
('camping', 'キャンプ', 'アウトドア', 14),
('yoga', 'ヨガ', 'スポーツ', 15),
('gym', 'ジム・筋トレ', 'スポーツ', 16),
('soccer', 'サッカー', 'スポーツ', 17),
('baseball', '野球', 'スポーツ', 18),
('basketball', 'バスケットボール', 'スポーツ', 19),
('tennis', 'テニス', 'スポーツ', 20),
('golf', 'ゴルフ', 'スポーツ', 21),

-- 音楽・エンタメ
('music', '音楽鑑賞', '音楽', 30),
('playing_music', '楽器演奏', '音楽', 31),
('singing', '歌・カラオケ', '音楽', 32),
('movies', '映画鑑賞', 'エンタメ', 33),
('anime', 'アニメ', 'エンタメ', 34),
('manga', '漫画', 'エンタメ', 35),
('games', 'ゲーム', 'エンタメ', 36),
('reading', '読書', 'エンタメ', 37),

-- 料理・食
('cooking', '料理', '料理・食', 50),
('baking', 'お菓子作り', '料理・食', 51),
('wine', 'ワイン', '料理・食', 52),
('beer', 'ビール・クラフトビール', '料理・食', 53),
('coffee', 'コーヒー', '料理・食', 54),
('dining_out', '外食・グルメ', '料理・食', 55),

-- 旅行・文化
('travel', '国内旅行', '旅行', 70),
('world_travel', '海外旅行', '旅行', 71),
('photography', '写真撮影', '芸術', 72),
('art', '美術・アート', '芸術', 73),
('theater', '演劇・舞台', '芸術', 74),
('fashion', 'ファッション', 'ライフスタイル', 75),

-- 学び・自己啓発
('language', '語学・外国語', '学び', 90),
('programming', 'プログラミング', '学び', 91),
('investing', '投資・資産運用', '学び', 92),
('business', 'ビジネス', '学び', 93),
('science', '科学', '学び', 94),
('history', '歴史', '学び', 95),

-- 趣味・クラフト
('gardening', 'ガーデニング', '趣味', 110),
('diy', 'DIY・日曜大工', '趣味', 111),
('drawing', '絵を描く', '趣味', 112),
('handicraft', '手芸・クラフト', '趣味', 113),
('collecting', 'コレクション', '趣味', 114),

-- ペット・動物
('dogs', '犬', 'ペット・動物', 130),
('cats', '猫', 'ペット・動物', 131),
('aquarium', '熱帯魚・アクアリウム', 'ペット・動物', 132),
('wildlife', '野生動物観察', 'ペット・動物', 133),

-- その他
('volunteering', 'ボランティア', 'その他', 150),
('meditation', '瞑想・マインドフルネス', 'その他', 151),
('dancing', 'ダンス', 'その他', 152),
('cars', '車・ドライブ', 'その他', 153),
('motorcycles', 'バイク', 'その他', 154);

-- カテゴリのインデックス
CREATE INDEX idx_interests_category ON interests (category_id);
CREATE INDEX idx_interests_name ON interests (name);

-- カテゴリマスターテーブルの作成
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

-- interestsテーブルの修正
-- 変更前: category VARCHAR(50) -> 変更後: category_id INTEGER REFERENCES categories(id)

-- 既存のinterestsテーブルがある場合はカラム変更
ALTER TABLE interests
    DROP COLUMN IF EXISTS category;
ALTER TABLE interests
    ADD COLUMN category_id INTEGER REFERENCES categories (id);
COMMENT ON COLUMN interests.category_id IS '趣味のカテゴリID（categoriesテーブルへの参照）';

-- インデックスの作成
CREATE INDEX idx_interests_category_id ON interests (category_id);

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

-- サンプルの趣味データをカテゴリと紐付けて挿入するクエリ例
-- 既存のinterestsテーブルにデータを入れる場合は、先にcategoriesテーブルにデータを入れてから行う
INSERT INTO interests (name, display_name, category_id, sort_order)
VALUES
-- スポーツカテゴリ (id=1)
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

-- アウトドアカテゴリ (id=2)
('hiking', 'ハイキング', (SELECT id FROM categories WHERE name = 'outdoor'), 10),
('camping', 'キャンプ', (SELECT id FROM categories WHERE name = 'outdoor'), 11),
('fishing', '釣り', (SELECT id FROM categories WHERE name = 'outdoor'), 12),
('mountain_climbing', '登山', (SELECT id FROM categories WHERE name = 'outdoor'), 13),
('surfing', 'サーフィン', (SELECT id FROM categories WHERE name = 'outdoor'), 14),

-- 音楽カテゴリ (id=3)
('music', '音楽鑑賞', (SELECT id FROM categories WHERE name = 'music'), 10),
('playing_music', '楽器演奏', (SELECT id FROM categories WHERE name = 'music'), 11),
('singing', '歌・カラオケ', (SELECT id FROM categories WHERE name = 'music'), 12),
('concerts', 'コンサート・ライブ', (SELECT id FROM categories WHERE name = 'music'), 13),
('classical_music', 'クラシック音楽', (SELECT id FROM categories WHERE name = 'music'), 14),

-- エンタメカテゴリ (id=4)
('movies', '映画鑑賞', (SELECT id FROM categories WHERE name = 'entertainment'), 10),
('anime', 'アニメ', (SELECT id FROM categories WHERE name = 'entertainment'), 11),
('manga', '漫画', (SELECT id FROM categories WHERE name = 'entertainment'), 12),
('games', 'ゲーム', (SELECT id FROM categories WHERE name = 'entertainment'), 13),
('reading', '読書', (SELECT id FROM categories WHERE name = 'entertainment'), 14),

-- 料理・食カテゴリ (id=5)
('cooking', '料理', (SELECT id FROM categories WHERE name = 'food_drink'), 10),
('baking', 'お菓子作り', (SELECT id FROM categories WHERE name = 'food_drink'), 11),
('wine', 'ワイン', (SELECT id FROM categories WHERE name = 'food_drink'), 12),
('beer', 'ビール・クラフトビール', (SELECT id FROM categories WHERE name = 'food_drink'), 13),
('coffee', 'コーヒー', (SELECT id FROM categories WHERE name = 'food_drink'), 14),
('dining_out', '外食・グルメ', (SELECT id FROM categories WHERE name = 'food_drink'), 15),

-- 旅行カテゴリ (id=6)
('travel', '国内旅行', (SELECT id FROM categories WHERE name = 'travel'), 10),
('world_travel', '海外旅行', (SELECT id FROM categories WHERE name = 'travel'), 11),
('backpacking', 'バックパッキング', (SELECT id FROM categories WHERE name = 'travel'), 12),
('hot_springs', '温泉巡り', (SELECT id FROM categories WHERE name = 'travel'), 13),

-- 芸術カテゴリ (id=7)
('photography', '写真撮影', (SELECT id FROM categories WHERE name = 'art'), 10),
('art', '美術・アート', (SELECT id FROM categories WHERE name = 'art'), 11),
('theater', '演劇・舞台', (SELECT id FROM categories WHERE name = 'art'), 12),
('drawing', '絵を描く', (SELECT id FROM categories WHERE name = 'art'), 13),
('calligraphy', '書道', (SELECT id FROM categories WHERE name = 'art'), 14),

-- ライフスタイルカテゴリ (id=8)
('fashion', 'ファッション', (SELECT id FROM categories WHERE name = 'lifestyle'), 10),
('interior', 'インテリア', (SELECT id FROM categories WHERE name = 'lifestyle'), 11),
('minimalism', 'ミニマリスト', (SELECT id FROM categories WHERE name = 'lifestyle'), 12),

-- 学びカテゴリ (id=9)
('language', '語学・外国語', (SELECT id FROM categories WHERE name = 'learning'), 10),
('programming', 'プログラミング', (SELECT id FROM categories WHERE name = 'learning'), 11),
('investing', '投資・資産運用', (SELECT id FROM categories WHERE name = 'learning'), 12),
('business', 'ビジネス', (SELECT id FROM categories WHERE name = 'learning'), 13),
('science', '科学', (SELECT id FROM categories WHERE name = 'learning'), 14),
('history', '歴史', (SELECT id FROM categories WHERE name = 'learning'), 15),

-- 趣味・クラフトカテゴリ (id=10)
('gardening', 'ガーデニング', (SELECT id FROM categories WHERE name = 'crafts'), 10),
('diy', 'DIY・日曜大工', (SELECT id FROM categories WHERE name = 'crafts'), 11),
('handicraft', '手芸・クラフト', (SELECT id FROM categories WHERE name = 'crafts'), 12),
('collecting', 'コレクション', (SELECT id FROM categories WHERE name = 'crafts'), 13),
('pottery', '陶芸', (SELECT id FROM categories WHERE name = 'crafts'), 14),

-- ペット・動物カテゴリ (id=11)
('dogs', '犬', (SELECT id FROM categories WHERE name = 'pets_animals'), 10),
('cats', '猫', (SELECT id FROM categories WHERE name = 'pets_animals'), 11),
('aquarium', '熱帯魚・アクアリウム', (SELECT id FROM categories WHERE name = 'pets_animals'), 12),
('wildlife', '野生動物観察', (SELECT id FROM categories WHERE name = 'pets_animals'), 13),

-- 健康・癒しカテゴリ (id=12)
('meditation', '瞑想・マインドフルネス', (SELECT id FROM categories WHERE name = 'wellness'), 10),
('aromatherapy', 'アロマテラピー', (SELECT id FROM categories WHERE name = 'wellness'), 11),
('hot_yoga', 'ホットヨガ', (SELECT id FROM categories WHERE name = 'wellness'), 12),
('spa', 'スパ・マッサージ', (SELECT id FROM categories WHERE name = 'wellness'), 13),

-- テクノロジーカテゴリ (id=13)
('gadgets', 'ガジェット', (SELECT id FROM categories WHERE name = 'technology'), 10),
('vr', 'VR・AR', (SELECT id FROM categories WHERE name = 'technology'), 11),
('drones', 'ドローン', (SELECT id FROM categories WHERE name = 'technology'), 12),
('ai', 'AI・人工知能', (SELECT id FROM categories WHERE name = 'technology'), 13),

-- 乗り物カテゴリ (id=14)
('cars', '車・ドライブ', (SELECT id FROM categories WHERE name = 'vehicles'), 10),
('motorcycles', 'バイク', (SELECT id FROM categories WHERE name = 'vehicles'), 11),
('bicycles', '自転車', (SELECT id FROM categories WHERE name = 'vehicles'), 12),

-- その他カテゴリ (id=15)
('volunteering', 'ボランティア', (SELECT id FROM categories WHERE name = 'other'), 10),
('dancing', 'ダンス', (SELECT id FROM categories WHERE name = 'other'), 11),
('astrology', '占星術・星座', (SELECT id FROM categories WHERE name = 'other'), 12),
('board_games', 'ボードゲーム', (SELECT id FROM categories WHERE name = 'other'), 13);

-- ユーザーと趣味の関連を管理する中間テーブル
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
CREATE INDEX idx_user_interests_user_id ON user_interests (user_id);
CREATE INDEX idx_user_interests_interest_id ON user_interests (interest_id);

-- コミュニティテーブル
CREATE TABLE communities
(
    id                   UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    name                 VARCHAR(100)             NOT NULL,
    slug                 VARCHAR(100)             NOT NULL UNIQUE,
    description          TEXT,
    profile_image_url    VARCHAR(255),
    cover_image_url      VARCHAR(255),
    category_id          INTEGER REFERENCES categories(id),
    is_private           BOOLEAN                  NOT NULL DEFAULT FALSE,
    is_verified          BOOLEAN                  NOT NULL DEFAULT FALSE,
    member_count         INTEGER                  NOT NULL DEFAULT 0,
    creator_id           UUID REFERENCES users(id),
    rules                TEXT,
    location             VARCHAR(100),
    website_url          VARCHAR(255),
    created_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at           TIMESTAMP WITH TIME ZONE
);

-- テーブルコメント
COMMENT ON TABLE communities IS 'コミュニティ情報を格納するテーブル';

-- カラムコメント
COMMENT ON COLUMN communities.id IS 'コミュニティの一意識別子（UUIDv4）';
COMMENT ON COLUMN communities.name IS 'コミュニティ名';
COMMENT ON COLUMN communities.slug IS 'URLに使用されるコミュニティの一意識別子（英数字とハイフンのみ）';
COMMENT ON COLUMN communities.description IS 'コミュニティの説明';
COMMENT ON COLUMN communities.profile_image_url IS 'コミュニティのプロフィール画像URL';
COMMENT ON COLUMN communities.cover_image_url IS 'コミュニティのカバー画像URL';
COMMENT ON COLUMN communities.category_id IS 'コミュニティのカテゴリID';
COMMENT ON COLUMN communities.is_private IS 'プライベートコミュニティフラグ（参加に承認が必要）';
COMMENT ON COLUMN communities.is_verified IS '公式コミュニティフラグ';
COMMENT ON COLUMN communities.member_count IS 'メンバー数（キャッシュ）';
COMMENT ON COLUMN communities.creator_id IS 'コミュニティ作成者のユーザーID';
COMMENT ON COLUMN communities.rules IS 'コミュニティルール';
COMMENT ON COLUMN communities.location IS 'コミュニティの地域（オプション）';
COMMENT ON COLUMN communities.website_url IS 'コミュニティの外部ウェブサイト（オプション）';
COMMENT ON COLUMN communities.created_at IS 'レコード作成日時';
COMMENT ON COLUMN communities.updated_at IS 'レコード更新日時';
COMMENT ON COLUMN communities.deleted_at IS '論理削除日時（NULLは有効なレコードを示す）';

-- インデックス
CREATE INDEX idx_communities_category_id ON communities(category_id);
CREATE INDEX idx_communities_creator_id ON communities(creator_id);
CREATE INDEX idx_communities_created_at ON communities(created_at);
CREATE INDEX idx_communities_is_private ON communities(is_private);
CREATE INDEX idx_communities_deleted_at ON communities(deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX idx_communities_slug ON communities(slug);

-- コミュニティと趣味の関連テーブル
CREATE TABLE community_interests
(
    community_id UUID REFERENCES communities(id) ON DELETE CASCADE,
    interest_id  INTEGER REFERENCES interests(id) ON DELETE CASCADE,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY (community_id, interest_id)
);

-- テーブルコメント
COMMENT ON TABLE community_interests IS 'コミュニティと趣味の関連を管理するテーブル';

-- コミュニティメンバーシップテーブル
CREATE TABLE community_members
(
    community_id UUID REFERENCES communities(id) ON DELETE CASCADE,
    user_id      UUID REFERENCES users(id) ON DELETE CASCADE,
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
CREATE INDEX idx_community_members_user_id ON community_members(user_id);
CREATE INDEX idx_community_members_role ON community_members(role);
CREATE INDEX idx_community_members_is_approved ON community_members(is_approved);

-- コミュニティ投稿テーブル
CREATE TABLE community_posts
(
    id           UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    community_id UUID REFERENCES communities(id) ON DELETE CASCADE,
    user_id      UUID REFERENCES users(id) ON DELETE SET NULL,
    title        VARCHAR(200),
    content      TEXT                     NOT NULL,
    media_urls   TEXT[],
    like_count   INTEGER                  NOT NULL DEFAULT 0,
    comment_count INTEGER                 NOT NULL DEFAULT 0,
    is_pinned    BOOLEAN                  NOT NULL DEFAULT FALSE,
    is_approved  BOOLEAN                  NOT NULL DEFAULT TRUE,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at   TIMESTAMP WITH TIME ZONE
);

-- テーブルコメント
COMMENT ON TABLE community_posts IS 'コミュニティ内の投稿を管理するテーブル';

-- カラムコメント
COMMENT ON COLUMN community_posts.id IS '投稿ID';
COMMENT ON COLUMN community_posts.community_id IS 'コミュニティID';
COMMENT ON COLUMN community_posts.user_id IS '投稿者のユーザーID';
COMMENT ON COLUMN community_posts.title IS '投稿タイトル（オプション）';
COMMENT ON COLUMN community_posts.content IS '投稿内容';
COMMENT ON COLUMN community_posts.media_urls IS '添付メディアのURL配列（画像や動画）';
COMMENT ON COLUMN community_posts.like_count IS 'いいね数（キャッシュ）';
COMMENT ON COLUMN community_posts.comment_count IS 'コメント数（キャッシュ）';
COMMENT ON COLUMN community_posts.is_pinned IS 'ピン留め投稿フラグ';
COMMENT ON COLUMN community_posts.is_approved IS '承認済みフラグ（モデレーション用）';
COMMENT ON COLUMN community_posts.created_at IS '投稿作成日時';
COMMENT ON COLUMN community_posts.updated_at IS '投稿更新日時';
COMMENT ON COLUMN community_posts.deleted_at IS '論理削除日時';

-- インデックス
CREATE INDEX idx_community_posts_community_id ON community_posts(community_id);
CREATE INDEX idx_community_posts_user_id ON community_posts(user_id);
CREATE INDEX idx_community_posts_created_at ON community_posts(created_at);
CREATE INDEX idx_community_posts_is_pinned ON community_posts(is_pinned);
CREATE INDEX idx_community_posts_deleted_at ON community_posts(deleted_at) WHERE deleted_at IS NULL;

-- コミュニティ投稿のコメントテーブル
CREATE TABLE community_post_comments
(
    id           UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    post_id      UUID REFERENCES community_posts(id) ON DELETE CASCADE,
    user_id      UUID REFERENCES users(id) ON DELETE SET NULL,
    parent_id    UUID REFERENCES community_post_comments(id) ON DELETE CASCADE,
    content      TEXT                     NOT NULL,
    media_urls   TEXT[],
    like_count   INTEGER                  NOT NULL DEFAULT 0,
    is_approved  BOOLEAN                  NOT NULL DEFAULT TRUE,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at   TIMESTAMP WITH TIME ZONE
);

-- テーブルコメント
COMMENT ON TABLE community_post_comments IS 'コミュニティ投稿へのコメントを管理するテーブル';

-- カラムコメント
COMMENT ON COLUMN community_post_comments.id IS 'コメントID';
COMMENT ON COLUMN community_post_comments.post_id IS '投稿ID';
COMMENT ON COLUMN community_post_comments.user_id IS 'コメント投稿者のユーザーID';
COMMENT ON COLUMN community_post_comments.parent_id IS '親コメントID（返信の場合）';
COMMENT ON COLUMN community_post_comments.content IS 'コメント内容';
COMMENT ON COLUMN community_post_comments.media_urls IS '添付メディアのURL配列';
COMMENT ON COLUMN community_post_comments.like_count IS 'いいね数（キャッシュ）';
COMMENT ON COLUMN community_post_comments.is_approved IS '承認済みフラグ（モデレーション用）';
COMMENT ON COLUMN community_post_comments.created_at IS 'コメント作成日時';
COMMENT ON COLUMN community_post_comments.updated_at IS 'コメント更新日時';
COMMENT ON COLUMN community_post_comments.deleted_at IS '論理削除日時';

-- インデックス
CREATE INDEX idx_community_post_comments_post_id ON community_post_comments(post_id);
CREATE INDEX idx_community_post_comments_user_id ON community_post_comments(user_id);
CREATE INDEX idx_community_post_comments_parent_id ON community_post_comments(parent_id);
CREATE INDEX idx_community_post_comments_created_at ON community_post_comments(created_at);
CREATE INDEX idx_community_post_comments_deleted_at ON community_post_comments(deleted_at) WHERE deleted_at IS NULL;

-- いいねテーブル
CREATE TABLE community_likes
(
    id           UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    user_id      UUID REFERENCES users(id) ON DELETE CASCADE,
    post_id      UUID REFERENCES community_posts(id) ON DELETE CASCADE,
    comment_id   UUID REFERENCES community_post_comments(id) ON DELETE CASCADE,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CHECK ((post_id IS NULL) != (comment_id IS NULL)) -- post_idまたはcomment_idのどちらか一方のみが非NULL
);

-- テーブルコメント
COMMENT ON TABLE community_likes IS 'コミュニティの投稿やコメントへのいいねを管理するテーブル';

-- ユニーク制約（1ユーザーが1投稿/コメントに1いいねのみ可能）
CREATE UNIQUE INDEX idx_community_likes_post_user ON community_likes(post_id, user_id) WHERE post_id IS NOT NULL;
CREATE UNIQUE INDEX idx_community_likes_comment_user ON community_likes(comment_id, user_id) WHERE comment_id IS NOT NULL;

-- イベントテーブル
CREATE TABLE community_events
(
    id            UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    community_id  UUID REFERENCES communities(id) ON DELETE CASCADE,
    creator_id    UUID REFERENCES users(id) ON DELETE SET NULL,
    title         VARCHAR(200)             NOT NULL,
    description   TEXT,
    location      VARCHAR(255),
    is_online     BOOLEAN                  NOT NULL DEFAULT FALSE,
    start_time    TIMESTAMP WITH TIME ZONE NOT NULL,
    end_time      TIMESTAMP WITH TIME ZONE,
    max_attendees INTEGER,
    image_url     VARCHAR(255),
    created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMP WITH TIME ZONE
);

-- テーブルコメント
COMMENT ON TABLE community_events IS 'コミュニティのイベント情報を管理するテーブル';

-- インデックス
CREATE INDEX idx_community_events_community_id ON community_events(community_id);
CREATE INDEX idx_community_events_creator_id ON community_events(creator_id);