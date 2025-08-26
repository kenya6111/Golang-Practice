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



#### user_service.goがDB実装に依存してる
usecaseは本来DBを知らなくてもいいのに知っている。

```go
type UserService struct {
    repo *postgres.UserRepository  // ← ここが具体的型
}
```
これはアプリケーション層がインフラ層に依存している


「もし MySQL で保存したい」となったら UserService を書き換えないといけない。
DB処理をモック化してテストするのも難しい。

DDDの原則に照らすと

ドメイン層：ビジネスルールの中心（エンティティ・値オブジェクト・リポジトリのインターフェース）

アプリケーション層：ユースケースを組み立てる（インターフェースを使うだけ）

インフラ層：実装詳細（PostgresやMySQL）はここに押し込める


#### postgres ってディレクトリ名にしてあるのは「これはPostgresに依存する実装ですよっていう意味
infrastructure/
├── postgres/
│   └── user_repository.go   // Postgres専用の実装

確かに、dbがpostgresなので依存してる
```go
package postgres

import (
	"database/sql"
	"ddd_gpt_learning/domain/user"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
```

理解した。
```txt
あーーーなるほど、postgres以外のDBを使った時に（例えばmysqlとか、）新たに、infradtructure/mysqlってディレクトリきってそこにuserRepository.goって作るから、その場合そのuserRepositoryは
type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
ってなっててmysqlを持つから、

そうなった時にユースケースが

type UserService struct {
	repo *postgres.UserRepository
}

func NewUserService(repo *postgres.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

ってなってると、mysqlのリポジトリ使いたい時に、UserServiceも修正入れなきゃいけなくなるので、
userServiceをpostgresに依存させるんじゃなくて、
interfaceに依存させることで、
mysqlとかのリポジトリが増えた時に、修正が楽ってことか。
```