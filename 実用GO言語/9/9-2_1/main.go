package kyu_2_1

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type txAdmin struct {
	*sql.DB
}

type Service struct {
	tx txAdmin
}

func (t *txAdmin) Transaction(ctx context.Context, f func(ctx context.Context) (err error)) error {
	tx, err := t.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := f(ctx); err != nil {
		return fmt.Errorf("transaction query failed %w", err)
	}
	return tx.Commit()
}
func (s *Service) updateUser(ctx context.Context, userId int, userName string) error {
	updateFunc := func(ctx context.Context) error {
		if _, err := s.tx.ExecContext(ctx, "update users set username = $1 where id =1", userName); err != nil {
			return err
		}
		if _, err := s.tx.ExecContext(ctx, "update users set username = $1 where id =2", userName); err != nil {
			return err
		}
		return nil
	}
	return s.tx.Transaction(ctx, updateFunc)
}

func Test_9_2_1() {

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

	txadmin := &txAdmin{
		DB: db,
	}

	service := &Service{
		tx: *txadmin,
	}

	ctx := context.Background()

	err2 := service.updateUser(ctx, 1, "update user by transaction 2")
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(err2)
}
