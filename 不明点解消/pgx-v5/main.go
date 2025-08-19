package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println(111)
	ctx := context.Background()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load .env file: ", err)
	}

	user := os.Getenv("DBUser")
	password := os.Getenv("DBPassword")
	host := os.Getenv("DBHost")
	port := os.Getenv("DBPort")
	database := os.Getenv("DBName")

	fmt.Println(user, password, host, port, database)

	conn, err := pgx.Connect(ctx, fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database))
	if err != nil {
		log.Fatal("connect error:", err)
	}
	defer conn.Close(ctx)

	tx, err := conn.Begin(ctx)

	if err != nil {
		log.Fatal("begin error:", err)
	}
	defer tx.Rollback(ctx)

	// tx.Exec
	_, err = tx.Exec(ctx, "UPDATE users SET username=$1 WHERE id=$2", "new_name", 1)
	if err != nil {
		log.Fatal("update error:", err)
	}

	// tx.Commit
	if err := tx.Commit(ctx); err != nil {
		log.Fatal("commit error:", err)
	}

	// QueryRow
	row := conn.QueryRow(ctx, "select username from users")
	fmt.Println(row)
	var username string
	row.Scan(&username)
	fmt.Println("username:", username)

	// Query
	rows, err := conn.Query(ctx, "select username from users")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("username", name)
	}
	rows.Close()

	// Exec
	ct, _ := conn.Exec(ctx, "UPDATE users SET username=$1 WHERE id=$2", "new_name!!", 1)
	fmt.Println("rows affected:", ct.RowsAffected())

	// conn.Begin
	// tx2, err2 := conn.Begin(ctx)
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }
	// // tx.Rollback
	// defer tx2.Rollback(ctx)

	// // tx.Exec
	// _, _ = tx2.Exec(ctx, "INSERT INTO users(username, email) VALUES ($1, $2)", "temp", "temp@example.com")

	// var id int
	// _ = tx2.QueryRow(ctx, "SELECT id FROM users WHERE email=$1", "temp@example.com").Scan(&id)
	// fmt.Println("inserted id:", id)

	// // tx.commit
	// if err := tx2.Commit(ctx); err != nil {
	// 	log.Fatal(err)
	// }

	rows2, _ := conn.Query(ctx, "SELECT id, username FROM users LIMIT 5")

	type User struct {
		ID   int
		Name string
	}

	users, err := pgx.CollectRows(rows2, func(row pgx.CollectableRow) (User, error) {
		var u User
		err := row.Scan(&u.ID, &u.Name)
		return u, err
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("users:", users)

	fmt.Println(pgx.TextFormatCode)
	fmt.Println(pgx.BinaryFormatCode)

	rows3, _ := conn.Query(ctx, "SELECT username FROM users")
	// fmt.Println(rows3)
	names, _ := pgx.AppendRows([]string{}, rows3, func(row pgx.CollectableRow) (string, error) {
		var name string
		err := row.Scan(&name)
		return name, err
	})

	fmt.Println(names)

	// BeginFunc
	ctx3 := context.Background()
	pgx.BeginFunc(ctx3, conn, func(tx pgx.Tx) error {
		_, err := tx.Exec(ctx, "UPDATE users SET username=$1 WHERE id=$2", "new_name!!", 1)
		if err != nil {
			return err
		}

		_, err = tx.Exec(ctx, "UPDATE users SET username=$1 WHERE id=$2", "new_name!!!!!!!", 2)
		return err
		//	            BeginFunc が conn.Begin(ctx) を呼ぶ
		//	            fn(tx) が実行される
		//            	fn が nil を返した → Commit
		//	            fn が error を返した → Rollback
		//              普通に書くよりも短縮してかけて見通しが良いのがメリット
	})

	//                  if err != nil {
	// 	                  log.Fatal("transaction failed:", err)
	//                  }
	rows4, err := conn.Query(ctx, "SELECT id, username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	for rows4.Next() {
		var id int
		var name string

		rows4.Scan(&id, &name)
		fmt.Println(id, name)
	}
	fmt.Println("------")

	rows5, err := conn.Query(ctx, "SELECT id, username FROM users where id = $1", 1)
	if err != nil {
		log.Fatal(err)
	}

	// CollectExactlyRow
	user3, err := pgx.CollectExactlyOneRow(rows5, func(row pgx.CollectableRow) (User, error) {
		var u User
		err := row.Scan(&u.ID, &u.Name)
		fmt.Println(u)
		return u, err
	})

	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}

	fmt.Println("user3: ", user3)

}
