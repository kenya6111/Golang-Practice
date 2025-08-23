package kyu_1

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type User struct {
	UserID   string
	UserName sql.NullString
	Email    string
}

func Test_9_1() {
	fmt.Println(111)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed t")
	}

	user := os.Getenv("DBUser")
	password := os.Getenv("DBPassword")
	host := os.Getenv("DBHost")
	port := os.Getenv("DBPort")
	database := os.Getenv("DBName")
	fmt.Println(user, password, host, port)

	db, err := sql.Open("pgx",
		fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database))

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	rows, err := db.QueryContext(ctx, "select id, username, email from users order by id;")
	if err != nil {
		fmt.Println(err)
	}

	var users []*User

	for rows.Next() {
		var (
			userID, email string
			userName      sql.NullString
		)
		fmt.Println(1)
		if err := rows.Scan(&userID, &userName, &email); err != nil {
			log.Fatalf("scan the user: %v", err)
		}

		users = append(users, &User{
			UserID:   userID,
			UserName: userName,
			Email:    email,
		})

		if userName.Valid {
			fmt.Println("is not null string", userName.String, userName.Valid)
		} else {
			fmt.Println("is null string", userName.String, userName.Valid)

		}
	}
	fmt.Println(users)
	if err := rows.Close(); err != nil {
		log.Fatalf("rows close : %v", err)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("scan users : %v", err)
	}

	var (
		userID, userName, email string
	)
	row := db.QueryRowContext(ctx, "select id, username, email from users where id =1 order by id;")
	err2 := row.Scan(&userID, &userName, &email)
	if err2 != nil {
		fmt.Println(err)
	}
	fmt.Println(userID, userName, email)

}
