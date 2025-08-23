package main

import (
	"encoding/json"
	"fmt"
)

// type Book struct{
// 	Title string `json:"title"`
// 	Author Author `json:"author"`
// }

// type Author struct {
// 	Name string `json:"name"`
// 	Age int `json:"age"`
// 	Developer bool `json:"is_developer"`
// }


type SensorReading struct{
	Name string `json:"name"`
	Capacity int `json:"capacity"`
	Time string `json:"time"`
	Information Info `json:"info"`
}


type Info struct{
	Description string `json:"desc"`
}
func main(){

	// author := Author{Name:"kenya " , Age:22, Developer: true}
	// book := Book{Title:"test title", Author:author}

	// fmt.Printf("%+v\n", book)

	// byteArray, err := json.MarshalIndent(book, "","   ")

	// if err != nil{
	// 	fmt.Println(err)
	// }

	// fmt.Println(string(byteArray))
	jsonString := `{"name":"kenya","capacity":40,"time":"2025-11-11", "info":{"desc":"aaaaaaa"}}`

	var reading SensorReading
	err := json.Unmarshal([]byte(jsonString), &reading)

	if err != nil{
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", reading)

}