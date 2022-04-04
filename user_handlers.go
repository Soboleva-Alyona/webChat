package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

var userSessionId int

func indexHandler(c *gin.Context) {
	if userSessionId <= 0 {
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				//"title": "Home Page",
			},
		)
		return
	}
	// Call the HTML method of the Context to render a template
	c.Set("friend_id", c.PostForm("friend_id"))
	if c.PostForm("friend_id") != "" {
		c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("chat/%d", c.PostForm("friend_id")))
		return
	}
	render(c, gin.H{}, "index.html")

}

func enterHandler(c *gin.Context) {

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"enterPage.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			//"title": "Home Page",
		},
	)

}

func registerHandler(c *gin.Context) {
	if userSessionId != 0 {
		c.Redirect(http.StatusFound, "/")
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusFound,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				//"title": "Home Page",
				"Message": "You are logged in",
			},
		)
	}
	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"registerPage.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			//"title": "Home Page",
		},
	)

}

func doEnterHandler(c *gin.Context) {
	username := c.PostForm("login")
	password := c.PostForm("password")

	var sameSiteCookie http.SameSite
	user := getUsersByLoginAndPass(username, password)
	// Check if the username/password combination is valid
	if user.Id > 0 {
		userSessionId = user.Id
		// If the username/password is valid set the token in a cookie
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.SetSameSite(sameSiteCookie)

		c.Set("is_logged_in", true)

		render(c, gin.H{
			"title": "Successful Login", "ErrorMessage": "Invalid credentials provided"}, "index.html")

		//c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/user/%d", user.Id))

	} else {
		// If the username/password combination is invalid,
		// show the error message on the login page

		c.HTML(http.StatusBadRequest, "index.html", /*gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"}*/gin.H{"Message": "Login Failed"})

	}
}

func generateSessionToken() string {
	// We're using a random 16 character string as the session token
	// This is NOT a secure way of generating session tokens
	// DO NOT USE THIS IN PRODUCTION
	return strconv.FormatInt(rand.Int63(), 16)
}

func doRegisterHandler(c *gin.Context) {
	username := c.PostForm("login")
	password := c.PostForm("password")

	var sameSiteCookie http.SameSite

	if !userAlreadyExists(username) {
		_, err := db.Exec(`INSERT INTO "Users"  (login, password) VALUES ($1, $2)`, username, password)
		if err != nil {
			fmt.Println("Can't insert user into table")
		}

		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.SetSameSite(sameSiteCookie)

		c.Set("is_logged_in", true)

		render(c, gin.H{
			"title": "Successful Register", "ErrorMessage": "Invalid credentials provided"}, "index.html")

		//c.Redirect(http.StatusFound, "/")

	} else {
		//c.Set("Message", "You are already registered, try login")
		c.HTML(http.StatusBadRequest, "index.html", /*gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"}*/gin.H{"Message": "You are already registered, try to login"})

		//c.Redirect(http.StatusFound, "")

	}

}

func getUsersByLoginAndPass(login, password string) UserData {
	for _, u := range getUsersList() {
		if u.Login == login && u.Password == password {
			return u
		}
	}
	return UserData{}
}

func userAlreadyExists(login string) bool {
	res, e := db.Query(`SELECT FROM "Users" WHERE login=$1`, login)
	if e != nil {
		fmt.Println("Problems with db")
		return false
	}

	if !res.Next() {
		return false
	}
	return true
}

func doLogOutHandler(c *gin.Context) {
	c.Set("is_logged_in", false)
	userSessionId = 0
	c.Redirect(http.StatusMovedPermanently, "/")
}
