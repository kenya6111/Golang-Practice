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

- ロジックの散在を防ぐ
オブジェクトのユーザ名3文字以上というチェックは、userオブジェクトのインスタンス化の箇所のみに書いておいて
ユーザの新規作成の時はそのインスタンス化処理を使うのでチェックがしっかりされるし、更新の時も、「値は不変」の原則からユーザオブジェクトを更新したい値で新規作成、つまりインスタンス化して「代入」するので、更新する時もユーザ名チェックがされて良い。
つまりは「値は不変」。更新の時は置き換え。という原則を守ればいい。