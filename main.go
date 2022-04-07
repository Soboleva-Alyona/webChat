package main

import (
	"fmt"
	_ "github.com/gin-gonic/contrib/sessions"
	_ "github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

var router *gin.Engine

func main() {
	//http.HandleFunc("/", indexHandler)
	//http.HandleFunc("/enter", enterHandler)
	//log.Fatal(http.ListenAndServe("localhost:8080", nil))
	var e error
	e, db = connectAndGet()
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println("hi")

	res, err := db.Exec(`INSERT INTO "Users" ("login", "password") VALUES ('mee', '123456')`)
	if err != nil {
		fmt.Println(err.Error())
	}
	if res == nil {
		fmt.Println("User already exists")
	}

	router = gin.Default()
	router.LoadHTMLGlob("webapp/*")

	initializeRouts()

	router.Run("localhost:8181")

}

var friendSessionId int

func render(c *gin.Context, data gin.H, templateName string) {
	if userSessionId > 0 {
		c.Set("is_logged_in", true)
	}
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	data["friends"] = getUserFriendsList()
	data["user_me"] = userSessionId
	data["Users"] = getUsersList()
	data["friend_id"], _ = c.Get("friend_id")
	data["messages"] = getMessages(friendSessionId, userSessionId)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)

	}

}
