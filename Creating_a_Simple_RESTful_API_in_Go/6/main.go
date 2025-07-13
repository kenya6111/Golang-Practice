package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func helloworld (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"hello world")

}
func handleRequest (){
	myRouter :=mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloworld).Methods("GET")
	myRouter.HandleFunc("/users",AllUsers).Methods("GET")
	myRouter.HandleFunc("/users/{name}/{email}",NewUser).Methods("POST")
	myRouter.HandleFunc("/users/{name}",DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/users/{name}/{email}",UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081",myRouter))
}



func main(){
	fmt.Println("go orm tutorial ")
	InitialMigration()
	
	handleRequest()
}