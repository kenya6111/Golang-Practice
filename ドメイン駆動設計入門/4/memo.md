#### ドメインサービスとは
エンティティや値オブジェクトだけでは表現しづらい「ドメイン固有の振る舞い」を担うオブジェクト

エンティティや値オブジェクトは、それぞれ自分に密接に関わるルールや振る舞いを持つのが理想です。
でも、「複数のエンティティや値オブジェクトにまたがる」処理や、どれにも属さないルールがあります。
そういう「宙に浮いたビジネスロジック」を置く場所がドメインサービスです。

名前は 動詞的 になりやすい（UserRegistrationService、PaymentServiceなど）

ドメインサービスの必要な場面
例1: ユーザーを新しいルームに招待する
UserにもRoomにも直接は属さない
招待ロジックは「ユーザーの権限チェック」「ルームの定員チェック」「通知送信」など複数のエンティティを横断

```go
package domain

type InvitationService struct {
    userRepo UserRepository
    roomRepo RoomRepository
}

func NewInvitationService(uRepo UserRepository, rRepo RoomRepository) InvitationService {
    return InvitationService{userRepo: uRepo, roomRepo: rRepo}
}

func (s InvitationService) InviteUserToRoom(userID, roomID string) error {
    user, err := s.userRepo.FindByID(userID)
    if err != nil {
        return err
    }
    room, err := s.roomRepo.FindByID(roomID)
    if err != nil {
        return err
    }

    if !user.CanJoin(room) {
        return errors.New("user is not allowed to join this room")
    }
    return room.AddMember(user)
}

```

|                | エンティティ（Entity）      | 値オブジェクト（VO）    | ドメインサービス                      |
|----------------|-----------------------------|-------------------------|----------------------------------------|
| **同一性**     | IDで識別                     | 値で識別                 | 持たない                               |
| **状態**       | 持つ（可変）                  | 不変                     | 原則持たない                           |
| **主な役割**   | 自分に関する振る舞い          | 値に関するルール         | 複数モデル横断の振る舞い               |
| **例**         | User, Room, Message          | Email, Money, DateRange | UserInvitationService, PaymentCalculationService |





- ドメイン駆動設計で取り上げられるサービスは２種類ある。「ドメインのためのサービス」と「アプリケーションのためのサービス」

- エンティティや値オブジェクトに記載すると不自然になる振る舞いはドメインサービスに記載する。
- 逆に不自然なものだけを、サービスに記載する。実はすべての振る舞いはドメインサービスに記載できてしまうから。
全部ドメインサービスに記載してしまうとどの記載してしまうとどのクラス定義を見ただけでどんな振る舞いをするのか全然わからない。
結果、そのクラス（モデル）が意味の薄い無口なモデルになってしまう。ドメインドメインモデル貧血症。

- 不変条件（invariant）は、「そのモデル（EntityやValue Object）が存在する限り、常に成り立っていなければならないルールや制約」のことです。
ルールが1回でも破られたら、そのオブジェクトは“おかしい状態”になる
モデルのコンストラクタや状態変更メソッドの中で必ず守る
呼び出し元や外部のコードに依存せず、モデル自身が自分を守る

2. 例（現実世界）
銀行口座→残高は常に0以上。口座番号は必ず一意。など。
チャットルーム→定員を超えた人数は入れない。同じユーザーが二重参加しない。など

```go
package domain

import "errors"

type UserID string

var (
	ErrRoomFull      = errors.New("room is full")
	ErrAlreadyMember = errors.New("user already in room")
)

type Room struct {
	id       string
	capacity int
	members  map[UserID]struct{}
}

func NewRoom(id string, capacity int) *Room {
	return &Room{
		id:       id,
		capacity: capacity,
		members:  make(map[UserID]struct{}),
	}
}

func (r *Room) AddMember(uid UserID) error {
	// 不変条件のチェック
	if len(r.members) >= r.capacity {
		return ErrRoomFull
	}
	if _, exists := r.members[uid]; exists {
		return ErrAlreadyMember
	}

	r.members[uid] = struct{}{}
	return nil
}

// ここでは**「定員超過禁止」「重複禁止」が不変条件。
// AddMemberを通さず直接r.membersをいじれば壊せますが、それはドメインルール違反**。
```

```txt
実務での重要性
不変条件をモデルの外に置くと、呼び出すたびに同じチェックを書くことになり、抜け漏れやバグの原因になる
モデル内部で守れば、「壊れた状態」自体を作らせない
```
