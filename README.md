# Golang-Practice
- 明示的な定義
  - varを利用して変数を定義する場合、型の指定が必要です。varを利用すると複数の異なる型の変数を () で囲み、まとめて定義できます。
  - デフォルト値ですが、数値系は 0、文字列は 空文字、ポインタは nil になっている
  ```go
  	var (
      i2 int    = 200
      s2 string = "golung"
    )

    fmt.Println(i2, s2)// 200golung
  ```
- 暗黙的な定義
  - 型の指定をしないでも良い。
  - 「:=」で定義する
  ```go
  	i10 := 300
    fmt.Print(i10)// 300
    fmt.Print("\n")
    fmt.Printf("i10=%T", i10)// int

    // 一度 :=で定義した変数には異なる型の値は代入できない
    i10="abc"// cannot use "aaaaa" (untyped string constant) as int value in assignmentcompilerIncompatibleAssign
  ```
  - :=は関数内でしか定義できない
    - 例えばグローバルで :=で変数定義すると、以下のエラー分がでる
    - section3/main.go:5:1: syntax error: non-declaration statement outside function body
    - 逆に明示的な定義であればグローバルで定義できる
    - 基本的に関数内か関数が以下で使い分けるが、できれば明示的あ定義を基本使うこと
- int型
  - int8: 符号付き 8ビット整数
    - 取り得る値: -128 ～ 127
    - メモリ使用量: 1バイト
  - int16: 符号付き 16ビット整数
    - 取り得る値: -32768 ～ 32767
    - メモリ使用量: 2バイト

  - int32: 符号付き 32ビット整数
    - 取り得る値: -2,147,483,648 ～ 2,147,483,647
    - メモリ使用量: 4バイト

  - int64: 符号付き 64ビット整数
    - 取り得る値: -9,223,372,036,854,775,808 ～ 9,223,372,036,854,775,807
    - メモリ使用量: 8バイト
  - Go には int8, int16, int32, int64 とは別に int という型もあります。
    - int はその環境(実行環境)の デフォルトのワードサイズ になります
    - 32ビット環境なら int は 32ビット
    - 64ビット環境なら int は 64ビット
  - 使い分け
    - 取り扱う数値の範囲が非常に大きい or 小さいことが明確な場合
    - → int64 や int8 などを明示的に使う

  - 普通に整数演算をするだけで、特に気にする必要がない場合
    - → int を使えば OK (Go においては、ほとんどの場合 int で十分)

- float
  - 
- byte型
  ```go

	byteA := []byte{72, 73}
	fmt.Println(byteA) // [72 73]が出る

	fmt.Println(string(byteA)) // HI が出る

  c := []byte("HI") // バイト配列に直す
	fmt.Println(c)    // [72 73]が出力される
  ```
- 配列型
  ```go
	// GOの配列型は、あとから要素数を変更できない。増減できない
	fmt.Println("----------")
	var arr1 [3]int
	fmt.Println(arr1)        // [0 0 0]
	fmt.Printf("%T\n", arr1) // [3]int

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2) // [A B ]

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3) // [1 2 3]

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)        // [C D]
	fmt.Printf("%T\n", arr4) // [2]string

	fmt.Println(arr1[0]) // 0
	fmt.Println(arr2[0]) // A

	arr2[2] = "E"
	fmt.Println(arr2[2]) // E

	fmt.Println(len(arr1)) // 3
  ```
- interface
  ```go
  	var x interface{} // あらゆる型と互換性がある
    fmt.Println(x)    // nil 初期値はnilとなっている。pythonでいうところのNoneになる。

    x = "aaa"
    fmt.Println(x) /// aaa
    x = 111
    fmt.Println(x) // 111
    x = true
    fmt.Println(x) // true

    x = [3]int{1, 2, 3}
    fmt.Println(x) //[1 2 3]
  ```
  - あくまでinterface型はすべての型を汎用的に表す手段であって、演算の対象としては利用できないことに注意。
