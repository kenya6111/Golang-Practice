package kyu_5

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

var _ pgx.Logger = (*logger)(nil)

type logger struct{}
type PgTable struct {
	schemaName string
	TableName  string
}

func (l *logger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	if msg == "Query" {
		log.Printf("SQL:\n%v\nARGS:%v\n", data["sql"], data["args"])
	}
}

func Test_9_5() {
	ctx := context.Background()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed t")
	}

	user := os.Getenv("DBUser")
	password := os.Getenv("DBPassword")
	host := os.Getenv("DBHost")
	port := os.Getenv("DBPort")
	database := os.Getenv("DBName")

	config, err := pgx.ParseConfig(fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database))
	if err != nil {
		fmt.Println(err)
	}

	config.Logger = &logger{}

	conn, err := pgx.ConnectConfig(ctx, config)

	if err != nil {
		log.Fatalf("connect: %v\n", err)
	}

	sql := `select schemaname,tablename from pg_tables where schemaname = $1;`
	args := `information_schema`

	rows, err := conn.Query(ctx, sql, args)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var pgtables []PgTable

	for rows.Next() {
		var s, t string
		if err := rows.Scan(&s, &t); err != nil {
			fmt.Println(err)
		}
		pgtables = append(pgtables, PgTable{schemaName: s, TableName: t})
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// 検証ーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーー
	fmt.Println("ーーーーーーーーーーーーーーー")
	fmt.Println(s.Speak())

	x := 10
	px := &x        // x のアドレスを px に入れる
	fmt.Println(x)  // 10
	fmt.Println(px) // 0x14000122018  ← メモリ上の住所っぽい数値

	fmt.Println("ーーーーーーーーーーーーーーー")
	xx := 11
	pxx := &xx

	fmt.Println(xx)
	fmt.Println(pxx)
	fmt.Println(*pxx)
	*pxx = 22
	fmt.Println(*pxx)
	fmt.Println(pxx)
	var _ Speaker = (*Person)(nil) // var _ インターフェース型 = (型)(値)
	// var _ Speaker = (*Animal)(nil) // Speak がないからエラーになる

	var s Stringer = MyInt(0)
	fmt.Println(s)

	var _ Stringer = (*MyInt)(nil)
	var s2 Stringer
	s2 = MyInt(42)
	fmt.Println(s2.String())

	// 	var _ インターフェース = 型(値)
	// → 値レシーバーなら MyInt(0) でも代入できる

	// var _ インターフェース = (*型)(nil)
	// → ポインタレシーバーの場合はこっち
}

type Speaker interface {
	Speak() string
}

type Person struct{}

func (p *Person) Speak() string {
	return "hello"
}

var s Speaker = &Person{}

type Animal struct{}

type Stringer interface {
	String() string
}

type MyInt int

func (m MyInt) String() string {
	return fmt.Sprintf("MyInt value: %d", m)
}
