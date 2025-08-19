package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load .env file: ", err)
	}

	user := os.Getenv("DBUser")
	password := os.Getenv("DBPassword")
	host := os.Getenv("DBHost")
	port := os.Getenv("DBPort")
	database := os.Getenv("DBName")

	db, err = sql.Open("postgres",
		fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database))

	if err != nil {
		log.Fatal("OpenError: ", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("PingError: ", err)
	}
	fmt.Println("connected!")

	// Drives
	fmt.Println(sql.Drivers())

	rows, err := db.Query("SELECT id, username, email FROM users")
	if err != nil {
		log.Fatalln(err)
	}

	// columnTypes
	cols, _ := rows.ColumnTypes()
	for _, c := range cols {
		fmt.Println("column name:", c.Name())
	}

	// DatabaseTypeName
	for _, v := range cols {
		fmt.Println(v.DatabaseTypeName())

		// DecimalSize
		fmt.Println(v.DecimalSize())

		// Length
		fmt.Println(v.Length())

		// Name
		fmt.Println(v.Name())

		// ScanType
		fmt.Println(v.ScanType())

		// Nullable
		fmt.Println(v.Nullable())
		fmt.Println("---")
	}

	//Conn (Type)
	//                  Conn = Connection（接続） の略。
	//                  つまり DBへの物理的な接続1本 を表す型
	//.                  DB（sql.DB）が「接続プール」を管理するのに対して、Conn は「その中の1つの実際のコネクション」を直接扱います。
	//                   接続プール？
	//                   プール＝コネクションを貯めておいて、みんなで必要なときに借りて使う仕組み
	//                   GOのプールはGo の *sql.DB はまさにこの「接続プール」を管理しています。
	//                   DBにアクセスするときは「コネクション（Connection＝物理的な接続）」を張る必要がある
	//                   でもコネクションの確立はコストが高い（ネットワーク確立、認証、初期化など）
	//                    あらかじめ複数のコネクションを作って「プール（＝貯めておく場所）」に入れておき、必要なときに使い回す仕組み が「接続プール」です。
	//                    Go の *sql.DB はまさにこの「接続プール」を管理しています。

	// DB
	// fmt.Println(db)
	fmt.Printf("%T\n", db)

	//. Begin
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	// Exec
	// _, err = tx.Exec("INSERT INTO users(id,username,email,password) VALUES(7,'Taro', 'aaadsda@gmail.comdsa','aaaaaaaaa')")
	//                Execは結果を返さないクエリ（INSERT、DELETE、CREATE TABLE, DROP TABLEなど）を実行できる
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	// _, err = tx.Exec("INSERT INTO users(id,username,email,password) VALUES(8,'Taro', 'aaa111@gmail.cssom','aaaaaaaaa')")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// Commit
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	// BeginTx              The provided context is used until the transaction is committed or rolled back.
	//                      コンテキストやオプションを指定できるBeginの上位版
	tx2, err2 := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err2 != nil {
		log.Fatal(err2)
	}
	id := 37
	_, execErr := tx2.Exec(`UPDATE users SET username = $1 WHERE id = $2`, "paid", id)
	if execErr != nil {
		_ = tx2.Rollback()
		log.Fatal(execErr)
	}
	if err := tx2.Commit(); err != nil {
		log.Fatal(err)
	}

	// Conn                 接続プールから「1本のコネクション」を取り出すメソッド
	//                      通常の db.Query() や db.Exec() は「プールから自動でコネクションを借りて返す」仕組みです。
	//
	ctx2 := context.Background()
	conn, err := db.Conn(ctx2)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//                      Conn のBeginTx 普通にトランザクションを始める
	tx2, err3 := conn.BeginTx(ctx2, nil)
	if err3 != nil {
		log.Fatal(err)
	}
	tx2.ExecContext(ctx2, "UPDATE users SET username = $1 WHERE id = $2", "update name", "1")
	if rows, err := tx2.QueryContext(ctx2, "select id,username from users where id < $1", 3); err == nil {
		for rows.Next() {
			var id int
			var name string

			rows.Scan(&id, &name)
			fmt.Println("id=", id, "name=", name)
		}
	}
	stmt, err := tx2.Prepare("UPDATE users SET username = $1 WHERE id = $2")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec("prepare test", 1)
	if err := tx2.Commit(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("トランザクション成功")

}
