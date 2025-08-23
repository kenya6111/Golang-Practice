package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey =[]byte("mysupersecretphrase")
// var mySigningKey = os.Getenv("JWT_TOKEN")


func homepage (w http.ResponseWriter, r *http.Request){
	validToken ,err := GenerateJWT()

	if err != nil{
		fmt.Println(w,err.Error())
	}
	fmt.Fprintf(w,validToken)
}

func GenerateJWT()(string,error){
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Elliot forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil{
		fmt.Errorf("somthing went wrong %s", err.Error())
		return "", err
	}

	return tokenString , nil

}


func handleRequests(){
	http.HandleFunc("/",homepage)
	log.Fatal(http.ListenAndServe(":9001", nil))

}


func main(){

	// tokenString , err := GenerateJWT()

	// if err != nil{
	// 	fmt.Println("error generate token string")
	// }

	// fmt.Println(tokenString)

	handleRequests()
}