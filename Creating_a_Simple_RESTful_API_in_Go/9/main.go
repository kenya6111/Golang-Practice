package main

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context){
	c.JSON(200, gin.H{
		"message":"hello world!",
	})
}

func PostHomePage(c *gin.Context){
	body := c.Request.Body
	value,err := io.ReadAll(body)

	if err != nil{
		fmt.Println(err.Error())
	}

	c.JSON(200, gin.H{
		"message":string(value),
	})
}

func QueryStrings (c *gin.Context){
	name:= c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name":name,
		"age":age,
	})
}

func PathParameters (c *gin.Context){
	name:= c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name":name,
		"age":age,
	})
}
func main(){
	r := gin.Default()

	// r.GET("/", func(c *gin.Context){
	// 	c.JSON(200, gin.H{
	// 		"message":"hello world",
	// 	})
	// })
	r.GET("/",HomePage)
	r.POST("/",PostHomePage)
	r.GET("/query",QueryStrings)
	r.GET("/path/:name/:age",PathParameters)
	r.Run()
}