- 型変換
  ```go
  	var i_3 int = 1
    fmt.Printf("%T\n", i_3) //int
    fl64_2 := float64(i_3)
    fmt.Println(fl64_2)        // 1
    fmt.Printf("%T\n", fl64_2) // float64

    inti_3 := int(fl64_2)
    fmt.Println(inti_3)      // 1
    fmt.Printf("%T", inti_3) // int

    	var s_4 string = "100"
      fmt.Printf("%T\n", s_4) // string

      // strconvで文字列から数値に変換できる
      i, _ = strconv.Atoi(s_4) // アンダースコアとなることで、関数から帰ってくる2つ目の値を使わないとできる。
      fmt.Printf("%T\n", i) // int
      fmt.Print(i)          // 100
  ```


- 定数
  ```go
  // 頭文字を大文字にすると他パッケージから呼び出せる
  const Pi = 3.14

  // まとめて定数定義する
  const (
    URL      = "http://xxx.co.jp"
    SiteName = "test"
  )

  // 同じ値をまとめて定数定義する
  const (
    A = 1
    B
    C
    D = "A"
    E
    F
  )
  fmt.Println(A, B, C, D, E, F)// 1,1,1,A,A,A

  // iotaは連番を生成する
  const (
    c0 = iota
    c1
    c2
  )
  fmt.Println(c0, c1, c2)// 0 1 2
  ```

- fmt.Print()
  - 与えられた引数をスペースで区切って標準出力に出力する。末尾に改行は追加されない
  ```go
  fmt.Print(1, 2, 3)// 1 2 3
  ```
- fmt.Printf()
  - 書式指定子付きの文字列と、与えられた引数を標準出力に出力する。末尾に改行は追加されない
  ```go
	name := "ALice"
	age := 30
	fmt.Printf("名前は%vです。年齢は%vです", name, age)// 名前はALiceです。年齢は30です
  ```
- fmt.Println()
  - 与えられた引数をスペースで区切って標準出力に出力し、最後に改行を追加します。

- varブロック
  ```go
  var (
    i2 int    = 200
    s2 string = "golung"
  )
  ```
  - 複数の var 宣言をまとめて書く ための書き方
  - これは以下のように書いたのと同じ
  ```go
    var i2 int    = 200
    var s2 string = "golung"
  ```
  - メリット
    - 関連する変数をひとまとめに宣言できる
      - たとえばパッケージレベル (グローバルスコープ) で初期化する定数・変数が多い場合など、それらを視覚的にまとめられるメリットがあります。
  - 可読性の向上
    - 関連した変数をブロックにまとめると可読性が上がり、あとから見返したときにも整理しやすくなります。

- %v
  - 変数の値を表現する書式指定子
  ```go
	name := "ALice"
	age := 30
	fmt.Printf("名前は%vです。年齢は%vです", name, age)// 名前はALiceです。年齢は30です
  ```
- %T 
  - 値の型を表現する書式指定子
  ```go
  fmt.Printf("%Tと%Tと%T", 1, 2, 3)// intとintとint
  ```



- 
	// 暗黙的な定義（明示的な定義と比べて、型指定の必要がない）
	//:= は、「短変数宣言 (short variable declaration)」 と呼ばれる構文です。
	// 変数の型を明示的に指定せずに、新しい変数を宣言しながら初期化する際に使われます。
  // 基本的には明示的な型指定をする明示的な定義を使った方が良いとされている
	// 型指定をすることでバグを抑えるように元々設計された言語なので。



func (t Time) Year() int の (t Time) は レシーバ（receiver） と呼ばれます。
Go ではメソッドを「どの型にぶら下げるか」を明示するために、関数の先頭に (変数名 型) を付けます。

func (t Time) Year() int { … }
Time 型に Year というメソッド を定義する。

メソッド内で t という名前で受け取るのは呼び出し元の値（コピー）。

呼び出しは someTime.Year() の形になる。
内部では Go が Year(someTime) のように変換して呼んでいるだけ。

https://cs.opensource.google/go/go/+/refs/tags/go1.24.3:src/time/time.go;l=140#:~:text=//%20clock%20reading.-,type%20Time%20struct%20%7B,-//%20wall%20and%20ext
