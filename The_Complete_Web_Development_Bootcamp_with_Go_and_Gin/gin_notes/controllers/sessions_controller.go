package controllers

import (
	"fmt"
	"gin_note/helpers"
	"gin_note/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func LoginPage(c *gin.Context){
	c.HTML(
		http.StatusOK,
		"home/login.html",
		gin.H{},
	)
}
func SignUpPage(c *gin.Context){

	log.Println("signuppage")
	c.HTML(
		http.StatusOK,
		"home/signup.html",
		gin.H{},
	)
}

func SignUp(c *gin.Context){
	log.Println(11)
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")

	available := models.UserCheckAvailability(email)

	fmt.Println(available)

	if !available{
		c.HTML(
			http.StatusIMUsed,
			"home/signup.html",
			gin.H{
				"alert":"Email aleady exists!",
			},
		)
		return
	}

	if password != confirmPassword{
		c.HTML(
			http.StatusNotAcceptable,
			"home/signup.html",
			gin.H{
				"alert":"Password missmatch",
			},
		)
		return
	}
	user := models.UserCreate(email,password)

	if user.ID == 0{
		c.HTML(
			http.StatusNotAcceptable,
			"home/signup.html",
			gin.H{
				"alert":"unable to create user!",
			},
		)
	} else{
		helpers.SessionSet(c, user.ID)
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}


func Login(c *gin.Context){
	email := c.PostForm("email")
	password := c.PostForm("password")
	user := models.UserCheck(email,password)

	if user != nil{
		helpers.SessionSet(c, user.ID)
		c.Redirect(http.StatusMovedPermanently,"/")
	} else {
		c.HTML(
			http.StatusOK,
			"home/login.html",
			gin.H{
				"alert":"email and/or password mismatch!",
			},
		)
	}
}

func Logout(c *gin.Context){
	helpers.SessionClear(c)
	c.HTML(
		http.StatusOK,
		"home/login.html",
		gin.H{
			"alert":"logout",
		},
	)
}
