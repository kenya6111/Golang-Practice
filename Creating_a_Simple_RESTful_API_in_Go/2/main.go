package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles (w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Article{Title:"test title", Desc:"test description", Content:"hello world"},
		Article{Title:"test title", Desc:"test description", Content:"hello world"},
		Article{Title:"test title", Desc:"test description", Content:"hello world"},
		Article{Title:"test title", Desc:"test description", Content:"hello world"},
	}
	fmt.Println("endpoint hit al articles endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "homepage endpoint hit")
}

func testPostArticles(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "test POST endpoint worked")
}

func handleRequest (){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main(){
	handleRequest()

}