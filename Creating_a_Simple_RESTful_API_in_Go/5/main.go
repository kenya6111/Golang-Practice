package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	CreatedAt string `json:created_at`
	email string `json:email`

}

func main() {
	fmt.Println("Go MySQL tutorial")

	// MySQL に接続
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 接続確認
	if err := db.Ping(); err != nil {
		panic(err)
	}

	// // INSERT 実行
	// query := `
	// 	INSERT INTO user (id, name, password, created_at, email)
	// 	VALUES (?, ?, ?, ?, ?)
	// `
	// _, err = db.Exec(query, 20, "ELLIOT", "3", "2020-11-11", "elliot@example.com")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("✅ Successfully inserted into MySQL database")

	results , err := db.Query("select * from user")

	if err != nil{
		panic(err.Error())
	}

	for results.Next(){
		var user User

		err = results.Scan(&user.Id, &user.Name, &user.Password, &user.CreatedAt, &user.email)

		if err != nil{
			panic(err.Error())
		}
		fmt.Println(user.Name)
	}

	
	

}
