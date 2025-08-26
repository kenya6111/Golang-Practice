#### ドメイン
アプリが対象とする「問題領域」。
例: Twitterクローンなら「ユーザー」「ツイート」「フォロー」がドメインの一部。


#### エンティティ (Entity)
ドメインの中で「IDを持つオブジェクト」。
値が変わっても「同じもの」として扱われる。
例: ユーザーは名前を変更しても同じユーザーIDを持っているから「同一人物」。



 INSERT INTO users (id, username, email)
 VALUES
     (gen_random_uuid(), 'kenya', 'kenya@example.com'),
     (gen_random_uuid(), 'taro', 'taro@example.com');

#### エンティティはIDを持って同一性で区別される概念ごとに作成する
SQLのテーブル = エンティティ ではない（けど近いことも多い）。

例：Twitterっぽいアプリなら
User (ユーザー) → IDを持つ
Tweet (ツイート) → IDを持つ

Follow (フォロー関係) → これもIDを持つ or 複合キー

つまり type User struct {...} は エンティティ定義 です。
「ユーザー」という概念をコードで表現している。


#### ドメインロジックとは
👉 ビジネスルールをコード化したもの。
データの入れ物じゃなくて、そのデータがどう振る舞うかを表す。

例:
func (u *User) ChangeUsername(newName string) {
    if newName == "" {
        panic("ユーザー名は空にできません") // ルール違反
    }
    u.Username = newName
}

「ユーザー名は空文字にできない」 → これがドメインロジック。

もしDBに直接 UPDATE users SET username='' って書いたらルール破れる。
だから「エンティティが自分で自分を守る」イメージ。

#### リポジトリ (Repository) はどの単位で作る？

👉 エンティティごとに作るのが基本。
なぜならエンティティ単位で「保存」「取得」するから。

例：
UserRepository → ユーザーを保存・取得する
TweetRepository → ツイートを保存・取得する

リポジトリは「DBテーブルに直接対応する」というより、
「エンティティを永続化するための抽象インターフェース」って感じ。