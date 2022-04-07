package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func doAddFriendHandler(c *gin.Context) {
	if userSessionId > 0 {
		userFriend := c.PostForm("user_id")
		_, e := db.Exec(`INSERT INTO "Friends"(
			friend_one, friend_two, status)
		VALUES ($1, $2, $3)`, userSessionId, userFriend, 2)

		_, err := db.Exec(`INSERT INTO "Friends"(
			friend_one, friend_two, status)
		VALUES ($2, $1, $3)`, userSessionId, userFriend, 2)

		if e != nil && err != nil {
			fmt.Println(e.Error())
		}
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}

func chatWithFriendHandler(c *gin.Context) {
	// find messages
	//friendId, _ := c.Get("friend_id")

	//chatId := c.GetInt("chat_id")
	//c.Set("friend_id", friendId)
	//c.Set("chat_id", chatId)
	friendId := c.Param("friend_id")
	friendSessionId, _ = strconv.Atoi(friendId)

	c.Set("friend_id", friendId)
	friendIdInt, _ := strconv.Atoi(friendId)

	getMessages(friendIdInt, userSessionId)

	render(c, gin.H{}, "chat.html")
}

func sendMessageToFriendHandler(c *gin.Context) {
	messageContent := c.PostForm("message_content")
	authorId := userSessionId
	friendId := c.PostForm("friend_id")
	c.Set("friend_id", friendId)
	_, err := db.Query(`INSERT INTO "Messages" (friend_id, content, author) VALUES ($1, $2, $3)`,
		friendId, messageContent, authorId)
	if err != nil {
		fmt.Println("Запись не добавлена")
	}

	render(c, gin.H{}, "chat.html")

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/chat/%s", friendId))
}
