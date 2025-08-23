package main

import (
	"fmt"
	"gin_note/controllers"
	controller_helpers "gin_note/controllers/helpers"
	"gin_note/middlewares"
	"gin_note/models"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)


func main(){
	log.Println("111")
	r := gin.Default()
	r.Use(gin.Logger())

	r.Static("/vendor", "./static/vendor")
	r.LoadHTMLGlob("templates/**/**")

	models.ConnectDatabase()
	models.DBMigrate()

	// session init
	log.Println("222")

	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("notes",store))
	r.Use(middlewares.AuthenticateUser())

	notes := r.Group("/notes")
	{
		r.GET("/",controllers.NotesIndex)
		r.GET("/new",controllers.NotesNew)
		r.POST("/",controllers.NotesCreate)
		r.GET("/:id",controllers.NotesShow)
		r.GET("/edit/:id",controllers.NotesEditPage)
		r.POST("/:id",controllers.NotesUpdate)
		r.DELETE("/:id",controllers.NotesDelete)
	}
	fmt.Println(notes)

	// r.GET("/notes",controllers.NotesIndex)
	// r.GET("/notes/new",controllers.NotesNew)
	// r.POST("/notes",controllers.NotesCreate)
	// r.GET("/notes/:id",controllers.NotesShow)
	// r.GET("/notes/edit/:id",controllers.NotesEditPage)
	// r.POST("/notes/:id",controllers.NotesUpdate)
	// r.DELETE("/notes/:id",controllers.NotesDelete)

	r.GET("/login",controllers.LoginPage)
	r.GET("/signup",controllers.SignUpPage)

	r.POST("/signup",controllers.SignUp)
	r.POST("/login",controllers.Login)
	r.POST("/logout",controllers.Logout)


	r.GET("/",func(c *gin.Context){
		c.HTML(http.StatusOK,
			"home/index.html",
			gin.H{
			"title":"gin notes app",
			"content":"this is my first gin project with index.html",
			"logged_in": controller_helpers.IsUserLoggedIn(c),
		})
	})

	log.Println("server started")
	r.Run()
}