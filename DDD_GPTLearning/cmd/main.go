package main

import (
	"ddd_gpt_learning/application"
	"ddd_gpt_learning/infrastructure/config"
	"ddd_gpt_learning/infrastructure/postgres"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// DB接続
	db, err := config.NewDB()
	if err != nil {
		log.Fatal("failed to connect db:", err)
	}
	defer db.Close()

	// Repository作成
	repo := postgres.NewUserRepository(db)
	service := application.NewUserService(repo)

	// ユースケース: 登録
	newUser, err := service.ResisterUser("kenya_service", "kenya_service@example.com")
	if err != nil {
		log.Fatal("【insert error】:", err)
	}
	fmt.Println("新規ユーザーを保存しました:", newUser)

	// ユースケース: 一覧取得
	users, err := service.GetUsers()
	if err != nil {
		log.Fatal("select error:", err)
	}

	fmt.Println("DBに存在するユーザー一覧:")
	for _, u := range users {
		fmt.Printf("- %s (%s)\n", u.Username, u.Email)
	}

}
