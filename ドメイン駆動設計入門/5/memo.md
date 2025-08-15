#### リポジトリはデータを永続化しリポジトリはデータを永続化市再構築するといった市再構築するといった処理を抽象的に扱うための処理を抽象的に扱うためのオブジェクト。
「永続化」＝ DBやファイルに保存すること
「再構築」＝ 保存されているデータからドメインオブジェクトを作り直すこと
「抽象化」＝ 呼び出し側はDBの種類やSQLの書き方を知らなくても使えるようにすること



1. リポジトリがない場合（直接DBアクセス）
```go
func ChangeUserName(db *sql.DB, id string, newName string) error {
    _, err := db.Exec(`UPDATE users SET name = ? WHERE id = ?`, newName, id)
    return err
}
```
特徴
DBの存在を関数が直接知っている

SQLがアプリケーションロジックに混ざる

DBの種類や構造が変わると、アプリケーションの広い範囲を修正する必要がある

テスト時にDBが必須（モック化が難しい）


2. リポジトリがある場合（永続化を抽象化）
```go
// ドメイン層
type UserRepository interface {
    FindByID(id UserID) (*User, error)
    Save(user *User) error
}

// アプリケーション層
func (svc *UserService) ChangeUserName(id UserID, newName UserName) error {
    user, err := svc.repo.FindByID(id)
    if err != nil {
        return err
    }
    user.ChangeName(newName) // ドメインルール適用
    return svc.repo.Save(user)
}
```
DB操作の詳細（SQLなど）はインフラ層で隠す
アプリケーション層・ドメイン層は「保管庫から取ってくる」「保管庫に戻す」という感覚で使える
実装を差し替え可能（MySQL → PostgreSQL → メモリ など）

サービス側はDBを操作する「メソッド」を呼ぶだけでよくなるってこと。


#### リポジトリをインターフェースにする理由
ドメイン層を具体的な永続化技術から切り離すため」

#### 永続化
. 生活の例でいうと…
メモ帳に書いた買い物リスト → アプリ終了したら消える（揮発）

ノートやクラウドに保存した買い物リスト → 閉じても残る（永続）

つまり「永続化」は、今ある情報を長期間残すために保存すること。
電源が落ちても、アプリが終了してもデータが残る状態にする。

