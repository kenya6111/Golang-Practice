package kyu_6

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
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

func Test_9_6() {
	// ctx := context.Background()

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("failed t")
	// }

	// user := os.Getenv("DBUser")
	// password := os.Getenv("DBPassword")
	// host := os.Getenv("DBHost")
	// port := os.Getenv("DBPort")
	// database := os.Getenv("DBName")

	// config, err := pgx.ParseConfig(fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// config.Logger = &logger{}

	// conn, err := pgx.ConnectConfig(ctx, config)

	// if err != nil {
	// 	log.Fatalf("connect: %v\n", err)
	// }
	// tx, err := conn.Begin(ctx)
	// stmt, err := tx.Prepare(ctx, "insert", "insert into users (id ,username) values($1,$2);")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// sql := `select schemaname,tablename from pg_tables where schemaname = $1;`
	// args := `information_schema`

	// rows, err := conn.Query(ctx, sql, args)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer rows.Close()

	// var pgtables []PgTable

	// for rows.Next() {
	// 	var s, t string
	// 	if err := rows.Scan(&s, &t); err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	pgtables = append(pgtables, PgTable{schemaName: s, TableName: t})
	// }
	// if err := rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }

}
