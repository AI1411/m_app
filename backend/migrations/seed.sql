TRUNCATE TABLE users CASCADE;
TRUNCATE TABLE user_interests CASCADE;
-- ダミーユーザーデータの挿入
INSERT INTO users (email,
                   password_hash,
                   name,
                   nickname,
                   birth_date,
                   gender,
                   profile_image_url,
                   about_me,
                   job_title,
                   company,
                   education_id,
                   prefecture_id,
                   looking_for,
                   is_verified,
                   is_premium)
VALUES
-- ユーザー1
('tanaka.taro@example.com',
 '$2a$10$8KxS6hT3MRlxU/A.CIm/VOzVVHdO.Dvj5k1/KJa.xHgVvz3zKzUBC', -- 'password123'のハッシュ値
 '田中 太郎',
 'タロー',
 '1990-05-15',
 '男性',
 'https://example.com/profiles/tanaka.jpg',
 '東京在住のエンジニアです。趣味はプログラミングと旅行です。',
 'ソフトウェアエンジニア',
 'テック株式会社',
 5, -- 大卒
 (SELECT id FROM prefectures WHERE name = '東京都'), -- 東京都
 '同じ趣味を持つ友達を探しています。',
 TRUE,
 TRUE),
-- ユーザー2
('yamada.hanako@example.com',
 '$2a$10$8KxS6hT3MRlxU/A.CIm/VOzVVHdO.Dvj5k1/KJa.xHgVvz3zKzUBC', -- 'password123'のハッシュ値
 '山田 花子',
 'ハナコ',
 '1992-08-23',
 '女性',
 'https://example.com/profiles/yamada.jpg',
 '大阪在住のデザイナーです。アート、音楽、ファッションに興味があります。',
 'UIデザイナー',
 'クリエイティブ株式会社',
 5, -- 大卒
 (SELECT id FROM prefectures WHERE name = '大阪府'), -- 大阪府
 'クリエイティブな友人や仕事のつながりを探しています。',
 TRUE,
 FALSE),
-- ユーザー3
('suzuki.kenji@example.com',
 '$2a$10$8KxS6hT3MRlxU/A.CIm/VOzVVHdO.Dvj5k1/KJa.xHgVvz3zKzUBC', -- 'password123'のハッシュ値
 '鈴木 健二',
 'ケンジ',
 '1985-12-10',
 '男性',
 'https://example.com/profiles/suzuki.jpg',
 '名古屋在住の営業マンです。スポーツ観戦と料理が趣味です。',
 '営業マネージャー',
 '商事株式会社',
 5, -- 大卒
 (SELECT id FROM prefectures WHERE name = '愛知県'), -- 愛知県
 'スポーツや食事を一緒に楽しめる友達を探しています。',
 FALSE,
 FALSE),
-- ユーザー4
('sato.yuki@example.com',
 '$2a$10$8KxS6hT3MRlxU/A.CIm/VOzVVHdO.Dvj5k1/KJa.xHgVvz3zKzUBC', -- 'password123'のハッシュ値
 '佐藤 ゆき',
 'ユキ',
 '1995-03-20',
 '女性',
 'https://example.com/profiles/sato.jpg',
 '福岡在住の学生です。読書とカフェ巡りが好きです。',
 '大学院生',
 '福岡大学',
 4, -- 短大卒
 (SELECT id FROM prefectures WHERE name = '福岡県'), -- 福岡県
 '同じ趣味を持つ友達や勉強仲間を探しています。',
 TRUE,
 FALSE),
-- ユーザー5
('watanabe.hiroshi@example.com',
 '$2a$10$8KxS6hT3MRlxU/A.CIm/VOzVVHdO.Dvj5k1/KJa.xHgVvz3zKzUBC', -- 'password123'のハッシュ値
 '渡辺 浩',
 'ヒロシ',
 '1980-07-05',
 '男性',
 'https://example.com/profiles/watanabe.jpg',
 '札幌在住のシェフです。料理と登山が趣味です。',
 'シェフ',
 'グルメレストラン',
 3, -- 専門卒
 (SELECT id FROM prefectures WHERE name = '北海道'), -- 北海道
 'アウトドア活動を一緒に楽しめる友達を探しています。',
 FALSE,
 TRUE);

-- ユーザーの趣味データを関連付ける
-- 実際のinterestsテーブルに存在する趣味名を確認して使用
INSERT INTO user_interests (user_id, interest_id)
VALUES ((SELECT id FROM users WHERE email = 'tanaka.taro@example.com'), 1),
       ((SELECT id FROM users WHERE email = 'tanaka.taro@example.com'), 2),
       ((SELECT id FROM users WHERE email = 'tanaka.taro@example.com'), 3),

       ((SELECT id FROM users WHERE email = 'yamada.hanako@example.com'), 4),
       ((SELECT id FROM users WHERE email = 'yamada.hanako@example.com'), 5),
       ((SELECT id FROM users WHERE email = 'yamada.hanako@example.com'), 6),

       ((SELECT id FROM users WHERE email = 'suzuki.kenji@example.com'), 7),
       ((SELECT id FROM users WHERE email = 'suzuki.kenji@example.com'), 8),
       ((SELECT id FROM users WHERE email = 'suzuki.kenji@example.com'), 9),

       ((SELECT id FROM users WHERE email = 'sato.yuki@example.com'), 10),
       ((SELECT id FROM users WHERE email = 'sato.yuki@example.com'), 11),
       ((SELECT id FROM users WHERE email = 'sato.yuki@example.com'), 12),
       ((SELECT id FROM users WHERE email = 'watanabe.hiroshi@example.com'), 13),
       ((SELECT id FROM users WHERE email = 'watanabe.hiroshi@example.com'), 14),
       ((SELECT id FROM users WHERE email = 'watanabe.hiroshi@example.com'), 15);