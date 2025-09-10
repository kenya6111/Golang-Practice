## 値オブジェクト

#### 値オブジェクトとは
システム固有の値を表したオブジェクトを値オブジェクトと呼ぶ
このFullNameはその名のとおり氏名を表現したオブジェクトで、値の表現。
```txt
class FullName
{
  public FullName(string firstName, string lastName)
  {
    FirstName = firstName;
    LastName = lastName;
  }

  public string FirstName { get; }
  public string LastName { get; }
}

```

<br/>

#### 値の性質
```txt
・不変である
・交換が可能である
・等価性によって比較される
```

<br/>

- 不変である
```txt
var greet = "こんにちは";
Console.WriteLine(greet); // こんにちは が表示される
greet = "Hello";
Console.WriteLine(greet); // Hello が表示される
```
  - 上記では「値」が変更されているわけではなく、変更されているのは「変数の内容」。なので値は不変。代入によって変数の内容を変更している。
  - なぜ不変が大事か
    - バグを減らせる
      - いつの間にか値が変わっていた。。ってのを防ぐ

- 交換が可能である
```txt
var fullName = new FullName("masanobu", "naruse");
fullName = new FullName("masanobu", "sato");
```
  - fullNameプロパティの内部を書き換えるのではなく、そのまま置き換える


- 等価性によって評価される

#### 値オブジェクトのメリット
```txt
・表現力を増す
・不正な値を存在させない
・誤った代入を防ぐ
・ロジックの散在を防ぐ
```

- 表現力を増す

製品番号を持つプロパティがあるとして、「void Method(string modelNumber)」てメソッドがあったとして、どんな文字列が入ってくるのかよくわからんよね。

```txt
var modelNumber = "a20421-100-1";
```

なので以下みたいに製品番号をさらに分割プロパティしてModelNumberのプロパティとして持たせることをすれば、つまり値オブジェクトによって自分がどういうものか自己文書化できる

```txt
class ModelNumber
{
  private readonly string productCode;
  private readonly string branch;
  private readonly string lot;

  public ModelNumber(string productCode, string branch, string lot)
  {
    if (productCode == null) throw new ArgumentNullException(nameof(productCode));
    if (branch == null) throw new ArgumentNullException(nameof(branch));
    if (lot == null) throw new ArgumentNullException(nameof(lot));

    this.productCode = productCode;
    this.branch = branch;
    this.lot = lot;
  }

  public override string ToString()
  {
    return productCode + "-" + branch + "-" + lot;
  }
}
```

- 不正な値を存在させない
ユーザ名は3文字以上」というルールがある場合、例えば2文字の名前が登録されてしまうと、その後ユーザ名を使う時に毎回いろんな箇所で三文字以上かチェックする必要がある。

何度も同じチェックを書くことになるし一ヶ所でも間違えばシステムが破綻する

なのでインスタンス生成時に「this.value=value」の前に「if(value.length()>3)」のようなチェック処理を入れておけば良い

**不正な値は遅効性の毒みたいなもの**

- ロジックの散在を防ぐ
オブジェクトのユーザ名3文字以上というチェックは、userオブジェクトのインスタンス化の箇所のみに書いておいて
ユーザの新規作成の時はそのインスタンス化処理を使うのでチェックがしっかりされるし、更新の時も、「値は不変」の原則からユーザオブジェクトを更新したい値で新規作成、つまりインスタンス化して「代入」するので、更新する時もユーザ名チェックがされて良い。
つまりは「値は不変」。更新の時は置き換え。という原則を守ればいい。



#### GOでの説明
DDDでは 「同一性（ID）」を持たず、値そのものが等しければ同一とみなせるオブジェクト を 値オブジェクト と呼びます。

voと略されがち

特徴

- IDを持たない
  - エンティティのように一意な識別子が不要
  - 値が同じなら同一とみなせる

- 不変（Immutable）であることが望ましい
  - 値オブジェクトは生成後に変更しない
  - 変更が必要なら新しいインスタンスを作る

- 概念を表現する
  - 単なるプリミティブ型（string, intなど）ではなく、意味を持った型 にする


```go
package user

import (
    "errors"
    "regexp"
)

type Email struct {
    value string
}

// コンストラクタ（不正な値は作れない）
func NewEmail(value string) (Email, error) {
    if !isValidEmail(value) {
        return Email{}, errors.New("invalid email format")
    }
    return Email{value: value}, nil
}

func (e Email) String() string {
    return e.value
}

// 値が同じなら同一とみなす
func (e Email) Equals(other Email) bool {
    return e.value == other.value
}

func isValidEmail(email string) bool {
    re := regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)
    return re.MatchString(email)
}

```

```txt
表現力が増す
  string や int ではなく、Email や Money という型でドメインの意味を直接表現できる

不変性で安全
  値の整合性が保たれる

バリデーションを閉じ込められる
  VOのコンストラクタ内で不正値を防げる
```