package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"

	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"tutorial.sqlc.dev/app/tutorial"
)

// var Conn *sql.DB

func main() {
	fmt.Println(11)
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
	fmt.Println(user, password, host, port, database)

	ctx := context.Background()
	// conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database))

	// queries := tutorial.New(conn)
	// Conn, err = pgx.Open("postgres",
	// 	fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database))

	// if err != nil {
	// 	log.Fatal("OpenError: ", err)
	// }

	pool, err := pgxpool.New(ctx, fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database))

	queries := tutorial.New(pool)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(authors)

	// create an author
	result, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio: pgtype.Text{
			String: "Co-author klof The C Programming Language and The Go Programming Language",
			Valid:  true,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(result.RowsAffected())

	var insertedAuthorID = 1
	fetchedAuthor, err := queries.GetAuthor(ctx)
	if err != nil {
		fmt.Println(err)
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthorID, fetchedAuthor.ID))
	fmt.Println(err)

}
