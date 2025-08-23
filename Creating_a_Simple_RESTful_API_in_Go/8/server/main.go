package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey =[]byte("mysupersecretphrase")

func HomePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"super secret information")
}

func isAuthorised(endpoint func(http.ResponseWriter, *http.Request)) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if r.Header["Token"] != nil{
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token)(interface{},error){
				if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok{//SigningMethodHS256
					return nil,fmt.Errorf("there was an error")
				}
				return mySigningKey, nil
			})

			if err != nil{
				fmt.Println(111)
				fmt.Fprintf(w,err.Error())
			}

			if token.Valid{
				endpoint(w,r)
			}
		}else{
			fmt.Fprintf(w,"not authorised")
		}
	})
}


func handleRequests(){
	http.Handle("/",isAuthorised(HomePage))
	log.Fatal(http.ListenAndServe(":9000", nil))
}


func main(){
	fmt.Println("simpel server")
	handleRequests()

}