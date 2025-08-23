package kyu_2

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Service struct {
	db *sql.DB
}

func (s *Service) updateUser(ctx context.Context, userId int, userName string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, "update users set username = $1 where id =1", userName); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
func Test_9_2() {

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

	service := &Service{
		db: db,
	}
	ctx := context.Background()

	err2 := service.updateUser(ctx, 1, "transactionTest: username")
	if err2 != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

}
