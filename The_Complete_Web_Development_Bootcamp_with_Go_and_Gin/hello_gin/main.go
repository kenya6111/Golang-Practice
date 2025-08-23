package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main(){
	r := gin.Default()
	r.Use(gin.Logger())

	r.LoadHTMLGlob("templates/**/**")
	r.Static("/vendor", "./static/vendor")

	r.GET("/",func(c *gin.Context){
		c.HTML(http.StatusOK, "views/index.html",gin.H{
			"title":"hello gin",
			"content":"this is my first gin project with index.html",
		})
	})

	log.Println("server started")
	r.Run()
}