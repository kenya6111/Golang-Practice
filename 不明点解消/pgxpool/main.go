package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// pgxpool →postgresへの接続を効率よくする仕組み

// 			例え話
// 			レストランで毎回お客さんが来るたびに「シェフを新しく雇って → 面接して → 教育して → 料理を作らせる」だと効率悪いですよね。
// 			→ 普通は「シェフを常駐させておいて、注文が来たらすぐ対応」します。
// 			データベース接続も同じ
// 			DB に接続するには「ハンドシェイク（挨拶の儀式）」が必要で、毎回これをやるのは重い処理
// 			だから「プール」＝ あらかじめ接続をいくつか作ってためておく
// 			使い終わったら閉じずに プールに返して再利用する
// 			👉 これが コネクションプール の考え方です。

func main() {
	// ctx := context.Background()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load .env file: ", err)
	}

	user := os.Getenv("DBUser")
	password := os.Getenv("DBPassword")
	host := os.Getenv("DBHost")
	port := os.Getenv("DBPort")
	database := os.Getenv("DBName")
	poolMaxConns := os.Getenv("pool_max_conns")
	poolMaxConnLifetime := os.Getenv("pool_max_conn_lifetime")

	fmt.Println(user, password, host, port, database)

	// ParseConfig
	config, err := pgxpool.ParseConfig(fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable&pool_max_conns=%s&pool_max_conn_lifetime=%s", user, password, host, port, database, poolMaxConns, poolMaxConnLifetime))
	if err != nil {
		fmt.Println(err)
	}
	// 								pgxpool.ParseConfig は実務だと 「接続文字列から一度 Config を作って、
	// 								そこに細かい設定やフックを追加してから pgxpool.NewWithConfig でプールを作る」 という使い方をします。
	//								ParseConfig は「文字列 + α（コードで制御可能なオプション）」を持った設定を作るための入り口。
	//								ただの文字列で済むなら pgxpool.New で十分、でも実務では「ちょっとカスタムしたい」が出てくるので ParseConfig が便利。

	//								-- ただの文字列
	//									dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
	//								    	user, pass, host, port, dbname)
	//									pool, _ := pgxpool.New(ctx, dsn) // 文字列ベース

	//								-- ParseConfigを使う
	//									config, _ := pgxpool.ParseConfig(dsn)

	//									文字列では無理な設定を追加
	//									config.MaxConns = 20
	//									config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
	//										_, err := conn.Exec(ctx, "SET TIME ZONE 'Asia/Tokyo'")
	//										return err
	//									}
	// 									pool, _ := pgxpool.NewWithConfig(ctx, config) // 柔軟に設定できる

	// Config
	fmt.Println("---------")
	fmt.Println("config: ", config)
	fmt.Println("---------")
	fmt.Println(config.MaxConns)
	fmt.Println(config.MinConns)
	fmt.Println(config.MinIdleConns)
	fmt.Println(config.AfterConnect)  // 新しい接続が確立された直後に呼ばれるフック
	fmt.Println(config.AfterRelease)  // 接続をプールに返すときに呼ばれる。
	fmt.Println(config.BeforeAcquire) // プールから接続を借りる直前に呼ばれる。
	fmt.Println(config.BeforeClose)   // 接続がプールから完全に削除される直前に呼ばれる。
	fmt.Println(config.BeforeConnect)
	fmt.Println(config.ConnConfig)
	fmt.Println(config.HealthCheckPeriod)
	fmt.Println(config.MaxConnIdleTime)
	fmt.Println(config.MaxConnLifetime)
	fmt.Println(config.MaxConnLifetimeJitter)
	fmt.Println(config.ConnString())
	fmt.Println(config.Copy())
	fmt.Println(config.ConnConfig.Config)
	fmt.Println(config.ConnConfig.Host)
	fmt.Println(config.ConnConfig.Port)
	//						クライアント
	//					│ Acquire (借りる)
	//					▼
	//					[プール] ← BeforeAcquire → (OK? NGなら破棄)
	//					│
	//					│ AfterConnect (接続直後の初期化)
	//					│
	//					▼
	//					[接続]   ← AfterRelease → (返却or破棄)
	//					│
	//					▼ BeforeClose (完全削除前)
	//					要は接続プールと接続プールと各コネクションのライフサイクルを管理するための仕組み

	//					プールが新しい接続を作るとき
	//						BeforeConnect / AfterConnect
	//						👉 ここで「接続先を変える」「初期化SQLを叩く」みたいなことができる。

	//					プールから接続を貸し出すとき
	//						BeforeAcquire
	//						👉 借りようとした接続が「壊れてないか？」をチェックできる。
	//						壊れてたら破棄して別の接続を探す。

	//					プールに接続を返すとき
	//						AfterRelease
	//						👉 正常なら再利用、問題あれば捨てる。
	//						（例：トランザクションが中途半端に残ってるなら破棄する）

	//					接続を閉じるとき
	//						BeforeClose
	//						👉 完全に捨てる直前に「ログを出す」とか「監視に通知する」。

	// 					プール全体のサイズや寿命を決める
	//						MaxConns, MinConns, MinIdleConns, MaxConnLifetime, MaxConnIdleTime など
	//						👉 どれくらいの本数の接続を常備して、どのくらいで入れ替えるかを決める。

	//					壊れてないか定期チェック
	//						HealthCheckPeriod
	//						👉 アイドル中の接続を定期的に ping して、死んでたら閉じる。

	// Config
	fmt.Println(config.ConnString()) // 接続文字列を返す。
	// 									ログ出力やデバッグに使ったりする

	// Copy
	fmt.Println(config.Copy().ConnString()) // *Configの構造体をそのまま再取得できる
	//										設定をまるごとコピーして新しいプールを作りたいときに使う。

	// New
	ctx := context.Background()
	pool, err := pgxpool.New(
		ctx,
		fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable&pool_max_conns=%s&pool_max_conn_lifetime=%s", user, password, host, port, database, poolMaxConns, poolMaxConnLifetime),
	)
	fmt.Println(pool)
	fmt.Printf("%T\n", pool)

	pool2, err := pgxpool.NewWithConfig(ctx, config)
	fmt.Println(pool2)
	fmt.Printf("%T\n", pool2)

	// Acquire (Acquire はプールから接続 (*Conn) を返す。)
	conn, err := pool.Acquire(ctx)

	fmt.Println(conn)

}
