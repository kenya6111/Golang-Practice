package main

import (
	"app/todo_app/app/models"
	"fmt"
	"log"
)

func main() {
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)

	log.Println("testsasasasas")
	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.Password = "password"
	fmt.Println(u)

	u.CreateUser()
}
